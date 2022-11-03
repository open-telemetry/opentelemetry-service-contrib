// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
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

// MetricsSettings provides settings for hostmetricsreceiver/process metrics.
type MetricsSettings struct {
	ProcessContextSwitches     MetricSettings `mapstructure:"process.context_switches"`
	ProcessCPUTime             MetricSettings `mapstructure:"process.cpu.time"`
	ProcessDiskIo              MetricSettings `mapstructure:"process.disk.io"`
	ProcessMemoryPhysicalUsage MetricSettings `mapstructure:"process.memory.physical_usage"`
	ProcessMemoryVirtualUsage  MetricSettings `mapstructure:"process.memory.virtual_usage"`
	ProcessOpenFileDescriptors MetricSettings `mapstructure:"process.open_file_descriptors"`
	ProcessPagingFaults        MetricSettings `mapstructure:"process.paging.faults"`
	ProcessThreads             MetricSettings `mapstructure:"process.threads"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ProcessContextSwitches: MetricSettings{
			Enabled: false,
		},
		ProcessCPUTime: MetricSettings{
			Enabled: true,
		},
		ProcessDiskIo: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryPhysicalUsage: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryVirtualUsage: MetricSettings{
			Enabled: true,
		},
		ProcessOpenFileDescriptors: MetricSettings{
			Enabled: false,
		},
		ProcessPagingFaults: MetricSettings{
			Enabled: false,
		},
		ProcessThreads: MetricSettings{
			Enabled: false,
		},
	}
}

// AttributeContextSwitchType specifies the a value context_switch_type attribute.
type AttributeContextSwitchType int

const (
	_ AttributeContextSwitchType = iota
	AttributeContextSwitchTypeInvoluntary
	AttributeContextSwitchTypeVoluntary
)

// String returns the string representation of the AttributeContextSwitchType.
func (av AttributeContextSwitchType) String() string {
	switch av {
	case AttributeContextSwitchTypeInvoluntary:
		return "involuntary"
	case AttributeContextSwitchTypeVoluntary:
		return "voluntary"
	}
	return ""
}

// MapAttributeContextSwitchType is a helper map of string to AttributeContextSwitchType attribute value.
var MapAttributeContextSwitchType = map[string]AttributeContextSwitchType{
	"involuntary": AttributeContextSwitchTypeInvoluntary,
	"voluntary":   AttributeContextSwitchTypeVoluntary,
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

// AttributePagingFaultType specifies the a value paging_fault_type attribute.
type AttributePagingFaultType int

const (
	_ AttributePagingFaultType = iota
	AttributePagingFaultTypeMajor
	AttributePagingFaultTypeMinor
)

// String returns the string representation of the AttributePagingFaultType.
func (av AttributePagingFaultType) String() string {
	switch av {
	case AttributePagingFaultTypeMajor:
		return "major"
	case AttributePagingFaultTypeMinor:
		return "minor"
	}
	return ""
}

// MapAttributePagingFaultType is a helper map of string to AttributePagingFaultType attribute value.
var MapAttributePagingFaultType = map[string]AttributePagingFaultType{
	"major": AttributePagingFaultTypeMajor,
	"minor": AttributePagingFaultTypeMinor,
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

type metricProcessContextSwitches struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.context_switches metric with initial data.
func (m *metricProcessContextSwitches) init() {
	m.data.SetName("process.context_switches")
	m.data.SetDescription("Number of times the process has been context switched.")
	m.data.SetUnit("{count}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessContextSwitches) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, contextSwitchTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", contextSwitchTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessContextSwitches) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessContextSwitches) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessContextSwitches(settings MetricSettings) metricProcessContextSwitches {
	m := metricProcessContextSwitches{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
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
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
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
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessDiskIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
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

type metricProcessMemoryPhysicalUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.physical_usage metric with initial data.
func (m *metricProcessMemoryPhysicalUsage) init() {
	m.data.SetName("process.memory.physical_usage")
	m.data.SetDescription("The amount of physical memory in use")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessMemoryPhysicalUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
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
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessMemoryVirtualUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
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

type metricProcessOpenFileDescriptors struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.open_file_descriptors metric with initial data.
func (m *metricProcessOpenFileDescriptors) init() {
	m.data.SetName("process.open_file_descriptors")
	m.data.SetDescription("Number of file descriptors in use by the process.")
	m.data.SetUnit("{count}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessOpenFileDescriptors) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessOpenFileDescriptors) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessOpenFileDescriptors) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessOpenFileDescriptors(settings MetricSettings) metricProcessOpenFileDescriptors {
	m := metricProcessOpenFileDescriptors{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessPagingFaults struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.paging.faults metric with initial data.
func (m *metricProcessPagingFaults) init() {
	m.data.SetName("process.paging.faults")
	m.data.SetDescription("Number of page faults the process has made. This metric is only available on Linux.")
	m.data.SetUnit("{faults}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessPagingFaults) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, pagingFaultTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", pagingFaultTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessPagingFaults) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessPagingFaults) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessPagingFaults(settings MetricSettings) metricProcessPagingFaults {
	m := metricProcessPagingFaults{settings: settings}
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
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessThreads) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
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
	metricProcessContextSwitches     metricProcessContextSwitches
	metricProcessCPUTime             metricProcessCPUTime
	metricProcessDiskIo              metricProcessDiskIo
	metricProcessMemoryPhysicalUsage metricProcessMemoryPhysicalUsage
	metricProcessMemoryVirtualUsage  metricProcessMemoryVirtualUsage
	metricProcessOpenFileDescriptors metricProcessOpenFileDescriptors
	metricProcessPagingFaults        metricProcessPagingFaults
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
		metricProcessContextSwitches:     newMetricProcessContextSwitches(settings.ProcessContextSwitches),
		metricProcessCPUTime:             newMetricProcessCPUTime(settings.ProcessCPUTime),
		metricProcessDiskIo:              newMetricProcessDiskIo(settings.ProcessDiskIo),
		metricProcessMemoryPhysicalUsage: newMetricProcessMemoryPhysicalUsage(settings.ProcessMemoryPhysicalUsage),
		metricProcessMemoryVirtualUsage:  newMetricProcessMemoryVirtualUsage(settings.ProcessMemoryVirtualUsage),
		metricProcessOpenFileDescriptors: newMetricProcessOpenFileDescriptors(settings.ProcessOpenFileDescriptors),
		metricProcessPagingFaults:        newMetricProcessPagingFaults(settings.ProcessPagingFaults),
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
		rm.Resource().Attributes().PutStr("process.command", val)
	}
}

// WithProcessCommandLine sets provided value as "process.command_line" attribute for current resource.
func WithProcessCommandLine(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("process.command_line", val)
	}
}

// WithProcessExecutableName sets provided value as "process.executable.name" attribute for current resource.
func WithProcessExecutableName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("process.executable.name", val)
	}
}

// WithProcessExecutablePath sets provided value as "process.executable.path" attribute for current resource.
func WithProcessExecutablePath(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("process.executable.path", val)
	}
}

// WithProcessOwner sets provided value as "process.owner" attribute for current resource.
func WithProcessOwner(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("process.owner", val)
	}
}

// WithProcessParentPid sets provided value as "process.parent_pid" attribute for current resource.
func WithProcessParentPid(val int64) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutInt("process.parent_pid", val)
	}
}

// WithProcessPid sets provided value as "process.pid" attribute for current resource.
func WithProcessPid(val int64) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutInt("process.pid", val)
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
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/process")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricProcessContextSwitches.emit(ils.Metrics())
	mb.metricProcessCPUTime.emit(ils.Metrics())
	mb.metricProcessDiskIo.emit(ils.Metrics())
	mb.metricProcessMemoryPhysicalUsage.emit(ils.Metrics())
	mb.metricProcessMemoryVirtualUsage.emit(ils.Metrics())
	mb.metricProcessOpenFileDescriptors.emit(ils.Metrics())
	mb.metricProcessPagingFaults.emit(ils.Metrics())
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

// RecordProcessContextSwitchesDataPoint adds a data point to process.context_switches metric.
func (mb *MetricsBuilder) RecordProcessContextSwitchesDataPoint(ts pcommon.Timestamp, val int64, contextSwitchTypeAttributeValue AttributeContextSwitchType) {
	mb.metricProcessContextSwitches.recordDataPoint(mb.startTime, ts, val, contextSwitchTypeAttributeValue.String())
}

// RecordProcessCPUTimeDataPoint adds a data point to process.cpu.time metric.
func (mb *MetricsBuilder) RecordProcessCPUTimeDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricProcessCPUTime.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordProcessDiskIoDataPoint adds a data point to process.disk.io metric.
func (mb *MetricsBuilder) RecordProcessDiskIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricProcessDiskIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordProcessMemoryPhysicalUsageDataPoint adds a data point to process.memory.physical_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryPhysicalUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryPhysicalUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryVirtualUsageDataPoint adds a data point to process.memory.virtual_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryVirtualUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryVirtualUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessOpenFileDescriptorsDataPoint adds a data point to process.open_file_descriptors metric.
func (mb *MetricsBuilder) RecordProcessOpenFileDescriptorsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessOpenFileDescriptors.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessPagingFaultsDataPoint adds a data point to process.paging.faults metric.
func (mb *MetricsBuilder) RecordProcessPagingFaultsDataPoint(ts pcommon.Timestamp, val int64, pagingFaultTypeAttributeValue AttributePagingFaultType) {
	mb.metricProcessPagingFaults.recordDataPoint(mb.startTime, ts, val, pagingFaultTypeAttributeValue.String())
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
