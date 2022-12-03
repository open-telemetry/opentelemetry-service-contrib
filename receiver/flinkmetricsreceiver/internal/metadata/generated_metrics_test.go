// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), component.BuildInfo{}, WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["flink.job.checkpoint.count"] = true
	mb.RecordFlinkJobCheckpointCountDataPoint(ts, "1", AttributeCheckpoint(1))

	enabledMetrics["flink.job.checkpoint.in_progress"] = true
	mb.RecordFlinkJobCheckpointInProgressDataPoint(ts, "1")

	enabledMetrics["flink.job.last_checkpoint.size"] = true
	mb.RecordFlinkJobLastCheckpointSizeDataPoint(ts, "1")

	enabledMetrics["flink.job.last_checkpoint.time"] = true
	mb.RecordFlinkJobLastCheckpointTimeDataPoint(ts, "1")

	enabledMetrics["flink.job.restart.count"] = true
	mb.RecordFlinkJobRestartCountDataPoint(ts, "1")

	enabledMetrics["flink.jvm.class_loader.classes_loaded"] = true
	mb.RecordFlinkJvmClassLoaderClassesLoadedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.cpu.load"] = true
	mb.RecordFlinkJvmCPULoadDataPoint(ts, "1")

	enabledMetrics["flink.jvm.cpu.time"] = true
	mb.RecordFlinkJvmCPUTimeDataPoint(ts, "1")

	enabledMetrics["flink.jvm.gc.collections.count"] = true
	mb.RecordFlinkJvmGcCollectionsCountDataPoint(ts, "1", AttributeGarbageCollectorName(1))

	enabledMetrics["flink.jvm.gc.collections.time"] = true
	mb.RecordFlinkJvmGcCollectionsTimeDataPoint(ts, "1", AttributeGarbageCollectorName(1))

	enabledMetrics["flink.jvm.memory.direct.total_capacity"] = true
	mb.RecordFlinkJvmMemoryDirectTotalCapacityDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.direct.used"] = true
	mb.RecordFlinkJvmMemoryDirectUsedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.heap.committed"] = true
	mb.RecordFlinkJvmMemoryHeapCommittedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.heap.max"] = true
	mb.RecordFlinkJvmMemoryHeapMaxDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.heap.used"] = true
	mb.RecordFlinkJvmMemoryHeapUsedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.mapped.total_capacity"] = true
	mb.RecordFlinkJvmMemoryMappedTotalCapacityDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.mapped.used"] = true
	mb.RecordFlinkJvmMemoryMappedUsedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.metaspace.committed"] = true
	mb.RecordFlinkJvmMemoryMetaspaceCommittedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.metaspace.max"] = true
	mb.RecordFlinkJvmMemoryMetaspaceMaxDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.metaspace.used"] = true
	mb.RecordFlinkJvmMemoryMetaspaceUsedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.nonheap.committed"] = true
	mb.RecordFlinkJvmMemoryNonheapCommittedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.nonheap.max"] = true
	mb.RecordFlinkJvmMemoryNonheapMaxDataPoint(ts, "1")

	enabledMetrics["flink.jvm.memory.nonheap.used"] = true
	mb.RecordFlinkJvmMemoryNonheapUsedDataPoint(ts, "1")

	enabledMetrics["flink.jvm.threads.count"] = true
	mb.RecordFlinkJvmThreadsCountDataPoint(ts, "1")

	enabledMetrics["flink.memory.managed.total"] = true
	mb.RecordFlinkMemoryManagedTotalDataPoint(ts, "1")

	enabledMetrics["flink.memory.managed.used"] = true
	mb.RecordFlinkMemoryManagedUsedDataPoint(ts, "1")

	enabledMetrics["flink.operator.record.count"] = true
	mb.RecordFlinkOperatorRecordCountDataPoint(ts, "1", "attr-val", AttributeRecord(1))

	enabledMetrics["flink.operator.watermark.output"] = true
	mb.RecordFlinkOperatorWatermarkOutputDataPoint(ts, "1", "attr-val")

	enabledMetrics["flink.task.record.count"] = true
	mb.RecordFlinkTaskRecordCountDataPoint(ts, "1", AttributeRecord(1))

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		FlinkJobCheckpointCount:           MetricSettings{Enabled: true},
		FlinkJobCheckpointInProgress:      MetricSettings{Enabled: true},
		FlinkJobLastCheckpointSize:        MetricSettings{Enabled: true},
		FlinkJobLastCheckpointTime:        MetricSettings{Enabled: true},
		FlinkJobRestartCount:              MetricSettings{Enabled: true},
		FlinkJvmClassLoaderClassesLoaded:  MetricSettings{Enabled: true},
		FlinkJvmCPULoad:                   MetricSettings{Enabled: true},
		FlinkJvmCPUTime:                   MetricSettings{Enabled: true},
		FlinkJvmGcCollectionsCount:        MetricSettings{Enabled: true},
		FlinkJvmGcCollectionsTime:         MetricSettings{Enabled: true},
		FlinkJvmMemoryDirectTotalCapacity: MetricSettings{Enabled: true},
		FlinkJvmMemoryDirectUsed:          MetricSettings{Enabled: true},
		FlinkJvmMemoryHeapCommitted:       MetricSettings{Enabled: true},
		FlinkJvmMemoryHeapMax:             MetricSettings{Enabled: true},
		FlinkJvmMemoryHeapUsed:            MetricSettings{Enabled: true},
		FlinkJvmMemoryMappedTotalCapacity: MetricSettings{Enabled: true},
		FlinkJvmMemoryMappedUsed:          MetricSettings{Enabled: true},
		FlinkJvmMemoryMetaspaceCommitted:  MetricSettings{Enabled: true},
		FlinkJvmMemoryMetaspaceMax:        MetricSettings{Enabled: true},
		FlinkJvmMemoryMetaspaceUsed:       MetricSettings{Enabled: true},
		FlinkJvmMemoryNonheapCommitted:    MetricSettings{Enabled: true},
		FlinkJvmMemoryNonheapMax:          MetricSettings{Enabled: true},
		FlinkJvmMemoryNonheapUsed:         MetricSettings{Enabled: true},
		FlinkJvmThreadsCount:              MetricSettings{Enabled: true},
		FlinkMemoryManagedTotal:           MetricSettings{Enabled: true},
		FlinkMemoryManagedUsed:            MetricSettings{Enabled: true},
		FlinkOperatorRecordCount:          MetricSettings{Enabled: true},
		FlinkOperatorWatermarkOutput:      MetricSettings{Enabled: true},
		FlinkTaskRecordCount:              MetricSettings{Enabled: true},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))

	mb.RecordFlinkJobCheckpointCountDataPoint(ts, "1", AttributeCheckpoint(1))
	mb.RecordFlinkJobCheckpointInProgressDataPoint(ts, "1")
	mb.RecordFlinkJobLastCheckpointSizeDataPoint(ts, "1")
	mb.RecordFlinkJobLastCheckpointTimeDataPoint(ts, "1")
	mb.RecordFlinkJobRestartCountDataPoint(ts, "1")
	mb.RecordFlinkJvmClassLoaderClassesLoadedDataPoint(ts, "1")
	mb.RecordFlinkJvmCPULoadDataPoint(ts, "1")
	mb.RecordFlinkJvmCPUTimeDataPoint(ts, "1")
	mb.RecordFlinkJvmGcCollectionsCountDataPoint(ts, "1", AttributeGarbageCollectorName(1))
	mb.RecordFlinkJvmGcCollectionsTimeDataPoint(ts, "1", AttributeGarbageCollectorName(1))
	mb.RecordFlinkJvmMemoryDirectTotalCapacityDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryDirectUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMappedTotalCapacityDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMappedUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmThreadsCountDataPoint(ts, "1")
	mb.RecordFlinkMemoryManagedTotalDataPoint(ts, "1")
	mb.RecordFlinkMemoryManagedUsedDataPoint(ts, "1")
	mb.RecordFlinkOperatorRecordCountDataPoint(ts, "1", "attr-val", AttributeRecord(1))
	mb.RecordFlinkOperatorWatermarkOutputDataPoint(ts, "1", "attr-val")
	mb.RecordFlinkTaskRecordCountDataPoint(ts, "1", AttributeRecord(1))

	metrics := mb.Emit(WithFlinkJobName("attr-val"), WithFlinkResourceTypeJobmanager, WithFlinkSubtaskIndex("attr-val"), WithFlinkTaskName("attr-val"), WithFlinkTaskmanagerID("attr-val"), WithHostName("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("flink.job.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("flink.resource.type")
	assert.True(t, ok)
	assert.Equal(t, "jobmanager", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("flink.subtask.index")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("flink.task.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("flink.taskmanager.id")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("host.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "flink.job.checkpoint.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of checkpoints completed or failed.", ms.At(i).Description())
			assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("checkpoint")
			assert.True(t, ok)
			assert.Equal(t, "completed", attrVal.Str())
			validatedMetrics["flink.job.checkpoint.count"] = struct{}{}
		case "flink.job.checkpoint.in_progress":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of checkpoints in progress.", ms.At(i).Description())
			assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.job.checkpoint.in_progress"] = struct{}{}
		case "flink.job.last_checkpoint.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total size of the last checkpoint.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.job.last_checkpoint.size"] = struct{}{}
		case "flink.job.last_checkpoint.time":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The end to end duration of the last checkpoint.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.job.last_checkpoint.time"] = struct{}{}
		case "flink.job.restart.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of restarts since this job was submitted, including full restarts and fine-grained restarts.", ms.At(i).Description())
			assert.Equal(t, "{restarts}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.job.restart.count"] = struct{}{}
		case "flink.jvm.class_loader.classes_loaded":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of classes loaded since the start of the JVM.", ms.At(i).Description())
			assert.Equal(t, "{classes}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.class_loader.classes_loaded"] = struct{}{}
		case "flink.jvm.cpu.load":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The CPU usage of the JVM for a jobmanager or taskmanager.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["flink.jvm.cpu.load"] = struct{}{}
		case "flink.jvm.cpu.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The CPU time used by the JVM for a jobmanager or taskmanager.", ms.At(i).Description())
			assert.Equal(t, "ns", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.cpu.time"] = struct{}{}
		case "flink.jvm.gc.collections.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of collections that have occurred.", ms.At(i).Description())
			assert.Equal(t, "{collections}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("name")
			assert.True(t, ok)
			assert.Equal(t, "PS_MarkSweep", attrVal.Str())
			validatedMetrics["flink.jvm.gc.collections.count"] = struct{}{}
		case "flink.jvm.gc.collections.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total time spent performing garbage collection.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("name")
			assert.True(t, ok)
			assert.Equal(t, "PS_MarkSweep", attrVal.Str())
			validatedMetrics["flink.jvm.gc.collections.time"] = struct{}{}
		case "flink.jvm.memory.direct.total_capacity":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total capacity of all buffers in the direct buffer pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.direct.total_capacity"] = struct{}{}
		case "flink.jvm.memory.direct.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of memory used by the JVM for the direct buffer pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.direct.used"] = struct{}{}
		case "flink.jvm.memory.heap.committed":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of heap memory guaranteed to be available to the JVM.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.heap.committed"] = struct{}{}
		case "flink.jvm.memory.heap.max":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The maximum amount of heap memory that can be used for memory management.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.heap.max"] = struct{}{}
		case "flink.jvm.memory.heap.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of heap memory currently used.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.heap.used"] = struct{}{}
		case "flink.jvm.memory.mapped.total_capacity":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of buffers in the mapped buffer pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.mapped.total_capacity"] = struct{}{}
		case "flink.jvm.memory.mapped.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of memory used by the JVM for the mapped buffer pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.mapped.used"] = struct{}{}
		case "flink.jvm.memory.metaspace.committed":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of memory guaranteed to be available to the JVM in the Metaspace memory pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.metaspace.committed"] = struct{}{}
		case "flink.jvm.memory.metaspace.max":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The maximum amount of memory that can be used in the Metaspace memory pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.metaspace.max"] = struct{}{}
		case "flink.jvm.memory.metaspace.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of memory currently used in the Metaspace memory pool.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.metaspace.used"] = struct{}{}
		case "flink.jvm.memory.nonheap.committed":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of non-heap memory guaranteed to be available to the JVM.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.nonheap.committed"] = struct{}{}
		case "flink.jvm.memory.nonheap.max":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The maximum amount of non-heap memory that can be used for memory management.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.nonheap.max"] = struct{}{}
		case "flink.jvm.memory.nonheap.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of non-heap memory currently used.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.memory.nonheap.used"] = struct{}{}
		case "flink.jvm.threads.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of live threads.", ms.At(i).Description())
			assert.Equal(t, "{threads}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.jvm.threads.count"] = struct{}{}
		case "flink.memory.managed.total":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total amount of managed memory.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.memory.managed.total"] = struct{}{}
		case "flink.memory.managed.used":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of managed memory currently used.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["flink.memory.managed.used"] = struct{}{}
		case "flink.operator.record.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of records an operator has.", ms.At(i).Description())
			assert.Equal(t, "{records}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("name")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("record")
			assert.True(t, ok)
			assert.Equal(t, "in", attrVal.Str())
			validatedMetrics["flink.operator.record.count"] = struct{}{}
		case "flink.operator.watermark.output":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The last watermark this operator has emitted.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("name")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["flink.operator.watermark.output"] = struct{}{}
		case "flink.task.record.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of records a task has.", ms.At(i).Description())
			assert.Equal(t, "{records}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("record")
			assert.True(t, ok)
			assert.Equal(t, "in", attrVal.Str())
			validatedMetrics["flink.task.record.count"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		FlinkJobCheckpointCount:           MetricSettings{Enabled: false},
		FlinkJobCheckpointInProgress:      MetricSettings{Enabled: false},
		FlinkJobLastCheckpointSize:        MetricSettings{Enabled: false},
		FlinkJobLastCheckpointTime:        MetricSettings{Enabled: false},
		FlinkJobRestartCount:              MetricSettings{Enabled: false},
		FlinkJvmClassLoaderClassesLoaded:  MetricSettings{Enabled: false},
		FlinkJvmCPULoad:                   MetricSettings{Enabled: false},
		FlinkJvmCPUTime:                   MetricSettings{Enabled: false},
		FlinkJvmGcCollectionsCount:        MetricSettings{Enabled: false},
		FlinkJvmGcCollectionsTime:         MetricSettings{Enabled: false},
		FlinkJvmMemoryDirectTotalCapacity: MetricSettings{Enabled: false},
		FlinkJvmMemoryDirectUsed:          MetricSettings{Enabled: false},
		FlinkJvmMemoryHeapCommitted:       MetricSettings{Enabled: false},
		FlinkJvmMemoryHeapMax:             MetricSettings{Enabled: false},
		FlinkJvmMemoryHeapUsed:            MetricSettings{Enabled: false},
		FlinkJvmMemoryMappedTotalCapacity: MetricSettings{Enabled: false},
		FlinkJvmMemoryMappedUsed:          MetricSettings{Enabled: false},
		FlinkJvmMemoryMetaspaceCommitted:  MetricSettings{Enabled: false},
		FlinkJvmMemoryMetaspaceMax:        MetricSettings{Enabled: false},
		FlinkJvmMemoryMetaspaceUsed:       MetricSettings{Enabled: false},
		FlinkJvmMemoryNonheapCommitted:    MetricSettings{Enabled: false},
		FlinkJvmMemoryNonheapMax:          MetricSettings{Enabled: false},
		FlinkJvmMemoryNonheapUsed:         MetricSettings{Enabled: false},
		FlinkJvmThreadsCount:              MetricSettings{Enabled: false},
		FlinkMemoryManagedTotal:           MetricSettings{Enabled: false},
		FlinkMemoryManagedUsed:            MetricSettings{Enabled: false},
		FlinkOperatorRecordCount:          MetricSettings{Enabled: false},
		FlinkOperatorWatermarkOutput:      MetricSettings{Enabled: false},
		FlinkTaskRecordCount:              MetricSettings{Enabled: false},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))
	mb.RecordFlinkJobCheckpointCountDataPoint(ts, "1", AttributeCheckpoint(1))
	mb.RecordFlinkJobCheckpointInProgressDataPoint(ts, "1")
	mb.RecordFlinkJobLastCheckpointSizeDataPoint(ts, "1")
	mb.RecordFlinkJobLastCheckpointTimeDataPoint(ts, "1")
	mb.RecordFlinkJobRestartCountDataPoint(ts, "1")
	mb.RecordFlinkJvmClassLoaderClassesLoadedDataPoint(ts, "1")
	mb.RecordFlinkJvmCPULoadDataPoint(ts, "1")
	mb.RecordFlinkJvmCPUTimeDataPoint(ts, "1")
	mb.RecordFlinkJvmGcCollectionsCountDataPoint(ts, "1", AttributeGarbageCollectorName(1))
	mb.RecordFlinkJvmGcCollectionsTimeDataPoint(ts, "1", AttributeGarbageCollectorName(1))
	mb.RecordFlinkJvmMemoryDirectTotalCapacityDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryDirectUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryHeapUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMappedTotalCapacityDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMappedUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryMetaspaceUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapCommittedDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapMaxDataPoint(ts, "1")
	mb.RecordFlinkJvmMemoryNonheapUsedDataPoint(ts, "1")
	mb.RecordFlinkJvmThreadsCountDataPoint(ts, "1")
	mb.RecordFlinkMemoryManagedTotalDataPoint(ts, "1")
	mb.RecordFlinkMemoryManagedUsedDataPoint(ts, "1")
	mb.RecordFlinkOperatorRecordCountDataPoint(ts, "1", "attr-val", AttributeRecord(1))
	mb.RecordFlinkOperatorWatermarkOutputDataPoint(ts, "1", "attr-val")
	mb.RecordFlinkTaskRecordCountDataPoint(ts, "1", AttributeRecord(1))

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
