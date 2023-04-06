// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"fmt"
	"strconv"
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

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for mesosreceiver metrics.
type MetricsSettings struct {
	MesosCPUUtilization          MetricSettings `mapstructure:"mesos.cpu.utilization"`
	MesosGpuUtilization          MetricSettings `mapstructure:"mesos.gpu.utilization"`
	MesosMemLimit                MetricSettings `mapstructure:"mesos.mem.limit"`
	MesosMemUtilization          MetricSettings `mapstructure:"mesos.mem.utilization"`
	MesosSlavesActiveCount       MetricSettings `mapstructure:"mesos.slaves.active.count"`
	MesosSlavesConnectedCount    MetricSettings `mapstructure:"mesos.slaves.connected.count"`
	MesosSlavesDisconnectedCount MetricSettings `mapstructure:"mesos.slaves.disconnected.count"`
	MesosSlavesInactiveCount     MetricSettings `mapstructure:"mesos.slaves.inactive.count"`
	MesosTasksFailedCount        MetricSettings `mapstructure:"mesos.tasks.failed.count"`
	MesosTasksFinishedCount      MetricSettings `mapstructure:"mesos.tasks.finished.count"`
	MesosUptime                  MetricSettings `mapstructure:"mesos.uptime"`
	SystemLoad15m                MetricSettings `mapstructure:"system.load.15m"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		MesosCPUUtilization: MetricSettings{
			Enabled: true,
		},
		MesosGpuUtilization: MetricSettings{
			Enabled: true,
		},
		MesosMemLimit: MetricSettings{
			Enabled: true,
		},
		MesosMemUtilization: MetricSettings{
			Enabled: true,
		},
		MesosSlavesActiveCount: MetricSettings{
			Enabled: true,
		},
		MesosSlavesConnectedCount: MetricSettings{
			Enabled: true,
		},
		MesosSlavesDisconnectedCount: MetricSettings{
			Enabled: true,
		},
		MesosSlavesInactiveCount: MetricSettings{
			Enabled: true,
		},
		MesosTasksFailedCount: MetricSettings{
			Enabled: true,
		},
		MesosTasksFinishedCount: MetricSettings{
			Enabled: true,
		},
		MesosUptime: MetricSettings{
			Enabled: true,
		},
		SystemLoad15m: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular metric.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesSettings provides settings for mesosreceiver metrics.
type ResourceAttributesSettings struct {
	MesosServerName ResourceAttributeSettings `mapstructure:"mesos.server.name"`
	MesosServerPort ResourceAttributeSettings `mapstructure:"mesos.server.port"`
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{
		MesosServerName: ResourceAttributeSettings{
			Enabled: false,
		},
		MesosServerPort: ResourceAttributeSettings{
			Enabled: false,
		},
	}
}

type metricMesosCPUUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.cpu.utilization metric with initial data.
func (m *metricMesosCPUUtilization) init() {
	m.data.SetName("mesos.cpu.utilization")
	m.data.SetDescription("Master percentage of allocated CPUs expressed as ratio [0-1].")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricMesosCPUUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosCPUUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosCPUUtilization(settings MetricSettings) metricMesosCPUUtilization {
	m := metricMesosCPUUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosGpuUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.gpu.utilization metric with initial data.
func (m *metricMesosGpuUtilization) init() {
	m.data.SetName("mesos.gpu.utilization")
	m.data.SetDescription("Master percentage of allocated GPUs expressed as ratio [0-1].")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricMesosGpuUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosGpuUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosGpuUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosGpuUtilization(settings MetricSettings) metricMesosGpuUtilization {
	m := metricMesosGpuUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosMemLimit struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.mem.limit metric with initial data.
func (m *metricMesosMemLimit) init() {
	m.data.SetName("mesos.mem.limit")
	m.data.SetDescription("Master total memory in bytes.")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricMesosMemLimit) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosMemLimit) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosMemLimit) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosMemLimit(settings MetricSettings) metricMesosMemLimit {
	m := metricMesosMemLimit{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosMemUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.mem.utilization metric with initial data.
func (m *metricMesosMemUtilization) init() {
	m.data.SetName("mesos.mem.utilization")
	m.data.SetDescription("Master percentage of allocated memory expressed as ratio [0-1].")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricMesosMemUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosMemUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosMemUtilization) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosMemUtilization(settings MetricSettings) metricMesosMemUtilization {
	m := metricMesosMemUtilization{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosSlavesActiveCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.slaves.active.count metric with initial data.
func (m *metricMesosSlavesActiveCount) init() {
	m.data.SetName("mesos.slaves.active.count")
	m.data.SetDescription("The number of active agents/slaves running for the master.")
	m.data.SetUnit("{Slave}")
	m.data.SetEmptyGauge()
}

func (m *metricMesosSlavesActiveCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosSlavesActiveCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosSlavesActiveCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosSlavesActiveCount(settings MetricSettings) metricMesosSlavesActiveCount {
	m := metricMesosSlavesActiveCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosSlavesConnectedCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.slaves.connected.count metric with initial data.
func (m *metricMesosSlavesConnectedCount) init() {
	m.data.SetName("mesos.slaves.connected.count")
	m.data.SetDescription("The number of agents/slaves connected to the master.")
	m.data.SetUnit("{Slave}")
	m.data.SetEmptyGauge()
}

func (m *metricMesosSlavesConnectedCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosSlavesConnectedCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosSlavesConnectedCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosSlavesConnectedCount(settings MetricSettings) metricMesosSlavesConnectedCount {
	m := metricMesosSlavesConnectedCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosSlavesDisconnectedCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.slaves.disconnected.count metric with initial data.
func (m *metricMesosSlavesDisconnectedCount) init() {
	m.data.SetName("mesos.slaves.disconnected.count")
	m.data.SetDescription("The number of agents/slaves disconnected from the master")
	m.data.SetUnit("{Slave}")
	m.data.SetEmptyGauge()
}

func (m *metricMesosSlavesDisconnectedCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosSlavesDisconnectedCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosSlavesDisconnectedCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosSlavesDisconnectedCount(settings MetricSettings) metricMesosSlavesDisconnectedCount {
	m := metricMesosSlavesDisconnectedCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosSlavesInactiveCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.slaves.inactive.count metric with initial data.
func (m *metricMesosSlavesInactiveCount) init() {
	m.data.SetName("mesos.slaves.inactive.count")
	m.data.SetDescription("The number of agents/slaves that are inactive.")
	m.data.SetUnit("{Slave}")
	m.data.SetEmptyGauge()
}

func (m *metricMesosSlavesInactiveCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosSlavesInactiveCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosSlavesInactiveCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosSlavesInactiveCount(settings MetricSettings) metricMesosSlavesInactiveCount {
	m := metricMesosSlavesInactiveCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosTasksFailedCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.tasks.failed.count metric with initial data.
func (m *metricMesosTasksFailedCount) init() {
	m.data.SetName("mesos.tasks.failed.count")
	m.data.SetDescription("The number of tasks that were failed.")
	m.data.SetUnit("{Task}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMesosTasksFailedCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosTasksFailedCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosTasksFailedCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosTasksFailedCount(settings MetricSettings) metricMesosTasksFailedCount {
	m := metricMesosTasksFailedCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosTasksFinishedCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.tasks.finished.count metric with initial data.
func (m *metricMesosTasksFinishedCount) init() {
	m.data.SetName("mesos.tasks.finished.count")
	m.data.SetDescription("The number of tasks that were finished.")
	m.data.SetUnit("{Task}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMesosTasksFinishedCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosTasksFinishedCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosTasksFinishedCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosTasksFinishedCount(settings MetricSettings) metricMesosTasksFinishedCount {
	m := metricMesosTasksFinishedCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMesosUptime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mesos.uptime metric with initial data.
func (m *metricMesosUptime) init() {
	m.data.SetName("mesos.uptime")
	m.data.SetDescription("Time master has been up in seconds.")
	m.data.SetUnit("s")
	m.data.SetEmptyGauge()
}

func (m *metricMesosUptime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMesosUptime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMesosUptime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMesosUptime(settings MetricSettings) metricMesosUptime {
	m := metricMesosUptime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemLoad15m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.load.15m metric with initial data.
func (m *metricSystemLoad15m) init() {
	m.data.SetName("system.load.15m")
	m.data.SetDescription("the average system load during the last 15 minutes expressed as ratio [0-1].")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricSystemLoad15m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemLoad15m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemLoad15m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemLoad15m(settings MetricSettings) metricSystemLoad15m {
	m := metricSystemLoad15m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:"resource_attributes"`
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                          pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                    int                 // maximum observed number of metrics per resource.
	resourceCapacity                   int                 // maximum observed number of resource attributes.
	metricsBuffer                      pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                          component.BuildInfo // contains version information
	resourceAttributesSettings         ResourceAttributesSettings
	metricMesosCPUUtilization          metricMesosCPUUtilization
	metricMesosGpuUtilization          metricMesosGpuUtilization
	metricMesosMemLimit                metricMesosMemLimit
	metricMesosMemUtilization          metricMesosMemUtilization
	metricMesosSlavesActiveCount       metricMesosSlavesActiveCount
	metricMesosSlavesConnectedCount    metricMesosSlavesConnectedCount
	metricMesosSlavesDisconnectedCount metricMesosSlavesDisconnectedCount
	metricMesosSlavesInactiveCount     metricMesosSlavesInactiveCount
	metricMesosTasksFailedCount        metricMesosTasksFailedCount
	metricMesosTasksFinishedCount      metricMesosTasksFinishedCount
	metricMesosUptime                  metricMesosUptime
	metricSystemLoad15m                metricSystemLoad15m
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsSettings(),
		ResourceAttributes: DefaultResourceAttributesSettings(),
	}
}

