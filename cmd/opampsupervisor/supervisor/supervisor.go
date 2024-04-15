// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"text/template"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/knadh/koanf/maps"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/oklog/ulid/v2"
	"github.com/open-telemetry/opamp-go/client"
	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
	"github.com/open-telemetry/opamp-go/server"
	serverTypes "github.com/open-telemetry/opamp-go/server/types"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	semconv "go.opentelemetry.io/collector/semconv/v1.21.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/commander"
	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/healthchecker"
)

var (
	//go:embed templates/bootstrap_pipeline.yaml
	bootstrapConfTpl string

	//go:embed templates/extraconfig.yaml
	extraConfigTpl string

	//go:embed templates/opampextension.yaml
	opampextensionTpl string

	//go:embed templates/owntelemetry.yaml
	ownTelemetryTpl string
)

// Supervisor implements supervising of OpenTelemetry Collector and uses OpAMPClient
// to work with an OpAMP Server.
type Supervisor struct {
	logger *zap.Logger

	// Commander that starts/stops the Agent process.
	commander *commander.Commander

	startedAt time.Time

	healthCheckTicker  *backoff.Ticker
	healthChecker      *healthchecker.HTTPHealthChecker
	lastHealthCheckErr error

	// Supervisor's own config.
	config config.Supervisor

	agentDescription *protobufs.AgentDescription

	// Agent's instance id.
	instanceID ulid.ULID

	bootstrapTemplate      *template.Template
	opampextensionTemplate *template.Template
	extraConfigTemplate    *template.Template
	ownTelemetryTemplate   *template.Template

	// A config section to be added to the Collector's config to fetch its own metrics.
	// TODO: store this persistently so that when starting we can compose the effective
	// config correctly.
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21078
	agentConfigOwnMetricsSection *atomic.Value

	// agentHealthCheckEndpoint is the endpoint the Collector's health check extension
	// will listen on for health check requests from the Supervisor.
	agentHealthCheckEndpoint string

	// Supervisor-assembled config to be given to the Collector.
	mergedConfig *atomic.Value

	// Final effective config of the Collector.
	effectiveConfig *atomic.Value

	// Location of the effective config file.
	effectiveConfigFilePath string

	// Last received remote config.
	remoteConfig *protobufs.AgentRemoteConfig

	// A channel to indicate there is a new config to apply.
	hasNewConfig chan struct{}

	// The OpAMP client to connect to the OpAMP Server.
	opampClient client.OpAMPClient

	shuttingDown *atomic.Bool
	supervisorWG sync.WaitGroup

	agentHasStarted               bool
	agentStartHealthCheckAttempts int

	connectedToOpAMPServer chan struct{}

	// The OpAMP server to communicate with the Collector's OpAMP extension
	opampServer     server.OpAMPServer
	opampServerPort int
}

func NewSupervisor(logger *zap.Logger, configFile string) (*Supervisor, error) {
	s := &Supervisor{
		logger:                       logger,
		hasNewConfig:                 make(chan struct{}, 1),
		effectiveConfigFilePath:      "effective.yaml",
		shuttingDown:                 &atomic.Bool{},
		agentConfigOwnMetricsSection: &atomic.Value{},
		mergedConfig:                 &atomic.Value{},
		connectedToOpAMPServer:       make(chan struct{}),
		effectiveConfig:              &atomic.Value{},
	}
	if err := s.createTemplates(); err != nil {
		return nil, err
	}

	if err := s.loadConfig(configFile); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	id, err := s.createInstanceID()
	if err != nil {
		return nil, err
	}

	s.instanceID = id

	if err = s.getBootstrapInfo(); err != nil {
		return nil, fmt.Errorf("could not get bootstrap info from the Collector: %w", err)
	}

	port, err := s.findRandomPort()

	if err != nil {
		return nil, fmt.Errorf("could not find port for health check: %w", err)
	}

	s.agentHealthCheckEndpoint = fmt.Sprintf("localhost:%d", port)

	logger.Debug("Supervisor starting",
		zap.String("id", s.instanceID.String()))

	s.loadAgentEffectiveConfig()

	if err = s.startOpAMP(); err != nil {
		return nil, fmt.Errorf("cannot start OpAMP client: %w", err)
	}

	if connErr := s.waitForOpAMPConnection(); connErr != nil {
		return nil, fmt.Errorf("failed to connect to the OpAMP server: %w", err)
	}

	s.commander, err = commander.NewCommander(
		s.logger,
		s.config.Agent,
		"--config", s.effectiveConfigFilePath,
	)
	if err != nil {
		return nil, err
	}

	s.startHealthCheckTicker()

	s.supervisorWG.Add(1)
	go func() {
		defer s.supervisorWG.Done()
		s.runAgentProcess()
	}()

	return s, nil
}

