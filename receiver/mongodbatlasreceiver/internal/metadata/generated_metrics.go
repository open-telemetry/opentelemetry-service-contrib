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

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "mongoatlasreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	MongodbatlasDbCounts                                  MetricIntf
	MongodbatlasDbSize                                    MetricIntf
	MongodbatlasDiskPartitionIopsAverage                  MetricIntf
	MongodbatlasDiskPartitionIopsMax                      MetricIntf
	MongodbatlasDiskPartitionLatencyAverage               MetricIntf
	MongodbatlasDiskPartitionLatencyMax                   MetricIntf
	MongodbatlasDiskPartitionSpaceAverage                 MetricIntf
	MongodbatlasDiskPartitionSpaceMax                     MetricIntf
	MongodbatlasDiskPartitionUsageAverage                 MetricIntf
	MongodbatlasDiskPartitionUsageMax                     MetricIntf
	MongodbatlasDiskPartitionUtilizationAverage           MetricIntf
	MongodbatlasDiskPartitionUtilizationMax               MetricIntf
	MongodbatlasProcessAsserts                            MetricIntf
	MongodbatlasProcessBackgroundFlush                    MetricIntf
	MongodbatlasProcessCacheIo                            MetricIntf
	MongodbatlasProcessCacheSize                          MetricIntf
	MongodbatlasProcessConnections                        MetricIntf
	MongodbatlasProcessCPUChildrenNormalizedUsageAverage  MetricIntf
	MongodbatlasProcessCPUChildrenNormalizedUsageMax      MetricIntf
	MongodbatlasProcessCPUChildrenUsageAverage            MetricIntf
	MongodbatlasProcessCPUChildrenUsageMax                MetricIntf
	MongodbatlasProcessCPUNormalizedUsageAverage          MetricIntf
	MongodbatlasProcessCPUNormalizedUsageMax              MetricIntf
	MongodbatlasProcessCPUUsageAverage                    MetricIntf
	MongodbatlasProcessCPUUsageMax                        MetricIntf
	MongodbatlasProcessCursors                            MetricIntf
	MongodbatlasProcessDbDocumentRate                     MetricIntf
	MongodbatlasProcessDbOperationsRate                   MetricIntf
	MongodbatlasProcessDbOperationsTime                   MetricIntf
	MongodbatlasProcessDbQueryExecutorScanned             MetricIntf
	MongodbatlasProcessDbQueryTargetingScannedPerReturned MetricIntf
	MongodbatlasProcessDbStorage                          MetricIntf
	MongodbatlasProcessFtsCPUUsage                        MetricIntf
	MongodbatlasProcessGlobalLock                         MetricIntf
	MongodbatlasProcessIndexBtreeMissRatio                MetricIntf
	MongodbatlasProcessIndexCounters                      MetricIntf
	MongodbatlasProcessJournalingCommits                  MetricIntf
	MongodbatlasProcessJournalingDataFiles                MetricIntf
	MongodbatlasProcessJournalingWritten                  MetricIntf
	MongodbatlasProcessMemoryUsage                        MetricIntf
	MongodbatlasProcessNetworkIo                          MetricIntf
	MongodbatlasProcessNetworkRequests                    MetricIntf
	MongodbatlasProcessOplogRate                          MetricIntf
	MongodbatlasProcessOplogTime                          MetricIntf
	MongodbatlasProcessPageFaults                         MetricIntf
	MongodbatlasProcessRestarts                           MetricIntf
	MongodbatlasProcessTickets                            MetricIntf
	MongodbatlasSystemCPUNormalizedUsageAverage           MetricIntf
	MongodbatlasSystemCPUNormalizedUsageMax               MetricIntf
	MongodbatlasSystemCPUUsageAverage                     MetricIntf
	MongodbatlasSystemCPUUsageMax                         MetricIntf
	MongodbatlasSystemFtsCPUNormalizedUsage               MetricIntf
	MongodbatlasSystemFtsCPUUsage                         MetricIntf
	MongodbatlasSystemFtsDiskUsed                         MetricIntf
	MongodbatlasSystemFtsMemoryUsage                      MetricIntf
	MongodbatlasSystemMemoryUsageAverage                  MetricIntf
	MongodbatlasSystemMemoryUsageMax                      MetricIntf
	MongodbatlasSystemNetworkIoAverage                    MetricIntf
	MongodbatlasSystemNetworkIoMax                        MetricIntf
	MongodbatlasSystemPagingIoAverage                     MetricIntf
	MongodbatlasSystemPagingIoMax                         MetricIntf
	MongodbatlasSystemPagingUsageAverage                  MetricIntf
	MongodbatlasSystemPagingUsageMax                      MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"mongodbatlas.db.counts",
		"mongodbatlas.db.size",
		"mongodbatlas.disk.partition.iops.average",
		"mongodbatlas.disk.partition.iops.max",
		"mongodbatlas.disk.partition.latency.average",
		"mongodbatlas.disk.partition.latency.max",
		"mongodbatlas.disk.partition.space.average",
		"mongodbatlas.disk.partition.space.max",
		"mongodbatlas.disk.partition.usage.average",
		"mongodbatlas.disk.partition.usage.max",
		"mongodbatlas.disk.partition.utilization.average",
		"mongodbatlas.disk.partition.utilization.max",
		"mongodbatlas.process.asserts",
		"mongodbatlas.process.background_flush",
		"mongodbatlas.process.cache.io",
		"mongodbatlas.process.cache.size",
		"mongodbatlas.process.connections",
		"mongodbatlas.process.cpu.children.normalized.usage.average",
		"mongodbatlas.process.cpu.children.normalized.usage.max",
		"mongodbatlas.process.cpu.children.usage.average",
		"mongodbatlas.process.cpu.children.usage.max",
		"mongodbatlas.process.cpu.normalized.usage.average",
		"mongodbatlas.process.cpu.normalized.usage.max",
		"mongodbatlas.process.cpu.usage.average",
		"mongodbatlas.process.cpu.usage.max",
		"mongodbatlas.process.cursors",
		"mongodbatlas.process.db.document.rate",
		"mongodbatlas.process.db.operations.rate",
		"mongodbatlas.process.db.operations.time",
		"mongodbatlas.process.db.query_executor.scanned",
		"mongodbatlas.process.db.query_targeting.scanned_per_returned",
		"mongodbatlas.process.db.storage",
		"mongodbatlas.process.fts.cpu.usage",
		"mongodbatlas.process.global_lock",
		"mongodbatlas.process.index.btree_miss_ratio",
		"mongodbatlas.process.index.counters",
		"mongodbatlas.process.journaling.commits",
		"mongodbatlas.process.journaling.data_files",
		"mongodbatlas.process.journaling.written",
		"mongodbatlas.process.memory.usage",
		"mongodbatlas.process.network.io",
		"mongodbatlas.process.network.requests",
		"mongodbatlas.process.oplog.rate",
		"mongodbatlas.process.oplog.time",
		"mongodbatlas.process.page_faults",
		"mongodbatlas.process.restarts",
		"mongodbatlas.process.tickets",
		"mongodbatlas.system.cpu.normalized.usage.average",
		"mongodbatlas.system.cpu.normalized.usage.max",
		"mongodbatlas.system.cpu.usage.average",
		"mongodbatlas.system.cpu.usage.max",
		"mongodbatlas.system.fts.cpu.normalized.usage",
		"mongodbatlas.system.fts.cpu.usage",
		"mongodbatlas.system.fts.disk.used",
		"mongodbatlas.system.fts.memory.usage",
		"mongodbatlas.system.memory.usage.average",
		"mongodbatlas.system.memory.usage.max",
		"mongodbatlas.system.network.io.average",
		"mongodbatlas.system.network.io.max",
		"mongodbatlas.system.paging.io.average",
		"mongodbatlas.system.paging.io.max",
		"mongodbatlas.system.paging.usage.average",
		"mongodbatlas.system.paging.usage.max",
	}
}