func NewMetricsBuilderConfig(ms MetricsSettings, ras ResourceAttributesSettings) MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            ms,
		ResourceAttributes: ras,
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                          pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                      pmetric.NewMetrics(),
		buildInfo:                          settings.BuildInfo,
		resourceAttributesSettings:         mbc.ResourceAttributes,
		metricMesosCPUUtilization:          newMetricMesosCPUUtilization(mbc.Metrics.MesosCPUUtilization),
		metricMesosGpuUtilization:          newMetricMesosGpuUtilization(mbc.Metrics.MesosGpuUtilization),
		metricMesosMemLimit:                newMetricMesosMemLimit(mbc.Metrics.MesosMemLimit),
		metricMesosMemUtilization:          newMetricMesosMemUtilization(mbc.Metrics.MesosMemUtilization),
		metricMesosSlavesActiveCount:       newMetricMesosSlavesActiveCount(mbc.Metrics.MesosSlavesActiveCount),
		metricMesosSlavesConnectedCount:    newMetricMesosSlavesConnectedCount(mbc.Metrics.MesosSlavesConnectedCount),
		metricMesosSlavesDisconnectedCount: newMetricMesosSlavesDisconnectedCount(mbc.Metrics.MesosSlavesDisconnectedCount),
		metricMesosSlavesInactiveCount:     newMetricMesosSlavesInactiveCount(mbc.Metrics.MesosSlavesInactiveCount),
		metricMesosTasksFailedCount:        newMetricMesosTasksFailedCount(mbc.Metrics.MesosTasksFailedCount),
		metricMesosTasksFinishedCount:      newMetricMesosTasksFinishedCount(mbc.Metrics.MesosTasksFinishedCount),
		metricMesosUptime:                  newMetricMesosUptime(mbc.Metrics.MesosUptime),
		metricSystemLoad15m:                newMetricSystemLoad15m(mbc.Metrics.SystemLoad15m),
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
type ResourceMetricsOption func(ResourceAttributesSettings, pmetric.ResourceMetrics)

