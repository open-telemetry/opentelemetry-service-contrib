// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

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

// MetricsConfig provides config for hostmetricsreceiver/memory metrics.
type MetricsConfig struct {
	SystemLinuxMemoryAvailable MetricConfig `mapstructure:"system.linux.memory.available"`
	SystemMemoryLimit          MetricConfig `mapstructure:"system.memory.limit"`
	SystemMemoryUsage          MetricConfig `mapstructure:"system.memory.usage"`
	SystemMemoryUtilization    MetricConfig `mapstructure:"system.memory.utilization"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SystemLinuxMemoryAvailable: MetricConfig{
			Enabled: false,
		},
		SystemMemoryLimit: MetricConfig{
			Enabled: false,
		},
		SystemMemoryUsage: MetricConfig{
			Enabled: true,
		},
		SystemMemoryUtilization: MetricConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/memory metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
