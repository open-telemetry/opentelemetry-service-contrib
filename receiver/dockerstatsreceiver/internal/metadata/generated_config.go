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

// MetricsConfig provides config for docker_stats metrics.
type MetricsConfig struct {
	ContainerBlockioIoMergedRecursive          MetricConfig `mapstructure:"container.blockio.io_merged_recursive"`
	ContainerBlockioIoQueuedRecursive          MetricConfig `mapstructure:"container.blockio.io_queued_recursive"`
	ContainerBlockioIoServiceBytesRecursive    MetricConfig `mapstructure:"container.blockio.io_service_bytes_recursive"`
	ContainerBlockioIoServiceTimeRecursive     MetricConfig `mapstructure:"container.blockio.io_service_time_recursive"`
	ContainerBlockioIoServicedRecursive        MetricConfig `mapstructure:"container.blockio.io_serviced_recursive"`
	ContainerBlockioIoTimeRecursive            MetricConfig `mapstructure:"container.blockio.io_time_recursive"`
	ContainerBlockioIoWaitTimeRecursive        MetricConfig `mapstructure:"container.blockio.io_wait_time_recursive"`
	ContainerBlockioSectorsRecursive           MetricConfig `mapstructure:"container.blockio.sectors_recursive"`
	ContainerCPUPercent                        MetricConfig `mapstructure:"container.cpu.percent"`
	ContainerCPUThrottlingDataPeriods          MetricConfig `mapstructure:"container.cpu.throttling_data.periods"`
	ContainerCPUThrottlingDataThrottledPeriods MetricConfig `mapstructure:"container.cpu.throttling_data.throttled_periods"`
	ContainerCPUThrottlingDataThrottledTime    MetricConfig `mapstructure:"container.cpu.throttling_data.throttled_time"`
	ContainerCPUUsageKernelmode                MetricConfig `mapstructure:"container.cpu.usage.kernelmode"`
	ContainerCPUUsagePercpu                    MetricConfig `mapstructure:"container.cpu.usage.percpu"`
	ContainerCPUUsageSystem                    MetricConfig `mapstructure:"container.cpu.usage.system"`
	ContainerCPUUsageTotal                     MetricConfig `mapstructure:"container.cpu.usage.total"`
	ContainerCPUUsageUsermode                  MetricConfig `mapstructure:"container.cpu.usage.usermode"`
	ContainerCPUUtilization                    MetricConfig `mapstructure:"container.cpu.utilization"`
	ContainerMemoryActiveAnon                  MetricConfig `mapstructure:"container.memory.active_anon"`
	ContainerMemoryActiveFile                  MetricConfig `mapstructure:"container.memory.active_file"`
	ContainerMemoryAnon                        MetricConfig `mapstructure:"container.memory.anon"`
	ContainerMemoryCache                       MetricConfig `mapstructure:"container.memory.cache"`
	ContainerMemoryDirty                       MetricConfig `mapstructure:"container.memory.dirty"`
	ContainerMemoryFile                        MetricConfig `mapstructure:"container.memory.file"`
	ContainerMemoryHierarchicalMemoryLimit     MetricConfig `mapstructure:"container.memory.hierarchical_memory_limit"`
	ContainerMemoryHierarchicalMemswLimit      MetricConfig `mapstructure:"container.memory.hierarchical_memsw_limit"`
	ContainerMemoryInactiveAnon                MetricConfig `mapstructure:"container.memory.inactive_anon"`
	ContainerMemoryInactiveFile                MetricConfig `mapstructure:"container.memory.inactive_file"`
	ContainerMemoryMappedFile                  MetricConfig `mapstructure:"container.memory.mapped_file"`
	ContainerMemoryPercent                     MetricConfig `mapstructure:"container.memory.percent"`
	ContainerMemoryPgfault                     MetricConfig `mapstructure:"container.memory.pgfault"`
	ContainerMemoryPgmajfault                  MetricConfig `mapstructure:"container.memory.pgmajfault"`
	ContainerMemoryPgpgin                      MetricConfig `mapstructure:"container.memory.pgpgin"`
	ContainerMemoryPgpgout                     MetricConfig `mapstructure:"container.memory.pgpgout"`
	ContainerMemoryRss                         MetricConfig `mapstructure:"container.memory.rss"`
	ContainerMemoryRssHuge                     MetricConfig `mapstructure:"container.memory.rss_huge"`
	ContainerMemoryTotalActiveAnon             MetricConfig `mapstructure:"container.memory.total_active_anon"`
	ContainerMemoryTotalActiveFile             MetricConfig `mapstructure:"container.memory.total_active_file"`
	ContainerMemoryTotalCache                  MetricConfig `mapstructure:"container.memory.total_cache"`
	ContainerMemoryTotalDirty                  MetricConfig `mapstructure:"container.memory.total_dirty"`
	ContainerMemoryTotalInactiveAnon           MetricConfig `mapstructure:"container.memory.total_inactive_anon"`
	ContainerMemoryTotalInactiveFile           MetricConfig `mapstructure:"container.memory.total_inactive_file"`
	ContainerMemoryTotalMappedFile             MetricConfig `mapstructure:"container.memory.total_mapped_file"`
	ContainerMemoryTotalPgfault                MetricConfig `mapstructure:"container.memory.total_pgfault"`
	ContainerMemoryTotalPgmajfault             MetricConfig `mapstructure:"container.memory.total_pgmajfault"`
	ContainerMemoryTotalPgpgin                 MetricConfig `mapstructure:"container.memory.total_pgpgin"`
	ContainerMemoryTotalPgpgout                MetricConfig `mapstructure:"container.memory.total_pgpgout"`
	ContainerMemoryTotalRss                    MetricConfig `mapstructure:"container.memory.total_rss"`
	ContainerMemoryTotalRssHuge                MetricConfig `mapstructure:"container.memory.total_rss_huge"`
	ContainerMemoryTotalUnevictable            MetricConfig `mapstructure:"container.memory.total_unevictable"`
	ContainerMemoryTotalWriteback              MetricConfig `mapstructure:"container.memory.total_writeback"`
	ContainerMemoryUnevictable                 MetricConfig `mapstructure:"container.memory.unevictable"`
	ContainerMemoryUsageLimit                  MetricConfig `mapstructure:"container.memory.usage.limit"`
	ContainerMemoryUsageMax                    MetricConfig `mapstructure:"container.memory.usage.max"`
	ContainerMemoryUsageTotal                  MetricConfig `mapstructure:"container.memory.usage.total"`
	ContainerMemoryWriteback                   MetricConfig `mapstructure:"container.memory.writeback"`
	ContainerNetworkIoUsageRxBytes             MetricConfig `mapstructure:"container.network.io.usage.rx_bytes"`
	ContainerNetworkIoUsageRxDropped           MetricConfig `mapstructure:"container.network.io.usage.rx_dropped"`
	ContainerNetworkIoUsageRxErrors            MetricConfig `mapstructure:"container.network.io.usage.rx_errors"`
	ContainerNetworkIoUsageRxPackets           MetricConfig `mapstructure:"container.network.io.usage.rx_packets"`
	ContainerNetworkIoUsageTxBytes             MetricConfig `mapstructure:"container.network.io.usage.tx_bytes"`
	ContainerNetworkIoUsageTxDropped           MetricConfig `mapstructure:"container.network.io.usage.tx_dropped"`
	ContainerNetworkIoUsageTxErrors            MetricConfig `mapstructure:"container.network.io.usage.tx_errors"`
	ContainerNetworkIoUsageTxPackets           MetricConfig `mapstructure:"container.network.io.usage.tx_packets"`
	ContainerPidsCount                         MetricConfig `mapstructure:"container.pids.count"`
	ContainerPidsLimit                         MetricConfig `mapstructure:"container.pids.limit"`
	ContainerUptime                            MetricConfig `mapstructure:"container.uptime"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		ContainerBlockioIoMergedRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioIoQueuedRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioIoServiceBytesRecursive: MetricConfig{
			Enabled: true,
		},
		ContainerBlockioIoServiceTimeRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioIoServicedRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioIoTimeRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioIoWaitTimeRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerBlockioSectorsRecursive: MetricConfig{
			Enabled: false,
		},
		ContainerCPUPercent: MetricConfig{
			Enabled: true,
		},
		ContainerCPUThrottlingDataPeriods: MetricConfig{
			Enabled: false,
		},
		ContainerCPUThrottlingDataThrottledPeriods: MetricConfig{
			Enabled: false,
		},
		ContainerCPUThrottlingDataThrottledTime: MetricConfig{
			Enabled: false,
		},
		ContainerCPUUsageKernelmode: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsagePercpu: MetricConfig{
			Enabled: false,
		},
		ContainerCPUUsageSystem: MetricConfig{
			Enabled: false,
		},
		ContainerCPUUsageTotal: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsageUsermode: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUtilization: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryActiveAnon: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryActiveFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryAnon: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryCache: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryDirty: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryFile: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryHierarchicalMemoryLimit: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryHierarchicalMemswLimit: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryInactiveAnon: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryInactiveFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryMappedFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryPercent: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryPgfault: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryPgmajfault: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryPgpgin: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryPgpgout: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryRss: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryRssHuge: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalActiveAnon: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalActiveFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalCache: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryTotalDirty: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalInactiveAnon: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalInactiveFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalMappedFile: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalPgfault: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalPgmajfault: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalPgpgin: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalPgpgout: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalRss: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalRssHuge: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalUnevictable: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryTotalWriteback: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryUnevictable: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryUsageLimit: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryUsageMax: MetricConfig{
			Enabled: false,
		},
		ContainerMemoryUsageTotal: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryWriteback: MetricConfig{
			Enabled: false,
		},
		ContainerNetworkIoUsageRxBytes: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageRxDropped: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageRxErrors: MetricConfig{
			Enabled: false,
		},
		ContainerNetworkIoUsageRxPackets: MetricConfig{
			Enabled: false,
		},
		ContainerNetworkIoUsageTxBytes: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageTxDropped: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageTxErrors: MetricConfig{
			Enabled: false,
		},
		ContainerNetworkIoUsageTxPackets: MetricConfig{
			Enabled: false,
		},
		ContainerPidsCount: MetricConfig{
			Enabled: false,
		},
		ContainerPidsLimit: MetricConfig{
			Enabled: false,
		},
		ContainerUptime: MetricConfig{
			Enabled: false,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for docker_stats resource attributes.
type ResourceAttributesConfig struct {
	ContainerHostname  ResourceAttributeConfig `mapstructure:"container.hostname"`
	ContainerID        ResourceAttributeConfig `mapstructure:"container.id"`
	ContainerImageName ResourceAttributeConfig `mapstructure:"container.image.name"`
	ContainerName      ResourceAttributeConfig `mapstructure:"container.name"`
	ContainerRuntime   ResourceAttributeConfig `mapstructure:"container.runtime"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		ContainerHostname: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerID: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerImageName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerRuntime: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for docker_stats metrics builder.
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
