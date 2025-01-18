// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/models"
)

var errClientNotInit = errors.New("client not initialized")

// Names of metrics in message_stats
const (
	deliverStat        = "deliver"
	publishStat        = "publish"
	ackStat            = "ack"
	dropUnroutableStat = "drop_unroutable"
)

// Metrics to gather from queue message_stats structure
var messageStatMetrics = []string{
	deliverStat,
	publishStat,
	ackStat,
	dropUnroutableStat,
}

// rabbitmqScraper handles scraping of RabbitMQ metrics
type rabbitmqScraper struct {
	client   client
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *rabbitmqScraper {
	return &rabbitmqScraper{
		logger:   logger,
		cfg:      cfg,
		settings: settings.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
	}
}

// start starts the scraper by creating a new HTTP Client
func (r *rabbitmqScraper) start(ctx context.Context, host component.Host) error {
	if r.cfg.Endpoint == "" || r.cfg.Username == "" || r.cfg.Password == "" {
		return fmt.Errorf("invalid configuration: missing endpoint, username, or password")
	}

	client, err := newClient(ctx, r.cfg, host, r.settings, r.logger)
	if err != nil {
		return fmt.Errorf("failed to initialize RabbitMQ client: %w", err)
	}

	r.client = client
	return nil
}

// scrape collects metrics from the RabbitMQ API
func (r *rabbitmqScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	now := pcommon.NewTimestampFromTime(time.Now())

	// Validate client initialization
	if r.client == nil {
		return pmetric.NewMetrics(), errClientNotInit
	}

	// Collect queue metrics
	if err := r.collectQueueMetrics(ctx, now); err != nil {
		r.logger.Error("failed to collect queue metrics", zap.Error(err))
		return pmetric.NewMetrics(), err
	}

	// Collect node metrics if enabled
	if r.cfg.EnableNodeMetrics {
		if err := r.collectNodeMetrics(ctx, now); err != nil {
			r.logger.Error("failed to collect node metrics", zap.Error(err))
			return pmetric.NewMetrics(), err
		}
	}

	return r.mb.Emit(), nil
}

func (r *rabbitmqScraper) collectQueueMetrics(ctx context.Context, now pcommon.Timestamp) error {
	queues, err := r.client.GetQueues(ctx)
	if err != nil {
		return err
	}

	for _, queue := range queues {
		r.collectQueue(queue, now)
	}
	return nil
}

func (r *rabbitmqScraper) collectNodeMetrics(ctx context.Context, now pcommon.Timestamp) error {
	nodes, err := r.client.GetNodes(ctx)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		r.collectNode(node, now)
	}
	return nil
}

// collectQueue collects metrics for a specific queue
func (r *rabbitmqScraper) collectQueue(queue *models.Queue, now pcommon.Timestamp) {
	r.mb.RecordRabbitmqConsumerCountDataPoint(now, queue.Consumers)
	r.mb.RecordRabbitmqMessageCurrentDataPoint(now, queue.UnacknowledgedMessages, metadata.AttributeMessageStateUnacknowledged)
	r.mb.RecordRabbitmqMessageCurrentDataPoint(now, queue.ReadyMessages, metadata.AttributeMessageStateReady)

	for _, messageStatMetric := range messageStatMetrics {
		val, ok := queue.MessageStats[messageStatMetric]
		if !ok {
			r.logger.Debug("metric not found", zap.String("metric", messageStatMetric), zap.String("queue", queue.Name))
			continue
		}

		val64, ok := convertValToInt64(val)
		if !ok {
			r.logger.Warn("metric not int64", zap.String("metric", messageStatMetric), zap.String("queue", queue.Name))
			continue
		}

		switch messageStatMetric {
		case deliverStat:
			r.mb.RecordRabbitmqMessageDeliveredDataPoint(now, val64)
		case publishStat:
			r.mb.RecordRabbitmqMessagePublishedDataPoint(now, val64)
		case ackStat:
			r.mb.RecordRabbitmqMessageAcknowledgedDataPoint(now, val64)
		case dropUnroutableStat:
			r.mb.RecordRabbitmqMessageDroppedDataPoint(now, val64)
		}
	}

	rb := r.mb.NewResourceBuilder()
	rb.SetRabbitmqQueueName(queue.Name)
	rb.SetRabbitmqNodeName(queue.Node)
	rb.SetRabbitmqVhostName(queue.VHost)
	r.mb.EmitForResource(metadata.WithResource(rb.Emit()))
}

func (r *rabbitmqScraper) collectNode(node *models.Node, now pcommon.Timestamp) {
	// Record node-specific metrics
	r.mb.RecordRabbitmqNodeDiskFreeDataPoint(now, node.DiskFree)
	r.mb.RecordRabbitmqNodeFdUsedDataPoint(now, node.FDUsed)
	r.mb.RecordRabbitmqNodeMemLimitDataPoint(now, node.MemLimit)
	r.mb.RecordRabbitmqNodeMemUsedDataPoint(now, node.MemUsed)

	// Build resource attributes for the node
	rb := r.mb.NewResourceBuilder()
	rb.SetRabbitmqNodeName(node.Name)

	// Emit resource and log attributes for debugging
	resource := rb.Emit()
	r.logger.Debug("Emitting resource for node", zap.String("attributes", formatResourceAttributes(resource)))

	// Emit metrics for the node
	r.mb.EmitForResource(metadata.WithResource(resource))
}

// formatResourceAttributes converts resource attributes to a string for logging
func formatResourceAttributes(resource pcommon.Resource) string {
	attrs := resource.Attributes()
	var attributes []string
	attrs.Range(func(k string, v pcommon.Value) bool {
		attributes = append(attributes, k+"="+v.AsString())
		return true
	})
	return "{" + strings.Join(attributes, ", ") + "}"
}

// convertValToInt64 converts a value to int64
func convertValToInt64(val any) (int64, bool) {
	f64Val, ok := val.(float64)
	if !ok {
		return 0, false
	}
	return int64(f64Val), true
}
