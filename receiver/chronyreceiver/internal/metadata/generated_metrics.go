// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
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

// MetricsSettings provides settings for chrony receiver metrics.
type MetricsSettings struct {
	NtpFrequencyOffset MetricSettings `mapstructure:"ntp.frequency.offset"`
	NtpSkew            MetricSettings `mapstructure:"ntp.skew"`
	NtpStratum         MetricSettings `mapstructure:"ntp.stratum"`
	NtpTimeCorrection  MetricSettings `mapstructure:"ntp.time.correction"`
	NtpTimeLastOffset  MetricSettings `mapstructure:"ntp.time.last_offset"`
	NtpTimeRmsOffset   MetricSettings `mapstructure:"ntp.time.rms_offset"`
	NtpTimeRootDelay   MetricSettings `mapstructure:"ntp.time.root_delay"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		NtpFrequencyOffset: MetricSettings{
			Enabled: false,
		},
		NtpSkew: MetricSettings{
			Enabled: true,
		},
		NtpStratum: MetricSettings{
			Enabled: false,
		},
		NtpTimeCorrection: MetricSettings{
			Enabled: true,
		},
		NtpTimeLastOffset: MetricSettings{
			Enabled: true,
		},
		NtpTimeRmsOffset: MetricSettings{
			Enabled: false,
		},
		NtpTimeRootDelay: MetricSettings{
			Enabled: false,
		},
	}
}

// AttributeLeapStatus specifies the a value leap.status attribute.
type AttributeLeapStatus int

const (
	_ AttributeLeapStatus = iota
	AttributeLeapStatusNormal
	AttributeLeapStatusInsertSecond
	AttributeLeapStatusDeleteSecond
	AttributeLeapStatusUnsynchronised
)

// String returns the string representation of the AttributeLeapStatus.
func (av AttributeLeapStatus) String() string {
	switch av {
	case AttributeLeapStatusNormal:
		return "normal"
	case AttributeLeapStatusInsertSecond:
		return "insert_second"
	case AttributeLeapStatusDeleteSecond:
		return "delete_second"
	case AttributeLeapStatusUnsynchronised:
		return "unsynchronised"
	}
	return ""
}

// MapAttributeLeapStatus is a helper map of string to AttributeLeapStatus attribute value.
var MapAttributeLeapStatus = map[string]AttributeLeapStatus{
	"normal":         AttributeLeapStatusNormal,
	"insert_second":  AttributeLeapStatusInsertSecond,
	"delete_second":  AttributeLeapStatusDeleteSecond,
	"unsynchronised": AttributeLeapStatusUnsynchronised,
}

type metricNtpFrequencyOffset struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.frequency.offset metric with initial data.
func (m *metricNtpFrequencyOffset) init() {
	m.data.SetName("ntp.frequency.offset")
	m.data.SetDescription("The frequency is the rate by which the system s clock would be wrong if chronyd was not correcting it.")
	m.data.SetUnit("ppm")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNtpFrequencyOffset) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, leapStatusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("leap.status", leapStatusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpFrequencyOffset) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpFrequencyOffset) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpFrequencyOffset(settings MetricSettings) metricNtpFrequencyOffset {
	m := metricNtpFrequencyOffset{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpSkew struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.skew metric with initial data.
func (m *metricNtpSkew) init() {
	m.data.SetName("ntp.skew")
	m.data.SetDescription("This is the estimated error bound on the frequency.")
	m.data.SetUnit("ppm")
	m.data.SetEmptyGauge()
}

func (m *metricNtpSkew) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpSkew) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpSkew) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpSkew(settings MetricSettings) metricNtpSkew {
	m := metricNtpSkew{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpStratum struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.stratum metric with initial data.
func (m *metricNtpStratum) init() {
	m.data.SetName("ntp.stratum")
	m.data.SetDescription("The number of hops away from the reference system keeping the reference time")
	m.data.SetUnit("{count}")
	m.data.SetEmptyGauge()
}

func (m *metricNtpStratum) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpStratum) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpStratum) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpStratum(settings MetricSettings) metricNtpStratum {
	m := metricNtpStratum{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpTimeCorrection struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.time.correction metric with initial data.
func (m *metricNtpTimeCorrection) init() {
	m.data.SetName("ntp.time.correction")
	m.data.SetDescription("The number of seconds difference between the system's clock and the reference clock")
	m.data.SetUnit("seconds")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNtpTimeCorrection) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, leapStatusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("leap.status", leapStatusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpTimeCorrection) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpTimeCorrection) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpTimeCorrection(settings MetricSettings) metricNtpTimeCorrection {
	m := metricNtpTimeCorrection{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpTimeLastOffset struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.time.last_offset metric with initial data.
func (m *metricNtpTimeLastOffset) init() {
	m.data.SetName("ntp.time.last_offset")
	m.data.SetDescription("The estimated local offset on the last clock update")
	m.data.SetUnit("seconds")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNtpTimeLastOffset) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, leapStatusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("leap.status", leapStatusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpTimeLastOffset) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpTimeLastOffset) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpTimeLastOffset(settings MetricSettings) metricNtpTimeLastOffset {
	m := metricNtpTimeLastOffset{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpTimeRmsOffset struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.time.rms_offset metric with initial data.
func (m *metricNtpTimeRmsOffset) init() {
	m.data.SetName("ntp.time.rms_offset")
	m.data.SetDescription("the long term average of the offset value")
	m.data.SetUnit("seconds")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNtpTimeRmsOffset) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, leapStatusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("leap.status", leapStatusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpTimeRmsOffset) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpTimeRmsOffset) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpTimeRmsOffset(settings MetricSettings) metricNtpTimeRmsOffset {
	m := metricNtpTimeRmsOffset{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNtpTimeRootDelay struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills ntp.time.root_delay metric with initial data.
func (m *metricNtpTimeRootDelay) init() {
	m.data.SetName("ntp.time.root_delay")
	m.data.SetDescription("This is the total of the network path delays to the stratum-1 system from which the system is ultimately synchronised.")
	m.data.SetUnit("seconds")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNtpTimeRootDelay) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, leapStatusAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("leap.status", leapStatusAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNtpTimeRootDelay) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNtpTimeRootDelay) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNtpTimeRootDelay(settings MetricSettings) metricNtpTimeRootDelay {
	m := metricNtpTimeRootDelay{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity          int                 // maximum observed number of metrics per resource.
	resourceCapacity         int                 // maximum observed number of resource attributes.
	metricsBuffer            pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                component.BuildInfo // contains version information
	metricNtpFrequencyOffset metricNtpFrequencyOffset
	metricNtpSkew            metricNtpSkew
	metricNtpStratum         metricNtpStratum
	metricNtpTimeCorrection  metricNtpTimeCorrection
	metricNtpTimeLastOffset  metricNtpTimeLastOffset
	metricNtpTimeRmsOffset   metricNtpTimeRmsOffset
	metricNtpTimeRootDelay   metricNtpTimeRootDelay
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(ms MetricsSettings, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:            pmetric.NewMetrics(),
		buildInfo:                settings.BuildInfo,
		metricNtpFrequencyOffset: newMetricNtpFrequencyOffset(ms.NtpFrequencyOffset),
		metricNtpSkew:            newMetricNtpSkew(ms.NtpSkew),
		metricNtpStratum:         newMetricNtpStratum(ms.NtpStratum),
		metricNtpTimeCorrection:  newMetricNtpTimeCorrection(ms.NtpTimeCorrection),
		metricNtpTimeLastOffset:  newMetricNtpTimeLastOffset(ms.NtpTimeLastOffset),
		metricNtpTimeRmsOffset:   newMetricNtpTimeRmsOffset(ms.NtpTimeRmsOffset),
		metricNtpTimeRootDelay:   newMetricNtpTimeRootDelay(ms.NtpTimeRootDelay),
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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/chrony receiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricNtpFrequencyOffset.emit(ils.Metrics())
	mb.metricNtpSkew.emit(ils.Metrics())
	mb.metricNtpStratum.emit(ils.Metrics())
	mb.metricNtpTimeCorrection.emit(ils.Metrics())
	mb.metricNtpTimeLastOffset.emit(ils.Metrics())
	mb.metricNtpTimeRmsOffset.emit(ils.Metrics())
	mb.metricNtpTimeRootDelay.emit(ils.Metrics())
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

// RecordNtpFrequencyOffsetDataPoint adds a data point to ntp.frequency.offset metric.
func (mb *MetricsBuilder) RecordNtpFrequencyOffsetDataPoint(ts pcommon.Timestamp, val float64, leapStatusAttributeValue AttributeLeapStatus) {
	mb.metricNtpFrequencyOffset.recordDataPoint(mb.startTime, ts, val, leapStatusAttributeValue.String())
}

// RecordNtpSkewDataPoint adds a data point to ntp.skew metric.
func (mb *MetricsBuilder) RecordNtpSkewDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricNtpSkew.recordDataPoint(mb.startTime, ts, val)
}

// RecordNtpStratumDataPoint adds a data point to ntp.stratum metric.
func (mb *MetricsBuilder) RecordNtpStratumDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNtpStratum.recordDataPoint(mb.startTime, ts, val)
}

// RecordNtpTimeCorrectionDataPoint adds a data point to ntp.time.correction metric.
func (mb *MetricsBuilder) RecordNtpTimeCorrectionDataPoint(ts pcommon.Timestamp, val float64, leapStatusAttributeValue AttributeLeapStatus) {
	mb.metricNtpTimeCorrection.recordDataPoint(mb.startTime, ts, val, leapStatusAttributeValue.String())
}

// RecordNtpTimeLastOffsetDataPoint adds a data point to ntp.time.last_offset metric.
func (mb *MetricsBuilder) RecordNtpTimeLastOffsetDataPoint(ts pcommon.Timestamp, val float64, leapStatusAttributeValue AttributeLeapStatus) {
	mb.metricNtpTimeLastOffset.recordDataPoint(mb.startTime, ts, val, leapStatusAttributeValue.String())
}

// RecordNtpTimeRmsOffsetDataPoint adds a data point to ntp.time.rms_offset metric.
func (mb *MetricsBuilder) RecordNtpTimeRmsOffsetDataPoint(ts pcommon.Timestamp, val float64, leapStatusAttributeValue AttributeLeapStatus) {
	mb.metricNtpTimeRmsOffset.recordDataPoint(mb.startTime, ts, val, leapStatusAttributeValue.String())
}

// RecordNtpTimeRootDelayDataPoint adds a data point to ntp.time.root_delay metric.
func (mb *MetricsBuilder) RecordNtpTimeRootDelayDataPoint(ts pcommon.Timestamp, val float64, leapStatusAttributeValue AttributeLeapStatus) {
	mb.metricNtpTimeRootDelay.recordDataPoint(mb.startTime, ts, val, leapStatusAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
