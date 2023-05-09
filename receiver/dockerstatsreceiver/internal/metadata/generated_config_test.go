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
					ContainerBlockioIoMergedRecursive:          MetricConfig{Enabled: true},
					ContainerBlockioIoQueuedRecursive:          MetricConfig{Enabled: true},
					ContainerBlockioIoServiceBytesRecursive:    MetricConfig{Enabled: true},
					ContainerBlockioIoServiceTimeRecursive:     MetricConfig{Enabled: true},
					ContainerBlockioIoServicedRecursive:        MetricConfig{Enabled: true},
					ContainerBlockioIoTimeRecursive:            MetricConfig{Enabled: true},
					ContainerBlockioIoWaitTimeRecursive:        MetricConfig{Enabled: true},
					ContainerBlockioSectorsRecursive:           MetricConfig{Enabled: true},
					ContainerCPUPercent:                        MetricConfig{Enabled: true},
					ContainerCPUThrottlingDataPeriods:          MetricConfig{Enabled: true},
					ContainerCPUThrottlingDataThrottledPeriods: MetricConfig{Enabled: true},
					ContainerCPUThrottlingDataThrottledTime:    MetricConfig{Enabled: true},
					ContainerCPUUsageKernelmode:                MetricConfig{Enabled: true},
					ContainerCPUUsagePercpu:                    MetricConfig{Enabled: true},
					ContainerCPUUsageSystem:                    MetricConfig{Enabled: true},
					ContainerCPUUsageTotal:                     MetricConfig{Enabled: true},
					ContainerCPUUsageUsermode:                  MetricConfig{Enabled: true},
					ContainerMemoryActiveAnon:                  MetricConfig{Enabled: true},
					ContainerMemoryActiveFile:                  MetricConfig{Enabled: true},
					ContainerMemoryCache:                       MetricConfig{Enabled: true},
					ContainerMemoryDirty:                       MetricConfig{Enabled: true},
					ContainerMemoryHierarchicalMemoryLimit:     MetricConfig{Enabled: true},
					ContainerMemoryHierarchicalMemswLimit:      MetricConfig{Enabled: true},
					ContainerMemoryInactiveAnon:                MetricConfig{Enabled: true},
					ContainerMemoryInactiveFile:                MetricConfig{Enabled: true},
					ContainerMemoryMappedFile:                  MetricConfig{Enabled: true},
					ContainerMemoryPercent:                     MetricConfig{Enabled: true},
					ContainerMemoryPgfault:                     MetricConfig{Enabled: true},
					ContainerMemoryPgmajfault:                  MetricConfig{Enabled: true},
					ContainerMemoryPgpgin:                      MetricConfig{Enabled: true},
					ContainerMemoryPgpgout:                     MetricConfig{Enabled: true},
					ContainerMemoryRss:                         MetricConfig{Enabled: true},
					ContainerMemoryRssHuge:                     MetricConfig{Enabled: true},
					ContainerMemorySwap:                        MetricConfig{Enabled: true},
					ContainerMemoryTotalActiveAnon:             MetricConfig{Enabled: true},
					ContainerMemoryTotalActiveFile:             MetricConfig{Enabled: true},
					ContainerMemoryTotalCache:                  MetricConfig{Enabled: true},
					ContainerMemoryTotalDirty:                  MetricConfig{Enabled: true},
					ContainerMemoryTotalInactiveAnon:           MetricConfig{Enabled: true},
					ContainerMemoryTotalInactiveFile:           MetricConfig{Enabled: true},
					ContainerMemoryTotalMappedFile:             MetricConfig{Enabled: true},
					ContainerMemoryTotalPgfault:                MetricConfig{Enabled: true},
					ContainerMemoryTotalPgmajfault:             MetricConfig{Enabled: true},
					ContainerMemoryTotalPgpgin:                 MetricConfig{Enabled: true},
					ContainerMemoryTotalPgpgout:                MetricConfig{Enabled: true},
					ContainerMemoryTotalRss:                    MetricConfig{Enabled: true},
					ContainerMemoryTotalRssHuge:                MetricConfig{Enabled: true},
					ContainerMemoryTotalSwap:                   MetricConfig{Enabled: true},
					ContainerMemoryTotalUnevictable:            MetricConfig{Enabled: true},
					ContainerMemoryTotalWriteback:              MetricConfig{Enabled: true},
					ContainerMemoryUnevictable:                 MetricConfig{Enabled: true},
					ContainerMemoryUsageLimit:                  MetricConfig{Enabled: true},
					ContainerMemoryUsageMax:                    MetricConfig{Enabled: true},
					ContainerMemoryUsageTotal:                  MetricConfig{Enabled: true},
					ContainerMemoryWriteback:                   MetricConfig{Enabled: true},
					ContainerNetworkIoUsageRxBytes:             MetricConfig{Enabled: true},
					ContainerNetworkIoUsageRxDropped:           MetricConfig{Enabled: true},
					ContainerNetworkIoUsageRxErrors:            MetricConfig{Enabled: true},
					ContainerNetworkIoUsageRxPackets:           MetricConfig{Enabled: true},
					ContainerNetworkIoUsageTxBytes:             MetricConfig{Enabled: true},
					ContainerNetworkIoUsageTxDropped:           MetricConfig{Enabled: true},
					ContainerNetworkIoUsageTxErrors:            MetricConfig{Enabled: true},
					ContainerNetworkIoUsageTxPackets:           MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerHostname:  ResourceAttributeConfig{Enabled: true},
					ContainerID:        ResourceAttributeConfig{Enabled: true},
					ContainerImageName: ResourceAttributeConfig{Enabled: true},
					ContainerName:      ResourceAttributeConfig{Enabled: true},
					ContainerRuntime:   ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ContainerBlockioIoMergedRecursive:          MetricConfig{Enabled: false},
					ContainerBlockioIoQueuedRecursive:          MetricConfig{Enabled: false},
					ContainerBlockioIoServiceBytesRecursive:    MetricConfig{Enabled: false},
					ContainerBlockioIoServiceTimeRecursive:     MetricConfig{Enabled: false},
					ContainerBlockioIoServicedRecursive:        MetricConfig{Enabled: false},
					ContainerBlockioIoTimeRecursive:            MetricConfig{Enabled: false},
					ContainerBlockioIoWaitTimeRecursive:        MetricConfig{Enabled: false},
					ContainerBlockioSectorsRecursive:           MetricConfig{Enabled: false},
					ContainerCPUPercent:                        MetricConfig{Enabled: false},
					ContainerCPUThrottlingDataPeriods:          MetricConfig{Enabled: false},
					ContainerCPUThrottlingDataThrottledPeriods: MetricConfig{Enabled: false},
					ContainerCPUThrottlingDataThrottledTime:    MetricConfig{Enabled: false},
					ContainerCPUUsageKernelmode:                MetricConfig{Enabled: false},
					ContainerCPUUsagePercpu:                    MetricConfig{Enabled: false},
					ContainerCPUUsageSystem:                    MetricConfig{Enabled: false},
					ContainerCPUUsageTotal:                     MetricConfig{Enabled: false},
					ContainerCPUUsageUsermode:                  MetricConfig{Enabled: false},
					ContainerMemoryActiveAnon:                  MetricConfig{Enabled: false},
					ContainerMemoryActiveFile:                  MetricConfig{Enabled: false},
					ContainerMemoryCache:                       MetricConfig{Enabled: false},
					ContainerMemoryDirty:                       MetricConfig{Enabled: false},
					ContainerMemoryHierarchicalMemoryLimit:     MetricConfig{Enabled: false},
					ContainerMemoryHierarchicalMemswLimit:      MetricConfig{Enabled: false},
					ContainerMemoryInactiveAnon:                MetricConfig{Enabled: false},
					ContainerMemoryInactiveFile:                MetricConfig{Enabled: false},
					ContainerMemoryMappedFile:                  MetricConfig{Enabled: false},
					ContainerMemoryPercent:                     MetricConfig{Enabled: false},
					ContainerMemoryPgfault:                     MetricConfig{Enabled: false},
					ContainerMemoryPgmajfault:                  MetricConfig{Enabled: false},
					ContainerMemoryPgpgin:                      MetricConfig{Enabled: false},
					ContainerMemoryPgpgout:                     MetricConfig{Enabled: false},
					ContainerMemoryRss:                         MetricConfig{Enabled: false},
					ContainerMemoryRssHuge:                     MetricConfig{Enabled: false},
					ContainerMemorySwap:                        MetricConfig{Enabled: false},
					ContainerMemoryTotalActiveAnon:             MetricConfig{Enabled: false},
					ContainerMemoryTotalActiveFile:             MetricConfig{Enabled: false},
					ContainerMemoryTotalCache:                  MetricConfig{Enabled: false},
					ContainerMemoryTotalDirty:                  MetricConfig{Enabled: false},
					ContainerMemoryTotalInactiveAnon:           MetricConfig{Enabled: false},
					ContainerMemoryTotalInactiveFile:           MetricConfig{Enabled: false},
					ContainerMemoryTotalMappedFile:             MetricConfig{Enabled: false},
					ContainerMemoryTotalPgfault:                MetricConfig{Enabled: false},
					ContainerMemoryTotalPgmajfault:             MetricConfig{Enabled: false},
					ContainerMemoryTotalPgpgin:                 MetricConfig{Enabled: false},
					ContainerMemoryTotalPgpgout:                MetricConfig{Enabled: false},
					ContainerMemoryTotalRss:                    MetricConfig{Enabled: false},
					ContainerMemoryTotalRssHuge:                MetricConfig{Enabled: false},
					ContainerMemoryTotalSwap:                   MetricConfig{Enabled: false},
					ContainerMemoryTotalUnevictable:            MetricConfig{Enabled: false},
					ContainerMemoryTotalWriteback:              MetricConfig{Enabled: false},
					ContainerMemoryUnevictable:                 MetricConfig{Enabled: false},
					ContainerMemoryUsageLimit:                  MetricConfig{Enabled: false},
					ContainerMemoryUsageMax:                    MetricConfig{Enabled: false},
					ContainerMemoryUsageTotal:                  MetricConfig{Enabled: false},
					ContainerMemoryWriteback:                   MetricConfig{Enabled: false},
					ContainerNetworkIoUsageRxBytes:             MetricConfig{Enabled: false},
					ContainerNetworkIoUsageRxDropped:           MetricConfig{Enabled: false},
					ContainerNetworkIoUsageRxErrors:            MetricConfig{Enabled: false},
					ContainerNetworkIoUsageRxPackets:           MetricConfig{Enabled: false},
					ContainerNetworkIoUsageTxBytes:             MetricConfig{Enabled: false},
					ContainerNetworkIoUsageTxDropped:           MetricConfig{Enabled: false},
					ContainerNetworkIoUsageTxErrors:            MetricConfig{Enabled: false},
					ContainerNetworkIoUsageTxPackets:           MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerHostname:  ResourceAttributeConfig{Enabled: false},
					ContainerID:        ResourceAttributeConfig{Enabled: false},
					ContainerImageName: ResourceAttributeConfig{Enabled: false},
					ContainerName:      ResourceAttributeConfig{Enabled: false},
					ContainerRuntime:   ResourceAttributeConfig{Enabled: false},
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