func (s *Supervisor) createTemplates() error {
	var err error

	if s.bootstrapTemplate, err = template.New("bootstrap").Parse(bootstrapConfTpl); err != nil {
		return err
	}
	if s.extraConfigTemplate, err = template.New("extraconfig").Parse(extraConfigTpl); err != nil {
		return err
	}
	if s.opampextensionTemplate, err = template.New("opampextension").Parse(opampextensionTpl); err != nil {
		return err
	}
	if s.ownTelemetryTemplate, err = template.New("owntelemetry").Parse(ownTelemetryTpl); err != nil {
		return err
	}

	return nil
}

func (s *Supervisor) loadConfig(configFile string) error {
	if configFile == "" {
		return errors.New("path to config file cannot be empty")
	}

	k := koanf.New("::")
	if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
		return err
	}

	decodeConf := koanf.UnmarshalConf{
		Tag: "mapstructure",
	}

	if err := k.UnmarshalWithConf("", &s.config, decodeConf); err != nil {
		return fmt.Errorf("cannot parse %v: %w", configFile, err)
	}

	return nil
}

// getBootstrapInfo obtains the Collector's agent description by
// starting a Collector with a specific config that only start's
// an OpAMP extension, obtains the agent description, then
// shuts down the Collector. This only needs to happen
// once per Collector binary.
func (s *Supervisor) getBootstrapInfo() (err error) {
	port, err := s.findRandomPort()
	if err != nil {
		return err
	}

	s.opampServerPort, err = s.findRandomPort()
	if err != nil {
		return err
	}

	var k = koanf.New(".")

	var cfg bytes.Buffer
	err = s.bootstrapTemplate.Execute(&cfg, map[string]any{
		"EndpointPort":   port,
		"InstanceUid":    s.instanceID.String(),
		"SupervisorPort": s.opampServerPort,
	})
	if err != nil {
		return err
	}

	if err = k.Load(rawbytes.Provider(cfg.Bytes()), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc)); err != nil {
		return err
	}
	if err = k.Load(rawbytes.Provider(s.composeOpAMPExtensionConfig()), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc)); err != nil {
		return err
	}

	effectiveConfigBytes, err := k.Marshal(yaml.Parser())
	if err != nil {
		return err
	}

	s.writeEffectiveConfigToFile(string(effectiveConfigBytes), s.effectiveConfigFilePath)

	srv := server.New(newLoggerFromZap(s.logger))

	done := make(chan error, 1)
	var connected atomic.Bool

	// Start a one-shot server to get the Collector's agent description
	// using the Collector's OpAMP extension.
	err = srv.Start(newServerSettings(flattenedSettings{
		endpoint: fmt.Sprintf("localhost:%d", s.opampServerPort),
		onConnectingFunc: func(_ *http.Request) {
			connected.Store(true)

		},
		onMessageFunc: func(_ serverTypes.Connection, message *protobufs.AgentToServer) {
			if message.AgentDescription != nil {
				instanceIDSeen := false
				s.agentDescription = message.AgentDescription
				identAttr := s.agentDescription.IdentifyingAttributes

				for _, attr := range identAttr {
					if attr.Key == semconv.AttributeServiceInstanceID {
						// TODO: Consider whether to attempt restarting the Collector.
						// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/29864
						if attr.Value.GetStringValue() != s.instanceID.String() {
							done <- fmt.Errorf(
								"the Collector's instance ID (%s) does not match with the instance ID set by the Supervisor (%s)",
								attr.Value.GetStringValue(),
								s.instanceID.String())
							return
						}
						instanceIDSeen = true
					}
				}

				if !instanceIDSeen {
					done <- errors.New("the Collector did not specify an instance ID in its AgentDescription message")
					return
				}

				done <- nil
			}
		},
	}))
	if err != nil {
		return err
	}

	cmd, err := commander.NewCommander(
		s.logger,
		s.config.Agent,
		"--config", s.effectiveConfigFilePath,
	)
	if err != nil {
		return err
	}

	if err = cmd.Start(context.Background()); err != nil {
		return err
	}

	select {
	// TODO make timeout configurable
	case <-time.After(3 * time.Second):
		if connected.Load() {
			return errors.New("collector connected but never responded with an AgentDescription message")
		} else {
			return errors.New("collector's OpAMP client never connected to the Supervisor")
		}
	case err = <-done:
		if err != nil {
			return err
		}
	}

	if err = cmd.Stop(context.Background()); err != nil {
		return err
	}

	if err = srv.Stop(context.Background()); err != nil {
		return err
	}

	return nil
}

