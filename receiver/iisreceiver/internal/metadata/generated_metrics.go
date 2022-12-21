// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledProvidedByUser bool
}

// IsEnabledProvidedByUser returns true if `enabled` option is explicitly set in user settings to any value.
func (ms *MetricSettings) IsEnabledProvidedByUser() bool {
	return ms.enabledProvidedByUser
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for iisreceiver metrics.
type MetricsSettings struct {
	IisConnectionActive       MetricSettings `mapstructure:"iis.connection.active"`
	IisConnectionAnonymous    MetricSettings `mapstructure:"iis.connection.anonymous"`
	IisConnectionAttemptCount MetricSettings `mapstructure:"iis.connection.attempt.count"`
	IisNetworkBlocked         MetricSettings `mapstructure:"iis.network.blocked"`
	IisNetworkFileCount       MetricSettings `mapstructure:"iis.network.file.count"`
	IisNetworkIo              MetricSettings `mapstructure:"iis.network.io"`
	IisRequestCount           MetricSettings `mapstructure:"iis.request.count"`
	IisRequestQueueAgeMax     MetricSettings `mapstructure:"iis.request.queue.age.max"`
	IisRequestQueueCount      MetricSettings `mapstructure:"iis.request.queue.count"`
	IisRequestRejected        MetricSettings `mapstructure:"iis.request.rejected"`
	IisThreadActive           MetricSettings `mapstructure:"iis.thread.active"`
	IisUptime                 MetricSettings `mapstructure:"iis.uptime"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		IisConnectionActive: MetricSettings{
			Enabled: true,
		},
		IisConnectionAnonymous: MetricSettings{
			Enabled: true,
		},
		IisConnectionAttemptCount: MetricSettings{
			Enabled: true,
		},
		IisNetworkBlocked: MetricSettings{
			Enabled: true,
		},
		IisNetworkFileCount: MetricSettings{
			Enabled: true,
		},
		IisNetworkIo: MetricSettings{
			Enabled: true,
		},
		IisRequestCount: MetricSettings{
			Enabled: true,
		},
		IisRequestQueueAgeMax: MetricSettings{
			Enabled: true,
		},
		IisRequestQueueCount: MetricSettings{
			Enabled: true,
		},
		IisRequestRejected: MetricSettings{
			Enabled: true,
		},
		IisThreadActive: MetricSettings{
			Enabled: true,
		},
		IisUptime: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionSent
	AttributeDirectionReceived
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionSent:
		return "sent"
	case AttributeDirectionReceived:
		return "received"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"sent":     AttributeDirectionSent,
	"received": AttributeDirectionReceived,
}

// AttributeRequest specifies the a value request attribute.
type AttributeRequest int

const (
	_ AttributeRequest = iota
	AttributeRequestDelete
	AttributeRequestGet
	AttributeRequestHead
	AttributeRequestOptions
	AttributeRequestPost
	AttributeRequestPut
	AttributeRequestTrace
)

// String returns the string representation of the AttributeRequest.
func (av AttributeRequest) String() string {
	switch av {
	case AttributeRequestDelete:
		return "delete"
	case AttributeRequestGet:
		return "get"
	case AttributeRequestHead:
		return "head"
	case AttributeRequestOptions:
		return "options"
	case AttributeRequestPost:
		return "post"
	case AttributeRequestPut:
		return "put"
	case AttributeRequestTrace:
		return "trace"
	}
	return ""
}

// MapAttributeRequest is a helper map of string to AttributeRequest attribute value.
var MapAttributeRequest = map[string]AttributeRequest{
	"delete":  AttributeRequestDelete,
	"get":     AttributeRequestGet,
	"head":    AttributeRequestHead,
	"options": AttributeRequestOptions,
	"post":    AttributeRequestPost,
	"put":     AttributeRequestPut,
	"trace":   AttributeRequestTrace,
}

type metricIisConnectionActive struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.connection.active metric with initial data.
func (m *metricIisConnectionActive) init() {
	m.data.SetName("iis.connection.active")
	m.data.SetDescription("Number of active connections.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisConnectionActive) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisConnectionActive) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisConnectionActive) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisConnectionActive(settings MetricSettings) metricIisConnectionActive {
	m := metricIisConnectionActive{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisConnectionAnonymous struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.connection.anonymous metric with initial data.
func (m *metricIisConnectionAnonymous) init() {
	m.data.SetName("iis.connection.anonymous")
	m.data.SetDescription("Number of connections established anonymously.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisConnectionAnonymous) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisConnectionAnonymous) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisConnectionAnonymous) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisConnectionAnonymous(settings MetricSettings) metricIisConnectionAnonymous {
	m := metricIisConnectionAnonymous{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisConnectionAttemptCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.connection.attempt.count metric with initial data.
func (m *metricIisConnectionAttemptCount) init() {
	m.data.SetName("iis.connection.attempt.count")
	m.data.SetDescription("Total number of attempts to connect to the server.")
	m.data.SetUnit("{attempts}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisConnectionAttemptCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisConnectionAttemptCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisConnectionAttemptCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisConnectionAttemptCount(settings MetricSettings) metricIisConnectionAttemptCount {
	m := metricIisConnectionAttemptCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisNetworkBlocked struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.network.blocked metric with initial data.
func (m *metricIisNetworkBlocked) init() {
	m.data.SetName("iis.network.blocked")
	m.data.SetDescription("Number of bytes blocked due to bandwidth throttling.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisNetworkBlocked) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisNetworkBlocked) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisNetworkBlocked) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisNetworkBlocked(settings MetricSettings) metricIisNetworkBlocked {
	m := metricIisNetworkBlocked{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisNetworkFileCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.network.file.count metric with initial data.
func (m *metricIisNetworkFileCount) init() {
	m.data.SetName("iis.network.file.count")
	m.data.SetDescription("Number of transmitted files.")
	m.data.SetUnit("{files}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricIisNetworkFileCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisNetworkFileCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisNetworkFileCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisNetworkFileCount(settings MetricSettings) metricIisNetworkFileCount {
	m := metricIisNetworkFileCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisNetworkIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.network.io metric with initial data.
func (m *metricIisNetworkIo) init() {
	m.data.SetName("iis.network.io")
	m.data.SetDescription("Total amount of bytes sent and received.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricIisNetworkIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisNetworkIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisNetworkIo) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisNetworkIo(settings MetricSettings) metricIisNetworkIo {
	m := metricIisNetworkIo{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisRequestCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.request.count metric with initial data.
func (m *metricIisRequestCount) init() {
	m.data.SetName("iis.request.count")
	m.data.SetDescription("Total number of requests of a given type.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricIisRequestCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("request", requestAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisRequestCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisRequestCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisRequestCount(settings MetricSettings) metricIisRequestCount {
	m := metricIisRequestCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisRequestQueueAgeMax struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.request.queue.age.max metric with initial data.
func (m *metricIisRequestQueueAgeMax) init() {
	m.data.SetName("iis.request.queue.age.max")
	m.data.SetDescription("Age of oldest request in the queue.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricIisRequestQueueAgeMax) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisRequestQueueAgeMax) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisRequestQueueAgeMax) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisRequestQueueAgeMax(settings MetricSettings) metricIisRequestQueueAgeMax {
	m := metricIisRequestQueueAgeMax{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisRequestQueueCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.request.queue.count metric with initial data.
func (m *metricIisRequestQueueCount) init() {
	m.data.SetName("iis.request.queue.count")
	m.data.SetDescription("Current number of requests in the queue.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisRequestQueueCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisRequestQueueCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisRequestQueueCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisRequestQueueCount(settings MetricSettings) metricIisRequestQueueCount {
	m := metricIisRequestQueueCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisRequestRejected struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.request.rejected metric with initial data.
func (m *metricIisRequestRejected) init() {
	m.data.SetName("iis.request.rejected")
	m.data.SetDescription("Total number of requests rejected.")
	m.data.SetUnit("{requests}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisRequestRejected) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisRequestRejected) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisRequestRejected) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisRequestRejected(settings MetricSettings) metricIisRequestRejected {
	m := metricIisRequestRejected{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisThreadActive struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.thread.active metric with initial data.
func (m *metricIisThreadActive) init() {
	m.data.SetName("iis.thread.active")
	m.data.SetDescription("Current number of active threads.")
	m.data.SetUnit("{threads}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricIisThreadActive) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisThreadActive) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisThreadActive) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisThreadActive(settings MetricSettings) metricIisThreadActive {
	m := metricIisThreadActive{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricIisUptime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills iis.uptime metric with initial data.
func (m *metricIisUptime) init() {
	m.data.SetName("iis.uptime")
	m.data.SetDescription("The amount of time the server has been up.")
	m.data.SetUnit("s")
	m.data.SetEmptyGauge()
}

func (m *metricIisUptime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricIisUptime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricIisUptime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricIisUptime(settings MetricSettings) metricIisUptime {
	m := metricIisUptime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                       pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                 int                 // maximum observed number of metrics per resource.
	resourceCapacity                int                 // maximum observed number of resource attributes.
	metricsBuffer                   pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                       component.BuildInfo // contains version information
	metricIisConnectionActive       metricIisConnectionActive
	metricIisConnectionAnonymous    metricIisConnectionAnonymous
	metricIisConnectionAttemptCount metricIisConnectionAttemptCount
	metricIisNetworkBlocked         metricIisNetworkBlocked
	metricIisNetworkFileCount       metricIisNetworkFileCount
	metricIisNetworkIo              metricIisNetworkIo
	metricIisRequestCount           metricIisRequestCount
	metricIisRequestQueueAgeMax     metricIisRequestQueueAgeMax
	metricIisRequestQueueCount      metricIisRequestQueueCount
	metricIisRequestRejected        metricIisRequestRejected
	metricIisThreadActive           metricIisThreadActive
	metricIisUptime                 metricIisUptime
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(ms MetricsSettings, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                       pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                   pmetric.NewMetrics(),
		buildInfo:                       settings.BuildInfo,
		metricIisConnectionActive:       newMetricIisConnectionActive(ms.IisConnectionActive),
		metricIisConnectionAnonymous:    newMetricIisConnectionAnonymous(ms.IisConnectionAnonymous),
		metricIisConnectionAttemptCount: newMetricIisConnectionAttemptCount(ms.IisConnectionAttemptCount),
		metricIisNetworkBlocked:         newMetricIisNetworkBlocked(ms.IisNetworkBlocked),
		metricIisNetworkFileCount:       newMetricIisNetworkFileCount(ms.IisNetworkFileCount),
		metricIisNetworkIo:              newMetricIisNetworkIo(ms.IisNetworkIo),
		metricIisRequestCount:           newMetricIisRequestCount(ms.IisRequestCount),
		metricIisRequestQueueAgeMax:     newMetricIisRequestQueueAgeMax(ms.IisRequestQueueAgeMax),
		metricIisRequestQueueCount:      newMetricIisRequestQueueCount(ms.IisRequestQueueCount),
		metricIisRequestRejected:        newMetricIisRequestRejected(ms.IisRequestRejected),
		metricIisThreadActive:           newMetricIisThreadActive(ms.IisThreadActive),
		metricIisUptime:                 newMetricIisUptime(ms.IisUptime),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithIisApplicationPool sets provided value as "iis.application_pool" attribute for current resource.
func WithIisApplicationPool(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("iis.application_pool", val)
	}
}

// WithIisSite sets provided value as "iis.site" attribute for current resource.
func WithIisSite(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("iis.site", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/iisreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricIisConnectionActive.emit(ils.Metrics())
	mb.metricIisConnectionAnonymous.emit(ils.Metrics())
	mb.metricIisConnectionAttemptCount.emit(ils.Metrics())
	mb.metricIisNetworkBlocked.emit(ils.Metrics())
	mb.metricIisNetworkFileCount.emit(ils.Metrics())
	mb.metricIisNetworkIo.emit(ils.Metrics())
	mb.metricIisRequestCount.emit(ils.Metrics())
	mb.metricIisRequestQueueAgeMax.emit(ils.Metrics())
	mb.metricIisRequestQueueCount.emit(ils.Metrics())
	mb.metricIisRequestRejected.emit(ils.Metrics())
	mb.metricIisThreadActive.emit(ils.Metrics())
	mb.metricIisUptime.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordIisConnectionActiveDataPoint adds a data point to iis.connection.active metric.
func (mb *MetricsBuilder) RecordIisConnectionActiveDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisConnectionActive.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisConnectionAnonymousDataPoint adds a data point to iis.connection.anonymous metric.
func (mb *MetricsBuilder) RecordIisConnectionAnonymousDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisConnectionAnonymous.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisConnectionAttemptCountDataPoint adds a data point to iis.connection.attempt.count metric.
func (mb *MetricsBuilder) RecordIisConnectionAttemptCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisConnectionAttemptCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisNetworkBlockedDataPoint adds a data point to iis.network.blocked metric.
func (mb *MetricsBuilder) RecordIisNetworkBlockedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisNetworkBlocked.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisNetworkFileCountDataPoint adds a data point to iis.network.file.count metric.
func (mb *MetricsBuilder) RecordIisNetworkFileCountDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricIisNetworkFileCount.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordIisNetworkIoDataPoint adds a data point to iis.network.io metric.
func (mb *MetricsBuilder) RecordIisNetworkIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricIisNetworkIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordIisRequestCountDataPoint adds a data point to iis.request.count metric.
func (mb *MetricsBuilder) RecordIisRequestCountDataPoint(ts pcommon.Timestamp, val int64, requestAttributeValue AttributeRequest) {
	mb.metricIisRequestCount.recordDataPoint(mb.startTime, ts, val, requestAttributeValue.String())
}

// RecordIisRequestQueueAgeMaxDataPoint adds a data point to iis.request.queue.age.max metric.
func (mb *MetricsBuilder) RecordIisRequestQueueAgeMaxDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisRequestQueueAgeMax.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisRequestQueueCountDataPoint adds a data point to iis.request.queue.count metric.
func (mb *MetricsBuilder) RecordIisRequestQueueCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisRequestQueueCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisRequestRejectedDataPoint adds a data point to iis.request.rejected metric.
func (mb *MetricsBuilder) RecordIisRequestRejectedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisRequestRejected.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisThreadActiveDataPoint adds a data point to iis.thread.active metric.
func (mb *MetricsBuilder) RecordIisThreadActiveDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisThreadActive.recordDataPoint(mb.startTime, ts, val)
}

// RecordIisUptimeDataPoint adds a data point to iis.uptime metric.
func (mb *MetricsBuilder) RecordIisUptimeDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricIisUptime.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
