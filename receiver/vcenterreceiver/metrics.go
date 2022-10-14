// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"context"

	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

func (v *vcenterMetricScraper) recordHostSystemMemoryUsage(
	now pcommon.Timestamp,
	hs mo.HostSystem,
) {
	s := hs.Summary
	h := s.Hardware
	z := s.QuickStats

	memUsage := z.OverallMemoryUsage
	v.mb.RecordVcenterHostMemoryUsageDataPoint(now, int64(memUsage))

	memUtilization := 100 * float64(z.OverallMemoryUsage) / float64(h.MemorySize>>20)
	v.mb.RecordVcenterHostMemoryUtilizationDataPoint(now, memUtilization)

	ncpu := int32(h.NumCpuCores)
	cpuUsage := z.OverallCpuUsage
	cpuUtilization := 100 * float64(z.OverallCpuUsage) / float64(ncpu*h.CpuMhz)

	v.mb.RecordVcenterHostCPUUsageDataPoint(now, int64(cpuUsage))
	v.mb.RecordVcenterHostCPUUtilizationDataPoint(now, cpuUtilization)
}

func (v *vcenterMetricScraper) recordVMUsages(
	now pcommon.Timestamp,
	vm mo.VirtualMachine,
) {
	memUsage := vm.Summary.QuickStats.GuestMemoryUsage
	balloonedMem := vm.Summary.QuickStats.BalloonedMemory
	v.mb.RecordVcenterVMMemoryUsageDataPoint(now, int64(memUsage))
	v.mb.RecordVcenterVMMemoryBalloonedDataPoint(now, int64(balloonedMem))

	diskUsed := vm.Summary.Storage.Committed
	diskFree := vm.Summary.Storage.Uncommitted

	v.mb.RecordVcenterVMDiskUsageDataPoint(now, diskUsed, metadata.AttributeDiskStateUsed)
	v.mb.RecordVcenterVMDiskUsageDataPoint(now, diskFree, metadata.AttributeDiskStateAvailable)
	if diskFree != 0 {
		diskUtilization := float64(diskUsed) / float64(diskFree+diskUsed) * 100
		v.mb.RecordVcenterVMDiskUtilizationDataPoint(now, diskUtilization)
	}
}

func (v *vcenterMetricScraper) recordDatastoreProperties(
	now pcommon.Timestamp,
	ds mo.Datastore,
) {
	s := ds.Summary
	diskUsage := s.Capacity - s.FreeSpace
	diskUtilization := float64(diskUsage) / float64(s.Capacity) * 100
	v.mb.RecordVcenterDatastoreDiskUsageDataPoint(now, diskUsage, metadata.AttributeDiskStateUsed)
	v.mb.RecordVcenterDatastoreDiskUsageDataPoint(now, s.Capacity, metadata.AttributeDiskStateUsed)
	v.mb.RecordVcenterDatastoreDiskUtilizationDataPoint(now, diskUtilization)
}

func (v *vcenterMetricScraper) recordResourcePool(
	now pcommon.Timestamp,
	rp mo.ResourcePool,
) {
	s := rp.Summary.GetResourcePoolSummary()
	if s.QuickStats != nil {
		v.mb.RecordVcenterResourcePoolCPUUsageDataPoint(now, s.QuickStats.OverallCpuUsage)
		v.mb.RecordVcenterResourcePoolMemoryUsageDataPoint(now, s.QuickStats.GuestMemoryUsage)
	}

	v.mb.RecordVcenterResourcePoolCPUSharesDataPoint(now, int64(s.Config.CpuAllocation.Shares.Shares))
	v.mb.RecordVcenterResourcePoolMemorySharesDataPoint(now, int64(s.Config.MemoryAllocation.Shares.Shares))

}

var hostPerfMetricList = []string{
	// network metrics
	"net.bytesTx.average",
	"net.bytesRx.average",
	"net.packetsTx.summation",
	"net.packetsRx.summation",
	"net.usage.average",
	"net.errorsRx.summation",
	"net.errorsTx.summation",

	// disk metrics
	"virtualDisk.totalWriteLatency.average",
	"disk.deviceReadLatency.average",
	"disk.deviceWriteLatency.average",
	"disk.kernelReadLatency.average",
	"disk.kernelWriteLatency.average",
	"disk.maxTotalLatency.latest",
	"disk.read.average",
	"disk.write.average",
}

func (v *vcenterMetricScraper) recordHostPerformanceMetrics(
	ctx context.Context,
	host mo.HostSystem,
	errs *scrapererror.ScrapeErrors,
) {
	spec := types.PerfQuerySpec{
		Entity:    host.Reference(),
		MaxSample: 5,
		Format:    string(types.PerfFormatNormal),
		MetricId:  []types.PerfMetricId{{Instance: "*"}},
		// right now we are only grabbing real time metrics from the performance
		// manager
		IntervalId: int32(20),
	}

	info, err := v.client.performanceQuery(ctx, spec, hostPerfMetricList, []types.ManagedObjectReference{host.Reference()})
	if err != nil {
		errs.AddPartial(1, err)
		return
	}
	v.processHostPerformance(info.results)
}