func (s *Supervisor) Capabilities() protobufs.AgentCapabilities {
	var supportedCapabilities protobufs.AgentCapabilities
	if c := s.config.Capabilities; c != nil {
		// ReportsEffectiveConfig is set if unspecified or explicitly set to true.
		if (c.ReportsEffectiveConfig != nil && *c.ReportsEffectiveConfig) || c.ReportsEffectiveConfig == nil {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_ReportsEffectiveConfig
		}

		// ReportsHealth is set if unspecified or explicitly set to true.
		if (c.ReportsHealth != nil && *c.ReportsHealth) || c.ReportsHealth == nil {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_ReportsHealth
		}

		// ReportsOwnMetrics is set if unspecified or explicitly set to true.
		if (c.ReportsOwnMetrics != nil && *c.ReportsOwnMetrics) || c.ReportsOwnMetrics == nil {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_ReportsOwnMetrics
		}

		if c.AcceptsRemoteConfig != nil && *c.AcceptsRemoteConfig {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_AcceptsRemoteConfig
		}

		if c.ReportsRemoteConfig != nil && *c.ReportsRemoteConfig {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_ReportsRemoteConfig
		}

		if c.AcceptsOpAMPConnectionSettings != nil && *c.AcceptsOpAMPConnectionSettings {
			supportedCapabilities |= protobufs.AgentCapabilities_AgentCapabilities_AcceptsOpAMPConnectionSettings
		}
	}
	return supportedCapabilities
}

func (s *Supervisor) startOpAMP() error {
	s.opampClient = client.NewWebSocket(newLoggerFromZap(s.logger))

	tlsConfig, err := s.config.Server.TLSSetting.LoadTLSConfig()
	if err != nil {
		return err
	}

	s.logger.Debug("Connecting to OpAMP server...", zap.String("endpoint", s.config.Server.Endpoint), zap.Any("headers", s.config.Server.Headers))
	settings := types.StartSettings{
		OpAMPServerURL: s.config.Server.Endpoint,
		Header:         s.config.Server.Headers,
		TLSConfig:      tlsConfig,
		InstanceUid:    s.instanceID.String(),
		Callbacks: types.CallbacksStruct{
			OnConnectFunc: func(_ context.Context) {
				s.connectedToOpAMPServer <- struct{}{}
				s.logger.Debug("Connected to the server.")
			},
			OnConnectFailedFunc: func(_ context.Context, err error) {
				s.logger.Error("Failed to connect to the server", zap.Error(err))
			},
			OnErrorFunc: func(_ context.Context, err *protobufs.ServerErrorResponse) {
				s.logger.Error("Server returned an error response", zap.String("message", err.ErrorMessage))
			},
			OnMessageFunc: s.onMessage,
			OnOpampConnectionSettingsFunc: func(ctx context.Context, settings *protobufs.OpAMPConnectionSettings) error {
				//nolint:errcheck
				go s.onOpampConnectionSettings(ctx, settings)
				return nil
			},
			OnCommandFunc: func(_ context.Context, command *protobufs.ServerToAgentCommand) error {
				cmdType := command.GetType()
				if *cmdType.Enum() == protobufs.CommandType_CommandType_Restart {
					// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21077
					s.logger.Debug("Received restart command")
				}
				return nil
			},
			SaveRemoteConfigStatusFunc: func(_ context.Context, _ *protobufs.RemoteConfigStatus) {
				// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21079
			},
			GetEffectiveConfigFunc: func(_ context.Context) (*protobufs.EffectiveConfig, error) {
				return s.createEffectiveConfigMsg(), nil
			},
		},
		Capabilities: s.Capabilities(),
	}
	if err = s.opampClient.SetAgentDescription(s.agentDescription); err != nil {
		return err
	}

	if err = s.opampClient.SetHealth(&protobufs.ComponentHealth{Healthy: false}); err != nil {
		return err
	}

	s.logger.Debug("Starting OpAMP client...")
	if err = s.opampClient.Start(context.Background(), settings); err != nil {
		return err
	}
	s.logger.Debug("OpAMP client started.")

	s.logger.Debug("Starting OpAMP server...")
	if err = s.startOpAMPServer(); err != nil {
		return err
	}
	s.logger.Debug("OpAMP server started.")

	return nil
}

