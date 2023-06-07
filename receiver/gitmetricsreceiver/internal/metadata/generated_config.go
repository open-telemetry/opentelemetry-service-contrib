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
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for gitmetricsreceiver metrics.
type MetricsConfig struct {
	GitRepositoryBranchCommitCount MetricConfig `mapstructure:"git.repository.branch.commit.count"`
	GitRepositoryBranchCount       MetricConfig `mapstructure:"git.repository.branch.count"`
	GitRepositoryBranchTime        MetricConfig `mapstructure:"git.repository.branch.time"`
	GitRepositoryContributorCount  MetricConfig `mapstructure:"git.repository.contributor.count"`
	GitRepositoryCount             MetricConfig `mapstructure:"git.repository.count"`
	GitRepositoryPullRequestTime   MetricConfig `mapstructure:"git.repository.pull_request.time"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		GitRepositoryBranchCommitCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryBranchCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryBranchTime: MetricConfig{
			Enabled: true,
		},
		GitRepositoryContributorCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestTime: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for gitmetricsreceiver resource attributes.
type ResourceAttributesConfig struct {
	GitVendorName    ResourceAttributeConfig `mapstructure:"git.vendor.name"`
	OrganizationName ResourceAttributeConfig `mapstructure:"organization.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		GitVendorName: ResourceAttributeConfig{
			Enabled: true,
		},
		OrganizationName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for gitmetricsreceiver metrics builder.
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