// vmPerfMetricList may be customizable in the future but here is the full list of Virtual Machine Performance Counters
// https://docs.vmware.com/en/vRealize-Operations/8.6/com.vmware.vcom.metrics.doc/GUID-1322F5A4-DA1D-481F-BBEA-99B228E96AF2.html
var vmPerfMetricList = []string{
	// network metrics
	"net.packetsTx.summation",
	"net.packetsRx.summation",
	"net.bytesRx.average",
	"net.bytesTx.average",
	"net.usage.average",

	// disk metrics
	"disk.totalWriteLatency.average",
	"disk.totalReadLatency.average",
	"disk.maxTotalLatency.latest",
	"virtualDisk.totalWriteLatency.average",
	"virtualDisk.totalReadLatency.average",
}

func (v *vcenterMetricScraper) recordVMPerformance(
	ctx context.Context,
	vm mo.VirtualMachine,
	errs *scrapererror.ScrapeErrors,
) {
	spec := types.PerfQuerySpec{
		Entity: vm.Reference(),
		Format: string(types.PerfFormatNormal),
		// Just grabbing real time performance metrics of the current
		// supported metrics by this receiver. If more are added we may need
		// a system of making this user customizable or adapt to use a 5 minute interval per metric
		IntervalId: int32(20),
	}

	info, err := v.client.performanceQuery(ctx, spec, vmPerfMetricList, []types.ManagedObjectReference{vm.Reference()})
	if err != nil {
		errs.AddPartial(1, err)
		return
	}

	v.processVMPerformanceMetrics(info)
}

func (v *vcenterMetricScraper) processVMPerformanceMetrics(info *perfSampleResult) {
	for _, m := range info.results {
		for _, val := range m.Value {
			for j, nestedValue := range val.Value {
				si := m.SampleInfo[j]
				switch val.Name {
				// Performance monitoring level 1 metrics
				case "net.bytesTx.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMNetworkThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionTransmitted)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMNetworkThroughputTransmitDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.bytesRx.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMNetworkThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionReceived)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMNetworkThroughputReceiveDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.usage.average":
					v.mb.RecordVcenterVMNetworkUsageDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
				case "net.packetsTx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMNetworkPacketCountDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionTransmitted)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMNetworkPacketCountTransmitDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.packetsRx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMNetworkPacketCountDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionReceived)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMNetworkPacketCountReceiveDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}

				// Performance monitoring level 2 metrics required
				case "disk.totalReadLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionRead, metadata.AttributeDiskTypePhysical)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgReadDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskTypePhysical)
					}
				case "virtualDisk.totalReadLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionRead, metadata.AttributeDiskTypeVirtual)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgReadDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskTypeVirtual)
					}
				case "disk.totalWriteLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionWrite, metadata.AttributeDiskTypePhysical)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgWriteDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskTypePhysical)
					}
				case "virtualDisk.totalWriteLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionWrite, metadata.AttributeDiskTypeVirtual)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterVMDiskLatencyAvgWriteDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskTypeVirtual)
					}
				case "disk.maxTotalLatency.latest":
					v.mb.RecordVcenterVMDiskLatencyMaxDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
				}
			}
		}
	}
}

func (v *vcenterMetricScraper) processHostPerformance(metrics []performance.EntityMetric) {
	for _, m := range metrics {
		for _, val := range m.Value {
			for j, nestedValue := range val.Value {
				si := m.SampleInfo[j]
				switch val.Name {
				// Performance monitoring level 1 metrics
				case "net.usage.average":
					v.mb.RecordVcenterHostNetworkUsageDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
				case "net.bytesTx.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionTransmitted)
					}

					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkThroughputTransmitDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.bytesRx.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionReceived)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkThroughputReceiveDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.packetsTx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketCountDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionTransmitted)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketCountTransmitDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.packetsRx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketCountDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionReceived)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketCountReceiveDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}

				// Following requires performance level 2
				case "net.errorsRx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketErrorsDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionReceived)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketErrorsReceiveDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "net.errorsTx.summation":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketErrorsDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeThroughputDirectionTransmitted)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostNetworkPacketErrorsTransmitDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "disk.totalWriteLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionWrite)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostDiskLatencyAvgWriteDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "disk.totalReadLatency.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostDiskLatencyAvgDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionRead)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostDiskLatencyAvgReadDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "disk.maxTotalLatency.latest":
					v.mb.RecordVcenterHostDiskLatencyMaxDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)

				// Following requires performance level 4
				case "disk.read.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostDiskThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionRead)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostDiskThroughputReadDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				case "disk.write.average":
					if v.emitMetricsWithDirectionAttribute {
						v.mb.RecordVcenterHostDiskThroughputDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue, metadata.AttributeDiskDirectionWrite)
					}
					if v.emitMetricsWithoutDirectionAttribute {
						v.mb.RecordVcenterHostDiskThroughputWriteDataPoint(pcommon.NewTimestampFromTime(si.Timestamp), nestedValue)
					}
				}
			}
		}
	}
}
