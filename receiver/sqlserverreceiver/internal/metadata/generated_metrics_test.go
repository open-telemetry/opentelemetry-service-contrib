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

	enabledMetrics["sqlserver.batch.request.rate"] = true
	mb.RecordSqlserverBatchRequestRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.batch.sql_compilation.rate"] = true
	mb.RecordSqlserverBatchSQLCompilationRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.batch.sql_recompilation.rate"] = true
	mb.RecordSqlserverBatchSQLRecompilationRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.lock.wait.rate"] = true
	mb.RecordSqlserverLockWaitRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.lock.wait_time.avg"] = true
	mb.RecordSqlserverLockWaitTimeAvgDataPoint(ts, 1)

	enabledMetrics["sqlserver.page.buffer_cache.hit_ratio"] = true
	mb.RecordSqlserverPageBufferCacheHitRatioDataPoint(ts, 1)

	enabledMetrics["sqlserver.page.checkpoint.flush.rate"] = true
	mb.RecordSqlserverPageCheckpointFlushRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.page.lazy_write.rate"] = true
	mb.RecordSqlserverPageLazyWriteRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.page.life_expectancy"] = true
	mb.RecordSqlserverPageLifeExpectancyDataPoint(ts, 1)

	enabledMetrics["sqlserver.page.operation.rate"] = true
	mb.RecordSqlserverPageOperationRateDataPoint(ts, 1, AttributePageOperations(1))

	enabledMetrics["sqlserver.page.split.rate"] = true
	mb.RecordSqlserverPageSplitRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction.rate"] = true
	mb.RecordSqlserverTransactionRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction.write.rate"] = true
	mb.RecordSqlserverTransactionWriteRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.flush.data.rate"] = true
	mb.RecordSqlserverTransactionLogFlushDataRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.flush.rate"] = true
	mb.RecordSqlserverTransactionLogFlushRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.flush.wait.rate"] = true
	mb.RecordSqlserverTransactionLogFlushWaitRateDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.growth.count"] = true
	mb.RecordSqlserverTransactionLogGrowthCountDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.shrink.count"] = true
	mb.RecordSqlserverTransactionLogShrinkCountDataPoint(ts, 1)

	enabledMetrics["sqlserver.transaction_log.usage"] = true
	mb.RecordSqlserverTransactionLogUsageDataPoint(ts, 1)

	enabledMetrics["sqlserver.user.connection.count"] = true
	mb.RecordSqlserverUserConnectionCountDataPoint(ts, 1)

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
		SqlserverBatchRequestRate:            MetricSettings{Enabled: true},
		SqlserverBatchSQLCompilationRate:     MetricSettings{Enabled: true},
		SqlserverBatchSQLRecompilationRate:   MetricSettings{Enabled: true},
		SqlserverLockWaitRate:                MetricSettings{Enabled: true},
		SqlserverLockWaitTimeAvg:             MetricSettings{Enabled: true},
		SqlserverPageBufferCacheHitRatio:     MetricSettings{Enabled: true},
		SqlserverPageCheckpointFlushRate:     MetricSettings{Enabled: true},
		SqlserverPageLazyWriteRate:           MetricSettings{Enabled: true},
		SqlserverPageLifeExpectancy:          MetricSettings{Enabled: true},
		SqlserverPageOperationRate:           MetricSettings{Enabled: true},
		SqlserverPageSplitRate:               MetricSettings{Enabled: true},
		SqlserverTransactionRate:             MetricSettings{Enabled: true},
		SqlserverTransactionWriteRate:        MetricSettings{Enabled: true},
		SqlserverTransactionLogFlushDataRate: MetricSettings{Enabled: true},
		SqlserverTransactionLogFlushRate:     MetricSettings{Enabled: true},
		SqlserverTransactionLogFlushWaitRate: MetricSettings{Enabled: true},
		SqlserverTransactionLogGrowthCount:   MetricSettings{Enabled: true},
		SqlserverTransactionLogShrinkCount:   MetricSettings{Enabled: true},
		SqlserverTransactionLogUsage:         MetricSettings{Enabled: true},
		SqlserverUserConnectionCount:         MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordSqlserverBatchRequestRateDataPoint(ts, 1)
	mb.RecordSqlserverBatchSQLCompilationRateDataPoint(ts, 1)
	mb.RecordSqlserverBatchSQLRecompilationRateDataPoint(ts, 1)
	mb.RecordSqlserverLockWaitRateDataPoint(ts, 1)
	mb.RecordSqlserverLockWaitTimeAvgDataPoint(ts, 1)
	mb.RecordSqlserverPageBufferCacheHitRatioDataPoint(ts, 1)
	mb.RecordSqlserverPageCheckpointFlushRateDataPoint(ts, 1)
	mb.RecordSqlserverPageLazyWriteRateDataPoint(ts, 1)
	mb.RecordSqlserverPageLifeExpectancyDataPoint(ts, 1)
	mb.RecordSqlserverPageOperationRateDataPoint(ts, 1, AttributePageOperations(1))
	mb.RecordSqlserverPageSplitRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionWriteRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushDataRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushWaitRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogGrowthCountDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogShrinkCountDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogUsageDataPoint(ts, 1)
	mb.RecordSqlserverUserConnectionCountDataPoint(ts, 1)

	metrics := mb.Emit(WithSqlserverDatabaseName("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("sqlserver.database.name")
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
		case "sqlserver.batch.request.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of batch requests received by SQL Server.", ms.At(i).Description())
			assert.Equal(t, "{requests}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.batch.request.rate"] = struct{}{}
		case "sqlserver.batch.sql_compilation.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of SQL compilations needed.", ms.At(i).Description())
			assert.Equal(t, "{compilations}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.batch.sql_compilation.rate"] = struct{}{}
		case "sqlserver.batch.sql_recompilation.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of SQL recompilations needed.", ms.At(i).Description())
			assert.Equal(t, "{compilations}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.batch.sql_recompilation.rate"] = struct{}{}
		case "sqlserver.lock.wait.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of lock requests resulting in a wait.", ms.At(i).Description())
			assert.Equal(t, "{requests}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.lock.wait.rate"] = struct{}{}
		case "sqlserver.lock.wait_time.avg":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Average wait time for all lock requests that had to wait.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.lock.wait_time.avg"] = struct{}{}
		case "sqlserver.page.buffer_cache.hit_ratio":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Pages found in the buffer pool without having to read from disk.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.page.buffer_cache.hit_ratio"] = struct{}{}
		case "sqlserver.page.checkpoint.flush.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of pages flushed by operations requiring dirty pages to be flushed.", ms.At(i).Description())
			assert.Equal(t, "{pages}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.page.checkpoint.flush.rate"] = struct{}{}
		case "sqlserver.page.lazy_write.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of lazy writes moving dirty pages to disk.", ms.At(i).Description())
			assert.Equal(t, "{writes}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.page.lazy_write.rate"] = struct{}{}
		case "sqlserver.page.life_expectancy":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Time a page will stay in the buffer pool.", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["sqlserver.page.life_expectancy"] = struct{}{}
		case "sqlserver.page.operation.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of physical database page operations issued.", ms.At(i).Description())
			assert.Equal(t, "{operations}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "read", attrVal.Str())
			validatedMetrics["sqlserver.page.operation.rate"] = struct{}{}
		case "sqlserver.page.split.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of pages split as a result of overflowing index pages.", ms.At(i).Description())
			assert.Equal(t, "{pages}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.page.split.rate"] = struct{}{}
		case "sqlserver.transaction.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of transactions started for the database (not including XTP-only transactions).", ms.At(i).Description())
			assert.Equal(t, "{transactions}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.transaction.rate"] = struct{}{}
		case "sqlserver.transaction.write.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of transactions that wrote to the database and committed.", ms.At(i).Description())
			assert.Equal(t, "{transactions}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.transaction.write.rate"] = struct{}{}
		case "sqlserver.transaction_log.flush.data.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Total number of log bytes flushed.", ms.At(i).Description())
			assert.Equal(t, "By/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.transaction_log.flush.data.rate"] = struct{}{}
		case "sqlserver.transaction_log.flush.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of log flushes.", ms.At(i).Description())
			assert.Equal(t, "{flushes}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.transaction_log.flush.rate"] = struct{}{}
		case "sqlserver.transaction_log.flush.wait.rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of commits waiting for a transaction log flush.", ms.At(i).Description())
			assert.Equal(t, "{commits}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["sqlserver.transaction_log.flush.wait.rate"] = struct{}{}
		case "sqlserver.transaction_log.growth.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of transaction log expansions for a database.", ms.At(i).Description())
			assert.Equal(t, "{growths}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["sqlserver.transaction_log.growth.count"] = struct{}{}
		case "sqlserver.transaction_log.shrink.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of transaction log shrinks for a database.", ms.At(i).Description())
			assert.Equal(t, "{shrinks}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["sqlserver.transaction_log.shrink.count"] = struct{}{}
		case "sqlserver.transaction_log.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Percent of transaction log space used.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["sqlserver.transaction_log.usage"] = struct{}{}
		case "sqlserver.user.connection.count":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of users connected to the SQL Server.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["sqlserver.user.connection.count"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		SqlserverBatchRequestRate:            MetricSettings{Enabled: false},
		SqlserverBatchSQLCompilationRate:     MetricSettings{Enabled: false},
		SqlserverBatchSQLRecompilationRate:   MetricSettings{Enabled: false},
		SqlserverLockWaitRate:                MetricSettings{Enabled: false},
		SqlserverLockWaitTimeAvg:             MetricSettings{Enabled: false},
		SqlserverPageBufferCacheHitRatio:     MetricSettings{Enabled: false},
		SqlserverPageCheckpointFlushRate:     MetricSettings{Enabled: false},
		SqlserverPageLazyWriteRate:           MetricSettings{Enabled: false},
		SqlserverPageLifeExpectancy:          MetricSettings{Enabled: false},
		SqlserverPageOperationRate:           MetricSettings{Enabled: false},
		SqlserverPageSplitRate:               MetricSettings{Enabled: false},
		SqlserverTransactionRate:             MetricSettings{Enabled: false},
		SqlserverTransactionWriteRate:        MetricSettings{Enabled: false},
		SqlserverTransactionLogFlushDataRate: MetricSettings{Enabled: false},
		SqlserverTransactionLogFlushRate:     MetricSettings{Enabled: false},
		SqlserverTransactionLogFlushWaitRate: MetricSettings{Enabled: false},
		SqlserverTransactionLogGrowthCount:   MetricSettings{Enabled: false},
		SqlserverTransactionLogShrinkCount:   MetricSettings{Enabled: false},
		SqlserverTransactionLogUsage:         MetricSettings{Enabled: false},
		SqlserverUserConnectionCount:         MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordSqlserverBatchRequestRateDataPoint(ts, 1)
	mb.RecordSqlserverBatchSQLCompilationRateDataPoint(ts, 1)
	mb.RecordSqlserverBatchSQLRecompilationRateDataPoint(ts, 1)
	mb.RecordSqlserverLockWaitRateDataPoint(ts, 1)
	mb.RecordSqlserverLockWaitTimeAvgDataPoint(ts, 1)
	mb.RecordSqlserverPageBufferCacheHitRatioDataPoint(ts, 1)
	mb.RecordSqlserverPageCheckpointFlushRateDataPoint(ts, 1)
	mb.RecordSqlserverPageLazyWriteRateDataPoint(ts, 1)
	mb.RecordSqlserverPageLifeExpectancyDataPoint(ts, 1)
	mb.RecordSqlserverPageOperationRateDataPoint(ts, 1, AttributePageOperations(1))
	mb.RecordSqlserverPageSplitRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionWriteRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushDataRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogFlushWaitRateDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogGrowthCountDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogShrinkCountDataPoint(ts, 1)
	mb.RecordSqlserverTransactionLogUsageDataPoint(ts, 1)
	mb.RecordSqlserverUserConnectionCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
