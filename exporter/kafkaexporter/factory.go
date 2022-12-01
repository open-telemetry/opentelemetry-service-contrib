// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typeStr = "kafka"
	// The stability level of the exporter.
	stability           = component.StabilityLevelBeta
	defaultTracesTopic  = "otlp_spans"
	defaultMetricsTopic = "otlp_metrics"
	defaultLogsTopic    = "otlp_logs"
	defaultEncoding     = "otlp_proto"
	defaultBroker       = "localhost:9092"
	// default from sarama.NewConfig()
	defaultMetadataRetryMax = 3
	// default from sarama.NewConfig()
	defaultMetadataRetryBackoff = time.Millisecond * 250
	// default from sarama.NewConfig()
	defaultMetadataFull = true
	// default max.message.bytes for the producer
	defaultProducerMaxMessageBytes = 1000000
	// default required_acks for the producer
	defaultProducerRequiredAcks = sarama.WaitForLocal
	// default from sarama.NewConfig()
	defaultCompression = "none"
	// default from sarama.NewConfig()
	defaultFluxMaxMessages = 0
)

// FactoryOption applies changes to kafkaExporterFactory.
type FactoryOption func(factory *kafkaExporterFactory)

// WithTracesMarshalers adds tracesMarshalers to the exporter factory.
// This allows custom-built collectors to configure custom Kafka marshaler(s) for trace data.
// An example use case is keying the message by trace ID so downstream collectors could do tail-based sampling with horizontal scale
func WithTracesMarshalers(tracesMarshalers ...TracesMarshaler) FactoryOption {
	return func(factory *kafkaExporterFactory) {
		for _, marshaler := range tracesMarshalers {
			factory.tracesMarshalers[marshaler.Encoding()] = marshaler
		}
	}
}

// WithMetricsMarshalers adds metricsMarshalers to the exporter factory.
// This allows custom-built collectors to configure custom Kafka marshaler(s) for metric data.
// An example use case is keying the message by resource attribute values so downstream collectors can be horizontally scaled and still deliver metrics in the correct order to the backend.
// Another use case might be pre-aggregating metrics to reduce backend update throughput.
func WithMetricsMarshalers(metricsMarshalers ...MetricsMarshaler) FactoryOption {
	return func(factory *kafkaExporterFactory) {
		for _, marshaler := range metricsMarshalers {
			factory.metricsMarshalers[marshaler.Encoding()] = marshaler
		}
	}
}

// WithLogsMarshalers adds logsMarshalers to the exporter factory.
// This allows custom-built collectors to configure custom Kafka marshaler(s) for log data.
// An example use case is keying the message by resource attribute values so downstream collectors can be horizontally scaled and still deliver logs in the correct order to the backend.
func WithLogsMarshalers(logsMarshalers ...LogsMarshaler) FactoryOption {
	return func(factory *kafkaExporterFactory) {
		for _, marshaler := range logsMarshalers {
			factory.logsMarshalers[marshaler.Encoding()] = marshaler
		}
	}
}

// WithProducerFactory sets an alternative producer factory.  Primarily for injecting mocks/noop producers for testing
func WithProducerFactory(producerFactory ProducerFactoryFunc) FactoryOption {
	return func(factory *kafkaExporterFactory) {
		factory.producerFactory = producerFactory
	}
}

// NewFactory creates Kafka exporter factory.
func NewFactory(options ...FactoryOption) component.ExporterFactory {
	f := &kafkaExporterFactory{
		tracesMarshalers:  tracesMarshalers(),
		metricsMarshalers: metricsMarshalers(),
		logsMarshalers:    logsMarshalers(),
		producerFactory:   newSaramaProducer,
	}
	for _, o := range options {
		o(f)
	}
	return component.NewExporterFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesExporter(f.createTracesExporter, stability),
		component.WithMetricsExporter(f.createMetricsExporter, stability),
		component.WithLogsExporter(f.createLogsExporter, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		ExporterSettings: config.NewExporterSettings(component.NewID(typeStr)),
		TimeoutSettings:  exporterhelper.NewDefaultTimeoutSettings(),
		RetrySettings:    exporterhelper.NewDefaultRetrySettings(),
		QueueSettings:    exporterhelper.NewDefaultQueueSettings(),
		Brokers:          []string{defaultBroker},
		// using an empty topic to track when it has not been set by user, default is based on traces or metrics.
		Topic:    "",
		Encoding: defaultEncoding,
		Metadata: Metadata{
			Full: defaultMetadataFull,
			Retry: MetadataRetry{
				Max:     defaultMetadataRetryMax,
				Backoff: defaultMetadataRetryBackoff,
			},
		},
		Producer: Producer{
			MaxMessageBytes:  defaultProducerMaxMessageBytes,
			RequiredAcks:     defaultProducerRequiredAcks,
			Compression:      defaultCompression,
			FlushMaxMessages: defaultFluxMaxMessages,
		},
	}
}

// Factory function resposible for creating the SyncProducer
// Primarily used for injecting mock producers for testing
type ProducerFactoryFunc func(config Config) (sarama.SyncProducer, error)

type kafkaExporterFactory struct {
	tracesMarshalers  map[string]TracesMarshaler
	metricsMarshalers map[string]MetricsMarshaler
	logsMarshalers    map[string]LogsMarshaler
	producerFactory   ProducerFactoryFunc
}

func (f *kafkaExporterFactory) createTracesExporter(
	ctx context.Context,
	set component.ExporterCreateSettings,
	cfg component.Config,
) (component.TracesExporter, error) {
	oCfg := *(cfg.(*Config)) // Clone the config
	if oCfg.Topic == "" {
		oCfg.Topic = defaultTracesTopic
	}
	if oCfg.Encoding == "otlp_json" {
		set.Logger.Info("otlp_json is considered experimental and should not be used in a production environment")
	}
	exp, err := newTracesExporter(oCfg, set, f.producerFactory, f.tracesMarshalers)
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewTracesExporter(
		ctx,
		set,
		&oCfg,
		exp.tracesPusher,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
		// and will rely on the sarama Producer Timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithShutdown(exp.Close))
}

func (f *kafkaExporterFactory) createMetricsExporter(
	ctx context.Context,
	set component.ExporterCreateSettings,
	cfg component.Config,
) (component.MetricsExporter, error) {
	oCfg := *(cfg.(*Config)) // Clone the config
	if oCfg.Topic == "" {
		oCfg.Topic = defaultMetricsTopic
	}
	if oCfg.Encoding == "otlp_json" {
		set.Logger.Info("otlp_json is considered experimental and should not be used in a production environment")
	}
	exp, err := newMetricsExporter(oCfg, set, f.producerFactory, f.metricsMarshalers)
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewMetricsExporter(
		ctx,
		set,
		&oCfg,
		exp.metricsDataPusher,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
		// and will rely on the sarama Producer Timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithShutdown(exp.Close))
}

func (f *kafkaExporterFactory) createLogsExporter(
	ctx context.Context,
	set component.ExporterCreateSettings,
	cfg component.Config,
) (component.LogsExporter, error) {
	oCfg := *(cfg.(*Config)) // Clone the config
	if oCfg.Topic == "" {
		oCfg.Topic = defaultLogsTopic
	}
	if oCfg.Encoding == "otlp_json" {
		set.Logger.Info("otlp_json is considered experimental and should not be used in a production environment")
	}
	exp, err := newLogsExporter(oCfg, set, f.producerFactory, f.logsMarshalers)
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewLogsExporter(
		ctx,
		set,
		&oCfg,
		exp.logsDataPusher,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
		// and will rely on the sarama Producer Timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithShutdown(exp.Close))
}
