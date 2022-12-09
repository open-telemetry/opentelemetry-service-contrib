// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), componenttest.NewNopReceiverCreateSettings(), WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["couchdb.average_request_time"] = true
	mb.RecordCouchdbAverageRequestTimeDataPoint(ts, 1)

	enabledMetrics["couchdb.database.open"] = true
	mb.RecordCouchdbDatabaseOpenDataPoint(ts, 1)

	enabledMetrics["couchdb.database.operations"] = true
	mb.RecordCouchdbDatabaseOperationsDataPoint(ts, 1, AttributeOperation(1))

	enabledMetrics["couchdb.file_descriptor.open"] = true
	mb.RecordCouchdbFileDescriptorOpenDataPoint(ts, 1)

	enabledMetrics["couchdb.httpd.bulk_requests"] = true
	mb.RecordCouchdbHttpdBulkRequestsDataPoint(ts, 1)

	enabledMetrics["couchdb.httpd.requests"] = true
	mb.RecordCouchdbHttpdRequestsDataPoint(ts, 1, AttributeHTTPMethod(1))

	enabledMetrics["couchdb.httpd.responses"] = true
	mb.RecordCouchdbHttpdResponsesDataPoint(ts, 1, "attr-val")

	enabledMetrics["couchdb.httpd.views"] = true
	mb.RecordCouchdbHttpdViewsDataPoint(ts, 1, AttributeView(1))

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
		CouchdbAverageRequestTime: MetricSettings{Enabled: true},
		CouchdbDatabaseOpen:       MetricSettings{Enabled: true},
		CouchdbDatabaseOperations: MetricSettings{Enabled: true},
		CouchdbFileDescriptorOpen: MetricSettings{Enabled: true},
		CouchdbHttpdBulkRequests:  MetricSettings{Enabled: true},
		CouchdbHttpdRequests:      MetricSettings{Enabled: true},
		CouchdbHttpdResponses:     MetricSettings{Enabled: true},
		CouchdbHttpdViews:         MetricSettings{Enabled: true},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordCouchdbAverageRequestTimeDataPoint(ts, 1)
	mb.RecordCouchdbDatabaseOpenDataPoint(ts, 1)
	mb.RecordCouchdbDatabaseOperationsDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordCouchdbFileDescriptorOpenDataPoint(ts, 1)
	mb.RecordCouchdbHttpdBulkRequestsDataPoint(ts, 1)
	mb.RecordCouchdbHttpdRequestsDataPoint(ts, 1, AttributeHTTPMethod(1))
	mb.RecordCouchdbHttpdResponsesDataPoint(ts, 1, "attr-val")
	mb.RecordCouchdbHttpdViewsDataPoint(ts, 1, AttributeView(1))

	metrics := mb.Emit(WithCouchdbNodeName("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("couchdb.node.name")
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
		case "couchdb.average_request_time":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The average duration of a served request.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["couchdb.average_request_time"] = struct{}{}
		case "couchdb.database.open":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of open databases.", ms.At(i).Description())
			assert.Equal(t, "{databases}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["couchdb.database.open"] = struct{}{}
		case "couchdb.database.operations":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of database operations.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "writes", attrVal.Str())
			validatedMetrics["couchdb.database.operations"] = struct{}{}
		case "couchdb.file_descriptor.open":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of open file descriptors.", ms.At(i).Description())
			assert.Equal(t, "{files}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["couchdb.file_descriptor.open"] = struct{}{}
		case "couchdb.httpd.bulk_requests":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of bulk requests.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["couchdb.httpd.bulk_requests"] = struct{}{}
		case "couchdb.httpd.requests":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of HTTP requests by method.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("http.method")
			assert.True(t, ok)
			assert.Equal(t, "COPY", attrVal.Str())
			validatedMetrics["couchdb.httpd.requests"] = struct{}{}
		case "couchdb.httpd.responses":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of each HTTP status code.", ms.At(i).Description())
			assert.Equal(t, "{responses}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("http.status_code")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["couchdb.httpd.responses"] = struct{}{}
		case "couchdb.httpd.views":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of views read.", ms.At(i).Description())
			assert.Equal(t, "{views}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("view")
			assert.True(t, ok)
			assert.Equal(t, "temporary_view_reads", attrVal.Str())
			validatedMetrics["couchdb.httpd.views"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	metricsSettings := MetricsSettings{
		CouchdbAverageRequestTime: MetricSettings{Enabled: false},
		CouchdbDatabaseOpen:       MetricSettings{Enabled: false},
		CouchdbDatabaseOperations: MetricSettings{Enabled: false},
		CouchdbFileDescriptorOpen: MetricSettings{Enabled: false},
		CouchdbHttpdBulkRequests:  MetricSettings{Enabled: false},
		CouchdbHttpdRequests:      MetricSettings{Enabled: false},
		CouchdbHttpdResponses:     MetricSettings{Enabled: false},
		CouchdbHttpdViews:         MetricSettings{Enabled: false},
	}
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(metricsSettings, settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())
	mb.RecordCouchdbAverageRequestTimeDataPoint(ts, 1)
	mb.RecordCouchdbDatabaseOpenDataPoint(ts, 1)
	mb.RecordCouchdbDatabaseOperationsDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordCouchdbFileDescriptorOpenDataPoint(ts, 1)
	mb.RecordCouchdbHttpdBulkRequestsDataPoint(ts, 1)
	mb.RecordCouchdbHttpdRequestsDataPoint(ts, 1, AttributeHTTPMethod(1))
	mb.RecordCouchdbHttpdResponsesDataPoint(ts, 1, "attr-val")
	mb.RecordCouchdbHttpdViewsDataPoint(ts, 1, AttributeView(1))

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
