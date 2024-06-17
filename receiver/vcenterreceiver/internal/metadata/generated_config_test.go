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
					VcenterClusterCPUEffective:        MetricConfig{Enabled: true},
					VcenterClusterCPULimit:            MetricConfig{Enabled: true},
					VcenterClusterHostCount:           MetricConfig{Enabled: true},
					VcenterClusterMemoryEffective:     MetricConfig{Enabled: true},
					VcenterClusterMemoryLimit:         MetricConfig{Enabled: true},
					VcenterClusterVMCount:             MetricConfig{Enabled: true},
					VcenterClusterVMTemplateCount:     MetricConfig{Enabled: true},
					VcenterDatastoreDiskUsage:         MetricConfig{Enabled: true},
					VcenterDatastoreDiskUtilization:   MetricConfig{Enabled: true},
					VcenterHostCPUUsage:               MetricConfig{Enabled: true},
					VcenterHostCPUUtilization:         MetricConfig{Enabled: true},
					VcenterHostDiskLatencyAvg:         MetricConfig{Enabled: true},
					VcenterHostDiskLatencyMax:         MetricConfig{Enabled: true},
					VcenterHostDiskThroughput:         MetricConfig{Enabled: true},
					VcenterHostMemoryUsage:            MetricConfig{Enabled: true},
					VcenterHostMemoryUtilization:      MetricConfig{Enabled: true},
					VcenterHostNetworkPacketErrorRate: MetricConfig{Enabled: true},
					VcenterHostNetworkPacketRate:      MetricConfig{Enabled: true},
					VcenterHostNetworkThroughput:      MetricConfig{Enabled: true},
					VcenterHostNetworkUsage:           MetricConfig{Enabled: true},
					VcenterResourcePoolCPUShares:      MetricConfig{Enabled: true},
					VcenterResourcePoolCPUUsage:       MetricConfig{Enabled: true},
					VcenterResourcePoolMemoryShares:   MetricConfig{Enabled: true},
					VcenterResourcePoolMemoryUsage:    MetricConfig{Enabled: true},
					VcenterVMCPUReadiness:             MetricConfig{Enabled: true},
					VcenterVMCPUUsage:                 MetricConfig{Enabled: true},
					VcenterVMCPUUtilization:           MetricConfig{Enabled: true},
					VcenterVMDiskLatencyAvg:           MetricConfig{Enabled: true},
					VcenterVMDiskLatencyMax:           MetricConfig{Enabled: true},
					VcenterVMDiskThroughput:           MetricConfig{Enabled: true},
					VcenterVMDiskUsage:                MetricConfig{Enabled: true},
					VcenterVMDiskUtilization:          MetricConfig{Enabled: true},
					VcenterVMMemoryBallooned:          MetricConfig{Enabled: true},
					VcenterVMMemorySwapped:            MetricConfig{Enabled: true},
					VcenterVMMemorySwappedSsd:         MetricConfig{Enabled: true},
					VcenterVMMemoryUsage:              MetricConfig{Enabled: true},
					VcenterVMMemoryUtilization:        MetricConfig{Enabled: true},
					VcenterVMNetworkPacketDropRate:    MetricConfig{Enabled: true},
					VcenterVMNetworkPacketRate:        MetricConfig{Enabled: true},
					VcenterVMNetworkThroughput:        MetricConfig{Enabled: true},
					VcenterVMNetworkUsage:             MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					VcenterClusterName:               ResourceAttributeConfig{Enabled: true},
					VcenterDatacenterName:            ResourceAttributeConfig{Enabled: true},
					VcenterDatastoreName:             ResourceAttributeConfig{Enabled: true},
					VcenterHostName:                  ResourceAttributeConfig{Enabled: true},
					VcenterResourcePoolInventoryPath: ResourceAttributeConfig{Enabled: true},
					VcenterResourcePoolName:          ResourceAttributeConfig{Enabled: true},
					VcenterVirtualAppInventoryPath:   ResourceAttributeConfig{Enabled: true},
					VcenterVirtualAppName:            ResourceAttributeConfig{Enabled: true},
					VcenterVMID:                      ResourceAttributeConfig{Enabled: true},
					VcenterVMName:                    ResourceAttributeConfig{Enabled: true},
					VcenterVMTemplateID:              ResourceAttributeConfig{Enabled: true},
					VcenterVMTemplateName:            ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					VcenterClusterCPUEffective:        MetricConfig{Enabled: false},
					VcenterClusterCPULimit:            MetricConfig{Enabled: false},
					VcenterClusterHostCount:           MetricConfig{Enabled: false},
					VcenterClusterMemoryEffective:     MetricConfig{Enabled: false},
					VcenterClusterMemoryLimit:         MetricConfig{Enabled: false},
					VcenterClusterVMCount:             MetricConfig{Enabled: false},
					VcenterClusterVMTemplateCount:     MetricConfig{Enabled: false},
					VcenterDatastoreDiskUsage:         MetricConfig{Enabled: false},
					VcenterDatastoreDiskUtilization:   MetricConfig{Enabled: false},
					VcenterHostCPUUsage:               MetricConfig{Enabled: false},
					VcenterHostCPUUtilization:         MetricConfig{Enabled: false},
					VcenterHostDiskLatencyAvg:         MetricConfig{Enabled: false},
					VcenterHostDiskLatencyMax:         MetricConfig{Enabled: false},
					VcenterHostDiskThroughput:         MetricConfig{Enabled: false},
					VcenterHostMemoryUsage:            MetricConfig{Enabled: false},
					VcenterHostMemoryUtilization:      MetricConfig{Enabled: false},
					VcenterHostNetworkPacketErrorRate: MetricConfig{Enabled: false},
					VcenterHostNetworkPacketRate:      MetricConfig{Enabled: false},
					VcenterHostNetworkThroughput:      MetricConfig{Enabled: false},
					VcenterHostNetworkUsage:           MetricConfig{Enabled: false},
					VcenterResourcePoolCPUShares:      MetricConfig{Enabled: false},
					VcenterResourcePoolCPUUsage:       MetricConfig{Enabled: false},
					VcenterResourcePoolMemoryShares:   MetricConfig{Enabled: false},
					VcenterResourcePoolMemoryUsage:    MetricConfig{Enabled: false},
					VcenterVMCPUReadiness:             MetricConfig{Enabled: false},
					VcenterVMCPUUsage:                 MetricConfig{Enabled: false},
					VcenterVMCPUUtilization:           MetricConfig{Enabled: false},
					VcenterVMDiskLatencyAvg:           MetricConfig{Enabled: false},
					VcenterVMDiskLatencyMax:           MetricConfig{Enabled: false},
					VcenterVMDiskThroughput:           MetricConfig{Enabled: false},
					VcenterVMDiskUsage:                MetricConfig{Enabled: false},
					VcenterVMDiskUtilization:          MetricConfig{Enabled: false},
					VcenterVMMemoryBallooned:          MetricConfig{Enabled: false},
					VcenterVMMemorySwapped:            MetricConfig{Enabled: false},
					VcenterVMMemorySwappedSsd:         MetricConfig{Enabled: false},
					VcenterVMMemoryUsage:              MetricConfig{Enabled: false},
					VcenterVMMemoryUtilization:        MetricConfig{Enabled: false},
					VcenterVMNetworkPacketDropRate:    MetricConfig{Enabled: false},
					VcenterVMNetworkPacketRate:        MetricConfig{Enabled: false},
					VcenterVMNetworkThroughput:        MetricConfig{Enabled: false},
					VcenterVMNetworkUsage:             MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					VcenterClusterName:               ResourceAttributeConfig{Enabled: false},
					VcenterDatacenterName:            ResourceAttributeConfig{Enabled: false},
					VcenterDatastoreName:             ResourceAttributeConfig{Enabled: false},
					VcenterHostName:                  ResourceAttributeConfig{Enabled: false},
					VcenterResourcePoolInventoryPath: ResourceAttributeConfig{Enabled: false},
					VcenterResourcePoolName:          ResourceAttributeConfig{Enabled: false},
					VcenterVirtualAppInventoryPath:   ResourceAttributeConfig{Enabled: false},
					VcenterVirtualAppName:            ResourceAttributeConfig{Enabled: false},
					VcenterVMID:                      ResourceAttributeConfig{Enabled: false},
					VcenterVMName:                    ResourceAttributeConfig{Enabled: false},
					VcenterVMTemplateID:              ResourceAttributeConfig{Enabled: false},
					VcenterVMTemplateName:            ResourceAttributeConfig{Enabled: false},
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
				VcenterClusterName:               ResourceAttributeConfig{Enabled: true},
				VcenterDatacenterName:            ResourceAttributeConfig{Enabled: true},
				VcenterDatastoreName:             ResourceAttributeConfig{Enabled: true},
				VcenterHostName:                  ResourceAttributeConfig{Enabled: true},
				VcenterResourcePoolInventoryPath: ResourceAttributeConfig{Enabled: true},
				VcenterResourcePoolName:          ResourceAttributeConfig{Enabled: true},
				VcenterVirtualAppInventoryPath:   ResourceAttributeConfig{Enabled: true},
				VcenterVirtualAppName:            ResourceAttributeConfig{Enabled: true},
				VcenterVMID:                      ResourceAttributeConfig{Enabled: true},
				VcenterVMName:                    ResourceAttributeConfig{Enabled: true},
				VcenterVMTemplateID:              ResourceAttributeConfig{Enabled: true},
				VcenterVMTemplateName:            ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				VcenterClusterName:               ResourceAttributeConfig{Enabled: false},
				VcenterDatacenterName:            ResourceAttributeConfig{Enabled: false},
				VcenterDatastoreName:             ResourceAttributeConfig{Enabled: false},
				VcenterHostName:                  ResourceAttributeConfig{Enabled: false},
				VcenterResourcePoolInventoryPath: ResourceAttributeConfig{Enabled: false},
				VcenterResourcePoolName:          ResourceAttributeConfig{Enabled: false},
				VcenterVirtualAppInventoryPath:   ResourceAttributeConfig{Enabled: false},
				VcenterVirtualAppName:            ResourceAttributeConfig{Enabled: false},
				VcenterVMID:                      ResourceAttributeConfig{Enabled: false},
				VcenterVMName:                    ResourceAttributeConfig{Enabled: false},
				VcenterVMTemplateID:              ResourceAttributeConfig{Enabled: false},
				VcenterVMTemplateName:            ResourceAttributeConfig{Enabled: false},
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
