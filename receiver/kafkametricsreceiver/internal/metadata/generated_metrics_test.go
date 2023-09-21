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
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.consumer_fetch_count`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.consumer_fetch_rate`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.count`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.incoming_byte_rate`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.outgoing_byte_rate`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.request_latency`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.request_rate`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.request_size`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.requests_in_flight`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.response_rate`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			if test.configSet == testSetDefault {
				assert.Equal(t, "[WARNING] Please set `enabled` field explicitly for `messaging.kafka.broker.response_size`: This metric will be enabled by default in the next versions.", observedLogs.All()[expectedWarnings].Message)
				expectedWarnings++
			}
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			allMetricsCount++
			mb.RecordKafkaBrokersDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupLagDataPoint(ts, 1, "group-val", "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupLagSumDataPoint(ts, 1, "group-val", "topic-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupMembersDataPoint(ts, 1, "group-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupOffsetDataPoint(ts, 1, "group-val", "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaConsumerGroupOffsetSumDataPoint(ts, 1, "group-val", "topic-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionCurrentOffsetDataPoint(ts, 1, "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionOldestOffsetDataPoint(ts, 1, "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionReplicasDataPoint(ts, 1, "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaPartitionReplicasInSyncDataPoint(ts, 1, "topic-val", 9)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordKafkaTopicPartitionsDataPoint(ts, 1, "topic-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerConsumerFetchCountDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerConsumerFetchRateDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerIncomingByteRateDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerOutgoingByteRateDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerRequestLatencyDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerRequestRateDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerRequestSizeDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerRequestsInFlightDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerResponseRateDataPoint(ts, 1, 6)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMessagingKafkaBrokerResponseSizeDataPoint(ts, 1, 6)

			res := pcommon.NewResource()
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
				case "kafka.brokers":
					assert.False(t, validatedMetrics["kafka.brokers"], "Found a duplicate in the metrics slice: kafka.brokers")
					validatedMetrics["kafka.brokers"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "[DEPRACATED] Number of brokers in the cluster.", ms.At(i).Description())
					assert.Equal(t, "{broker}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "kafka.consumer_group.lag":
					assert.False(t, validatedMetrics["kafka.consumer_group.lag"], "Found a duplicate in the metrics slice: kafka.consumer_group.lag")
					validatedMetrics["kafka.consumer_group.lag"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current approximate lag of consumer group at partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "group-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.consumer_group.lag_sum":
					assert.False(t, validatedMetrics["kafka.consumer_group.lag_sum"], "Found a duplicate in the metrics slice: kafka.consumer_group.lag_sum")
					validatedMetrics["kafka.consumer_group.lag_sum"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current approximate sum of consumer group lag across all partitions of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "group-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
				case "kafka.consumer_group.members":
					assert.False(t, validatedMetrics["kafka.consumer_group.members"], "Found a duplicate in the metrics slice: kafka.consumer_group.members")
					validatedMetrics["kafka.consumer_group.members"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Count of members in the consumer group", ms.At(i).Description())
					assert.Equal(t, "{members}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "group-val", attrVal.Str())
				case "kafka.consumer_group.offset":
					assert.False(t, validatedMetrics["kafka.consumer_group.offset"], "Found a duplicate in the metrics slice: kafka.consumer_group.offset")
					validatedMetrics["kafka.consumer_group.offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current offset of the consumer group at partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "group-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.consumer_group.offset_sum":
					assert.False(t, validatedMetrics["kafka.consumer_group.offset_sum"], "Found a duplicate in the metrics slice: kafka.consumer_group.offset_sum")
					validatedMetrics["kafka.consumer_group.offset_sum"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Sum of consumer group offset across partitions of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("group")
					assert.True(t, ok)
					assert.EqualValues(t, "group-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
				case "kafka.partition.current_offset":
					assert.False(t, validatedMetrics["kafka.partition.current_offset"], "Found a duplicate in the metrics slice: kafka.partition.current_offset")
					validatedMetrics["kafka.partition.current_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Current offset of partition of topic.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.partition.oldest_offset":
					assert.False(t, validatedMetrics["kafka.partition.oldest_offset"], "Found a duplicate in the metrics slice: kafka.partition.oldest_offset")
					validatedMetrics["kafka.partition.oldest_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Oldest offset of partition of topic", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.partition.replicas":
					assert.False(t, validatedMetrics["kafka.partition.replicas"], "Found a duplicate in the metrics slice: kafka.partition.replicas")
					validatedMetrics["kafka.partition.replicas"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of replicas for partition of topic", ms.At(i).Description())
					assert.Equal(t, "{replicas}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.partition.replicas_in_sync":
					assert.False(t, validatedMetrics["kafka.partition.replicas_in_sync"], "Found a duplicate in the metrics slice: kafka.partition.replicas_in_sync")
					validatedMetrics["kafka.partition.replicas_in_sync"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of synchronized replicas of partition", ms.At(i).Description())
					assert.Equal(t, "{replicas}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("partition")
					assert.True(t, ok)
					assert.EqualValues(t, 9, attrVal.Int())
				case "kafka.topic.partitions":
					assert.False(t, validatedMetrics["kafka.topic.partitions"], "Found a duplicate in the metrics slice: kafka.topic.partitions")
					validatedMetrics["kafka.topic.partitions"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of partitions in topic.", ms.At(i).Description())
					assert.Equal(t, "{partitions}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("topic")
					assert.True(t, ok)
					assert.EqualValues(t, "topic-val", attrVal.Str())
				case "messaging.kafka.broker.consumer_fetch_count":
					assert.False(t, validatedMetrics["messaging.kafka.broker.consumer_fetch_count"], "Found a duplicate in the metrics slice: messaging.kafka.broker.consumer_fetch_count")
					validatedMetrics["messaging.kafka.broker.consumer_fetch_count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Count of consumer fetches", ms.At(i).Description())
					assert.Equal(t, "{fetches}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.consumer_fetch_rate":
					assert.False(t, validatedMetrics["messaging.kafka.broker.consumer_fetch_rate"], "Found a duplicate in the metrics slice: messaging.kafka.broker.consumer_fetch_rate")
					validatedMetrics["messaging.kafka.broker.consumer_fetch_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average consumer fetch Rate", ms.At(i).Description())
					assert.Equal(t, "{fetches}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.count":
					assert.False(t, validatedMetrics["messaging.kafka.broker.count"], "Found a duplicate in the metrics slice: messaging.kafka.broker.count")
					validatedMetrics["messaging.kafka.broker.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of brokers in the cluster.", ms.At(i).Description())
					assert.Equal(t, "{broker}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "messaging.kafka.broker.incoming_byte_rate":
					assert.False(t, validatedMetrics["messaging.kafka.broker.incoming_byte_rate"], "Found a duplicate in the metrics slice: messaging.kafka.broker.incoming_byte_rate")
					validatedMetrics["messaging.kafka.broker.incoming_byte_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average incoming Byte Rate in bytes/second", ms.At(i).Description())
					assert.Equal(t, "By/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.outgoing_byte_rate":
					assert.False(t, validatedMetrics["messaging.kafka.broker.outgoing_byte_rate"], "Found a duplicate in the metrics slice: messaging.kafka.broker.outgoing_byte_rate")
					validatedMetrics["messaging.kafka.broker.outgoing_byte_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average outgoing Byte Rate in bytes/second.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.request_latency":
					assert.False(t, validatedMetrics["messaging.kafka.broker.request_latency"], "Found a duplicate in the metrics slice: messaging.kafka.broker.request_latency")
					validatedMetrics["messaging.kafka.broker.request_latency"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average request latency in seconds", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.request_rate":
					assert.False(t, validatedMetrics["messaging.kafka.broker.request_rate"], "Found a duplicate in the metrics slice: messaging.kafka.broker.request_rate")
					validatedMetrics["messaging.kafka.broker.request_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average request rate per second.", ms.At(i).Description())
					assert.Equal(t, "{requests}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.request_size":
					assert.False(t, validatedMetrics["messaging.kafka.broker.request_size"], "Found a duplicate in the metrics slice: messaging.kafka.broker.request_size")
					validatedMetrics["messaging.kafka.broker.request_size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average request size in bytes", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.requests_in_flight":
					assert.False(t, validatedMetrics["messaging.kafka.broker.requests_in_flight"], "Found a duplicate in the metrics slice: messaging.kafka.broker.requests_in_flight")
					validatedMetrics["messaging.kafka.broker.requests_in_flight"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Requests in flight", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.response_rate":
					assert.False(t, validatedMetrics["messaging.kafka.broker.response_rate"], "Found a duplicate in the metrics slice: messaging.kafka.broker.response_rate")
					validatedMetrics["messaging.kafka.broker.response_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average response rate per second", ms.At(i).Description())
					assert.Equal(t, "{response}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				case "messaging.kafka.broker.response_size":
					assert.False(t, validatedMetrics["messaging.kafka.broker.response_size"], "Found a duplicate in the metrics slice: messaging.kafka.broker.response_size")
					validatedMetrics["messaging.kafka.broker.response_size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average response size in bytes", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("broker")
					assert.True(t, ok)
					assert.EqualValues(t, 6, attrVal.Int())
				}
			}
		})
	}
}
