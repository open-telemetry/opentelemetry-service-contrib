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

// MetricsSettings provides settings for hostmetricsreceiver/cpu metrics.
type MetricsSettings struct {
	SystemCPUTime        MetricSettings `mapstructure:"system.cpu.time"`
	SystemCPUUtilization MetricSettings `mapstructure:"system.cpu.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemCPUTime: MetricSettings{
			Enabled: true,
		},
		SystemCPUUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

type metricSystemCPUTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.time metric with initial data.
func (m *metricSystemCPUTime) init() {
	m.data.SetName("system.cpu.time")
	m.data.SetDescription("Total CPU seconds broken down by different states.")
	m.data.SetUnit("s")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemCPUTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Cpu, pdata.NewValueString(cpuAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUTime(settings MetricSettings) metricSystemCPUTime {
	m := metricSystemCPUTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPUUtilization struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.utilization metric with initial data.
func (m *metricSystemCPUUtilization) init() {
	m.data.SetName("system.cpu.utilization")
	m.data.SetDescription("Percentage of CPU time broken down by different states.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemCPUUtilization) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Cpu, pdata.NewValueString(cpuAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUUtilization) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUUtilization(settings MetricSettings) metricSystemCPUUtilization {
	m := metricSystemCPUUtilization{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                  pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity            int             // maximum observed number of metrics per resource.
	resourceCapacity           int             // maximum observed number of resource attributes.
	metricsBuffer              pdata.Metrics   // accumulates metrics data before emitting.
	metricSystemCPUTime        metricSystemCPUTime
	metricSystemCPUUtilization metricSystemCPUUtilization
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
		startTime:                  pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:              pdata.NewMetrics(),
		metricSystemCPUTime:        newMetricSystemCPUTime(settings.SystemCPUTime),
		metricSystemCPUUtilization: newMetricSystemCPUUtilization(settings.SystemCPUUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pdata.ResourceMetrics) {
	if mb.metricsCapacity < rm.InstrumentationLibraryMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.InstrumentationLibraryMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pdata.Resource)

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pdata.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ils.InstrumentationLibrary().SetName("otelcol/hostmetricsreceiver/cpu")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemCPUTime.emit(ils.Metrics())
	mb.metricSystemCPUUtilization.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pdata.Metrics {
	mb.EmitForResource(ro...)
	metrics := pdata.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemCPUTimeDataPoint adds a data point to system.cpu.time metric.
func (mb *MetricsBuilder) RecordSystemCPUTimeDataPoint(ts pdata.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	mb.metricSystemCPUTime.recordDataPoint(mb.startTime, ts, val, cpuAttributeValue, stateAttributeValue)
}

// RecordSystemCPUUtilizationDataPoint adds a data point to system.cpu.utilization metric.
func (mb *MetricsBuilder) RecordSystemCPUUtilizationDataPoint(ts pdata.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	mb.metricSystemCPUUtilization.recordDataPoint(mb.startTime, ts, val, cpuAttributeValue, stateAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Cpu (CPU number starting at 0.)
	Cpu string
	// State (Breakdown of CPU usage by type.)
	State string
}{
	"cpu",
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Idle      string
	Interrupt string
	Nice      string
	Softirq   string
	Steal     string
	System    string
	User      string
	Wait      string
}{
	"idle",
	"interrupt",
	"nice",
	"softirq",
	"steal",
	"system",
	"user",
	"wait",
}
