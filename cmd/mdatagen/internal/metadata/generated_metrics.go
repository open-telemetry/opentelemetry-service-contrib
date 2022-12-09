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

// MetricsSettings provides settings for testreceiver metrics.
type MetricsSettings struct {
	DefaultMetric  MetricSettings `mapstructure:"default.metric"`
	OptionalMetric MetricSettings `mapstructure:"optional.metric"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		DefaultMetric: MetricSettings{
			Enabled: true,
		},
		OptionalMetric: MetricSettings{
			Enabled: false,
		},
	}
}

// AttributeEnumAttr specifies the a value enum_attr attribute.
type AttributeEnumAttr int

const (
	_ AttributeEnumAttr = iota
	AttributeEnumAttrRed
	AttributeEnumAttrGreen
	AttributeEnumAttrBlue
)

// String returns the string representation of the AttributeEnumAttr.
func (av AttributeEnumAttr) String() string {
	switch av {
	case AttributeEnumAttrRed:
		return "red"
	case AttributeEnumAttrGreen:
		return "green"
	case AttributeEnumAttrBlue:
		return "blue"
	}
	return ""
}

// MapAttributeEnumAttr is a helper map of string to AttributeEnumAttr attribute value.
var MapAttributeEnumAttr = map[string]AttributeEnumAttr{
	"red":   AttributeEnumAttrRed,
	"green": AttributeEnumAttrGreen,
	"blue":  AttributeEnumAttrBlue,
}

type metricDefaultMetric struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills default.metric metric with initial data.
func (m *metricDefaultMetric) init() {
	m.data.SetName("default.metric")
	m.data.SetDescription("Monotonic cumulative sum int metric enabled by default.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDefaultMetric) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutInt("state", overriddenIntAttrAttributeValue)
	dp.Attributes().PutStr("enum_attr", enumAttrAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDefaultMetric) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDefaultMetric) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDefaultMetric(settings MetricSettings) metricDefaultMetric {
	m := metricDefaultMetric{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricOptionalMetric struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills optional.metric metric with initial data.
func (m *metricOptionalMetric) init() {
	m.data.SetName("optional.metric")
	m.data.SetDescription("[DEPRECATED] Gauge double metric disabled by default.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricOptionalMetric) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutBool("boolean_attr", booleanAttrAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricOptionalMetric) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricOptionalMetric) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricOptionalMetric(settings MetricSettings) metricOptionalMetric {
	m := metricOptionalMetric{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime            pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity      int                 // maximum observed number of metrics per resource.
	resourceCapacity     int                 // maximum observed number of resource attributes.
	metricsBuffer        pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo            component.BuildInfo // contains version information
	metricDefaultMetric  metricDefaultMetric
	metricOptionalMetric metricOptionalMetric
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(ms MetricsSettings, settings component.ReceiverCreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	if ms.OptionalMetric.Enabled {
		settings.Logger.Warn("[WARNING] `optional.metric` should not be enabled: This metric is deprecated and will be removed soon.")
	}

	mb := &MetricsBuilder{
		startTime:            pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:        pmetric.NewMetrics(),
		buildInfo:            settings.BuildInfo,
		metricDefaultMetric:  newMetricDefaultMetric(ms.DefaultMetric),
		metricOptionalMetric: newMetricOptionalMetric(ms.OptionalMetric),
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

// WithStringEnumResourceAttrOne sets "string.enum.resource.attr=one" attribute for current resource.
func WithStringEnumResourceAttrOne(rm pmetric.ResourceMetrics) {
	rm.Resource().Attributes().PutStr("string.enum.resource.attr", "one")
}

// WithStringEnumResourceAttrTwo sets "string.enum.resource.attr=two" attribute for current resource.
func WithStringEnumResourceAttrTwo(rm pmetric.ResourceMetrics) {
	rm.Resource().Attributes().PutStr("string.enum.resource.attr", "two")
}

// WithStringResourceAttr sets provided value as "string.resource.attr" attribute for current resource.
func WithStringResourceAttr(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("string.resource.attr", val)
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
	ils.Scope().SetName("otelcol/testreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricDefaultMetric.emit(ils.Metrics())
	mb.metricOptionalMetric.emit(ils.Metrics())
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

// RecordDefaultMetricDataPoint adds a data point to default.metric metric.
func (mb *MetricsBuilder) RecordDefaultMetricDataPoint(ts pcommon.Timestamp, val int64, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue AttributeEnumAttr) {
	mb.metricDefaultMetric.recordDataPoint(mb.startTime, ts, val, stringAttrAttributeValue, overriddenIntAttrAttributeValue, enumAttrAttributeValue.String())
}

// RecordOptionalMetricDataPoint adds a data point to optional.metric metric.
func (mb *MetricsBuilder) RecordOptionalMetricDataPoint(ts pcommon.Timestamp, val float64, stringAttrAttributeValue string, booleanAttrAttributeValue bool) {
	mb.metricOptionalMetric.recordDataPoint(mb.startTime, ts, val, stringAttrAttributeValue, booleanAttrAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
