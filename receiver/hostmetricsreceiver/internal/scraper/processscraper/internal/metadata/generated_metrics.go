// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/process metrics.
type MetricsSettings struct {
	ProcessCPUTime             MetricSettings `mapstructure:"process.cpu.time"`
	ProcessDiskIo              MetricSettings `mapstructure:"process.disk.io"`
	ProcessDiskIoRead          MetricSettings `mapstructure:"process.disk.io.read"`
	ProcessDiskIoWrite         MetricSettings `mapstructure:"process.disk.io.write"`
	ProcessMemoryPhysicalUsage MetricSettings `mapstructure:"process.memory.physical_usage"`
	ProcessMemoryVirtualUsage  MetricSettings `mapstructure:"process.memory.virtual_usage"`
	ProcessThreads             MetricSettings `mapstructure:"process.threads"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ProcessCPUTime: MetricSettings{
			Enabled: true,
		},
		ProcessDiskIo: MetricSettings{
			Enabled: true,
		},
		ProcessDiskIoRead: MetricSettings{
			Enabled: true,
		},
		ProcessDiskIoWrite: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryPhysicalUsage: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryVirtualUsage: MetricSettings{
			Enabled: true,
		},
		ProcessThreads: MetricSettings{
			Enabled: false,
		},
	}
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionRead
	AttributeDirectionWrite
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionRead:
		return "read"
	case AttributeDirectionWrite:
		return "write"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"read":  AttributeDirectionRead,
	"write": AttributeDirectionWrite,
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateSystem
	AttributeStateUser
	AttributeStateWait
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateSystem:
		return "system"
	case AttributeStateUser:
		return "user"
	case AttributeStateWait:
		return "wait"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"system": AttributeStateSystem,
	"user":   AttributeStateUser,
	"wait":   AttributeStateWait,
}

type metricProcessCPUTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.cpu.time metric with initial data.
func (m *metricProcessCPUTime) init() {
	m.data.SetName("process.cpu.time")
	m.data.SetDescription("Total CPU seconds broken down by different states.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().UpsertString("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessCPUTime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessCPUTime(settings MetricSettings) metricProcessCPUTime {
	m := metricProcessCPUTime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.io metric with initial data.
func (m *metricProcessDiskIo) init() {
	m.data.SetName("process.disk.io")
	m.data.SetDescription("Disk bytes transferred.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessDiskIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().UpsertString("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskIo) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskIo(settings MetricSettings) metricProcessDiskIo {
	m := metricProcessDiskIo{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskIoRead struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.io.read metric with initial data.
func (m *metricProcessDiskIoRead) init() {
	m.data.SetName("process.disk.io.read")
	m.data.SetDescription("Disk bytes read.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessDiskIoRead) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskIoRead) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskIoRead) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskIoRead(settings MetricSettings) metricProcessDiskIoRead {
	m := metricProcessDiskIoRead{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskIoWrite struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.io.write metric with initial data.
func (m *metricProcessDiskIoWrite) init() {
	m.data.SetName("process.disk.io.write")
	m.data.SetDescription("Disk bytes written.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessDiskIoWrite) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskIoWrite) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskIoWrite) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskIoWrite(settings MetricSettings) metricProcessDiskIoWrite {
	m := metricProcessDiskIoWrite{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryPhysicalUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.physical_usage metric with initial data.
func (m *metricProcessMemoryPhysicalUsage) init() {
	m.data.SetName("process.memory.physical_usage")
	m.data.SetDescription("The amount of physical memory in use.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessMemoryPhysicalUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryPhysicalUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryPhysicalUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryPhysicalUsage(settings MetricSettings) metricProcessMemoryPhysicalUsage {
	m := metricProcessMemoryPhysicalUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryVirtualUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.virtual_usage metric with initial data.
func (m *metricProcessMemoryVirtualUsage) init() {
	m.data.SetName("process.memory.virtual_usage")
	m.data.SetDescription("Virtual memory size.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessMemoryVirtualUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryVirtualUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryVirtualUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryVirtualUsage(settings MetricSettings) metricProcessMemoryVirtualUsage {
	m := metricProcessMemoryVirtualUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessThreads struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.threads metric with initial data.
func (m *metricProcessThreads) init() {
	m.data.SetName("process.threads")
	m.data.SetDescription("Process threads count.")
	m.data.SetUnit("{threads}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessThreads) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessThreads) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessThreads) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessThreads(settings MetricSettings) metricProcessThreads {
	m := metricProcessThreads{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                        pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                  int                 // maximum observed number of metrics per resource.
	resourceCapacity                 int                 // maximum observed number of resource attributes.
	metricsBuffer                    pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                        component.BuildInfo // contains version information
	metricProcessCPUTime             metricProcessCPUTime
	metricProcessDiskIo              metricProcessDiskIo
	metricProcessDiskIoRead          metricProcessDiskIoRead
	metricProcessDiskIoWrite         metricProcessDiskIoWrite
	metricProcessMemoryPhysicalUsage metricProcessMemoryPhysicalUsage
	metricProcessMemoryVirtualUsage  metricProcessMemoryVirtualUsage
	metricProcessThreads             metricProcessThreads
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                        pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                    pmetric.NewMetrics(),
		buildInfo:                        buildInfo,
		metricProcessCPUTime:             newMetricProcessCPUTime(settings.ProcessCPUTime),
		metricProcessDiskIo:              newMetricProcessDiskIo(settings.ProcessDiskIo),
		metricProcessDiskIoRead:          newMetricProcessDiskIoRead(settings.ProcessDiskIoRead),
		metricProcessDiskIoWrite:         newMetricProcessDiskIoWrite(settings.ProcessDiskIoWrite),
		metricProcessMemoryPhysicalUsage: newMetricProcessMemoryPhysicalUsage(settings.ProcessMemoryPhysicalUsage),
		metricProcessMemoryVirtualUsage:  newMetricProcessMemoryVirtualUsage(settings.ProcessMemoryVirtualUsage),
		metricProcessThreads:             newMetricProcessThreads(settings.ProcessThreads),
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

// WithProcessCommand sets provided value as "process.command" attribute for current resource.
func WithProcessCommand(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("process.command", val)
	}
}

// WithProcessCommandLine sets provided value as "process.command_line" attribute for current resource.
func WithProcessCommandLine(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("process.command_line", val)
	}
}

// WithProcessExecutableName sets provided value as "process.executable.name" attribute for current resource.
func WithProcessExecutableName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("process.executable.name", val)
	}
}

// WithProcessExecutablePath sets provided value as "process.executable.path" attribute for current resource.
func WithProcessExecutablePath(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("process.executable.path", val)
	}
}

// WithProcessOwner sets provided value as "process.owner" attribute for current resource.
func WithProcessOwner(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("process.owner", val)
	}
}

// WithProcessParentPid sets provided value as "process.parent_pid" attribute for current resource.
func WithProcessParentPid(val int64) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertInt("process.parent_pid", val)
	}
}

// WithProcessPid sets provided value as "process.pid" attribute for current resource.
func WithProcessPid(val int64) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertInt("process.pid", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).DataType() {
			case pmetric.MetricDataTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricDataTypeSum:
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
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/process")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricProcessCPUTime.emit(ils.Metrics())
	mb.metricProcessDiskIo.emit(ils.Metrics())
	mb.metricProcessDiskIoRead.emit(ils.Metrics())
	mb.metricProcessDiskIoWrite.emit(ils.Metrics())
	mb.metricProcessMemoryPhysicalUsage.emit(ils.Metrics())
	mb.metricProcessMemoryVirtualUsage.emit(ils.Metrics())
	mb.metricProcessThreads.emit(ils.Metrics())
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

// RecordProcessCPUTimeDataPoint adds a data point to process.cpu.time metric.
func (mb *MetricsBuilder) RecordProcessCPUTimeDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricProcessCPUTime.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordProcessDiskIoDataPoint adds a data point to process.disk.io metric.
func (mb *MetricsBuilder) RecordProcessDiskIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricProcessDiskIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordProcessDiskIoReadDataPoint adds a data point to process.disk.io.read metric.
func (mb *MetricsBuilder) RecordProcessDiskIoReadDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessDiskIoRead.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessDiskIoWriteDataPoint adds a data point to process.disk.io.write metric.
func (mb *MetricsBuilder) RecordProcessDiskIoWriteDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessDiskIoWrite.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryPhysicalUsageDataPoint adds a data point to process.memory.physical_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryPhysicalUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryPhysicalUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryVirtualUsageDataPoint adds a data point to process.memory.virtual_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryVirtualUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryVirtualUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessThreadsDataPoint adds a data point to process.threads metric.
func (mb *MetricsBuilder) RecordProcessThreadsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessThreads.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
