// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testMetricsSet int

const (
	testMetricsSetDefault testMetricsSet = iota
	testMetricsSetAll
	testMetricsSetNo
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name       string
		metricsSet testMetricsSet
	}{
		{
			name:       "default",
			metricsSet: testMetricsSetDefault,
		},
		{
			name:       "all_metrics",
			metricsSet: testMetricsSetAll,
		},
		{
			name:       "no_metrics",
			metricsSet: testMetricsSetNo,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			allMetricsCount++
			mb.RecordProcessContextSwitchesDataPoint(ts, 1, AttributeContextSwitchType(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessCPUTimeDataPoint(ts, 1, AttributeState(1))

			allMetricsCount++
			mb.RecordProcessCPUUtilizationDataPoint(ts, 1, AttributeState(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessDiskIoDataPoint(ts, 1, AttributeDirection(1))

			allMetricsCount++
			mb.RecordProcessDiskOperationsDataPoint(ts, 1, AttributeDirection(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessMemoryUsageDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessMemoryUtilizationDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordProcessMemoryVirtualDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessOpenFileDescriptorsDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessPagingFaultsDataPoint(ts, 1, AttributePagingFaultType(1))

			allMetricsCount++
			mb.RecordProcessSignalsPendingDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordProcessThreadsDataPoint(ts, 1)

			metrics := mb.Emit(WithProcessCommand("attr-val"), WithProcessCommandLine("attr-val"), WithProcessExecutableName("attr-val"), WithProcessExecutablePath("attr-val"), WithProcessOwner("attr-val"), WithProcessParentPid(1), WithProcessPid(1))

			if test.metricsSet == testMetricsSetNo {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("process.command")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessCommand.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessCommand.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.command_line")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessCommandLine.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessCommandLine.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.executable.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessExecutableName.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessExecutableName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.executable.path")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessExecutablePath.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessExecutablePath.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.owner")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessOwner.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessOwner.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.parent_pid")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessParentPid.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessParentPid.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, 1, attrVal.Int())
			}
			attrVal, ok = rm.Resource().Attributes().Get("process.pid")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ProcessPid.Enabled, ok)
			if mb.resourceAttributesSettings.ProcessPid.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, 1, attrVal.Int())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 7)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testMetricsSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testMetricsSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "process.context_switches":
					assert.False(t, validatedMetrics["process.context_switches"], "Found a duplicate in the metrics slice: process.context_switches")
					validatedMetrics["process.context_switches"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the process has been context switched.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "involuntary", attrVal.Str())
				case "process.cpu.time":
					assert.False(t, validatedMetrics["process.cpu.time"], "Found a duplicate in the metrics slice: process.cpu.time")
					validatedMetrics["process.cpu.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total CPU seconds broken down by different states.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "system", attrVal.Str())
				case "process.cpu.utilization":
					assert.False(t, validatedMetrics["process.cpu.utilization"], "Found a duplicate in the metrics slice: process.cpu.utilization")
					validatedMetrics["process.cpu.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of total CPU time used by the process since last scrape, expressed as a value between 0 and 1. On the first scrape, no data point is emitted for this metric.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "system", attrVal.Str())
				case "process.disk.io":
					assert.False(t, validatedMetrics["process.disk.io"], "Found a duplicate in the metrics slice: process.disk.io")
					validatedMetrics["process.disk.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Disk bytes transferred.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "read", attrVal.Str())
				case "process.disk.operations":
					assert.False(t, validatedMetrics["process.disk.operations"], "Found a duplicate in the metrics slice: process.disk.operations")
					validatedMetrics["process.disk.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of disk operations performed by the process.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.Equal(t, "read", attrVal.Str())
				case "process.memory.usage":
					assert.False(t, validatedMetrics["process.memory.usage"], "Found a duplicate in the metrics slice: process.memory.usage")
					validatedMetrics["process.memory.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of physical memory in use.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.memory.utilization":
					assert.False(t, validatedMetrics["process.memory.utilization"], "Found a duplicate in the metrics slice: process.memory.utilization")
					validatedMetrics["process.memory.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Percentage of total physical memory that is used by the process.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "process.memory.virtual":
					assert.False(t, validatedMetrics["process.memory.virtual"], "Found a duplicate in the metrics slice: process.memory.virtual")
					validatedMetrics["process.memory.virtual"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Virtual memory size.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.open_file_descriptors":
					assert.False(t, validatedMetrics["process.open_file_descriptors"], "Found a duplicate in the metrics slice: process.open_file_descriptors")
					validatedMetrics["process.open_file_descriptors"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of file descriptors in use by the process.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.paging.faults":
					assert.False(t, validatedMetrics["process.paging.faults"], "Found a duplicate in the metrics slice: process.paging.faults")
					validatedMetrics["process.paging.faults"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of page faults the process has made. This metric is only available on Linux.", ms.At(i).Description())
					assert.Equal(t, "{faults}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "major", attrVal.Str())
				case "process.signals_pending":
					assert.False(t, validatedMetrics["process.signals_pending"], "Found a duplicate in the metrics slice: process.signals_pending")
					validatedMetrics["process.signals_pending"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of pending signals for the process. This metric is only available on Linux.", ms.At(i).Description())
					assert.Equal(t, "{signals}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "process.threads":
					assert.False(t, validatedMetrics["process.threads"], "Found a duplicate in the metrics slice: process.threads")
					validatedMetrics["process.threads"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Process threads count.", ms.At(i).Description())
					assert.Equal(t, "{threads}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}

func loadConfig(t *testing.T, name string) MetricsSettings {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsSettings()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
