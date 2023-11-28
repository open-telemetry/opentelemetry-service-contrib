// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MysqlBufferPoolDataPages:     MetricConfig{Enabled: true},
					MysqlBufferPoolLimit:         MetricConfig{Enabled: true},
					MysqlBufferPoolOperations:    MetricConfig{Enabled: true},
					MysqlBufferPoolPageFlushes:   MetricConfig{Enabled: true},
					MysqlBufferPoolPages:         MetricConfig{Enabled: true},
					MysqlBufferPoolUsage:         MetricConfig{Enabled: true},
					MysqlClientNetworkIo:         MetricConfig{Enabled: true},
					MysqlCommands:                MetricConfig{Enabled: true},
					MysqlConnectionCount:         MetricConfig{Enabled: true},
					MysqlConnectionErrors:        MetricConfig{Enabled: true},
					MysqlDoubleWrites:            MetricConfig{Enabled: true},
					MysqlHandlers:                MetricConfig{Enabled: true},
					MysqlIndexIoWaitCount:        MetricConfig{Enabled: true},
					MysqlIndexIoWaitTime:         MetricConfig{Enabled: true},
					MysqlJoins:                   MetricConfig{Enabled: true},
					MysqlLocks:                   MetricConfig{Enabled: true},
					MysqlLogOperations:           MetricConfig{Enabled: true},
					MysqlMysqlxConnections:       MetricConfig{Enabled: true},
					MysqlMysqlxWorkerThreads:     MetricConfig{Enabled: true},
					MysqlOpenedResources:         MetricConfig{Enabled: true},
					MysqlOperations:              MetricConfig{Enabled: true},
					MysqlPageOperations:          MetricConfig{Enabled: true},
					MysqlPreparedStatements:      MetricConfig{Enabled: true},
					MysqlQcacheHits:              MetricConfig{Enabled: true},
					MysqlQcacheNotCached:         MetricConfig{Enabled: true},
					MysqlQcacheQueries:           MetricConfig{Enabled: true},
					MysqlQueryClientCount:        MetricConfig{Enabled: true},
					MysqlQueryCount:              MetricConfig{Enabled: true},
					MysqlQuerySlowCount:          MetricConfig{Enabled: true},
					MysqlReplicaSQLDelay:         MetricConfig{Enabled: true},
					MysqlReplicaTimeBehindSource: MetricConfig{Enabled: true},
					MysqlRowLocks:                MetricConfig{Enabled: true},
					MysqlRowOperations:           MetricConfig{Enabled: true},
					MysqlSorts:                   MetricConfig{Enabled: true},
					MysqlStatementEventCount:     MetricConfig{Enabled: true},
					MysqlStatementEventWaitTime:  MetricConfig{Enabled: true},
					MysqlTableIoWaitCount:        MetricConfig{Enabled: true},
					MysqlTableIoWaitTime:         MetricConfig{Enabled: true},
					MysqlTableLockWaitReadCount:  MetricConfig{Enabled: true},
					MysqlTableLockWaitReadTime:   MetricConfig{Enabled: true},
					MysqlTableLockWaitWriteCount: MetricConfig{Enabled: true},
					MysqlTableLockWaitWriteTime:  MetricConfig{Enabled: true},
					MysqlTableOpenCache:          MetricConfig{Enabled: true},
					MysqlThreads:                 MetricConfig{Enabled: true},
					MysqlTmpResources:            MetricConfig{Enabled: true},
					MysqlUptime:                  MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					MysqlInstanceEndpoint: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					MysqlBufferPoolDataPages:     MetricConfig{Enabled: false},
					MysqlBufferPoolLimit:         MetricConfig{Enabled: false},
					MysqlBufferPoolOperations:    MetricConfig{Enabled: false},
					MysqlBufferPoolPageFlushes:   MetricConfig{Enabled: false},
					MysqlBufferPoolPages:         MetricConfig{Enabled: false},
					MysqlBufferPoolUsage:         MetricConfig{Enabled: false},
					MysqlClientNetworkIo:         MetricConfig{Enabled: false},
					MysqlCommands:                MetricConfig{Enabled: false},
					MysqlConnectionCount:         MetricConfig{Enabled: false},
					MysqlConnectionErrors:        MetricConfig{Enabled: false},
					MysqlDoubleWrites:            MetricConfig{Enabled: false},
					MysqlHandlers:                MetricConfig{Enabled: false},
					MysqlIndexIoWaitCount:        MetricConfig{Enabled: false},
					MysqlIndexIoWaitTime:         MetricConfig{Enabled: false},
					MysqlJoins:                   MetricConfig{Enabled: false},
					MysqlLocks:                   MetricConfig{Enabled: false},
					MysqlLogOperations:           MetricConfig{Enabled: false},
					MysqlMysqlxConnections:       MetricConfig{Enabled: false},
					MysqlMysqlxWorkerThreads:     MetricConfig{Enabled: false},
					MysqlOpenedResources:         MetricConfig{Enabled: false},
					MysqlOperations:              MetricConfig{Enabled: false},
					MysqlPageOperations:          MetricConfig{Enabled: false},
					MysqlPreparedStatements:      MetricConfig{Enabled: false},
					MysqlQcacheHits:              MetricConfig{Enabled: false},
					MysqlQcacheNotCached:         MetricConfig{Enabled: false},
					MysqlQcacheQueries:           MetricConfig{Enabled: false},
					MysqlQueryClientCount:        MetricConfig{Enabled: false},
					MysqlQueryCount:              MetricConfig{Enabled: false},
					MysqlQuerySlowCount:          MetricConfig{Enabled: false},
					MysqlReplicaSQLDelay:         MetricConfig{Enabled: false},
					MysqlReplicaTimeBehindSource: MetricConfig{Enabled: false},
					MysqlRowLocks:                MetricConfig{Enabled: false},
					MysqlRowOperations:           MetricConfig{Enabled: false},
					MysqlSorts:                   MetricConfig{Enabled: false},
					MysqlStatementEventCount:     MetricConfig{Enabled: false},
					MysqlStatementEventWaitTime:  MetricConfig{Enabled: false},
					MysqlTableIoWaitCount:        MetricConfig{Enabled: false},
					MysqlTableIoWaitTime:         MetricConfig{Enabled: false},
					MysqlTableLockWaitReadCount:  MetricConfig{Enabled: false},
					MysqlTableLockWaitReadTime:   MetricConfig{Enabled: false},
					MysqlTableLockWaitWriteCount: MetricConfig{Enabled: false},
					MysqlTableLockWaitWriteTime:  MetricConfig{Enabled: false},
					MysqlTableOpenCache:          MetricConfig{Enabled: false},
					MysqlThreads:                 MetricConfig{Enabled: false},
					MysqlTmpResources:            MetricConfig{Enabled: false},
					MysqlUptime:                  MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					MysqlInstanceEndpoint: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				MysqlInstanceEndpoint: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				MysqlInstanceEndpoint: ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
