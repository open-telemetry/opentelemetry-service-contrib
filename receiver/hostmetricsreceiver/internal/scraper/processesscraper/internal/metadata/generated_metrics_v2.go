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

// MetricsSettings provides settings for hostmetricsreceiver/processes metrics.
type MetricsSettings struct {
	SystemProcessesCount   MetricSettings `mapstructure:"system.processes.count"`
	SystemProcessesCreated MetricSettings `mapstructure:"system.processes.created"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemProcessesCount: MetricSettings{
			Enabled: true,
		},
		SystemProcessesCreated: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeStatus specifies the a value status attribute.
type AttributeStatus int

const (
	_ AttributeStatus = iota
	AttributeStatusBlocked
	AttributeStatusDaemon
	AttributeStatusDetached
	AttributeStatusIdle
	AttributeStatusLocked
	AttributeStatusOrphan
	AttributeStatusPaging
	AttributeStatusRunning
	AttributeStatusSleeping
	AttributeStatusStopped
	AttributeStatusSystem
	AttributeStatusUnknown
	AttributeStatusZombies
)

// String returns the string representation of the AttributeStatus.
func (av AttributeStatus) String() string {
	switch av {
	case AttributeStatusBlocked:
		return "blocked"
	case AttributeStatusDaemon:
		return "daemon"
	case AttributeStatusDetached:
		return "detached"
	case AttributeStatusIdle:
		return "idle"
	case AttributeStatusLocked:
		return "locked"
	case AttributeStatusOrphan:
		return "orphan"
	case AttributeStatusPaging:
		return "paging"
	case AttributeStatusRunning:
		return "running"
	case AttributeStatusSleeping:
		return "sleeping"
	case AttributeStatusStopped:
		return "stopped"
	case AttributeStatusSystem:
		return "system"
	case AttributeStatusUnknown:
		return "unknown"
	case AttributeStatusZombies:
		return "zombies"
	}
	return ""
}

// MapAttributeStatus is a helper map of string to AttributeStatus attribute value.
var MapAttributeStatus = map[string]AttributeStatus{
	"blocked":  AttributeStatusBlocked,
	"daemon":   AttributeStatusDaemon,
	"detached": AttributeStatusDetached,
	"idle":     AttributeStatusIdle,
	"locked":   AttributeStatusLocked,
	"orphan":   AttributeStatusOrphan,
	"paging":   AttributeStatusPaging,
	"running":  AttributeStatusRunning,
	"sleeping": AttributeStatusSleeping,
	"stopped":  AttributeStatusStopped,
	"system":   AttributeStatusSystem,
	"unknown":  AttributeStatusUnknown,
	"zombies":  AttributeStatusZombies,
}

type metricSystemProcessesCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.processes.count metric with initial data.
func (m *metricSystemProcessesCount) init() {
	m.data.SetName("system.processes.count")
	m.data.SetDescription("Total number of processes in each state.")
	m.data.SetUnit("{processes}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemProcessesCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, statusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("status", pcommon.NewValueString(statusAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemProcessesCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemProcessesCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemProcessesCount(settings MetricSettings) metricSystemProcessesCount {
	m := metricSystemProcessesCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemProcessesCreated struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.processes.created metric with initial data.
func (m *metricSystemProcessesCreated) init() {
	m.data.SetName("system.processes.created")
	m.data.SetDescription("Total number of created processes.")
	m.data.SetUnit("{processes}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricSystemProcessesCreated) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemProcessesCreated) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemProcessesCreated) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemProcessesCreated(settings MetricSettings) metricSystemProcessesCreated {
	m := metricSystemProcessesCreated{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                    pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity              int               // maximum observed number of metrics per resource.
	resourceCapacity             int               // maximum observed number of resource attributes.
	metricsBuffer                pmetric.Metrics   // accumulates metrics data before emitting.
	metricSystemProcessesCount   metricSystemProcessesCount
	metricSystemProcessesCreated metricSystemProcessesCreated
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
		startTime:                    pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                pmetric.NewMetrics(),
		metricSystemProcessesCount:   newMetricSystemProcessesCount(settings.SystemProcessesCount),
		metricSystemProcessesCreated: newMetricSystemProcessesCreated(settings.SystemProcessesCreated),
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
	ils.Scope().SetName("otelcol/hostmetricsreceiver/processes")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemProcessesCount.emit(ils.Metrics())
	mb.metricSystemProcessesCreated.emit(ils.Metrics())
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

// RecordSystemProcessesCountDataPoint adds a data point to system.processes.count metric.
func (mb *MetricsBuilder) RecordSystemProcessesCountDataPoint(ts pcommon.Timestamp, val int64, statusAttributeValue AttributeStatus) {
	mb.metricSystemProcessesCount.recordDataPoint(mb.startTime, ts, val, statusAttributeValue.String())
}

// RecordSystemProcessesCreatedDataPoint adds a data point to system.processes.created metric.
func (mb *MetricsBuilder) RecordSystemProcessesCreatedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSystemProcessesCreated.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