// WithMesosServerName sets provided value as "mesos.server.name" attribute for current resource.
func WithMesosServerName(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.MesosServerName.Enabled {
			rm.Resource().Attributes().PutStr("mesos.server.name", val)
		}
	}
}

// WithMesosServerPort sets provided value as "mesos.server.port" attribute for current resource.
func WithMesosServerPort(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.MesosServerPort.Enabled {
			rm.Resource().Attributes().PutStr("mesos.server.port", val)
		}
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
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
	ils.Scope().SetName("otelcol/mesosreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricMesosCPUUtilization.emit(ils.Metrics())
	mb.metricMesosGpuUtilization.emit(ils.Metrics())
	mb.metricMesosMemLimit.emit(ils.Metrics())
	mb.metricMesosMemUtilization.emit(ils.Metrics())
	mb.metricMesosSlavesActiveCount.emit(ils.Metrics())
	mb.metricMesosSlavesConnectedCount.emit(ils.Metrics())
	mb.metricMesosSlavesDisconnectedCount.emit(ils.Metrics())
	mb.metricMesosSlavesInactiveCount.emit(ils.Metrics())
	mb.metricMesosTasksFailedCount.emit(ils.Metrics())
	mb.metricMesosTasksFinishedCount.emit(ils.Metrics())
	mb.metricMesosUptime.emit(ils.Metrics())
	mb.metricSystemLoad15m.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesSettings, rm)
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
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordMesosCPUUtilizationDataPoint adds a data point to mesos.cpu.utilization metric.
func (mb *MetricsBuilder) RecordMesosCPUUtilizationDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosCPUUtilization, value was %s: %w", inputVal, err)
	}
	mb.metricMesosCPUUtilization.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosGpuUtilizationDataPoint adds a data point to mesos.gpu.utilization metric.
