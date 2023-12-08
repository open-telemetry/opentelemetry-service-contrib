// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"
import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filter"

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

// MetricsConfig provides config for oracledb metrics.
type MetricsConfig struct {
	OracledbConsistentGets        MetricConfig `mapstructure:"oracledb.consistent_gets"`
	OracledbCPUTime               MetricConfig `mapstructure:"oracledb.cpu_time"`
	OracledbDbBlockGets           MetricConfig `mapstructure:"oracledb.db_block_gets"`
	OracledbDmlLocksLimit         MetricConfig `mapstructure:"oracledb.dml_locks.limit"`
	OracledbDmlLocksUsage         MetricConfig `mapstructure:"oracledb.dml_locks.usage"`
	OracledbEnqueueDeadlocks      MetricConfig `mapstructure:"oracledb.enqueue_deadlocks"`
	OracledbEnqueueLocksLimit     MetricConfig `mapstructure:"oracledb.enqueue_locks.limit"`
	OracledbEnqueueLocksUsage     MetricConfig `mapstructure:"oracledb.enqueue_locks.usage"`
	OracledbEnqueueResourcesLimit MetricConfig `mapstructure:"oracledb.enqueue_resources.limit"`
	OracledbEnqueueResourcesUsage MetricConfig `mapstructure:"oracledb.enqueue_resources.usage"`
	OracledbExchangeDeadlocks     MetricConfig `mapstructure:"oracledb.exchange_deadlocks"`
	OracledbExecutions            MetricConfig `mapstructure:"oracledb.executions"`
	OracledbHardParses            MetricConfig `mapstructure:"oracledb.hard_parses"`
	OracledbLogicalReads          MetricConfig `mapstructure:"oracledb.logical_reads"`
	OracledbParseCalls            MetricConfig `mapstructure:"oracledb.parse_calls"`
	OracledbPgaMemory             MetricConfig `mapstructure:"oracledb.pga_memory"`
	OracledbPhysicalReads         MetricConfig `mapstructure:"oracledb.physical_reads"`
	OracledbProcessesLimit        MetricConfig `mapstructure:"oracledb.processes.limit"`
	OracledbProcessesUsage        MetricConfig `mapstructure:"oracledb.processes.usage"`
	OracledbSessionsLimit         MetricConfig `mapstructure:"oracledb.sessions.limit"`
	OracledbSessionsUsage         MetricConfig `mapstructure:"oracledb.sessions.usage"`
	OracledbTablespaceSizeLimit   MetricConfig `mapstructure:"oracledb.tablespace_size.limit"`
	OracledbTablespaceSizeUsage   MetricConfig `mapstructure:"oracledb.tablespace_size.usage"`
	OracledbTransactionsLimit     MetricConfig `mapstructure:"oracledb.transactions.limit"`
	OracledbTransactionsUsage     MetricConfig `mapstructure:"oracledb.transactions.usage"`
	OracledbUserCommits           MetricConfig `mapstructure:"oracledb.user_commits"`
	OracledbUserRollbacks         MetricConfig `mapstructure:"oracledb.user_rollbacks"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		OracledbConsistentGets: MetricConfig{
			Enabled: false,
		},
		OracledbCPUTime: MetricConfig{
			Enabled: true,
		},
		OracledbDbBlockGets: MetricConfig{
			Enabled: false,
		},
		OracledbDmlLocksLimit: MetricConfig{
			Enabled: true,
		},
		OracledbDmlLocksUsage: MetricConfig{
			Enabled: true,
		},
		OracledbEnqueueDeadlocks: MetricConfig{
			Enabled: true,
		},
		OracledbEnqueueLocksLimit: MetricConfig{
			Enabled: true,
		},
		OracledbEnqueueLocksUsage: MetricConfig{
			Enabled: true,
		},
		OracledbEnqueueResourcesLimit: MetricConfig{
			Enabled: true,
		},
		OracledbEnqueueResourcesUsage: MetricConfig{
			Enabled: true,
		},
		OracledbExchangeDeadlocks: MetricConfig{
			Enabled: true,
		},
		OracledbExecutions: MetricConfig{
			Enabled: true,
		},
		OracledbHardParses: MetricConfig{
			Enabled: true,
		},
		OracledbLogicalReads: MetricConfig{
			Enabled: true,
		},
		OracledbParseCalls: MetricConfig{
			Enabled: true,
		},
		OracledbPgaMemory: MetricConfig{
			Enabled: true,
		},
		OracledbPhysicalReads: MetricConfig{
			Enabled: true,
		},
		OracledbProcessesLimit: MetricConfig{
			Enabled: true,
		},
		OracledbProcessesUsage: MetricConfig{
			Enabled: true,
		},
		OracledbSessionsLimit: MetricConfig{
			Enabled: true,
		},
		OracledbSessionsUsage: MetricConfig{
			Enabled: true,
		},
		OracledbTablespaceSizeLimit: MetricConfig{
			Enabled: true,
		},
		OracledbTablespaceSizeUsage: MetricConfig{
			Enabled: true,
		},
		OracledbTransactionsLimit: MetricConfig{
			Enabled: true,
		},
		OracledbTransactionsUsage: MetricConfig{
			Enabled: true,
		},
		OracledbUserCommits: MetricConfig{
			Enabled: true,
		},
		OracledbUserRollbacks: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool            `mapstructure:"enabled"`
	Include []filter.Config `mapstructure:"include"`
	Exclude []filter.Config `mapstructure:"exclude"`

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

// ResourceAttributesConfig provides config for oracledb resource attributes.
type ResourceAttributesConfig struct {
	OracledbInstanceName ResourceAttributeConfig `mapstructure:"oracledb.instance.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		OracledbInstanceName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for oracledb metrics builder.
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
