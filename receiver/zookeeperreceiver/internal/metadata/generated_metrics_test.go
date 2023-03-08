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
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperConnectionActiveDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperDataTreeEphemeralNodeCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperDataTreeSizeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperFileDescriptorLimitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperFileDescriptorOpenDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperFollowerCountDataPoint(ts, 1, AttributeState(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperFsyncExceededThresholdCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperLatencyAvgDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperLatencyMaxDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperLatencyMinDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperPacketCountDataPoint(ts, 1, AttributeDirection(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperRequestActiveDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperSyncPendingDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperWatchCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordZookeeperZnodeCountDataPoint(ts, 1)

			metrics := mb.Emit(WithServerState("attr-val"), WithZkVersion("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("server.state")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ServerState.Enabled, ok)
			if mb.resourceAttributesSettings.ServerState.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("zk.version")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.ZkVersion.Enabled, ok)
			if mb.resourceAttributesSettings.ZkVersion.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 2)

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
				case "zookeeper.connection.active":
					assert.False(t, validatedMetrics["zookeeper.connection.active"], "Found a duplicate in the metrics slice: zookeeper.connection.active")
					validatedMetrics["zookeeper.connection.active"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of active clients connected to a ZooKeeper server.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.data_tree.ephemeral_node.count":
					assert.False(t, validatedMetrics["zookeeper.data_tree.ephemeral_node.count"], "Found a duplicate in the metrics slice: zookeeper.data_tree.ephemeral_node.count")
					validatedMetrics["zookeeper.data_tree.ephemeral_node.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of ephemeral nodes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
					assert.Equal(t, "{nodes}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.data_tree.size":
					assert.False(t, validatedMetrics["zookeeper.data_tree.size"], "Found a duplicate in the metrics slice: zookeeper.data_tree.size")
					validatedMetrics["zookeeper.data_tree.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Size of data in bytes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.file_descriptor.limit":
					assert.False(t, validatedMetrics["zookeeper.file_descriptor.limit"], "Found a duplicate in the metrics slice: zookeeper.file_descriptor.limit")
					validatedMetrics["zookeeper.file_descriptor.limit"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum number of file descriptors that a ZooKeeper server can open.", ms.At(i).Description())
					assert.Equal(t, "{file_descriptors}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.file_descriptor.open":
					assert.False(t, validatedMetrics["zookeeper.file_descriptor.open"], "Found a duplicate in the metrics slice: zookeeper.file_descriptor.open")
					validatedMetrics["zookeeper.file_descriptor.open"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of file descriptors that a ZooKeeper server has open.", ms.At(i).Description())
					assert.Equal(t, "{file_descriptors}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.follower.count":
					assert.False(t, validatedMetrics["zookeeper.follower.count"], "Found a duplicate in the metrics slice: zookeeper.follower.count")
					validatedMetrics["zookeeper.follower.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of followers. Only exposed by the leader.", ms.At(i).Description())
					assert.Equal(t, "{followers}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "synced", attrVal.Str())
				case "zookeeper.fsync.exceeded_threshold.count":
					assert.False(t, validatedMetrics["zookeeper.fsync.exceeded_threshold.count"], "Found a duplicate in the metrics slice: zookeeper.fsync.exceeded_threshold.count")
					validatedMetrics["zookeeper.fsync.exceeded_threshold.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times fsync duration has exceeded warning threshold.", ms.At(i).Description())
					assert.Equal(t, "{events}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.latency.avg":
					assert.False(t, validatedMetrics["zookeeper.latency.avg"], "Found a duplicate in the metrics slice: zookeeper.latency.avg")
					validatedMetrics["zookeeper.latency.avg"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average time in milliseconds for requests to be processed.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.latency.max":
					assert.False(t, validatedMetrics["zookeeper.latency.max"], "Found a duplicate in the metrics slice: zookeeper.latency.max")
					validatedMetrics["zookeeper.latency.max"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Maximum time in milliseconds for requests to be processed.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.latency.min":
					assert.False(t, validatedMetrics["zookeeper.latency.min"], "Found a duplicate in the metrics slice: zookeeper.latency.min")
					validatedMetrics["zookeeper.latency.min"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Minimum time in milliseconds for requests to be processed.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.packet.count":
					assert.False(t, validatedMetrics["zookeeper.packet.count"], "Found a duplicate in the metrics slice: zookeeper.packet.count")
					validatedMetrics["zookeeper.packet.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of ZooKeeper packets received or sent by a server.", ms.At(i).Description())
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
					assert.Equal(t, "received", attrVal.Str())
				case "zookeeper.request.active":
					assert.False(t, validatedMetrics["zookeeper.request.active"], "Found a duplicate in the metrics slice: zookeeper.request.active")
					validatedMetrics["zookeeper.request.active"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of currently executing requests.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.sync.pending":
					assert.False(t, validatedMetrics["zookeeper.sync.pending"], "Found a duplicate in the metrics slice: zookeeper.sync.pending")
					validatedMetrics["zookeeper.sync.pending"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of pending syncs from the followers. Only exposed by the leader.", ms.At(i).Description())
					assert.Equal(t, "{syncs}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.watch.count":
					assert.False(t, validatedMetrics["zookeeper.watch.count"], "Found a duplicate in the metrics slice: zookeeper.watch.count")
					validatedMetrics["zookeeper.watch.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of watches placed on Z-Nodes on a ZooKeeper server.", ms.At(i).Description())
					assert.Equal(t, "{watches}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "zookeeper.znode.count":
					assert.False(t, validatedMetrics["zookeeper.znode.count"], "Found a duplicate in the metrics slice: zookeeper.znode.count")
					validatedMetrics["zookeeper.znode.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of z-nodes that a ZooKeeper server has in its data tree.", ms.At(i).Description())
					assert.Equal(t, "{znodes}", ms.At(i).Unit())
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

func loadConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
