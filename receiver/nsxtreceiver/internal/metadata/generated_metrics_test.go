// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeCPUUtilizationDataPoint(ts, 1, AttributeClassDatapath)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeFilesystemUsageDataPoint(ts, 1, AttributeDiskStateUsed)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeFilesystemUtilizationDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeMemoryCacheUsageDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeMemoryUsageDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeNetworkIoDataPoint(ts, 1, AttributeDirectionReceived)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNsxtNodeNetworkPacketCountDataPoint(ts, 1, AttributeDirectionReceived, AttributePacketTypeDropped)

			rb := mb.NewResourceBuilder()
			rb.SetDeviceID("device.id-val")
			rb.SetNsxtNodeID("nsxt.node.id-val")
			rb.SetNsxtNodeName("nsxt.node.name-val")
			rb.SetNsxtNodeType("nsxt.node.type-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "nsxt.node.cpu.utilization":
					assert.False(t, validatedMetrics["nsxt.node.cpu.utilization"], "Found a duplicate in the metrics slice: nsxt.node.cpu.utilization")
					validatedMetrics["nsxt.node.cpu.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The average amount of CPU being used by the node.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("class")
					assert.True(t, ok)
					assert.EqualValues(t, "datapath", attrVal.Str())
				case "nsxt.node.filesystem.usage":
					assert.False(t, validatedMetrics["nsxt.node.filesystem.usage"], "Found a duplicate in the metrics slice: nsxt.node.filesystem.usage")
					validatedMetrics["nsxt.node.filesystem.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of storage space used by the node.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.EqualValues(t, "used", attrVal.Str())
				case "nsxt.node.filesystem.utilization":
					assert.False(t, validatedMetrics["nsxt.node.filesystem.utilization"], "Found a duplicate in the metrics slice: nsxt.node.filesystem.utilization")
					validatedMetrics["nsxt.node.filesystem.utilization"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The percentage of storage space utilized.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "nsxt.node.memory.cache.usage":
					assert.False(t, validatedMetrics["nsxt.node.memory.cache.usage"], "Found a duplicate in the metrics slice: nsxt.node.memory.cache.usage")
					validatedMetrics["nsxt.node.memory.cache.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The size of the node's memory cache.", ms.At(i).Description())
					assert.Equal(t, "KBy", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "nsxt.node.memory.usage":
					assert.False(t, validatedMetrics["nsxt.node.memory.usage"], "Found a duplicate in the metrics slice: nsxt.node.memory.usage")
					validatedMetrics["nsxt.node.memory.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The memory usage of the node.", ms.At(i).Description())
					assert.Equal(t, "KBy", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "nsxt.node.network.io":
					assert.False(t, validatedMetrics["nsxt.node.network.io"], "Found a duplicate in the metrics slice: nsxt.node.network.io")
					validatedMetrics["nsxt.node.network.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of bytes which have flowed through the network interface.", ms.At(i).Description())
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
					assert.EqualValues(t, "received", attrVal.Str())
				case "nsxt.node.network.packet.count":
					assert.False(t, validatedMetrics["nsxt.node.network.packet.count"], "Found a duplicate in the metrics slice: nsxt.node.network.packet.count")
					validatedMetrics["nsxt.node.network.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of packets which have flowed through the network interface on the node.", ms.At(i).Description())
					assert.Equal(t, "{packets}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "received", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "dropped", attrVal.Str())
				}
			}
		})
	}
}
