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

// MetricsConfig provides config for mysql metrics.
type MetricsConfig struct {
	MysqlBufferPoolDataPages     MetricConfig `mapstructure:"mysql.buffer_pool.data_pages"`
	MysqlBufferPoolLimit         MetricConfig `mapstructure:"mysql.buffer_pool.limit"`
	MysqlBufferPoolOperations    MetricConfig `mapstructure:"mysql.buffer_pool.operations"`
	MysqlBufferPoolPageFlushes   MetricConfig `mapstructure:"mysql.buffer_pool.page_flushes"`
	MysqlBufferPoolPages         MetricConfig `mapstructure:"mysql.buffer_pool.pages"`
	MysqlBufferPoolUsage         MetricConfig `mapstructure:"mysql.buffer_pool.usage"`
	MysqlClientNetworkIo         MetricConfig `mapstructure:"mysql.client.network.io"`
	MysqlCommands                MetricConfig `mapstructure:"mysql.commands"`
	MysqlConnectionCount         MetricConfig `mapstructure:"mysql.connection.count"`
	MysqlConnectionErrors        MetricConfig `mapstructure:"mysql.connection.errors"`
	MysqlDoubleWrites            MetricConfig `mapstructure:"mysql.double_writes"`
	MysqlHandlers                MetricConfig `mapstructure:"mysql.handlers"`
	MysqlIndexIoWaitCount        MetricConfig `mapstructure:"mysql.index.io.wait.count"`
	MysqlIndexIoWaitTime         MetricConfig `mapstructure:"mysql.index.io.wait.time"`
	MysqlJoins                   MetricConfig `mapstructure:"mysql.joins"`
	MysqlLocks                   MetricConfig `mapstructure:"mysql.locks"`
	MysqlLogOperations           MetricConfig `mapstructure:"mysql.log_operations"`
	MysqlMysqlxConnections       MetricConfig `mapstructure:"mysql.mysqlx_connections"`
	MysqlMysqlxWorkerThreads     MetricConfig `mapstructure:"mysql.mysqlx_worker_threads"`
	MysqlOpenedResources         MetricConfig `mapstructure:"mysql.opened_resources"`
	MysqlOperations              MetricConfig `mapstructure:"mysql.operations"`
	MysqlPageOperations          MetricConfig `mapstructure:"mysql.page_operations"`
	MysqlPreparedStatements      MetricConfig `mapstructure:"mysql.prepared_statements"`
	MysqlQueryClientCount        MetricConfig `mapstructure:"mysql.query.client.count"`
	MysqlQueryCount              MetricConfig `mapstructure:"mysql.query.count"`
	MysqlQuerySlowCount          MetricConfig `mapstructure:"mysql.query.slow.count"`
	MysqlReplicaSQLDelay         MetricConfig `mapstructure:"mysql.replica.sql_delay"`
	MysqlReplicaTimeBehindSource MetricConfig `mapstructure:"mysql.replica.time_behind_source"`
	MysqlRowLocks                MetricConfig `mapstructure:"mysql.row_locks"`
	MysqlRowOperations           MetricConfig `mapstructure:"mysql.row_operations"`
	MysqlSorts                   MetricConfig `mapstructure:"mysql.sorts"`
	MysqlStatementEventCount     MetricConfig `mapstructure:"mysql.statement_event.count"`
	MysqlStatementEventWaitTime  MetricConfig `mapstructure:"mysql.statement_event.wait.time"`
	MysqlTableIoWaitCount        MetricConfig `mapstructure:"mysql.table.io.wait.count"`
	MysqlTableIoWaitTime         MetricConfig `mapstructure:"mysql.table.io.wait.time"`
	MysqlTableLockWaitReadCount  MetricConfig `mapstructure:"mysql.table.lock_wait.read.count"`
	MysqlTableLockWaitReadTime   MetricConfig `mapstructure:"mysql.table.lock_wait.read.time"`
	MysqlTableLockWaitWriteCount MetricConfig `mapstructure:"mysql.table.lock_wait.write.count"`
	MysqlTableLockWaitWriteTime  MetricConfig `mapstructure:"mysql.table.lock_wait.write.time"`
	MysqlTableOpenCache          MetricConfig `mapstructure:"mysql.table_open_cache"`
	MysqlThreads                 MetricConfig `mapstructure:"mysql.threads"`
	MysqlTmpResources            MetricConfig `mapstructure:"mysql.tmp_resources"`
	MysqlUptime                  MetricConfig `mapstructure:"mysql.uptime"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		MysqlBufferPoolDataPages: MetricConfig{
			Enabled: true,
		},
		MysqlBufferPoolLimit: MetricConfig{
			Enabled: true,
		},
		MysqlBufferPoolOperations: MetricConfig{
			Enabled: true,
		},
		MysqlBufferPoolPageFlushes: MetricConfig{
			Enabled: true,
		},
		MysqlBufferPoolPages: MetricConfig{
			Enabled: true,
		},
		MysqlBufferPoolUsage: MetricConfig{
			Enabled: true,
		},
		MysqlClientNetworkIo: MetricConfig{
			Enabled: false,
		},
		MysqlCommands: MetricConfig{
			Enabled: false,
		},
		MysqlConnectionCount: MetricConfig{
			Enabled: false,
		},
		MysqlConnectionErrors: MetricConfig{
			Enabled: false,
		},
		MysqlDoubleWrites: MetricConfig{
			Enabled: true,
		},
		MysqlHandlers: MetricConfig{
			Enabled: true,
		},
		MysqlIndexIoWaitCount: MetricConfig{
			Enabled: true,
		},
		MysqlIndexIoWaitTime: MetricConfig{
			Enabled: true,
		},
		MysqlJoins: MetricConfig{
			Enabled: false,
		},
		MysqlLocks: MetricConfig{
			Enabled: true,
		},
		MysqlLogOperations: MetricConfig{
			Enabled: true,
		},
		MysqlMysqlxConnections: MetricConfig{
			Enabled: true,
		},
		MysqlMysqlxWorkerThreads: MetricConfig{
			Enabled: false,
		},
		MysqlOpenedResources: MetricConfig{
			Enabled: true,
		},
		MysqlOperations: MetricConfig{
			Enabled: true,
		},
		MysqlPageOperations: MetricConfig{
			Enabled: true,
		},
		MysqlPreparedStatements: MetricConfig{
			Enabled: true,
		},
		MysqlQueryClientCount: MetricConfig{
			Enabled: false,
		},
		MysqlQueryCount: MetricConfig{
			Enabled: false,
		},
		MysqlQuerySlowCount: MetricConfig{
			Enabled: false,
		},
		MysqlReplicaSQLDelay: MetricConfig{
			Enabled: false,
		},
		MysqlReplicaTimeBehindSource: MetricConfig{
			Enabled: false,
		},
		MysqlRowLocks: MetricConfig{
			Enabled: true,
		},
		MysqlRowOperations: MetricConfig{
			Enabled: true,
		},
		MysqlSorts: MetricConfig{
			Enabled: true,
		},
		MysqlStatementEventCount: MetricConfig{
			Enabled: false,
		},
		MysqlStatementEventWaitTime: MetricConfig{
			Enabled: false,
		},
		MysqlTableIoWaitCount: MetricConfig{
			Enabled: true,
		},
		MysqlTableIoWaitTime: MetricConfig{
			Enabled: true,
		},
		MysqlTableLockWaitReadCount: MetricConfig{
			Enabled: false,
		},
		MysqlTableLockWaitReadTime: MetricConfig{
			Enabled: false,
		},
		MysqlTableLockWaitWriteCount: MetricConfig{
			Enabled: false,
		},
		MysqlTableLockWaitWriteTime: MetricConfig{
			Enabled: false,
		},
		MysqlTableOpenCache: MetricConfig{
			Enabled: false,
		},
		MysqlThreads: MetricConfig{
			Enabled: true,
		},
		MysqlTmpResources: MetricConfig{
			Enabled: true,
		},
		MysqlUptime: MetricConfig{
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

// ResourceAttributesConfig provides config for mysql resource attributes.
type ResourceAttributesConfig struct {
	MysqlInstanceEndpoint ResourceAttributeConfig `mapstructure:"mysql.instance.endpoint"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		MysqlInstanceEndpoint: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for mysql metrics builder.
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