// startOpAMPServer starts an OpAMP server that will communicate
// with an OpAMP extension running inside a Collector to receive
// data from inside the Collector. The internal server's lifetime is not
// matched to the Collector's process, but may be restarted
// depending on information received by the Supervisor from the remote
// OpAMP server.
func (s *Supervisor) startOpAMPServer() error {
	s.opampServer = server.New(newLoggerFromZap(s.logger))

	var err error
	s.opampServerPort, err = s.findRandomPort()
	if err != nil {
		return err
	}

	err = s.opampServer.Start(newServerSettings(flattenedSettings{
		endpoint:         fmt.Sprintf("localhost:%d", s.opampServerPort),
		onConnectingFunc: func(_ *http.Request) {},
		onMessageFunc: func(_ serverTypes.Connection, message *protobufs.AgentToServer) {
			s.logger.Debug("Received message")
			if message.AgentDescription != nil {
				s.agentDescription = message.AgentDescription
			}
			if message.EffectiveConfig != nil {
				s.logger.Debug("Setting confmap")
				s.effectiveConfig.Store(string(message.EffectiveConfig.ConfigMap.ConfigMap[""].Body))
				err = s.opampClient.UpdateEffectiveConfig(context.Background())
				if err != nil {
					s.logger.Error("The OpAMP client failed to update the effective config", zap.Error(err))
				}
			}
		},
	}))
	if err != nil {
		return err
	}

	return nil
}

func (s *Supervisor) stopOpAMP() error {
	s.logger.Debug("Stopping OpAMP client...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.opampClient.Stop(ctx)
	// TODO(srikanthccv): remove context.DeadlineExceeded after https://github.com/open-telemetry/opamp-go/pull/213
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	s.logger.Debug("OpAMP client stopped.")
	return nil
}

func (s *Supervisor) getHeadersFromSettings(protoHeaders *protobufs.Headers) http.Header {
	headers := make(http.Header)
	for _, header := range protoHeaders.Headers {
		headers.Add(header.Key, header.Value)
	}
	return headers
}

func (s *Supervisor) onOpampConnectionSettings(_ context.Context, settings *protobufs.OpAMPConnectionSettings) error {
	if settings == nil {
		s.logger.Debug("Received ConnectionSettings request with nil settings")
		return nil
	}

	newServerConfig := &config.OpAMPServer{}

	if settings.DestinationEndpoint != "" {
		newServerConfig.Endpoint = settings.DestinationEndpoint
	}
	if settings.Headers != nil {
		newServerConfig.Headers = s.getHeadersFromSettings(settings.Headers)
	}
	if settings.Certificate != nil {
		if len(settings.Certificate.CaPublicKey) != 0 {
			newServerConfig.TLSSetting.CAPem = configopaque.String(settings.Certificate.CaPublicKey)
		}
		if len(settings.Certificate.PublicKey) != 0 {
			newServerConfig.TLSSetting.CertPem = configopaque.String(settings.Certificate.PublicKey)
		}
		if len(settings.Certificate.PrivateKey) != 0 {
			newServerConfig.TLSSetting.KeyPem = configopaque.String(settings.Certificate.PrivateKey)
		}
	} else {
		newServerConfig.TLSSetting = configtls.ClientConfig{Insecure: true}
	}

	if err := s.stopOpAMP(); err != nil {
		s.logger.Error("Cannot stop the OpAMP client", zap.Error(err))
		return err
	}

	// take a copy of the current OpAMP server config
	oldServerConfig := s.config.Server
	// update the OpAMP server config
	s.config.Server = newServerConfig

	if err := s.startOpAMP(); err != nil {
		s.logger.Error("Cannot connect to the OpAMP server using the new settings", zap.Error(err))
		// revert the OpAMP server config
		s.config.Server = oldServerConfig
		// start the OpAMP client with the old settings
		if err := s.startOpAMP(); err != nil {
			s.logger.Error("Cannot reconnect to the OpAMP server after restoring old settings", zap.Error(err))
			return err
		}
	}
	return s.waitForOpAMPConnection()
}

func (s *Supervisor) waitForOpAMPConnection() error {
	// wait for the OpAMP client to connect to the server or timeout
	select {
	case <-s.connectedToOpAMPServer:
		return nil
	case <-time.After(10 * time.Second):
		return errors.New("timed out waiting for the server to connect")
	}
}

// TODO: Persist instance ID. https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21073
func (s *Supervisor) createInstanceID() (ulid.ULID, error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(0)), 0)
	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)

	if err != nil {
		return ulid.ULID{}, err
	}

	return id, nil

}

