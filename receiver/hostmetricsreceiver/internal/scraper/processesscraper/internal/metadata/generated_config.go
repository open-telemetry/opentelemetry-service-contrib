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

// MetricsConfig provides config for hostmetricsreceiver/processes metrics.
type MetricsConfig struct {
	SystemProcessesCount   MetricConfig `mapstructure:"system.processes.count"`
	SystemProcessesCreated MetricConfig `mapstructure:"system.processes.created"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SystemProcessesCount: MetricConfig{
			Enabled: true,
		},
		SystemProcessesCreated: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/processes metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
