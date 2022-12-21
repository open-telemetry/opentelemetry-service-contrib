// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), receivertest.NewNopCreateSettings(), WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["process.runtime.memstats.buck_hash_sys"] = true
	mb.RecordProcessRuntimeMemstatsBuckHashSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.frees"] = true
	mb.RecordProcessRuntimeMemstatsFreesDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.gc_cpu_fraction"] = true
	mb.RecordProcessRuntimeMemstatsGcCPUFractionDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.gc_sys"] = true
	mb.RecordProcessRuntimeMemstatsGcSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_alloc"] = true
	mb.RecordProcessRuntimeMemstatsHeapAllocDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_idle"] = true
	mb.RecordProcessRuntimeMemstatsHeapIdleDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_inuse"] = true
	mb.RecordProcessRuntimeMemstatsHeapInuseDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_objects"] = true
	mb.RecordProcessRuntimeMemstatsHeapObjectsDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_released"] = true
	mb.RecordProcessRuntimeMemstatsHeapReleasedDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.heap_sys"] = true
	mb.RecordProcessRuntimeMemstatsHeapSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.last_pause"] = true
	mb.RecordProcessRuntimeMemstatsLastPauseDataPoint(ts, 1)

	mb.RecordProcessRuntimeMemstatsLookupsDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.mallocs"] = true
	mb.RecordProcessRuntimeMemstatsMallocsDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.mcache_inuse"] = true
	mb.RecordProcessRuntimeMemstatsMcacheInuseDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.mcache_sys"] = true
	mb.RecordProcessRuntimeMemstatsMcacheSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.mspan_inuse"] = true
	mb.RecordProcessRuntimeMemstatsMspanInuseDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.mspan_sys"] = true
	mb.RecordProcessRuntimeMemstatsMspanSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.next_gc"] = true
	mb.RecordProcessRuntimeMemstatsNextGcDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.num_forced_gc"] = true
	mb.RecordProcessRuntimeMemstatsNumForcedGcDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.num_gc"] = true
	mb.RecordProcessRuntimeMemstatsNumGcDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.other_sys"] = true
	mb.RecordProcessRuntimeMemstatsOtherSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.pause_total"] = true
	mb.RecordProcessRuntimeMemstatsPauseTotalDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.stack_inuse"] = true
	mb.RecordProcessRuntimeMemstatsStackInuseDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.stack_sys"] = true
	mb.RecordProcessRuntimeMemstatsStackSysDataPoint(ts, 1)

	enabledMetrics["process.runtime.memstats.sys"] = true
	mb.RecordProcessRuntimeMemstatsSysDataPoint(ts, 1)

	mb.RecordProcessRuntimeMemstatsTotalAllocDataPoint(ts, 1)

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
	metricsSettings := MetricsSettings{
		ProcessRuntimeMemstatsBuckHashSys:   MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsFrees:         MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsGcCPUFraction: MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsGcSys:         MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapAlloc:     MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapIdle:      MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapInuse:     MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapObjects:   MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapReleased:  MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsHeapSys:       MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsLastPause:     MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsLookups:       MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsMallocs:       MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsMcacheInuse:   MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsMcacheSys:     MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsMspanInuse:    MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsMspanSys:      MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsNextGc:        MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsNumForcedGc:   MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsNumGc:         MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsOtherSys:      MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsPauseTotal:    MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsStackInuse:    MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsStackSys:      MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsSys:           MetricSettings{Enabled: true},
		ProcessRuntimeMemstatsTotalAlloc:    MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordProcessRuntimeMemstatsBuckHashSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsFreesDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsGcCPUFractionDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsGcSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapAllocDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapIdleDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapObjectsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapReleasedDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsLastPauseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsLookupsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMallocsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMcacheInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMcacheSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMspanInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMspanSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNextGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNumForcedGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNumGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsOtherSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsPauseTotalDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsStackInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsStackSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsTotalAllocDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "process.runtime.memstats.buck_hash_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of memory in profiling bucket hash tables.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.buck_hash_sys"] = struct{}{}
		case "process.runtime.memstats.frees":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cumulative count of heap objects freed.", ms.At(i).Description())
			assert.Equal(t, "{objects}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.frees"] = struct{}{}
		case "process.runtime.memstats.gc_cpu_fraction":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The fraction of this program's available CPU time used by the GC since the program started.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["process.runtime.memstats.gc_cpu_fraction"] = struct{}{}
		case "process.runtime.memstats.gc_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of memory in garbage collection metadata.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.gc_sys"] = struct{}{}
		case "process.runtime.memstats.heap_alloc":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of allocated heap objects.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_alloc"] = struct{}{}
		case "process.runtime.memstats.heap_idle":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes in idle (unused) spans.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_idle"] = struct{}{}
		case "process.runtime.memstats.heap_inuse":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes in in-use spans.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_inuse"] = struct{}{}
		case "process.runtime.memstats.heap_objects":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of allocated heap objects.", ms.At(i).Description())
			assert.Equal(t, "{objects}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_objects"] = struct{}{}
		case "process.runtime.memstats.heap_released":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of physical memory returned to the OS.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_released"] = struct{}{}
		case "process.runtime.memstats.heap_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of heap memory obtained by the OS.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.heap_sys"] = struct{}{}
		case "process.runtime.memstats.last_pause":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The most recent stop-the-world pause time.", ms.At(i).Description())
			assert.Equal(t, "ns", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.last_pause"] = struct{}{}
		case "process.runtime.memstats.lookups":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of pointer lookups performed by the runtime.", ms.At(i).Description())
			assert.Equal(t, "{lookups}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.lookups"] = struct{}{}
		case "process.runtime.memstats.mallocs":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cumulative count of heap objects allocated.", ms.At(i).Description())
			assert.Equal(t, "{objects}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.mallocs"] = struct{}{}
		case "process.runtime.memstats.mcache_inuse":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of allocated mcache structures.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.mcache_inuse"] = struct{}{}
		case "process.runtime.memstats.mcache_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of memory obtained from the OS for mcache structures.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.mcache_sys"] = struct{}{}
		case "process.runtime.memstats.mspan_inuse":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of allocated mspan structures.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.mspan_inuse"] = struct{}{}
		case "process.runtime.memstats.mspan_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of memory obtained from the OS for mspan structures.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.mspan_sys"] = struct{}{}
		case "process.runtime.memstats.next_gc":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The target heap size of the next GC cycle.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.next_gc"] = struct{}{}
		case "process.runtime.memstats.num_forced_gc":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of GC cycles that were forced by the application calling the GC function.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.num_forced_gc"] = struct{}{}
		case "process.runtime.memstats.num_gc":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of completed GC cycles.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.num_gc"] = struct{}{}
		case "process.runtime.memstats.other_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of memory in miscellaneous off-heap runtime allocations.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.other_sys"] = struct{}{}
		case "process.runtime.memstats.pause_total":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The cumulative nanoseconds in GC stop-the-world pauses since the program started.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.pause_total"] = struct{}{}
		case "process.runtime.memstats.stack_inuse":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes in stack spans.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.stack_inuse"] = struct{}{}
		case "process.runtime.memstats.stack_sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes of stack memory obtained from the OS.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.stack_sys"] = struct{}{}
		case "process.runtime.memstats.sys":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total bytes of memory obtained from the OS.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.sys"] = struct{}{}
		case "process.runtime.memstats.total_alloc":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cumulative bytes allocated for heap objects.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.runtime.memstats.total_alloc"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		ProcessRuntimeMemstatsBuckHashSys:   MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsFrees:         MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsGcCPUFraction: MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsGcSys:         MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapAlloc:     MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapIdle:      MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapInuse:     MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapObjects:   MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapReleased:  MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsHeapSys:       MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsLastPause:     MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsLookups:       MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsMallocs:       MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsMcacheInuse:   MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsMcacheSys:     MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsMspanInuse:    MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsMspanSys:      MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsNextGc:        MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsNumForcedGc:   MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsNumGc:         MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsOtherSys:      MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsPauseTotal:    MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsStackInuse:    MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsStackSys:      MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsSys:           MetricSettings{Enabled: false},
		ProcessRuntimeMemstatsTotalAlloc:    MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordProcessRuntimeMemstatsBuckHashSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsFreesDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsGcCPUFractionDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsGcSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapAllocDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapIdleDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapObjectsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapReleasedDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsHeapSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsLastPauseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsLookupsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMallocsDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMcacheInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMcacheSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMspanInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsMspanSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNextGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNumForcedGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsNumGcDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsOtherSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsPauseTotalDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsStackInuseDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsStackSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsSysDataPoint(ts, 1)
	mb.RecordProcessRuntimeMemstatsTotalAllocDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