func (s *Supervisor) composeExtraLocalConfig() []byte {
	var cfg bytes.Buffer
	resourceAttrs := map[string]string{}
	for _, attr := range s.agentDescription.IdentifyingAttributes {
		resourceAttrs[attr.Key] = attr.Value.GetStringValue()
	}
	for _, attr := range s.agentDescription.NonIdentifyingAttributes {
		resourceAttrs[attr.Key] = attr.Value.GetStringValue()
	}
	tplVars := map[string]any{
		"Healthcheck":        s.agentHealthCheckEndpoint,
		"ResourceAttributes": resourceAttrs,
		"SupervisorPort":     s.opampServerPort,
	}
	err := s.extraConfigTemplate.Execute(
		&cfg,
		tplVars,
	)
	if err != nil {
		s.logger.Error("Could not compose local config", zap.Error(err))
		return nil
	}

	return cfg.Bytes()
}

func (s *Supervisor) composeOpAMPExtensionConfig() []byte {
	var cfg bytes.Buffer
	tplVars := map[string]any{
		"InstanceUid":    s.instanceID.String(),
		"SupervisorPort": s.opampServerPort,
	}
	err := s.opampextensionTemplate.Execute(
		&cfg,
		tplVars,
	)
	if err != nil {
		s.logger.Error("Could not compose local config", zap.Error(err))
		return nil
	}

	return cfg.Bytes()
}

func (s *Supervisor) loadAgentEffectiveConfig() {
	var effectiveConfigBytes []byte

	effFromFile, err := os.ReadFile(s.effectiveConfigFilePath)
	if err == nil {
		// We have an effective config file.
		effectiveConfigBytes = effFromFile
	} else {
		// No effective config file, just use the initial config.
		effectiveConfigBytes = s.composeExtraLocalConfig()
	}

	s.mergedConfig.Store(string(effectiveConfigBytes))
}

// createEffectiveConfigMsg create an EffectiveConfig with the content of the
// current effective config.
func (s *Supervisor) createEffectiveConfigMsg() *protobufs.EffectiveConfig {
	cfgStr, ok := s.effectiveConfig.Load().(string)
	if !ok {
		cfgStr, ok = s.mergedConfig.Load().(string)
		if !ok {
			cfgStr = ""
		}
	}

	cfg := &protobufs.EffectiveConfig{
		ConfigMap: &protobufs.AgentConfigMap{
			ConfigMap: map[string]*protobufs.AgentConfigFile{
				"": {Body: []byte(cfgStr)},
			},
		},
	}

	return cfg
}

