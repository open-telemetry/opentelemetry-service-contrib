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

// MetricsConfig provides config for gitprovider metrics.
type MetricsConfig struct {
	GitRepositoryBranchCount             MetricConfig `mapstructure:"git.repository.branch.count"`
	GitRepositoryContributorCount        MetricConfig `mapstructure:"git.repository.contributor.count"`
	GitRepositoryCount                   MetricConfig `mapstructure:"git.repository.count"`
	GitRepositoryPullRequestApprovedTime MetricConfig `mapstructure:"git.repository.pull_request.approved.time"`
	GitRepositoryPullRequestMergedCount  MetricConfig `mapstructure:"git.repository.pull_request.merged.count"`
	GitRepositoryPullRequestMergedTime   MetricConfig `mapstructure:"git.repository.pull_request.merged.time"`
	GitRepositoryPullRequestOpenCount    MetricConfig `mapstructure:"git.repository.pull_request.open.count"`
	GitRepositoryPullRequestOpenTime     MetricConfig `mapstructure:"git.repository.pull_request.open.time"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		GitRepositoryBranchCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryContributorCount: MetricConfig{
			Enabled: false,
		},
		GitRepositoryCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestApprovedTime: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestMergedCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestMergedTime: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestOpenCount: MetricConfig{
			Enabled: true,
		},
		GitRepositoryPullRequestOpenTime: MetricConfig{
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
	err := parser.Unmarshal(rac, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for gitprovider resource attributes.
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

// MetricsBuilderConfig is a configuration for gitprovider metrics builder.
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