var metricsByName = map[string]MetricIntf{
	"mongodbatlas.db.counts":                                       Metrics.MongodbatlasDbCounts,
	"mongodbatlas.db.size":                                         Metrics.MongodbatlasDbSize,
	"mongodbatlas.disk.partition.iops.average":                     Metrics.MongodbatlasDiskPartitionIopsAverage,
	"mongodbatlas.disk.partition.iops.max":                         Metrics.MongodbatlasDiskPartitionIopsMax,
	"mongodbatlas.disk.partition.latency.average":                  Metrics.MongodbatlasDiskPartitionLatencyAverage,
	"mongodbatlas.disk.partition.latency.max":                      Metrics.MongodbatlasDiskPartitionLatencyMax,
	"mongodbatlas.disk.partition.space.average":                    Metrics.MongodbatlasDiskPartitionSpaceAverage,
	"mongodbatlas.disk.partition.space.max":                        Metrics.MongodbatlasDiskPartitionSpaceMax,
	"mongodbatlas.disk.partition.usage.average":                    Metrics.MongodbatlasDiskPartitionUsageAverage,
	"mongodbatlas.disk.partition.usage.max":                        Metrics.MongodbatlasDiskPartitionUsageMax,
	"mongodbatlas.disk.partition.utilization.average":              Metrics.MongodbatlasDiskPartitionUtilizationAverage,
	"mongodbatlas.disk.partition.utilization.max":                  Metrics.MongodbatlasDiskPartitionUtilizationMax,
	"mongodbatlas.process.asserts":                                 Metrics.MongodbatlasProcessAsserts,
	"mongodbatlas.process.background_flush":                        Metrics.MongodbatlasProcessBackgroundFlush,
	"mongodbatlas.process.cache.io":                                Metrics.MongodbatlasProcessCacheIo,
	"mongodbatlas.process.cache.size":                              Metrics.MongodbatlasProcessCacheSize,
	"mongodbatlas.process.connections":                             Metrics.MongodbatlasProcessConnections,
	"mongodbatlas.process.cpu.children.normalized.usage.average":   Metrics.MongodbatlasProcessCPUChildrenNormalizedUsageAverage,
	"mongodbatlas.process.cpu.children.normalized.usage.max":       Metrics.MongodbatlasProcessCPUChildrenNormalizedUsageMax,
	"mongodbatlas.process.cpu.children.usage.average":              Metrics.MongodbatlasProcessCPUChildrenUsageAverage,
	"mongodbatlas.process.cpu.children.usage.max":                  Metrics.MongodbatlasProcessCPUChildrenUsageMax,
	"mongodbatlas.process.cpu.normalized.usage.average":            Metrics.MongodbatlasProcessCPUNormalizedUsageAverage,
	"mongodbatlas.process.cpu.normalized.usage.max":                Metrics.MongodbatlasProcessCPUNormalizedUsageMax,
	"mongodbatlas.process.cpu.usage.average":                       Metrics.MongodbatlasProcessCPUUsageAverage,
	"mongodbatlas.process.cpu.usage.max":                           Metrics.MongodbatlasProcessCPUUsageMax,
	"mongodbatlas.process.cursors":                                 Metrics.MongodbatlasProcessCursors,
	"mongodbatlas.process.db.document.rate":                        Metrics.MongodbatlasProcessDbDocumentRate,
	"mongodbatlas.process.db.operations.rate":                      Metrics.MongodbatlasProcessDbOperationsRate,
	"mongodbatlas.process.db.operations.time":                      Metrics.MongodbatlasProcessDbOperationsTime,
	"mongodbatlas.process.db.query_executor.scanned":               Metrics.MongodbatlasProcessDbQueryExecutorScanned,
	"mongodbatlas.process.db.query_targeting.scanned_per_returned": Metrics.MongodbatlasProcessDbQueryTargetingScannedPerReturned,
	"mongodbatlas.process.db.storage":                              Metrics.MongodbatlasProcessDbStorage,
	"mongodbatlas.process.fts.cpu.usage":                           Metrics.MongodbatlasProcessFtsCPUUsage,
	"mongodbatlas.process.global_lock":                             Metrics.MongodbatlasProcessGlobalLock,
	"mongodbatlas.process.index.btree_miss_ratio":                  Metrics.MongodbatlasProcessIndexBtreeMissRatio,
	"mongodbatlas.process.index.counters":                          Metrics.MongodbatlasProcessIndexCounters,
	"mongodbatlas.process.journaling.commits":                      Metrics.MongodbatlasProcessJournalingCommits,
	"mongodbatlas.process.journaling.data_files":                   Metrics.MongodbatlasProcessJournalingDataFiles,
	"mongodbatlas.process.journaling.written":                      Metrics.MongodbatlasProcessJournalingWritten,
	"mongodbatlas.process.memory.usage":                            Metrics.MongodbatlasProcessMemoryUsage,
	"mongodbatlas.process.network.io":                              Metrics.MongodbatlasProcessNetworkIo,
	"mongodbatlas.process.network.requests":                        Metrics.MongodbatlasProcessNetworkRequests,
	"mongodbatlas.process.oplog.rate":                              Metrics.MongodbatlasProcessOplogRate,
	"mongodbatlas.process.oplog.time":                              Metrics.MongodbatlasProcessOplogTime,
	"mongodbatlas.process.page_faults":                             Metrics.MongodbatlasProcessPageFaults,
	"mongodbatlas.process.restarts":                                Metrics.MongodbatlasProcessRestarts,
	"mongodbatlas.process.tickets":                                 Metrics.MongodbatlasProcessTickets,
	"mongodbatlas.system.cpu.normalized.usage.average":             Metrics.MongodbatlasSystemCPUNormalizedUsageAverage,
	"mongodbatlas.system.cpu.normalized.usage.max":                 Metrics.MongodbatlasSystemCPUNormalizedUsageMax,
	"mongodbatlas.system.cpu.usage.average":                        Metrics.MongodbatlasSystemCPUUsageAverage,
	"mongodbatlas.system.cpu.usage.max":                            Metrics.MongodbatlasSystemCPUUsageMax,
	"mongodbatlas.system.fts.cpu.normalized.usage":                 Metrics.MongodbatlasSystemFtsCPUNormalizedUsage,
	"mongodbatlas.system.fts.cpu.usage":                            Metrics.MongodbatlasSystemFtsCPUUsage,
	"mongodbatlas.system.fts.disk.used":                            Metrics.MongodbatlasSystemFtsDiskUsed,
	"mongodbatlas.system.fts.memory.usage":                         Metrics.MongodbatlasSystemFtsMemoryUsage,
	"mongodbatlas.system.memory.usage.average":                     Metrics.MongodbatlasSystemMemoryUsageAverage,
	"mongodbatlas.system.memory.usage.max":                         Metrics.MongodbatlasSystemMemoryUsageMax,
	"mongodbatlas.system.network.io.average":                       Metrics.MongodbatlasSystemNetworkIoAverage,
	"mongodbatlas.system.network.io.max":                           Metrics.MongodbatlasSystemNetworkIoMax,
	"mongodbatlas.system.paging.io.average":                        Metrics.MongodbatlasSystemPagingIoAverage,
	"mongodbatlas.system.paging.io.max":                            Metrics.MongodbatlasSystemPagingIoMax,
	"mongodbatlas.system.paging.usage.average":                     Metrics.MongodbatlasSystemPagingUsageAverage,
	"mongodbatlas.system.paging.usage.max":                         Metrics.MongodbatlasSystemPagingUsageMax,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"mongodbatlas.db.counts",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.db.counts")
			metric.SetDescription("Database feature size")
			metric.SetUnit("{objects}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.db.size",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.db.size")
			metric.SetDescription("Database feature size")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.iops.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.iops.average")
			metric.SetDescription("Disk partition iops")
			metric.SetUnit("{ops}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.iops.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.iops.max")
			metric.SetDescription("Disk partition iops")
			metric.SetUnit("{ops}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.latency.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.latency.average")
			metric.SetDescription("Disk partition latency")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.latency.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.latency.max")
			metric.SetDescription("Disk partition latency")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.space.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.space.average")
			metric.SetDescription("Disk partition space")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.space.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.space.max")
			metric.SetDescription("Disk partition space")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.usage.average")
			metric.SetDescription("Disk partition usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.usage.max")
			metric.SetDescription("Disk partition usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.utilization.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.utilization.average")
			metric.SetDescription("Disk partition utilization (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.disk.partition.utilization.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.disk.partition.utilization.max")
			metric.SetDescription("Disk partition utilization (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.asserts",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.asserts")
			metric.SetDescription("Number of assertions per second")
			metric.SetUnit("{assertions}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.background_flush",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.background_flush")
			metric.SetDescription("Amount of data flushed in the background")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cache.io",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cache.io")
			metric.SetDescription("Cache throughput (per second)")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cache.size",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cache.size")
			metric.SetDescription("Cache sizes")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodbatlas.process.connections",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.connections")
			metric.SetDescription("Number of current connections")
			metric.SetUnit("{connections}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.children.normalized.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.children.normalized.usage.average")
			metric.SetDescription("CPU Usage for child processes, normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.children.normalized.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.children.normalized.usage.max")
			metric.SetDescription("CPU Usage for child processes, normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.children.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.children.usage.average")
			metric.SetDescription("CPU Usage for child processes (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.children.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.children.usage.max")
			metric.SetDescription("CPU Usage for child processes (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.normalized.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.normalized.usage.average")
			metric.SetDescription("CPU Usage, normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.normalized.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.normalized.usage.max")
			metric.SetDescription("CPU Usage, normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.usage.average")
			metric.SetDescription("CPU Usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cpu.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cpu.usage.max")
			metric.SetDescription("CPU Usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.cursors",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.cursors")
			metric.SetDescription("Number of cursors")
			metric.SetUnit("{cursors}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.document.rate",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.document.rate")
			metric.SetDescription("Document access rates")
			metric.SetUnit("{documents}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.operations.rate",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.operations.rate")
			metric.SetDescription("DB Operation Rates")
			metric.SetUnit("{operations}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.operations.time",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.operations.time")
			metric.SetDescription("DB Operation Times")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.query_executor.scanned",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.query_executor.scanned")
			metric.SetDescription("Scanned objects")
			metric.SetUnit("{objects}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.query_targeting.scanned_per_returned",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.query_targeting.scanned_per_returned")
			metric.SetDescription("Scanned objects per returned")
			metric.SetUnit("{scanned}/{returned}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.db.storage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.db.storage")
			metric.SetDescription("Storage used by the database")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.fts.cpu.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.fts.cpu.usage")
			metric.SetDescription("Full text search CPU (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.global_lock",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.global_lock")
			metric.SetDescription("Number and status of locks")
			metric.SetUnit("{locks}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.index.btree_miss_ratio",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.index.btree_miss_ratio")
			metric.SetDescription("Index miss ratio (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.index.counters",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.index.counters")
			metric.SetDescription("Indexes")
			metric.SetUnit("{indexes}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.journaling.commits",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.journaling.commits")
			metric.SetDescription("Journaling commits")
			metric.SetUnit("{commits}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.journaling.data_files",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.journaling.data_files")
			metric.SetDescription("Data file sizes")
			metric.SetUnit("MiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.journaling.written",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.journaling.written")
			metric.SetDescription("Journals written")
			metric.SetUnit("MiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.memory.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.memory.usage")
			metric.SetDescription("Memory Usage")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.network.io",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.network.io")
			metric.SetDescription("Network IO")
			metric.SetUnit("By/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.network.requests",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.network.requests")
			metric.SetDescription("Network requests")
			metric.SetUnit("{requests}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodbatlas.process.oplog.rate",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.oplog.rate")
			metric.SetDescription("Execution rate by operation")
			metric.SetUnit("GiBy/h")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.oplog.time",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.oplog.time")
			metric.SetDescription("Execution time by operation")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.page_faults",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.page_faults")
			metric.SetDescription("Page faults")
			metric.SetUnit("{faults}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.restarts",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.restarts")
			metric.SetDescription("Restarts in last hour")
			metric.SetUnit("{restarts}/h")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.process.tickets",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.process.tickets")
			metric.SetDescription("Tickets")
			metric.SetUnit("{tickets}")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.cpu.normalized.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.cpu.normalized.usage.average")
			metric.SetDescription("System CPU Normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.cpu.normalized.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.cpu.normalized.usage.max")
			metric.SetDescription("System CPU Normalized to pct")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.cpu.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.cpu.usage.average")
			metric.SetDescription("System CPU Usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.cpu.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.cpu.usage.max")
			metric.SetDescription("System CPU Usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.fts.cpu.normalized.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.fts.cpu.normalized.usage")
			metric.SetDescription("Full text search disk usage (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.fts.cpu.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.fts.cpu.usage")
			metric.SetDescription("Full-text search (%)")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.fts.disk.used",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.fts.disk.used")
			metric.SetDescription("Full text search disk usage")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.fts.memory.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.fts.memory.usage")
			metric.SetDescription("Full-text search")
			metric.SetUnit("MiBy")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodbatlas.system.memory.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.memory.usage.average")
			metric.SetDescription("System Memory Usage")
			metric.SetUnit("KiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.memory.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.memory.usage.max")
			metric.SetDescription("System Memory Usage")
			metric.SetUnit("KiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.network.io.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.network.io.average")
			metric.SetDescription("System Network IO")
			metric.SetUnit("By/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.network.io.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.network.io.max")
			metric.SetDescription("System Network IO")
			metric.SetUnit("By/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.paging.io.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.paging.io.average")
			metric.SetDescription("Swap IO")
			metric.SetUnit("{pages}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.paging.io.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.paging.io.max")
			metric.SetDescription("Swap IO")
			metric.SetUnit("{pages}/s")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.paging.usage.average",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.paging.usage.average")
			metric.SetDescription("Swap usage")
			metric.SetUnit("KiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"mongodbatlas.system.paging.usage.max",
		func(metric pdata.Metric) {
			metric.SetName("mongodbatlas.system.paging.usage.max")
			metric.SetDescription("Swap usage")
			metric.SetUnit("KiBy")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
}{}

// A is an alias for Attributes.
var A = Attributes