func (s *Supervisor) setupOwnMetrics(_ context.Context, settings *protobufs.TelemetryConnectionSettings) (configChanged bool) {
	var cfg bytes.Buffer
	if settings.DestinationEndpoint == "" {
		// No destination. Disable metric collection.
		s.logger.Debug("Disabling own metrics pipeline in the config")
	} else {
		s.logger.Debug("Enabling own metrics pipeline in the config")

		port, err := s.findRandomPort()

		if err != nil {
			s.logger.Error("Could not setup own metrics", zap.Error(err))
			return
		}

		err = s.ownTelemetryTemplate.Execute(
			&cfg,
			map[string]any{
				"PrometheusPort":  port,
				"MetricsEndpoint": settings.DestinationEndpoint,
			},
		)
		if err != nil {
			s.logger.Error("Could not setup own metrics", zap.Error(err))
			return
		}

	}
	s.agentConfigOwnMetricsSection.Store(cfg.String())

	// Need to recalculate the Agent config so that the metric config is included in it.
	configChanged, err := s.recalcEffectiveConfig()
	if err != nil {
		return
	}

	return configChanged
}

// composeEffectiveConfig composes the effective config from multiple sources:
// 1) the remote config from OpAMP Server
// 2) the own metrics config section
// 3) the local override config that is hard-coded in the Supervisor.
func (s *Supervisor) composeEffectiveConfig(config *protobufs.AgentRemoteConfig) (configChanged bool, err error) {
	var k = koanf.New(".")

	// Sort to make sure the order of merging is stable.
	var names []string
	for name := range config.Config.ConfigMap {
		if name == "" {
			// skip instance config
			continue
		}
		names = append(names, name)
	}

	sort.Strings(names)

	// Append instance config as the last item.
	names = append(names, "")

	// Merge received configs.
	for _, name := range names {
		item := config.Config.ConfigMap[name]
		var k2 = koanf.New(".")
		err = k2.Load(rawbytes.Provider(item.Body), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc))
		if err != nil {
			return false, fmt.Errorf("cannot parse config named %s: %w", name, err)
		}
		err = k.Merge(k2)
		if err != nil {
			return false, fmt.Errorf("cannot merge config named %s: %w", name, err)
		}
	}

	// Merge own metrics config.
	ownMetricsCfg, ok := s.agentConfigOwnMetricsSection.Load().(string)
	if ok {
		if err = k.Load(rawbytes.Provider([]byte(ownMetricsCfg)), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc)); err != nil {
			return false, err
		}
	}

	// Merge local config last since it has the highest precedence.
	if err = k.Load(rawbytes.Provider(s.composeExtraLocalConfig()), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc)); err != nil {
		return false, err
	}

	if err = k.Load(rawbytes.Provider(s.composeOpAMPExtensionConfig()), yaml.Parser(), koanf.WithMergeFunc(configMergeFunc)); err != nil {
		return false, err
	}

	// The merged final result is our effective config.
	effectiveConfigBytes, err := k.Marshal(yaml.Parser())
	if err != nil {
		return false, err
	}

	// Check if effective config is changed.
	newEffectiveConfig := string(effectiveConfigBytes)
	configChanged = false
	if s.mergedConfig.Load().(string) != newEffectiveConfig {
		s.logger.Debug("Effective config changed.")
		s.mergedConfig.Store(newEffectiveConfig)
		configChanged = true
	}

	return configChanged, nil
}

// Recalculate the Agent's effective config and if the config changes, signal to the
// background goroutine that the config needs to be applied to the Agent.
func (s *Supervisor) recalcEffectiveConfig() (configChanged bool, err error) {
	configChanged, err = s.composeEffectiveConfig(s.remoteConfig)
	if err != nil {
		s.logger.Error("Error composing effective config. Ignoring received config", zap.Error(err))
		return configChanged, err
	}

	return configChanged, nil
}

func (s *Supervisor) startAgent() {
	err := s.commander.Start(context.Background())
	if err != nil {
		s.logger.Error("Cannot start the agent", zap.Error(err))
		err = s.opampClient.SetHealth(&protobufs.ComponentHealth{Healthy: false, LastError: fmt.Sprintf("Cannot start the agent: %v", err)})

		if err != nil {
			s.logger.Error("Failed to report OpAMP client health", zap.Error(err))
		}

		return
	}

	s.agentHasStarted = false
	s.agentStartHealthCheckAttempts = 0
	s.startedAt = time.Now()
	s.startHealthCheckTicker()

	s.healthChecker = healthchecker.NewHTTPHealthChecker(fmt.Sprintf("http://%s", s.agentHealthCheckEndpoint))
}

