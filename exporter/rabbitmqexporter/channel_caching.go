package rabbitmqexporter

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"sync"
	"time"
)

type connectionConfig struct {
	logger            *zap.Logger
	connectionUrl     string
	connectionName    string
	channelPoolSize   int
	heartbeatInterval time.Duration
	connectionTimeout time.Duration
	confirmationMode  bool
	//durable           bool // TODO figure out what to do with this
}

type amqpChannelCacher struct {
	logger            *zap.Logger
	config            *connectionConfig
	connection        *amqp.Connection
	connLock          *sync.Mutex
	cachedChannelPool chan *amqpChannelWrapper
	connectionErrors  chan *amqp.Error
}

type amqpChannelWrapper struct {
	id         int
	channel    *amqp.Channel
	wasHealthy bool
	lock       *sync.Mutex
	logger     *zap.Logger
}

func newAmqpChannelCacher(config *connectionConfig) (*amqpChannelCacher, error) {
	acc := &amqpChannelCacher{
		logger:            config.logger,
		config:            config,
		connLock:          &sync.Mutex{},
		cachedChannelPool: make(chan *amqpChannelWrapper, config.channelPoolSize),
	}

	err := acc.connect()
	if err != nil {
		return nil, err
	}

	// Synchronously create and connect to channels
	for i := 0; i < config.channelPoolSize; i++ {
		acc.cachedChannelPool <- acc.createChannelWrapper(i, config.logger)
	}

	return acc, nil
}

func (acc *amqpChannelCacher) connect() error {

	// Compare, Lock, Recompare Strategy
	if acc.connection != nil && !acc.connection.IsClosed() /* <- atomic */ {
		acc.logger.Debug("Already connected before acquiring lock")
		return nil
	}

	acc.connLock.Lock() // Block all but one.
	defer acc.connLock.Unlock()

	// Recompare, check if an operation is still necessary after acquiring lock.
	if acc.connection != nil && !acc.connection.IsClosed() /* <- atomic */ {
		acc.logger.Debug("Already connected after acquiring lock")
		return nil
	}

	// Proceed with reconnectivity
	var amqpConn *amqp.Connection
	var err error

	// TODO TLS config
	amqpConn, err = amqp.DialConfig(acc.config.connectionUrl, amqp.Config{
		Heartbeat: acc.config.heartbeatInterval,
		Dial:      amqp.DefaultDial(acc.config.connectionTimeout),
		Properties: amqp.Table{
			"connection_name": acc.config.connectionName,
		},
	})
	if err != nil {
		return err
	}

	acc.connection = amqpConn

	// Goal is to lazily restore the connection so this needs to be buffered to avoid blocking on asynchronous amqp errors.
	// Also re-create this channel each time because apparently the amqp library can close it
	acc.connectionErrors = make(chan *amqp.Error, 1)
	acc.connection.NotifyClose(acc.connectionErrors)

	// TODO flow control callback
	//acc.Blockers = make(chan amqp.Blocking, 10)
	//acc.connection.NotifyBlocked(acc.Blockers)

	return nil
}

func (acc *amqpChannelCacher) restoreUnhealthyConnection() {
	healthy := true
	select {
	case err := <-acc.connectionErrors:
		healthy = false
		acc.logger.Debug("Received connection error, will retry restoring unhealthy connection", zap.Error(err))
	default:
		break
	}

	if !healthy || acc.connection.IsClosed() {
		// TODO, consider retrying multiple times with some timeout
		if err := acc.connect(); err != nil {
			acc.logger.Warn("Failed attempt at restoring unhealthy connection", zap.Error(err))
		} else {
			acc.logger.Info("Restored unhealthy connection")
		}
	}
}

func (acc *amqpChannelCacher) createChannelWrapper(id int, logger *zap.Logger) *amqpChannelWrapper {
	channelWrapper := &amqpChannelWrapper{id: id, logger: logger, lock: &sync.Mutex{}}
	channelWrapper.tryReplacingChannel(acc.connection, acc.config.confirmationMode)
	return channelWrapper
}

func (acw *amqpChannelWrapper) tryReplacingChannel(connection *amqp.Connection, confirmAcks bool) {
	// TODO consider confirmation mode

	acw.lock.Lock()
	defer acw.lock.Unlock()

	if acw.channel != nil {
		err := acw.channel.Close()
		if err != nil {
			acw.logger.Debug("Error closing existing channel", zap.Error(err))
			acw.wasHealthy = false
			return
		}
	}

	var err error
	acw.channel, err = connection.Channel()

	if err != nil {
		acw.logger.Warn("Channel creation error", zap.Error(err))
		acw.wasHealthy = false
		return
	}

	if confirmAcks {
		err := acw.channel.Confirm(false)
		if err != nil {
			acw.logger.Debug("Error entering confirm mode", zap.Error(err))
			acw.wasHealthy = false
			return
		}
	}

	// TODO consider error callbacks
	//ch.Errors = make(chan *amqp.Error, 100)
	//ch.Channel.NotifyClose(ch.Errors)

	acw.wasHealthy = true
}

func (acc *amqpChannelCacher) getChannelFromPool() *amqpChannelWrapper {
	return <-acc.cachedChannelPool
}

func (acc *amqpChannelCacher) returnChannelToPool(channel *amqpChannelWrapper, isUnhealthy bool) {
	if isUnhealthy {
		acc.reconnectChannel(channel)
	}
	acc.cachedChannelPool <- channel
	return
}

func (acc *amqpChannelCacher) reconnectChannel(channel *amqpChannelWrapper) {
	acc.restoreUnhealthyConnection()
	channel.tryReplacingChannel(acc.connection, acc.config.confirmationMode)
}

func (acc *amqpChannelCacher) close() error {
	err := acc.connection.Close()
	if err != nil {
		acc.logger.Debug("Received error from connection.Close()", zap.Error(err))
		if err != amqp.ErrClosed {
			return err
		}
	}
	return nil
}
