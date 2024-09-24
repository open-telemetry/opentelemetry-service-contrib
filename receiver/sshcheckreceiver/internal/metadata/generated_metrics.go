// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type metricSshcheckDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.duration metric with initial data.
func (m *metricSshcheckDuration) init() {
	m.data.SetName("sshcheck.duration")
	m.data.SetDescription("Measures the duration of SSH connection.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricSshcheckDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckDuration) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckDuration(cfg MetricConfig) metricSshcheckDuration {
	m := metricSshcheckDuration{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckError struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.error metric with initial data.
func (m *metricSshcheckError) init() {
	m.data.SetName("sshcheck.error")
	m.data.SetDescription("Records errors occurring during SSH check.")
	m.data.SetUnit("{error}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSshcheckError) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("error.message", errorMessageAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckError) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckError) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckError(cfg MetricConfig) metricSshcheckError {
	m := metricSshcheckError{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_duration metric with initial data.
func (m *metricSshcheckSftpDuration) init() {
	m.data.SetName("sshcheck.sftp_duration")
	m.data.SetDescription("Measures SFTP request duration.")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
}

func (m *metricSshcheckSftpDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpDuration) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpDuration(cfg MetricConfig) metricSshcheckSftpDuration {
	m := metricSshcheckSftpDuration{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpError struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_error metric with initial data.
func (m *metricSshcheckSftpError) init() {
	m.data.SetName("sshcheck.sftp_error")
	m.data.SetDescription("Records errors occurring during SFTP check.")
	m.data.SetUnit("{error}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSshcheckSftpError) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("error.message", errorMessageAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpError) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpError) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpError(cfg MetricConfig) metricSshcheckSftpError {
	m := metricSshcheckSftpError{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckSftpStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.sftp_status metric with initial data.
func (m *metricSshcheckSftpStatus) init() {
	m.data.SetName("sshcheck.sftp_status")
	m.data.SetDescription("1 if the SFTP server replied to request, otherwise 0.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSshcheckSftpStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckSftpStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckSftpStatus) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckSftpStatus(cfg MetricConfig) metricSshcheckSftpStatus {
	m := metricSshcheckSftpStatus{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSshcheckStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills sshcheck.status metric with initial data.
func (m *metricSshcheckStatus) init() {
	m.data.SetName("sshcheck.status")
	m.data.SetDescription("1 if the SSH client successfully connected, otherwise 0.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSshcheckStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSshcheckStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSshcheckStatus) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSshcheckStatus(cfg MetricConfig) metricSshcheckStatus {
	m := metricSshcheckStatus{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                         MetricsBuilderConfig // config of the metrics builder.
	startTime                      pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                int                  // maximum observed number of metrics per resource.
	metricsBuffer                  pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo  // contains version information.
	resourceAttributeIncludeFilter map[string]filter.Filter
	resourceAttributeExcludeFilter map[string]filter.Filter
	metricSshcheckDuration         metricSshcheckDuration
	metricSshcheckError            metricSshcheckError
	metricSshcheckSftpDuration     metricSshcheckSftpDuration
	metricSshcheckSftpError        metricSshcheckSftpError
	metricSshcheckSftpStatus       metricSshcheckSftpStatus
	metricSshcheckStatus           metricSshcheckStatus
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                         mbc,
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      settings.BuildInfo,
		metricSshcheckDuration:         newMetricSshcheckDuration(mbc.Metrics.SshcheckDuration),
		metricSshcheckError:            newMetricSshcheckError(mbc.Metrics.SshcheckError),
		metricSshcheckSftpDuration:     newMetricSshcheckSftpDuration(mbc.Metrics.SshcheckSftpDuration),
		metricSshcheckSftpError:        newMetricSshcheckSftpError(mbc.Metrics.SshcheckSftpError),
		metricSshcheckSftpStatus:       newMetricSshcheckSftpStatus(mbc.Metrics.SshcheckSftpStatus),
		metricSshcheckStatus:           newMetricSshcheckStatus(mbc.Metrics.SshcheckStatus),
		resourceAttributeIncludeFilter: make(map[string]filter.Filter),
		resourceAttributeExcludeFilter: make(map[string]filter.Filter),
	}
	if mbc.ResourceAttributes.SSHEndpoint.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["ssh.endpoint"] = filter.CreateFilter(mbc.ResourceAttributes.SSHEndpoint.MetricsInclude)
	}
	if mbc.ResourceAttributes.SSHEndpoint.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["ssh.endpoint"] = filter.CreateFilter(mbc.ResourceAttributes.SSHEndpoint.MetricsExclude)
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
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
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSshcheckDuration.emit(ils.Metrics())
	mb.metricSshcheckError.emit(ils.Metrics())
	mb.metricSshcheckSftpDuration.emit(ils.Metrics())
	mb.metricSshcheckSftpError.emit(ils.Metrics())
	mb.metricSshcheckSftpStatus.emit(ils.Metrics())
	mb.metricSshcheckStatus.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}
	for attr, filter := range mb.resourceAttributeIncludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range mb.resourceAttributeExcludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSshcheckDurationDataPoint adds a data point to sshcheck.duration metric.
func (mb *MetricsBuilder) RecordSshcheckDurationDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckDuration.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckErrorDataPoint adds a data point to sshcheck.error metric.
func (mb *MetricsBuilder) RecordSshcheckErrorDataPoint(ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	mb.metricSshcheckError.recordDataPoint(mb.startTime, ts, val, errorMessageAttributeValue)
}

// RecordSshcheckSftpDurationDataPoint adds a data point to sshcheck.sftp_duration metric.
func (mb *MetricsBuilder) RecordSshcheckSftpDurationDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckSftpDuration.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckSftpErrorDataPoint adds a data point to sshcheck.sftp_error metric.
func (mb *MetricsBuilder) RecordSshcheckSftpErrorDataPoint(ts pcommon.Timestamp, val int64, errorMessageAttributeValue string) {
	mb.metricSshcheckSftpError.recordDataPoint(mb.startTime, ts, val, errorMessageAttributeValue)
}

// RecordSshcheckSftpStatusDataPoint adds a data point to sshcheck.sftp_status metric.
func (mb *MetricsBuilder) RecordSshcheckSftpStatusDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckSftpStatus.recordDataPoint(mb.startTime, ts, val)
}

// RecordSshcheckStatusDataPoint adds a data point to sshcheck.status metric.
func (mb *MetricsBuilder) RecordSshcheckStatusDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSshcheckStatus.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