func (s *Supervisor) startHealthCheckTicker() {
	// Prepare health checker
	healthCheckBackoff := backoff.NewExponentialBackOff()
	healthCheckBackoff.MaxInterval = 60 * time.Second
	healthCheckBackoff.MaxElapsedTime = 0 // Never stop
	if s.healthCheckTicker != nil {
		s.healthCheckTicker.Stop()
	}
	s.healthCheckTicker = backoff.NewTicker(healthCheckBackoff)
}

func (s *Supervisor) healthCheck() {
	if !s.commander.IsRunning() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	err := s.healthChecker.Check(ctx)
	cancel()

	if errors.Is(err, s.lastHealthCheckErr) {
		// No difference from last check. Nothing new to report.
		return
	}

	// Prepare OpAMP health report.
	health := &protobufs.ComponentHealth{
		StartTimeUnixNano: uint64(s.startedAt.UnixNano()),
	}

	if err != nil {
		health.Healthy = false
		if !s.agentHasStarted && s.agentStartHealthCheckAttempts < 10 {
			health.LastError = "Agent is starting"
			s.agentStartHealthCheckAttempts++
		} else {
			health.LastError = err.Error()
			s.logger.Error("Agent is not healthy", zap.Error(err))
		}
	} else {
		s.agentHasStarted = true
		health.Healthy = true
		s.logger.Debug("Agent is healthy.")
	}

	// Report via OpAMP.
	if err2 := s.opampClient.SetHealth(health); err2 != nil {
		s.logger.Error("Could not report health to OpAMP server", zap.Error(err2))
		return
	}

	s.lastHealthCheckErr = err
}

func (s *Supervisor) runAgentProcess() {
	if _, err := os.Stat(s.effectiveConfigFilePath); err == nil {
		// We have an effective config file saved previously. Use it to start the agent.
		s.startAgent()
	}

	restartTimer := time.NewTimer(0)
	restartTimer.Stop()

	for {
		select {
		case <-s.hasNewConfig:
			restartTimer.Stop()
			s.stopAgentApplyConfig()
			s.startAgent()

		case <-s.commander.Done():
			if s.shuttingDown.Load() {
				return
			}

			s.logger.Debug("Agent process exited unexpectedly. Will restart in a bit...", zap.Int("pid", s.commander.Pid()), zap.Int("exit_code", s.commander.ExitCode()))
			errMsg := fmt.Sprintf(
				"Agent process PID=%d exited unexpectedly, exit code=%d. Will restart in a bit...",
				s.commander.Pid(), s.commander.ExitCode(),
			)
			err := s.opampClient.SetHealth(&protobufs.ComponentHealth{Healthy: false, LastError: errMsg})

			if err != nil {
				s.logger.Error("Could not report health to OpAMP server", zap.Error(err))
			}

			// TODO: decide why the agent stopped. If it was due to bad config, report it to server.
			// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21079

			// Wait 5 seconds before starting again.
			restartTimer.Stop()
			restartTimer.Reset(5 * time.Second)

		case <-restartTimer.C:
			s.startAgent()

		case <-s.healthCheckTicker.C:
			s.healthCheck()
		}
	}
}

func (s *Supervisor) stopAgentApplyConfig() {
	s.logger.Debug("Stopping the agent to apply new config")
	cfg := s.mergedConfig.Load().(string)
	err := s.commander.Stop(context.Background())

	if err != nil {
		s.logger.Error("Could not stop agent process", zap.Error(err))
	}

	s.writeEffectiveConfigToFile(cfg, s.effectiveConfigFilePath)
}

func (s *Supervisor) writeEffectiveConfigToFile(cfg string, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		s.logger.Error("Cannot create effective config file", zap.Error(err))
	}
	defer f.Close()

	_, err = f.WriteString(cfg)

	if err != nil {
		s.logger.Error("Cannot write effective config file", zap.Error(err))
	}
}