func (mb *MetricsBuilder) RecordMesosGpuUtilizationDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosGpuUtilization, value was %s: %w", inputVal, err)
	}
	mb.metricMesosGpuUtilization.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosMemLimitDataPoint adds a data point to mesos.mem.limit metric.
func (mb *MetricsBuilder) RecordMesosMemLimitDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosMemLimit, value was %s: %w", inputVal, err)
	}
	mb.metricMesosMemLimit.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosMemUtilizationDataPoint adds a data point to mesos.mem.utilization metric.
func (mb *MetricsBuilder) RecordMesosMemUtilizationDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosMemUtilization, value was %s: %w", inputVal, err)
	}
	mb.metricMesosMemUtilization.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosSlavesActiveCountDataPoint adds a data point to mesos.slaves.active.count metric.
func (mb *MetricsBuilder) RecordMesosSlavesActiveCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosSlavesActiveCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosSlavesActiveCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosSlavesConnectedCountDataPoint adds a data point to mesos.slaves.connected.count metric.
func (mb *MetricsBuilder) RecordMesosSlavesConnectedCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosSlavesConnectedCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosSlavesConnectedCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosSlavesDisconnectedCountDataPoint adds a data point to mesos.slaves.disconnected.count metric.
func (mb *MetricsBuilder) RecordMesosSlavesDisconnectedCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosSlavesDisconnectedCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosSlavesDisconnectedCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosSlavesInactiveCountDataPoint adds a data point to mesos.slaves.inactive.count metric.
func (mb *MetricsBuilder) RecordMesosSlavesInactiveCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosSlavesInactiveCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosSlavesInactiveCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosTasksFailedCountDataPoint adds a data point to mesos.tasks.failed.count metric.
func (mb *MetricsBuilder) RecordMesosTasksFailedCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosTasksFailedCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosTasksFailedCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosTasksFinishedCountDataPoint adds a data point to mesos.tasks.finished.count metric.
func (mb *MetricsBuilder) RecordMesosTasksFinishedCountDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosTasksFinishedCount, value was %s: %w", inputVal, err)
	}
	mb.metricMesosTasksFinishedCount.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordMesosUptimeDataPoint adds a data point to mesos.uptime metric.
func (mb *MetricsBuilder) RecordMesosUptimeDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for MesosUptime, value was %s: %w", inputVal, err)
	}
	mb.metricMesosUptime.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordSystemLoad15mDataPoint adds a data point to system.load.15m metric.
func (mb *MetricsBuilder) RecordSystemLoad15mDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float64 for SystemLoad15m, value was %s: %w", inputVal, err)
	}
	mb.metricSystemLoad15m.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
