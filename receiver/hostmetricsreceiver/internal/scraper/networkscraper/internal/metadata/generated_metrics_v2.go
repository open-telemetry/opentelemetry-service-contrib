// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/network metrics.
type MetricsSettings struct {
	SystemNetworkConnections MetricSettings `mapstructure:"system.network.connections"`
	SystemNetworkDropped     MetricSettings `mapstructure:"system.network.dropped"`
	SystemNetworkErrors      MetricSettings `mapstructure:"system.network.errors"`
	SystemNetworkIo          MetricSettings `mapstructure:"system.network.io"`
	SystemNetworkPackets     MetricSettings `mapstructure:"system.network.packets"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemNetworkConnections: MetricSettings{
			Enabled: true,
		},
		SystemNetworkDropped: MetricSettings{
			Enabled: true,
		},
		SystemNetworkErrors: MetricSettings{
			Enabled: true,
		},
		SystemNetworkIo: MetricSettings{
			Enabled: true,
		},
		SystemNetworkPackets: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionReceive
	AttributeDirectionTransmit
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionReceive:
		return "receive"
	case AttributeDirectionTransmit:
		return "transmit"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"receive":  AttributeDirectionReceive,
	"transmit": AttributeDirectionTransmit,
}

// AttributeProtocol specifies the a value protocol attribute.
type AttributeProtocol int

const (
	_ AttributeProtocol = iota
	AttributeProtocolTcp
)

// String returns the string representation of the AttributeProtocol.
func (av AttributeProtocol) String() string {
	switch av {
	case AttributeProtocolTcp:
		return "tcp"
	}
	return ""
}

// MapAttributeProtocol is a helper map of string to AttributeProtocol attribute value.
var MapAttributeProtocol = map[string]AttributeProtocol{
	"tcp": AttributeProtocolTcp,
}

type metricSystemNetworkConnections struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.connections metric with initial data.
func (m *metricSystemNetworkConnections) init() {
	m.data.SetName("system.network.connections")
	m.data.SetDescription("The number of connections.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkConnections) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, protocolAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("protocol", pcommon.NewValueString(protocolAttributeValue))
	dp.Attributes().Insert("state", pcommon.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkConnections) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkConnections) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkConnections(settings MetricSettings) metricSystemNetworkConnections {
	m := metricSystemNetworkConnections{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkDropped struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.dropped metric with initial data.
func (m *metricSystemNetworkDropped) init() {
	m.data.SetName("system.network.dropped")
	m.data.SetDescription("The number of packets dropped.")
	m.data.SetUnit("{packets}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkDropped) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("direction", pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkDropped) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkDropped) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkDropped(settings MetricSettings) metricSystemNetworkDropped {
	m := metricSystemNetworkDropped{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkErrors struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.errors metric with initial data.
func (m *metricSystemNetworkErrors) init() {
	m.data.SetName("system.network.errors")
	m.data.SetDescription("The number of errors encountered.")
	m.data.SetUnit("{errors}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkErrors) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("direction", pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkErrors) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkErrors) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkErrors(settings MetricSettings) metricSystemNetworkErrors {
	m := metricSystemNetworkErrors{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.io metric with initial data.
func (m *metricSystemNetworkIo) init() {
	m.data.SetName("system.network.io")
	m.data.SetDescription("The number of bytes transmitted and received.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("direction", pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkIo) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkIo(settings MetricSettings) metricSystemNetworkIo {
	m := metricSystemNetworkIo{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkPackets struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.packets metric with initial data.
func (m *metricSystemNetworkPackets) init() {
	m.data.SetName("system.network.packets")
	m.data.SetDescription("The number of packets transferred.")
	m.data.SetUnit("{packets}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkPackets) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("device", pcommon.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert("direction", pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkPackets) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkPackets) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkPackets(settings MetricSettings) metricSystemNetworkPackets {
	m := metricSystemNetworkPackets{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                int               // maximum observed number of metrics per resource.
	resourceCapacity               int               // maximum observed number of resource attributes.
	metricsBuffer                  pmetric.Metrics   // accumulates metrics data before emitting.
	metricSystemNetworkConnections metricSystemNetworkConnections
	metricSystemNetworkDropped     metricSystemNetworkDropped
	metricSystemNetworkErrors      metricSystemNetworkErrors
	metricSystemNetworkIo          metricSystemNetworkIo
	metricSystemNetworkPackets     metricSystemNetworkPackets
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		metricSystemNetworkConnections: newMetricSystemNetworkConnections(settings.SystemNetworkConnections),
		metricSystemNetworkDropped:     newMetricSystemNetworkDropped(settings.SystemNetworkDropped),
		metricSystemNetworkErrors:      newMetricSystemNetworkErrors(settings.SystemNetworkErrors),
		metricSystemNetworkIo:          newMetricSystemNetworkIo(settings.SystemNetworkIo),
		metricSystemNetworkPackets:     newMetricSystemNetworkPackets(settings.SystemNetworkPackets),
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

// ResourceOption applies changes to provided resource.
type ResourceOption func(pcommon.Resource)

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/network")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemNetworkConnections.emit(ils.Metrics())
	mb.metricSystemNetworkDropped.emit(ils.Metrics())
	mb.metricSystemNetworkErrors.emit(ils.Metrics())
	mb.metricSystemNetworkIo.emit(ils.Metrics())
	mb.metricSystemNetworkPackets.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pmetric.Metrics {
	mb.EmitForResource(ro...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemNetworkConnectionsDataPoint adds a data point to system.network.connections metric.
func (mb *MetricsBuilder) RecordSystemNetworkConnectionsDataPoint(ts pcommon.Timestamp, val int64, protocolAttributeValue AttributeProtocol, stateAttributeValue string) {
	mb.metricSystemNetworkConnections.recordDataPoint(mb.startTime, ts, val, protocolAttributeValue.String(), stateAttributeValue)
}

// RecordSystemNetworkDroppedDataPoint adds a data point to system.network.dropped metric.
func (mb *MetricsBuilder) RecordSystemNetworkDroppedDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemNetworkDropped.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemNetworkErrorsDataPoint adds a data point to system.network.errors metric.
func (mb *MetricsBuilder) RecordSystemNetworkErrorsDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemNetworkErrors.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemNetworkIoDataPoint adds a data point to system.network.io metric.
func (mb *MetricsBuilder) RecordSystemNetworkIoDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemNetworkIo.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// RecordSystemNetworkPacketsDataPoint adds a data point to system.network.packets metric.
func (mb *MetricsBuilder) RecordSystemNetworkPacketsDataPoint(ts pcommon.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue AttributeDirection) {
	mb.metricSystemNetworkPackets.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
