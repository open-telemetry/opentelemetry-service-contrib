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

// AttributeMessageState specifies the a value message.state attribute.
type AttributeMessageState int

const (
	_ AttributeMessageState = iota
	AttributeMessageStateReady
	AttributeMessageStateUnacknowledged
)

// String returns the string representation of the AttributeMessageState.
func (av AttributeMessageState) String() string {
	switch av {
	case AttributeMessageStateReady:
		return "ready"
	case AttributeMessageStateUnacknowledged:
		return "unacknowledged"
	}
	return ""
}

// MapAttributeMessageState is a helper map of string to AttributeMessageState attribute value.
var MapAttributeMessageState = map[string]AttributeMessageState{
	"ready":          AttributeMessageStateReady,
	"unacknowledged": AttributeMessageStateUnacknowledged,
}

type metricRabbitmqConsumerCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.consumer.count metric with initial data.
func (m *metricRabbitmqConsumerCount) init() {
	m.data.SetName("rabbitmq.consumer.count")
	m.data.SetDescription("The number of consumers currently reading from the queue.")
	m.data.SetUnit("{consumers}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqConsumerCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqConsumerCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqConsumerCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqConsumerCount(cfg MetricConfig) metricRabbitmqConsumerCount {
	m := metricRabbitmqConsumerCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageAcknowledged struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.acknowledged metric with initial data.
func (m *metricRabbitmqMessageAcknowledged) init() {
	m.data.SetName("rabbitmq.message.acknowledged")
	m.data.SetDescription("The number of messages acknowledged by consumers.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageAcknowledged) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageAcknowledged) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageAcknowledged) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageAcknowledged(cfg MetricConfig) metricRabbitmqMessageAcknowledged {
	m := metricRabbitmqMessageAcknowledged{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageCurrent struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.current metric with initial data.
func (m *metricRabbitmqMessageCurrent) init() {
	m.data.SetName("rabbitmq.message.current")
	m.data.SetDescription("The total number of messages currently in the queue.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRabbitmqMessageCurrent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, messageStateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", messageStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageCurrent) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageCurrent) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageCurrent(cfg MetricConfig) metricRabbitmqMessageCurrent {
	m := metricRabbitmqMessageCurrent{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageDelivered struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.delivered metric with initial data.
func (m *metricRabbitmqMessageDelivered) init() {
	m.data.SetName("rabbitmq.message.delivered")
	m.data.SetDescription("The number of messages delivered to consumers.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageDelivered) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageDelivered) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageDelivered) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageDelivered(cfg MetricConfig) metricRabbitmqMessageDelivered {
	m := metricRabbitmqMessageDelivered{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageDropped struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.dropped metric with initial data.
func (m *metricRabbitmqMessageDropped) init() {
	m.data.SetName("rabbitmq.message.dropped")
	m.data.SetDescription("The number of messages dropped as unroutable.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageDropped) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageDropped) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageDropped) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageDropped(cfg MetricConfig) metricRabbitmqMessageDropped {
	m := metricRabbitmqMessageDropped{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessagePublished struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.published metric with initial data.
func (m *metricRabbitmqMessagePublished) init() {
	m.data.SetName("rabbitmq.message.published")
	m.data.SetDescription("The number of messages published to a queue.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessagePublished) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessagePublished) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessagePublished) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessagePublished(cfg MetricConfig) metricRabbitmqMessagePublished {
	m := metricRabbitmqMessagePublished{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                            MetricsBuilderConfig // config of the metrics builder.
	startTime                         pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                   int                  // maximum observed number of metrics per resource.
	metricsBuffer                     pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo  // contains version information.
	resourceAttributeIncludeFilter    map[string]filter.Filter
	resourceAttributeExcludeFilter    map[string]filter.Filter
	metricRabbitmqConsumerCount       metricRabbitmqConsumerCount
	metricRabbitmqMessageAcknowledged metricRabbitmqMessageAcknowledged
	metricRabbitmqMessageCurrent      metricRabbitmqMessageCurrent
	metricRabbitmqMessageDelivered    metricRabbitmqMessageDelivered
	metricRabbitmqMessageDropped      metricRabbitmqMessageDropped
	metricRabbitmqMessagePublished    metricRabbitmqMessagePublished
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                            mbc,
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		metricRabbitmqConsumerCount:       newMetricRabbitmqConsumerCount(mbc.Metrics.RabbitmqConsumerCount),
		metricRabbitmqMessageAcknowledged: newMetricRabbitmqMessageAcknowledged(mbc.Metrics.RabbitmqMessageAcknowledged),
		metricRabbitmqMessageCurrent:      newMetricRabbitmqMessageCurrent(mbc.Metrics.RabbitmqMessageCurrent),
		metricRabbitmqMessageDelivered:    newMetricRabbitmqMessageDelivered(mbc.Metrics.RabbitmqMessageDelivered),
		metricRabbitmqMessageDropped:      newMetricRabbitmqMessageDropped(mbc.Metrics.RabbitmqMessageDropped),
		metricRabbitmqMessagePublished:    newMetricRabbitmqMessagePublished(mbc.Metrics.RabbitmqMessagePublished),
		resourceAttributeIncludeFilter:    make(map[string]filter.Filter),
		resourceAttributeExcludeFilter:    make(map[string]filter.Filter),
	}
	if mbc.ResourceAttributes.RabbitmqNodeName.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["rabbitmq.node.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqNodeName.MetricsInclude)
	}
	if mbc.ResourceAttributes.RabbitmqNodeName.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["rabbitmq.node.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqNodeName.MetricsExclude)
	}
	if mbc.ResourceAttributes.RabbitmqQueueName.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["rabbitmq.queue.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqQueueName.MetricsInclude)
	}
	if mbc.ResourceAttributes.RabbitmqQueueName.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["rabbitmq.queue.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqQueueName.MetricsExclude)
	}
	if mbc.ResourceAttributes.RabbitmqVhostName.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["rabbitmq.vhost.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqVhostName.MetricsInclude)
	}
	if mbc.ResourceAttributes.RabbitmqVhostName.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["rabbitmq.vhost.name"] = filter.CreateFilter(mbc.ResourceAttributes.RabbitmqVhostName.MetricsExclude)
	}

	for _, op := range options {
		op(mb)
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
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
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
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricRabbitmqConsumerCount.emit(ils.Metrics())
	mb.metricRabbitmqMessageAcknowledged.emit(ils.Metrics())
	mb.metricRabbitmqMessageCurrent.emit(ils.Metrics())
	mb.metricRabbitmqMessageDelivered.emit(ils.Metrics())
	mb.metricRabbitmqMessageDropped.emit(ils.Metrics())
	mb.metricRabbitmqMessagePublished.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
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
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordRabbitmqConsumerCountDataPoint adds a data point to rabbitmq.consumer.count metric.
func (mb *MetricsBuilder) RecordRabbitmqConsumerCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqConsumerCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageAcknowledgedDataPoint adds a data point to rabbitmq.message.acknowledged metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageAcknowledgedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageAcknowledged.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageCurrentDataPoint adds a data point to rabbitmq.message.current metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageCurrentDataPoint(ts pcommon.Timestamp, val int64, messageStateAttributeValue AttributeMessageState) {
	mb.metricRabbitmqMessageCurrent.recordDataPoint(mb.startTime, ts, val, messageStateAttributeValue.String())
}

// RecordRabbitmqMessageDeliveredDataPoint adds a data point to rabbitmq.message.delivered metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageDeliveredDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageDelivered.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageDroppedDataPoint adds a data point to rabbitmq.message.dropped metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageDroppedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageDropped.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessagePublishedDataPoint adds a data point to rabbitmq.message.published metric.
func (mb *MetricsBuilder) RecordRabbitmqMessagePublishedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessagePublished.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
