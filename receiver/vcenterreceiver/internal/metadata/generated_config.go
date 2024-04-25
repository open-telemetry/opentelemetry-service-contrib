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

// MetricsConfig provides config for vcenter metrics.
type MetricsConfig struct {
	VcenterClusterCPUEffective      MetricConfig `mapstructure:"vcenter.cluster.cpu.effective"`
	VcenterClusterCPULimit          MetricConfig `mapstructure:"vcenter.cluster.cpu.limit"`
	VcenterClusterHostCount         MetricConfig `mapstructure:"vcenter.cluster.host.count"`
	VcenterClusterMemoryEffective   MetricConfig `mapstructure:"vcenter.cluster.memory.effective"`
	VcenterClusterMemoryLimit       MetricConfig `mapstructure:"vcenter.cluster.memory.limit"`
	VcenterClusterMemoryUsed        MetricConfig `mapstructure:"vcenter.cluster.memory.used"`
	VcenterClusterVMCount           MetricConfig `mapstructure:"vcenter.cluster.vm.count"`
	VcenterDatastoreDiskUsage       MetricConfig `mapstructure:"vcenter.datastore.disk.usage"`
	VcenterDatastoreDiskUtilization MetricConfig `mapstructure:"vcenter.datastore.disk.utilization"`
	VcenterHostCPUUsage             MetricConfig `mapstructure:"vcenter.host.cpu.usage"`
	VcenterHostCPUUtilization       MetricConfig `mapstructure:"vcenter.host.cpu.utilization"`
	VcenterHostDiskLatencyAvg       MetricConfig `mapstructure:"vcenter.host.disk.latency.avg"`
	VcenterHostDiskLatencyMax       MetricConfig `mapstructure:"vcenter.host.disk.latency.max"`
	VcenterHostDiskThroughput       MetricConfig `mapstructure:"vcenter.host.disk.throughput"`
	VcenterHostMemoryUsage          MetricConfig `mapstructure:"vcenter.host.memory.usage"`
	VcenterHostMemoryUtilization    MetricConfig `mapstructure:"vcenter.host.memory.utilization"`
	VcenterHostNetworkPacketCount   MetricConfig `mapstructure:"vcenter.host.network.packet.count"`
	VcenterHostNetworkPacketErrors  MetricConfig `mapstructure:"vcenter.host.network.packet.errors"`
	VcenterHostNetworkThroughput    MetricConfig `mapstructure:"vcenter.host.network.throughput"`
	VcenterHostNetworkUsage         MetricConfig `mapstructure:"vcenter.host.network.usage"`
	VcenterResourcePoolCPUShares    MetricConfig `mapstructure:"vcenter.resource_pool.cpu.shares"`
	VcenterResourcePoolCPUUsage     MetricConfig `mapstructure:"vcenter.resource_pool.cpu.usage"`
	VcenterResourcePoolMemoryShares MetricConfig `mapstructure:"vcenter.resource_pool.memory.shares"`
	VcenterResourcePoolMemoryUsage  MetricConfig `mapstructure:"vcenter.resource_pool.memory.usage"`
	VcenterVMCPUUsage               MetricConfig `mapstructure:"vcenter.vm.cpu.usage"`
	VcenterVMCPUUtilization         MetricConfig `mapstructure:"vcenter.vm.cpu.utilization"`
	VcenterVMDiskLatencyAvg         MetricConfig `mapstructure:"vcenter.vm.disk.latency.avg"`
	VcenterVMDiskLatencyMax         MetricConfig `mapstructure:"vcenter.vm.disk.latency.max"`
	VcenterVMDiskThroughput         MetricConfig `mapstructure:"vcenter.vm.disk.throughput"`
	VcenterVMDiskUsage              MetricConfig `mapstructure:"vcenter.vm.disk.usage"`
	VcenterVMDiskUtilization        MetricConfig `mapstructure:"vcenter.vm.disk.utilization"`
	VcenterVMMemoryBallooned        MetricConfig `mapstructure:"vcenter.vm.memory.ballooned"`
	VcenterVMMemorySwapped          MetricConfig `mapstructure:"vcenter.vm.memory.swapped"`
	VcenterVMMemorySwappedSsd       MetricConfig `mapstructure:"vcenter.vm.memory.swapped_ssd"`
	VcenterVMMemoryUsage            MetricConfig `mapstructure:"vcenter.vm.memory.usage"`
	VcenterVMMemoryUtilization      MetricConfig `mapstructure:"vcenter.vm.memory.utilization"`
	VcenterVMNetworkPacketCount     MetricConfig `mapstructure:"vcenter.vm.network.packet.count"`
	VcenterVMNetworkThroughput      MetricConfig `mapstructure:"vcenter.vm.network.throughput"`
	VcenterVMNetworkUsage           MetricConfig `mapstructure:"vcenter.vm.network.usage"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		VcenterClusterCPUEffective: MetricConfig{
			Enabled: true,
		},
		VcenterClusterCPULimit: MetricConfig{
			Enabled: true,
		},
		VcenterClusterHostCount: MetricConfig{
			Enabled: true,
		},
		VcenterClusterMemoryEffective: MetricConfig{
			Enabled: true,
		},
		VcenterClusterMemoryLimit: MetricConfig{
			Enabled: true,
		},
		VcenterClusterMemoryUsed: MetricConfig{
			Enabled: true,
		},
		VcenterClusterVMCount: MetricConfig{
			Enabled: true,
		},
		VcenterDatastoreDiskUsage: MetricConfig{
			Enabled: true,
		},
		VcenterDatastoreDiskUtilization: MetricConfig{
			Enabled: true,
		},
		VcenterHostCPUUsage: MetricConfig{
			Enabled: true,
		},
		VcenterHostCPUUtilization: MetricConfig{
			Enabled: true,
		},
		VcenterHostDiskLatencyAvg: MetricConfig{
			Enabled: true,
		},
		VcenterHostDiskLatencyMax: MetricConfig{
			Enabled: true,
		},
		VcenterHostDiskThroughput: MetricConfig{
			Enabled: true,
		},
		VcenterHostMemoryUsage: MetricConfig{
			Enabled: true,
		},
		VcenterHostMemoryUtilization: MetricConfig{
			Enabled: true,
		},
		VcenterHostNetworkPacketCount: MetricConfig{
			Enabled: true,
		},
		VcenterHostNetworkPacketErrors: MetricConfig{
			Enabled: true,
		},
		VcenterHostNetworkThroughput: MetricConfig{
			Enabled: true,
		},
		VcenterHostNetworkUsage: MetricConfig{
			Enabled: true,
		},
		VcenterResourcePoolCPUShares: MetricConfig{
			Enabled: true,
		},
		VcenterResourcePoolCPUUsage: MetricConfig{
			Enabled: true,
		},
		VcenterResourcePoolMemoryShares: MetricConfig{
			Enabled: true,
		},
		VcenterResourcePoolMemoryUsage: MetricConfig{
			Enabled: true,
		},
		VcenterVMCPUUsage: MetricConfig{
			Enabled: true,
		},
		VcenterVMCPUUtilization: MetricConfig{
			Enabled: true,
		},
		VcenterVMDiskLatencyAvg: MetricConfig{
			Enabled: true,
		},
		VcenterVMDiskLatencyMax: MetricConfig{
			Enabled: true,
		},
		VcenterVMDiskThroughput: MetricConfig{
			Enabled: true,
		},
		VcenterVMDiskUsage: MetricConfig{
			Enabled: true,
		},
		VcenterVMDiskUtilization: MetricConfig{
			Enabled: true,
		},
		VcenterVMMemoryBallooned: MetricConfig{
			Enabled: true,
		},
		VcenterVMMemorySwapped: MetricConfig{
			Enabled: true,
		},
		VcenterVMMemorySwappedSsd: MetricConfig{
			Enabled: true,
		},
		VcenterVMMemoryUsage: MetricConfig{
			Enabled: true,
		},
		VcenterVMMemoryUtilization: MetricConfig{
			Enabled: false,
		},
		VcenterVMNetworkPacketCount: MetricConfig{
			Enabled: true,
		},
		VcenterVMNetworkThroughput: MetricConfig{
			Enabled: true,
		},
		VcenterVMNetworkUsage: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

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

// ResourceAttributesConfig provides config for vcenter resource attributes.
type ResourceAttributesConfig struct {
	VcenterClusterName               ResourceAttributeConfig `mapstructure:"vcenter.cluster.name"`
	VcenterDatacenterName            ResourceAttributeConfig `mapstructure:"vcenter.datacenter.name"`
	VcenterDatastoreName             ResourceAttributeConfig `mapstructure:"vcenter.datastore.name"`
	VcenterHostName                  ResourceAttributeConfig `mapstructure:"vcenter.host.name"`
	VcenterResourcePoolInventoryPath ResourceAttributeConfig `mapstructure:"vcenter.resource_pool.inventory_path"`
	VcenterResourcePoolName          ResourceAttributeConfig `mapstructure:"vcenter.resource_pool.name"`
	VcenterVMID                      ResourceAttributeConfig `mapstructure:"vcenter.vm.id"`
	VcenterVMName                    ResourceAttributeConfig `mapstructure:"vcenter.vm.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		VcenterClusterName: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterDatacenterName: ResourceAttributeConfig{
			Enabled: false,
		},
		VcenterDatastoreName: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterHostName: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterResourcePoolInventoryPath: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterResourcePoolName: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterVMID: ResourceAttributeConfig{
			Enabled: true,
		},
		VcenterVMName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for vcenter metrics builder.
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
