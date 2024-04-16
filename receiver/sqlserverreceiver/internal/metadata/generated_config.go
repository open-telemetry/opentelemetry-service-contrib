// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
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

// MetricsConfig provides config for sqlserver metrics.
type MetricsConfig struct {
	SqlserverBatchRequestRate            MetricConfig `mapstructure:"sqlserver.batch.request.rate"`
	SqlserverBatchSQLCompilationRate     MetricConfig `mapstructure:"sqlserver.batch.sql_compilation.rate"`
	SqlserverBatchSQLRecompilationRate   MetricConfig `mapstructure:"sqlserver.batch.sql_recompilation.rate"`
	SqlserverLockWaitRate                MetricConfig `mapstructure:"sqlserver.lock.wait.rate"`
	SqlserverLockWaitTimeAvg             MetricConfig `mapstructure:"sqlserver.lock.wait_time.avg"`
	SqlserverPageBufferCacheHitRatio     MetricConfig `mapstructure:"sqlserver.page.buffer_cache.hit_ratio"`
	SqlserverPageCheckpointFlushRate     MetricConfig `mapstructure:"sqlserver.page.checkpoint.flush.rate"`
	SqlserverPageLazyWriteRate           MetricConfig `mapstructure:"sqlserver.page.lazy_write.rate"`
	SqlserverPageLifeExpectancy          MetricConfig `mapstructure:"sqlserver.page.life_expectancy"`
	SqlserverPageOperationRate           MetricConfig `mapstructure:"sqlserver.page.operation.rate"`
	SqlserverPageSplitRate               MetricConfig `mapstructure:"sqlserver.page.split.rate"`
	SqlserverTransactionRate             MetricConfig `mapstructure:"sqlserver.transaction.rate"`
	SqlserverTransactionWriteRate        MetricConfig `mapstructure:"sqlserver.transaction.write.rate"`
	SqlserverTransactionLogFlushDataRate MetricConfig `mapstructure:"sqlserver.transaction_log.flush.data.rate"`
	SqlserverTransactionLogFlushRate     MetricConfig `mapstructure:"sqlserver.transaction_log.flush.rate"`
	SqlserverTransactionLogFlushWaitRate MetricConfig `mapstructure:"sqlserver.transaction_log.flush.wait.rate"`
	SqlserverTransactionLogGrowthCount   MetricConfig `mapstructure:"sqlserver.transaction_log.growth.count"`
	SqlserverTransactionLogShrinkCount   MetricConfig `mapstructure:"sqlserver.transaction_log.shrink.count"`
	SqlserverTransactionLogUsage         MetricConfig `mapstructure:"sqlserver.transaction_log.usage"`
	SqlserverUserConnectionCount         MetricConfig `mapstructure:"sqlserver.user.connection.count"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SqlserverBatchRequestRate: MetricConfig{
			Enabled: true,
		},
		SqlserverBatchSQLCompilationRate: MetricConfig{
			Enabled: true,
		},
		SqlserverBatchSQLRecompilationRate: MetricConfig{
			Enabled: true,
		},
		SqlserverLockWaitRate: MetricConfig{
			Enabled: true,
		},
		SqlserverLockWaitTimeAvg: MetricConfig{
			Enabled: true,
		},
		SqlserverPageBufferCacheHitRatio: MetricConfig{
			Enabled: true,
		},
		SqlserverPageCheckpointFlushRate: MetricConfig{
			Enabled: true,
		},
		SqlserverPageLazyWriteRate: MetricConfig{
			Enabled: true,
		},
		SqlserverPageLifeExpectancy: MetricConfig{
			Enabled: true,
		},
		SqlserverPageOperationRate: MetricConfig{
			Enabled: true,
		},
		SqlserverPageSplitRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionWriteRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogFlushDataRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogFlushRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogFlushWaitRate: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogGrowthCount: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogShrinkCount: MetricConfig{
			Enabled: true,
		},
		SqlserverTransactionLogUsage: MetricConfig{
			Enabled: true,
		},
		SqlserverUserConnectionCount: MetricConfig{
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

// ResourceAttributesConfig provides config for sqlserver resource attributes.
type ResourceAttributesConfig struct {
	SqlserverComputerName ResourceAttributeConfig `mapstructure:"sqlserver.computer.name"`
	SqlserverDatabaseName ResourceAttributeConfig `mapstructure:"sqlserver.database.name"`
	SqlserverInstanceName ResourceAttributeConfig `mapstructure:"sqlserver.instance.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		SqlserverComputerName: ResourceAttributeConfig{
			Enabled: false,
		},
		SqlserverDatabaseName: ResourceAttributeConfig{
			Enabled: true,
		},
		SqlserverInstanceName: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for sqlserver metrics builder.
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
