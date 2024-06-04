// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
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
					SnowflakeBillingCloudServiceTotal:              MetricConfig{Enabled: true},
					SnowflakeBillingTotalCreditTotal:               MetricConfig{Enabled: true},
					SnowflakeBillingVirtualWarehouseTotal:          MetricConfig{Enabled: true},
					SnowflakeBillingWarehouseCloudServiceTotal:     MetricConfig{Enabled: true},
					SnowflakeBillingWarehouseTotalCreditTotal:      MetricConfig{Enabled: true},
					SnowflakeBillingWarehouseVirtualWarehouseTotal: MetricConfig{Enabled: true},
					SnowflakeDatabaseBytesScannedAvg:               MetricConfig{Enabled: true},
					SnowflakeDatabaseQueryCount:                    MetricConfig{Enabled: true},
					SnowflakeLoginsTotal:                           MetricConfig{Enabled: true},
					SnowflakePipeCreditsUsedTotal:                  MetricConfig{Enabled: true},
					SnowflakeQueryBlocked:                          MetricConfig{Enabled: true},
					SnowflakeQueryBytesDeletedAvg:                  MetricConfig{Enabled: true},
					SnowflakeQueryBytesSpilledLocalAvg:             MetricConfig{Enabled: true},
					SnowflakeQueryBytesSpilledRemoteAvg:            MetricConfig{Enabled: true},
					SnowflakeQueryBytesWrittenAvg:                  MetricConfig{Enabled: true},
					SnowflakeQueryCompilationTimeAvg:               MetricConfig{Enabled: true},
					SnowflakeQueryDataScannedCacheAvg:              MetricConfig{Enabled: true},
					SnowflakeQueryExecuted:                         MetricConfig{Enabled: true},
					SnowflakeQueryExecutionTimeAvg:                 MetricConfig{Enabled: true},
					SnowflakeQueryPartitionsScannedAvg:             MetricConfig{Enabled: true},
					SnowflakeQueryQueuedOverload:                   MetricConfig{Enabled: true},
					SnowflakeQueryQueuedProvision:                  MetricConfig{Enabled: true},
					SnowflakeQueuedOverloadTimeAvg:                 MetricConfig{Enabled: true},
					SnowflakeQueuedProvisioningTimeAvg:             MetricConfig{Enabled: true},
					SnowflakeQueuedRepairTimeAvg:                   MetricConfig{Enabled: true},
					SnowflakeRowsDeletedAvg:                        MetricConfig{Enabled: true},
					SnowflakeRowsInsertedAvg:                       MetricConfig{Enabled: true},
					SnowflakeRowsProducedAvg:                       MetricConfig{Enabled: true},
					SnowflakeRowsUnloadedAvg:                       MetricConfig{Enabled: true},
					SnowflakeRowsUpdatedAvg:                        MetricConfig{Enabled: true},
					SnowflakeSessionIDCount:                        MetricConfig{Enabled: true},
					SnowflakeStorageFailsafeBytesTotal:             MetricConfig{Enabled: true},
					SnowflakeStorageStageBytesTotal:                MetricConfig{Enabled: true},
					SnowflakeStorageStorageBytesTotal:              MetricConfig{Enabled: true},
					SnowflakeTotalElapsedTimeAvg:                   MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SnowflakeAccountName: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SnowflakeBillingCloudServiceTotal:              MetricConfig{Enabled: false},
					SnowflakeBillingTotalCreditTotal:               MetricConfig{Enabled: false},
					SnowflakeBillingVirtualWarehouseTotal:          MetricConfig{Enabled: false},
					SnowflakeBillingWarehouseCloudServiceTotal:     MetricConfig{Enabled: false},
					SnowflakeBillingWarehouseTotalCreditTotal:      MetricConfig{Enabled: false},
					SnowflakeBillingWarehouseVirtualWarehouseTotal: MetricConfig{Enabled: false},
					SnowflakeDatabaseBytesScannedAvg:               MetricConfig{Enabled: false},
					SnowflakeDatabaseQueryCount:                    MetricConfig{Enabled: false},
					SnowflakeLoginsTotal:                           MetricConfig{Enabled: false},
					SnowflakePipeCreditsUsedTotal:                  MetricConfig{Enabled: false},
					SnowflakeQueryBlocked:                          MetricConfig{Enabled: false},
					SnowflakeQueryBytesDeletedAvg:                  MetricConfig{Enabled: false},
					SnowflakeQueryBytesSpilledLocalAvg:             MetricConfig{Enabled: false},
					SnowflakeQueryBytesSpilledRemoteAvg:            MetricConfig{Enabled: false},
					SnowflakeQueryBytesWrittenAvg:                  MetricConfig{Enabled: false},
					SnowflakeQueryCompilationTimeAvg:               MetricConfig{Enabled: false},
					SnowflakeQueryDataScannedCacheAvg:              MetricConfig{Enabled: false},
					SnowflakeQueryExecuted:                         MetricConfig{Enabled: false},
					SnowflakeQueryExecutionTimeAvg:                 MetricConfig{Enabled: false},
					SnowflakeQueryPartitionsScannedAvg:             MetricConfig{Enabled: false},
					SnowflakeQueryQueuedOverload:                   MetricConfig{Enabled: false},
					SnowflakeQueryQueuedProvision:                  MetricConfig{Enabled: false},
					SnowflakeQueuedOverloadTimeAvg:                 MetricConfig{Enabled: false},
					SnowflakeQueuedProvisioningTimeAvg:             MetricConfig{Enabled: false},
					SnowflakeQueuedRepairTimeAvg:                   MetricConfig{Enabled: false},
					SnowflakeRowsDeletedAvg:                        MetricConfig{Enabled: false},
					SnowflakeRowsInsertedAvg:                       MetricConfig{Enabled: false},
					SnowflakeRowsProducedAvg:                       MetricConfig{Enabled: false},
					SnowflakeRowsUnloadedAvg:                       MetricConfig{Enabled: false},
					SnowflakeRowsUpdatedAvg:                        MetricConfig{Enabled: false},
					SnowflakeSessionIDCount:                        MetricConfig{Enabled: false},
					SnowflakeStorageFailsafeBytesTotal:             MetricConfig{Enabled: false},
					SnowflakeStorageStageBytesTotal:                MetricConfig{Enabled: false},
					SnowflakeStorageStorageBytesTotal:              MetricConfig{Enabled: false},
					SnowflakeTotalElapsedTimeAvg:                   MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					SnowflakeAccountName: ResourceAttributeConfig{Enabled: false},
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
	require.NoError(t, sub.Unmarshal(&cfg))
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
				SnowflakeAccountName: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				SnowflakeAccountName: ResourceAttributeConfig{Enabled: false},
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
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
