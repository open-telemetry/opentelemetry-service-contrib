// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.18.0"
)

type metricK8sPodPhase struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.phase metric with initial data.
func (m *metricK8sPodPhase) init() {
	m.data.SetName("k8s.pod.phase")
	m.data.SetDescription("Current phase of the pod (1 - Pending, 2 - Running, 3 - Succeeded, 4 - Failed, 5 - Unknown)")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodPhase) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodPhase) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodPhase) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodPhase(cfg MetricConfig) metricK8sPodPhase {
	m := metricK8sPodPhase{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sPodStatusReasonEvicted struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.status_reason_evicted metric with initial data.
func (m *metricK8sPodStatusReasonEvicted) init() {
	m.data.SetName("k8s.pod.status_reason_evicted")
	m.data.SetDescription("Whether this pod status reason is Evicted (1), or not (0).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodStatusReasonEvicted) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodStatusReasonEvicted) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodStatusReasonEvicted) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodStatusReasonEvicted(cfg MetricConfig) metricK8sPodStatusReasonEvicted {
	m := metricK8sPodStatusReasonEvicted{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sPodStatusReasonNodeAffinity struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.status_reason_node_affinity metric with initial data.
func (m *metricK8sPodStatusReasonNodeAffinity) init() {
	m.data.SetName("k8s.pod.status_reason_node_affinity")
	m.data.SetDescription("Whether this pod status reason is NodeAffinity (1), or not (0).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodStatusReasonNodeAffinity) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodStatusReasonNodeAffinity) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodStatusReasonNodeAffinity) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodStatusReasonNodeAffinity(cfg MetricConfig) metricK8sPodStatusReasonNodeAffinity {
	m := metricK8sPodStatusReasonNodeAffinity{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sPodStatusReasonNodeLost struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.status_reason_node_lost metric with initial data.
func (m *metricK8sPodStatusReasonNodeLost) init() {
	m.data.SetName("k8s.pod.status_reason_node_lost")
	m.data.SetDescription("Whether this pod status reason is NodeLost (1), or not (0).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodStatusReasonNodeLost) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodStatusReasonNodeLost) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodStatusReasonNodeLost) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodStatusReasonNodeLost(cfg MetricConfig) metricK8sPodStatusReasonNodeLost {
	m := metricK8sPodStatusReasonNodeLost{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sPodStatusReasonShutdown struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.status_reason_shutdown metric with initial data.
func (m *metricK8sPodStatusReasonShutdown) init() {
	m.data.SetName("k8s.pod.status_reason_shutdown")
	m.data.SetDescription("Whether this pod status reason is Shutdown (1), or not (0).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodStatusReasonShutdown) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodStatusReasonShutdown) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodStatusReasonShutdown) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodStatusReasonShutdown(cfg MetricConfig) metricK8sPodStatusReasonShutdown {
	m := metricK8sPodStatusReasonShutdown{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricK8sPodStatusReasonUnexpectedAdmissionError struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills k8s.pod.status_reason_unexpected_admission_error metric with initial data.
func (m *metricK8sPodStatusReasonUnexpectedAdmissionError) init() {
	m.data.SetName("k8s.pod.status_reason_unexpected_admission_error")
	m.data.SetDescription("Whether this pod status reason is Unexpected Admission Error (1), or not (0).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricK8sPodStatusReasonUnexpectedAdmissionError) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricK8sPodStatusReasonUnexpectedAdmissionError) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricK8sPodStatusReasonUnexpectedAdmissionError) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricK8sPodStatusReasonUnexpectedAdmissionError(cfg MetricConfig) metricK8sPodStatusReasonUnexpectedAdmissionError {
	m := metricK8sPodStatusReasonUnexpectedAdmissionError{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	startTime                                        pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                                  int                 // maximum observed number of metrics per resource.
	resourceCapacity                                 int                 // maximum observed number of resource attributes.
	metricsBuffer                                    pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                                        component.BuildInfo // contains version information
	resourceAttributesConfig                         ResourceAttributesConfig
	metricK8sPodPhase                                metricK8sPodPhase
	metricK8sPodStatusReasonEvicted                  metricK8sPodStatusReasonEvicted
	metricK8sPodStatusReasonNodeAffinity             metricK8sPodStatusReasonNodeAffinity
	metricK8sPodStatusReasonNodeLost                 metricK8sPodStatusReasonNodeLost
	metricK8sPodStatusReasonShutdown                 metricK8sPodStatusReasonShutdown
	metricK8sPodStatusReasonUnexpectedAdmissionError metricK8sPodStatusReasonUnexpectedAdmissionError
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                                        pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                                    pmetric.NewMetrics(),
		buildInfo:                                        settings.BuildInfo,
		resourceAttributesConfig:                         mbc.ResourceAttributes,
		metricK8sPodPhase:                                newMetricK8sPodPhase(mbc.Metrics.K8sPodPhase),
		metricK8sPodStatusReasonEvicted:                  newMetricK8sPodStatusReasonEvicted(mbc.Metrics.K8sPodStatusReasonEvicted),
		metricK8sPodStatusReasonNodeAffinity:             newMetricK8sPodStatusReasonNodeAffinity(mbc.Metrics.K8sPodStatusReasonNodeAffinity),
		metricK8sPodStatusReasonNodeLost:                 newMetricK8sPodStatusReasonNodeLost(mbc.Metrics.K8sPodStatusReasonNodeLost),
		metricK8sPodStatusReasonShutdown:                 newMetricK8sPodStatusReasonShutdown(mbc.Metrics.K8sPodStatusReasonShutdown),
		metricK8sPodStatusReasonUnexpectedAdmissionError: newMetricK8sPodStatusReasonUnexpectedAdmissionError(mbc.Metrics.K8sPodStatusReasonUnexpectedAdmissionError),
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
type ResourceMetricsOption func(ResourceAttributesConfig, pmetric.ResourceMetrics)

// WithK8sNamespaceName sets provided value as "k8s.namespace.name" attribute for current resource.
func WithK8sNamespaceName(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sNamespaceName.Enabled {
			rm.Resource().Attributes().PutStr("k8s.namespace.name", val)
		}
	}
}

// WithK8sNodeName sets provided value as "k8s.node.name" attribute for current resource.
func WithK8sNodeName(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sNodeName.Enabled {
			rm.Resource().Attributes().PutStr("k8s.node.name", val)
		}
	}
}

// WithK8sPodName sets provided value as "k8s.pod.name" attribute for current resource.
func WithK8sPodName(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sPodName.Enabled {
			rm.Resource().Attributes().PutStr("k8s.pod.name", val)
		}
	}
}

// WithK8sPodUID sets provided value as "k8s.pod.uid" attribute for current resource.
func WithK8sPodUID(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.K8sPodUID.Enabled {
			rm.Resource().Attributes().PutStr("k8s.pod.uid", val)
		}
	}
}

// WithOpencensusResourcetype sets provided value as "opencensus.resourcetype" attribute for current resource.
func WithOpencensusResourcetype(val string) ResourceMetricsOption {
	return func(rac ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
		if rac.OpencensusResourcetype.Enabled {
			rm.Resource().Attributes().PutStr("opencensus.resourcetype", val)
		}
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(_ ResourceAttributesConfig, rm pmetric.ResourceMetrics) {
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
	ils.Scope().SetName("otelcol/k8sclusterreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricK8sPodPhase.emit(ils.Metrics())
	mb.metricK8sPodStatusReasonEvicted.emit(ils.Metrics())
	mb.metricK8sPodStatusReasonNodeAffinity.emit(ils.Metrics())
	mb.metricK8sPodStatusReasonNodeLost.emit(ils.Metrics())
	mb.metricK8sPodStatusReasonShutdown.emit(ils.Metrics())
	mb.metricK8sPodStatusReasonUnexpectedAdmissionError.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesConfig, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordK8sPodPhaseDataPoint adds a data point to k8s.pod.phase metric.
func (mb *MetricsBuilder) RecordK8sPodPhaseDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodPhase.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sPodStatusReasonEvictedDataPoint adds a data point to k8s.pod.status_reason_evicted metric.
func (mb *MetricsBuilder) RecordK8sPodStatusReasonEvictedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodStatusReasonEvicted.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sPodStatusReasonNodeAffinityDataPoint adds a data point to k8s.pod.status_reason_node_affinity metric.
func (mb *MetricsBuilder) RecordK8sPodStatusReasonNodeAffinityDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodStatusReasonNodeAffinity.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sPodStatusReasonNodeLostDataPoint adds a data point to k8s.pod.status_reason_node_lost metric.
func (mb *MetricsBuilder) RecordK8sPodStatusReasonNodeLostDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodStatusReasonNodeLost.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sPodStatusReasonShutdownDataPoint adds a data point to k8s.pod.status_reason_shutdown metric.
func (mb *MetricsBuilder) RecordK8sPodStatusReasonShutdownDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodStatusReasonShutdown.recordDataPoint(mb.startTime, ts, val)
}

// RecordK8sPodStatusReasonUnexpectedAdmissionErrorDataPoint adds a data point to k8s.pod.status_reason_unexpected_admission_error metric.
func (mb *MetricsBuilder) RecordK8sPodStatusReasonUnexpectedAdmissionErrorDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricK8sPodStatusReasonUnexpectedAdmissionError.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
