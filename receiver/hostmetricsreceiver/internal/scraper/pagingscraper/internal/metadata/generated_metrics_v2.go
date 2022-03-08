// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for paging metrics.
type MetricsSettings struct {
	SystemPagingFaults      MetricSettings `mapstructure:"system.paging.faults"`
	SystemPagingOperations  MetricSettings `mapstructure:"system.paging.operations"`
	SystemPagingUsage       MetricSettings `mapstructure:"system.paging.usage"`
	SystemPagingUtilization MetricSettings `mapstructure:"system.paging.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemPagingFaults: MetricSettings{
			Enabled: true,
		},
		SystemPagingOperations: MetricSettings{
			Enabled: true,
		},
		SystemPagingUsage: MetricSettings{
			Enabled: true,
		},
		SystemPagingUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

type metricSystemPagingFaults struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.faults metric with initial data.
func (m *metricSystemPagingFaults) init() {
	m.data.SetName("system.paging.faults")
	m.data.SetDescription("The number of page faults.")
	m.data.SetUnit("{faults}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingFaults) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Type, pdata.NewAttributeValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingFaults) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingFaults) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingFaults(settings MetricSettings) metricSystemPagingFaults {
	m := metricSystemPagingFaults{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.operations metric with initial data.
func (m *metricSystemPagingOperations) init() {
	m.data.SetName("system.paging.operations")
	m.data.SetDescription("The number of paging operations.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, directionAttributeValue string, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Direction, pdata.NewAttributeValueString(directionAttributeValue))
	dp.Attributes().Insert(A.Type, pdata.NewAttributeValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingOperations(settings MetricSettings) metricSystemPagingOperations {
	m := metricSystemPagingOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingUsage struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.usage metric with initial data.
func (m *metricSystemPagingUsage) init() {
	m.data.SetName("system.paging.usage")
	m.data.SetDescription("Swap (unix) or pagefile (windows) usage.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingUsage) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewAttributeValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingUsage) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingUsage(settings MetricSettings) metricSystemPagingUsage {
	m := metricSystemPagingUsage{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingUtilization struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.utilization metric with initial data.
func (m *metricSystemPagingUtilization) init() {
	m.data.SetName("system.paging.utilization")
	m.data.SetDescription("Swap (unix) or pagefile (windows) utilization.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingUtilization) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, deviceAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewAttributeValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingUtilization) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingUtilization(settings MetricSettings) metricSystemPagingUtilization {
	m := metricSystemPagingUtilization{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pdata.Timestamp
	metricSystemPagingFaults      metricSystemPagingFaults
	metricSystemPagingOperations  metricSystemPagingOperations
	metricSystemPagingUsage       metricSystemPagingUsage
	metricSystemPagingUtilization metricSystemPagingUtilization
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                     pdata.NewTimestampFromTime(time.Now()),
		metricSystemPagingFaults:      newMetricSystemPagingFaults(settings.SystemPagingFaults),
		metricSystemPagingOperations:  newMetricSystemPagingOperations(settings.SystemPagingOperations),
		metricSystemPagingUsage:       newMetricSystemPagingUsage(settings.SystemPagingUsage),
		metricSystemPagingUtilization: newMetricSystemPagingUtilization(settings.SystemPagingUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricSystemPagingFaults.emit(metrics)
	mb.metricSystemPagingOperations.emit(metrics)
	mb.metricSystemPagingUsage.emit(metrics)
	mb.metricSystemPagingUtilization.emit(metrics)
}

// RecordSystemPagingFaultsDataPoint adds a data point to system.paging.faults metric.
func (mb *MetricsBuilder) RecordSystemPagingFaultsDataPoint(ts pdata.Timestamp, val int64, typeAttributeValue string) {
	mb.metricSystemPagingFaults.recordDataPoint(mb.startTime, ts, val, typeAttributeValue)
}

// RecordSystemPagingOperationsDataPoint adds a data point to system.paging.operations metric.
func (mb *MetricsBuilder) RecordSystemPagingOperationsDataPoint(ts pdata.Timestamp, val int64, directionAttributeValue string, typeAttributeValue string) {
	mb.metricSystemPagingOperations.recordDataPoint(mb.startTime, ts, val, directionAttributeValue, typeAttributeValue)
}

// RecordSystemPagingUsageDataPoint adds a data point to system.paging.usage metric.
func (mb *MetricsBuilder) RecordSystemPagingUsageDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, stateAttributeValue string) {
	mb.metricSystemPagingUsage.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, stateAttributeValue)
}

// RecordSystemPagingUtilizationDataPoint adds a data point to system.paging.utilization metric.
func (mb *MetricsBuilder) RecordSystemPagingUtilizationDataPoint(ts pdata.Timestamp, val float64, deviceAttributeValue string, stateAttributeValue string) {
	mb.metricSystemPagingUtilization.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, stateAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// NewMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) NewMetricData() pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/paging")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Device (Name of the page file.)
	Device string
	// Direction (Page In or Page Out.)
	Direction string
	// State (Breakdown of paging usage by type.)
	State string
	// Type (Type of fault.)
	Type string
}{
	"device",
	"direction",
	"state",
	"type",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	PageIn  string
	PageOut string
}{
	"page_in",
	"page_out",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Cached string
	Free   string
	Used   string
}{
	"cached",
	"free",
	"used",
}

// AttributeType are the possible values that the attribute "type" can have.
var AttributeType = struct {
	Major string
	Minor string
}{
	"major",
	"minor",
}
