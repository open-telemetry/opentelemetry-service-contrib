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

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBackendsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterBuffersAllocatedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterBuffersWritesDataPoint(ts, 1, AttributeBgBufferSourceBackend)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterCheckpointCountDataPoint(ts, 1, AttributeBgCheckpointTypeRequested)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterDurationDataPoint(ts, 1, AttributeBgDurationTypeSync)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterMaxwrittenDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBlocksReadDataPoint(ts, 1, AttributeSourceHeapRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlCommitsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlConnectionMaxDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlDatabaseCountDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordPostgresqlDatabaseLocksDataPoint(ts, 1, "relation-val", "mode-val", "lock_type-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlDbSizeDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordPostgresqlDeadlocksDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlIndexScansDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlIndexSizeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlOperationsDataPoint(ts, 1, AttributeOperationIns)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlReplicationDataDelayDataPoint(ts, 1, "replication_client-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlRollbacksDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlRowsDataPoint(ts, 1, AttributeStateDead)

			allMetricsCount++
			mb.RecordPostgresqlSequentialScansDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableSizeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableVacuumCountDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordPostgresqlTempFilesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlWalAgeDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordPostgresqlWalDelayDataPoint(ts, 1, AttributeWalOperationLagFlush, "replication_client-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlWalLagDataPoint(ts, 1, AttributeWalOperationLagFlush, "replication_client-val")

			rb := mb.NewResourceBuilder()
			rb.SetPostgresqlDatabaseName("postgresql.database.name-val")
			rb.SetPostgresqlIndexName("postgresql.index.name-val")
			rb.SetPostgresqlSchemaName("postgresql.schema.name-val")
			rb.SetPostgresqlTableName("postgresql.table.name-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if test.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "postgresql.backends":
					assert.False(t, validatedMetrics["postgresql.backends"], "Found a duplicate in the metrics slice: postgresql.backends")
					validatedMetrics["postgresql.backends"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of backends.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.bgwriter.buffers.allocated":
					assert.False(t, validatedMetrics["postgresql.bgwriter.buffers.allocated"], "Found a duplicate in the metrics slice: postgresql.bgwriter.buffers.allocated")
					validatedMetrics["postgresql.bgwriter.buffers.allocated"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of buffers allocated.", ms.At(i).Description())
					assert.Equal(t, "{buffers}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.bgwriter.buffers.writes":
					assert.False(t, validatedMetrics["postgresql.bgwriter.buffers.writes"], "Found a duplicate in the metrics slice: postgresql.bgwriter.buffers.writes")
					validatedMetrics["postgresql.bgwriter.buffers.writes"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of buffers written.", ms.At(i).Description())
					assert.Equal(t, "{buffers}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("source")
					assert.True(t, ok)
					assert.EqualValues(t, "backend", attrVal.Str())
				case "postgresql.bgwriter.checkpoint.count":
					assert.False(t, validatedMetrics["postgresql.bgwriter.checkpoint.count"], "Found a duplicate in the metrics slice: postgresql.bgwriter.checkpoint.count")
					validatedMetrics["postgresql.bgwriter.checkpoint.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of checkpoints performed.", ms.At(i).Description())
					assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "requested", attrVal.Str())
				case "postgresql.bgwriter.duration":
					assert.False(t, validatedMetrics["postgresql.bgwriter.duration"], "Found a duplicate in the metrics slice: postgresql.bgwriter.duration")
					validatedMetrics["postgresql.bgwriter.duration"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total time spent writing and syncing files to disk by checkpoints.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "sync", attrVal.Str())
				case "postgresql.bgwriter.maxwritten":
					assert.False(t, validatedMetrics["postgresql.bgwriter.maxwritten"], "Found a duplicate in the metrics slice: postgresql.bgwriter.maxwritten")
					validatedMetrics["postgresql.bgwriter.maxwritten"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the background writer stopped a cleaning scan because it had written too many buffers.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.blocks_read":
					assert.False(t, validatedMetrics["postgresql.blocks_read"], "Found a duplicate in the metrics slice: postgresql.blocks_read")
					validatedMetrics["postgresql.blocks_read"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of blocks read.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("source")
					assert.True(t, ok)
					assert.EqualValues(t, "heap_read", attrVal.Str())
				case "postgresql.commits":
					assert.False(t, validatedMetrics["postgresql.commits"], "Found a duplicate in the metrics slice: postgresql.commits")
					validatedMetrics["postgresql.commits"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of commits.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.connection.max":
					assert.False(t, validatedMetrics["postgresql.connection.max"], "Found a duplicate in the metrics slice: postgresql.connection.max")
					validatedMetrics["postgresql.connection.max"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Configured maximum number of client connections allowed", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.database.count":
					assert.False(t, validatedMetrics["postgresql.database.count"], "Found a duplicate in the metrics slice: postgresql.database.count")
					validatedMetrics["postgresql.database.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of user databases.", ms.At(i).Description())
					assert.Equal(t, "{databases}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.database.locks":
					assert.False(t, validatedMetrics["postgresql.database.locks"], "Found a duplicate in the metrics slice: postgresql.database.locks")
					validatedMetrics["postgresql.database.locks"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of database locks.", ms.At(i).Description())
					assert.Equal(t, "{lock}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("relation")
					assert.True(t, ok)
					assert.EqualValues(t, "relation-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("mode")
					assert.True(t, ok)
					assert.EqualValues(t, "mode-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_type")
					assert.True(t, ok)
					assert.EqualValues(t, "lock_type-val", attrVal.Str())
				case "postgresql.db_size":
					assert.False(t, validatedMetrics["postgresql.db_size"], "Found a duplicate in the metrics slice: postgresql.db_size")
					validatedMetrics["postgresql.db_size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The database disk usage.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.deadlocks":
					assert.False(t, validatedMetrics["postgresql.deadlocks"], "Found a duplicate in the metrics slice: postgresql.deadlocks")
					validatedMetrics["postgresql.deadlocks"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of deadlocks.", ms.At(i).Description())
					assert.Equal(t, "{deadlock}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.index.scans":
					assert.False(t, validatedMetrics["postgresql.index.scans"], "Found a duplicate in the metrics slice: postgresql.index.scans")
					validatedMetrics["postgresql.index.scans"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of index scans on a table.", ms.At(i).Description())
					assert.Equal(t, "{scans}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.index.size":
					assert.False(t, validatedMetrics["postgresql.index.size"], "Found a duplicate in the metrics slice: postgresql.index.size")
					validatedMetrics["postgresql.index.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The size of the index on disk.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.operations":
					assert.False(t, validatedMetrics["postgresql.operations"], "Found a duplicate in the metrics slice: postgresql.operations")
					validatedMetrics["postgresql.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of db row operations.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "ins", attrVal.Str())
				case "postgresql.replication.data_delay":
					assert.False(t, validatedMetrics["postgresql.replication.data_delay"], "Found a duplicate in the metrics slice: postgresql.replication.data_delay")
					validatedMetrics["postgresql.replication.data_delay"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of data delayed in replication.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("replication_client")
					assert.True(t, ok)
					assert.EqualValues(t, "replication_client-val", attrVal.Str())
				case "postgresql.rollbacks":
					assert.False(t, validatedMetrics["postgresql.rollbacks"], "Found a duplicate in the metrics slice: postgresql.rollbacks")
					validatedMetrics["postgresql.rollbacks"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of rollbacks.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.rows":
					assert.False(t, validatedMetrics["postgresql.rows"], "Found a duplicate in the metrics slice: postgresql.rows")
					validatedMetrics["postgresql.rows"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of rows in the database.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.EqualValues(t, "dead", attrVal.Str())
				case "postgresql.sequential_scans":
					assert.False(t, validatedMetrics["postgresql.sequential_scans"], "Found a duplicate in the metrics slice: postgresql.sequential_scans")
					validatedMetrics["postgresql.sequential_scans"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of sequential scans.", ms.At(i).Description())
					assert.Equal(t, "{sequential_scan}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.table.count":
					assert.False(t, validatedMetrics["postgresql.table.count"], "Found a duplicate in the metrics slice: postgresql.table.count")
					validatedMetrics["postgresql.table.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of user tables in a database.", ms.At(i).Description())
					assert.Equal(t, "{table}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.table.size":
					assert.False(t, validatedMetrics["postgresql.table.size"], "Found a duplicate in the metrics slice: postgresql.table.size")
					validatedMetrics["postgresql.table.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Disk space used by a table.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.table.vacuum.count":
					assert.False(t, validatedMetrics["postgresql.table.vacuum.count"], "Found a duplicate in the metrics slice: postgresql.table.vacuum.count")
					validatedMetrics["postgresql.table.vacuum.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times a table has manually been vacuumed.", ms.At(i).Description())
					assert.Equal(t, "{vacuums}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.temp_files":
					assert.False(t, validatedMetrics["postgresql.temp_files"], "Found a duplicate in the metrics slice: postgresql.temp_files")
					validatedMetrics["postgresql.temp_files"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of temp files.", ms.At(i).Description())
					assert.Equal(t, "{temp_file}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.wal.age":
					assert.False(t, validatedMetrics["postgresql.wal.age"], "Found a duplicate in the metrics slice: postgresql.wal.age")
					validatedMetrics["postgresql.wal.age"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Age of the oldest WAL file.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.wal.delay":
					assert.False(t, validatedMetrics["postgresql.wal.delay"], "Found a duplicate in the metrics slice: postgresql.wal.delay")
					validatedMetrics["postgresql.wal.delay"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time between flushing recent WAL locally and receiving notification that the standby server has completed an operation with it.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "flush", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("replication_client")
					assert.True(t, ok)
					assert.EqualValues(t, "replication_client-val", attrVal.Str())
				case "postgresql.wal.lag":
					assert.False(t, validatedMetrics["postgresql.wal.lag"], "Found a duplicate in the metrics slice: postgresql.wal.lag")
					validatedMetrics["postgresql.wal.lag"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time between flushing recent WAL locally and receiving notification that the standby server has completed an operation with it.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "flush", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("replication_client")
					assert.True(t, ok)
					assert.EqualValues(t, "replication_client-val", attrVal.Str())
				}
			}
		})
	}
}