func (s *Supervisor) Shutdown() {
	s.logger.Debug("Supervisor shutting down...")
	s.shuttingDown.Store(true)
	if s.commander != nil {
		err := s.commander.Stop(context.Background())

		if err != nil {
			s.logger.Error("Could not stop agent process", zap.Error(err))
		}
	}

	if s.opampClient != nil {
		err := s.opampClient.SetHealth(
			&protobufs.ComponentHealth{
				Healthy: false, LastError: "Supervisor is shutdown",
			},
		)

		if err != nil {
			s.logger.Error("Could not report health to OpAMP server", zap.Error(err))
		}

		err = s.stopOpAMP()

		if err != nil {
			s.logger.Error("Could not stop the OpAMP client", zap.Error(err))
		}
	}

	if s.healthCheckTicker != nil {
		s.healthCheckTicker.Stop()
	}

	s.supervisorWG.Wait()
}

func (s *Supervisor) onMessage(ctx context.Context, msg *types.MessageData) {
	configChanged := false
	if msg.RemoteConfig != nil {
		s.remoteConfig = msg.RemoteConfig
		s.logger.Debug("Received remote config from server", zap.String("hash", fmt.Sprintf("%x", s.remoteConfig.ConfigHash)))

		var err error
		configChanged, err = s.recalcEffectiveConfig()
		if err != nil {
			err = s.opampClient.SetRemoteConfigStatus(&protobufs.RemoteConfigStatus{
				LastRemoteConfigHash: msg.RemoteConfig.ConfigHash,
				Status:               protobufs.RemoteConfigStatuses_RemoteConfigStatuses_FAILED,
				ErrorMessage:         err.Error(),
			})
			if err != nil {
				s.logger.Error("Could not report failed OpAMP remote config status", zap.Error(err))
			}
		} else {
			err = s.opampClient.SetRemoteConfigStatus(&protobufs.RemoteConfigStatus{
				LastRemoteConfigHash: msg.RemoteConfig.ConfigHash,
				Status:               protobufs.RemoteConfigStatuses_RemoteConfigStatuses_APPLIED,
			})
			if err != nil {
				s.logger.Error("Could not report applied OpAMP remote config status", zap.Error(err))
			}
		}
	}

	if msg.OwnMetricsConnSettings != nil {
		configChanged = s.setupOwnMetrics(ctx, msg.OwnMetricsConnSettings) || configChanged
	}

	if msg.AgentIdentification != nil {
		newInstanceID, err := ulid.Parse(msg.AgentIdentification.NewInstanceUid)
		if err != nil {
			s.logger.Error("Failed to parse instance ULID", zap.Error(err))
		}

		s.logger.Debug("Agent identity is changing",
			zap.String("old_id", s.instanceID.String()),
			zap.String("new_id", newInstanceID.String()))
		s.instanceID = newInstanceID
		err = s.opampClient.SetAgentDescription(s.agentDescription)
		if err != nil {
			s.logger.Error("Failed to send agent description to OpAMP server")
		}

		configChanged = true
	}

	if configChanged {
		err := s.opampClient.UpdateEffectiveConfig(ctx)
		if err != nil {
			s.logger.Error("The OpAMP client failed to update the effective config", zap.Error(err))
		}

		s.logger.Debug("Config is changed. Signal to restart the agent")
		// Signal that there is a new config.
		select {
		case s.hasNewConfig <- struct{}{}:
		default:
		}
	}
}

func (s *Supervisor) findRandomPort() (int, error) {
	l, err := net.Listen("tcp", "localhost:0")

	if err != nil {
		return 0, err
	}

	port := l.Addr().(*net.TCPAddr).Port

	err = l.Close()

	if err != nil {
		return 0, err
	}

	return port, nil
}

// The default koanf behavior is to override lists in the config.
// We want to override that here for extensions.
// Will be resolved by https://github.com/open-telemetry/opentelemetry-collector/issues/8754
func configMergeFunc(src, dest map[string]any) error {
	srcExtensions := maps.Search(src, []string{"service", "extensions"})
	destExtensions := maps.Search(dest, []string{"service", "extensions"})

	maps.Merge(src, dest)

	if destExt, ok := destExtensions.([]any); ok {
		if srcExt, ok := srcExtensions.([]any); ok {
			if service, ok := dest["service"].(map[string]any); ok {
				service["extensions"] = append(destExt, srcExt...)
			}
		}
	}

	return nil
}
