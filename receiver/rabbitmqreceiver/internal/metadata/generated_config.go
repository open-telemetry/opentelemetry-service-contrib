// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for rabbitmq metrics.
type MetricsConfig struct {
	RabbitmqConsumerCount       MetricConfig `mapstructure:"rabbitmq.consumer.count"`
	RabbitmqMessageAcknowledged MetricConfig `mapstructure:"rabbitmq.message.acknowledged"`
	RabbitmqMessageCurrent      MetricConfig `mapstructure:"rabbitmq.message.current"`
	RabbitmqMessageDelivered    MetricConfig `mapstructure:"rabbitmq.message.delivered"`
	RabbitmqMessageDropped      MetricConfig `mapstructure:"rabbitmq.message.dropped"`
	RabbitmqMessagePublished    MetricConfig `mapstructure:"rabbitmq.message.published"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		RabbitmqConsumerCount: MetricConfig{
			Enabled: true,
		},
		RabbitmqMessageAcknowledged: MetricConfig{
			Enabled: true,
		},
		RabbitmqMessageCurrent: MetricConfig{
			Enabled: true,
		},
		RabbitmqMessageDelivered: MetricConfig{
			Enabled: true,
		},
		RabbitmqMessageDropped: MetricConfig{
			Enabled: true,
		},
		RabbitmqMessagePublished: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for rabbitmq resource attributes.
type ResourceAttributesConfig struct {
	RabbitmqNodeName  ResourceAttributeConfig `mapstructure:"rabbitmq.node.name"`
	RabbitmqQueueName ResourceAttributeConfig `mapstructure:"rabbitmq.queue.name"`
	RabbitmqVhostName ResourceAttributeConfig `mapstructure:"rabbitmq.vhost.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		RabbitmqNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
		RabbitmqQueueName: ResourceAttributeConfig{
			Enabled: true,
		},
		RabbitmqVhostName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for rabbitmq metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
