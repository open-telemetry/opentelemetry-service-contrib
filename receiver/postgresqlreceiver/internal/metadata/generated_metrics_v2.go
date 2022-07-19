// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for postgresqlreceiver metrics.
type MetricsSettings struct {
	PostgresqlBackends   MetricSettings `mapstructure:"postgresql.backends"`
	PostgresqlBlocksRead MetricSettings `mapstructure:"postgresql.blocks_read"`
	PostgresqlCommits    MetricSettings `mapstructure:"postgresql.commits"`
	PostgresqlDbSize     MetricSettings `mapstructure:"postgresql.db_size"`
	PostgresqlOperations MetricSettings `mapstructure:"postgresql.operations"`
	PostgresqlRollbacks  MetricSettings `mapstructure:"postgresql.rollbacks"`
	PostgresqlRows       MetricSettings `mapstructure:"postgresql.rows"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		PostgresqlBackends: MetricSettings{
			Enabled: true,
		},
		PostgresqlBlocksRead: MetricSettings{
			Enabled: true,
		},
		PostgresqlCommits: MetricSettings{
			Enabled: true,
		},
		PostgresqlDbSize: MetricSettings{
			Enabled: true,
		},
		PostgresqlOperations: MetricSettings{
			Enabled: true,
		},
		PostgresqlRollbacks: MetricSettings{
			Enabled: true,
		},
		PostgresqlRows: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeOperation specifies the a value operation attribute.
type AttributeOperation int

const (
	_ AttributeOperation = iota
	AttributeOperationIns
	AttributeOperationUpd
	AttributeOperationDel
	AttributeOperationHotUpd
)

// String returns the string representation of the AttributeOperation.
func (av AttributeOperation) String() string {
	switch av {
	case AttributeOperationIns:
		return "ins"
	case AttributeOperationUpd:
		return "upd"
	case AttributeOperationDel:
		return "del"
	case AttributeOperationHotUpd:
		return "hot_upd"
	}
	return ""
}

// MapAttributeOperation is a helper map of string to AttributeOperation attribute value.
var MapAttributeOperation = map[string]AttributeOperation{
	"ins":     AttributeOperationIns,
	"upd":     AttributeOperationUpd,
	"del":     AttributeOperationDel,
	"hot_upd": AttributeOperationHotUpd,
}

// AttributeSource specifies the a value source attribute.
type AttributeSource int

const (
	_ AttributeSource = iota
	AttributeSourceHeapRead
	AttributeSourceHeapHit
	AttributeSourceIdxRead
	AttributeSourceIdxHit
	AttributeSourceToastRead
	AttributeSourceToastHit
	AttributeSourceTidxRead
	AttributeSourceTidxHit
)

// String returns the string representation of the AttributeSource.
func (av AttributeSource) String() string {
	switch av {
	case AttributeSourceHeapRead:
		return "heap_read"
	case AttributeSourceHeapHit:
		return "heap_hit"
	case AttributeSourceIdxRead:
		return "idx_read"
	case AttributeSourceIdxHit:
		return "idx_hit"
	case AttributeSourceToastRead:
		return "toast_read"
	case AttributeSourceToastHit:
		return "toast_hit"
	case AttributeSourceTidxRead:
		return "tidx_read"
	case AttributeSourceTidxHit:
		return "tidx_hit"
	}
	return ""
}

// MapAttributeSource is a helper map of string to AttributeSource attribute value.
var MapAttributeSource = map[string]AttributeSource{
	"heap_read":  AttributeSourceHeapRead,
	"heap_hit":   AttributeSourceHeapHit,
	"idx_read":   AttributeSourceIdxRead,
	"idx_hit":    AttributeSourceIdxHit,
	"toast_read": AttributeSourceToastRead,
	"toast_hit":  AttributeSourceToastHit,
	"tidx_read":  AttributeSourceTidxRead,
	"tidx_hit":   AttributeSourceTidxHit,
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateDead
	AttributeStateLive
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateDead:
		return "dead"
	case AttributeStateLive:
		return "live"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"dead": AttributeStateDead,
	"live": AttributeStateLive,
}

type metricPostgresqlBackends struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.backends metric with initial data.
func (m *metricPostgresqlBackends) init() {
	m.data.SetName("postgresql.backends")
	m.data.SetDescription("The number of backends.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBackends) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBackends) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBackends) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBackends(settings MetricSettings) metricPostgresqlBackends {
	m := metricPostgresqlBackends{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlBlocksRead struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.blocks_read metric with initial data.
func (m *metricPostgresqlBlocksRead) init() {
	m.data.SetName("postgresql.blocks_read")
	m.data.SetDescription("The number of blocks read.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBlocksRead) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert("table", pcommon.NewValueString(tableAttributeValue))
	dp.Attributes().Insert("source", pcommon.NewValueString(sourceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBlocksRead) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBlocksRead) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBlocksRead(settings MetricSettings) metricPostgresqlBlocksRead {
	m := metricPostgresqlBlocksRead{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlCommits struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.commits metric with initial data.
func (m *metricPostgresqlCommits) init() {
	m.data.SetName("postgresql.commits")
	m.data.SetDescription("The number of commits.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlCommits) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlCommits) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlCommits) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlCommits(settings MetricSettings) metricPostgresqlCommits {
	m := metricPostgresqlCommits{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlDbSize struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.db_size metric with initial data.
func (m *metricPostgresqlDbSize) init() {
	m.data.SetName("postgresql.db_size")
	m.data.SetDescription("The database disk usage.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlDbSize) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlDbSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlDbSize) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlDbSize(settings MetricSettings) metricPostgresqlDbSize {
	m := metricPostgresqlDbSize{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.operations metric with initial data.
func (m *metricPostgresqlOperations) init() {
	m.data.SetName("postgresql.operations")
	m.data.SetDescription("The number of db row operations.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert("table", pcommon.NewValueString(tableAttributeValue))
	dp.Attributes().Insert("operation", pcommon.NewValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlOperations) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlOperations(settings MetricSettings) metricPostgresqlOperations {
	m := metricPostgresqlOperations{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRollbacks struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rollbacks metric with initial data.
func (m *metricPostgresqlRollbacks) init() {
	m.data.SetName("postgresql.rollbacks")
	m.data.SetDescription("The number of rollbacks.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRollbacks) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRollbacks) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRollbacks) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRollbacks(settings MetricSettings) metricPostgresqlRollbacks {
	m := metricPostgresqlRollbacks{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRows struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rows metric with initial data.
func (m *metricPostgresqlRows) init() {
	m.data.SetName("postgresql.rows")
	m.data.SetDescription("The number of rows in the database.")
	m.data.SetUnit("1")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRows) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("database", pcommon.NewValueString(databaseAttributeValue))
	dp.Attributes().Insert("table", pcommon.NewValueString(tableAttributeValue))
	dp.Attributes().Insert("state", pcommon.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRows) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRows) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRows(settings MetricSettings) metricPostgresqlRows {
	m := metricPostgresqlRows{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                  pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity            int                 // maximum observed number of metrics per resource.
	resourceCapacity           int                 // maximum observed number of resource attributes.
	metricsBuffer              pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                  component.BuildInfo // contains version information
	metricPostgresqlBackends   metricPostgresqlBackends
	metricPostgresqlBlocksRead metricPostgresqlBlocksRead
	metricPostgresqlCommits    metricPostgresqlCommits
	metricPostgresqlDbSize     metricPostgresqlDbSize
	metricPostgresqlOperations metricPostgresqlOperations
	metricPostgresqlRollbacks  metricPostgresqlRollbacks
	metricPostgresqlRows       metricPostgresqlRows
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
		startTime:                  pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:              pmetric.NewMetrics(),
		buildInfo:                  buildInfo,
		metricPostgresqlBackends:   newMetricPostgresqlBackends(settings.PostgresqlBackends),
		metricPostgresqlBlocksRead: newMetricPostgresqlBlocksRead(settings.PostgresqlBlocksRead),
		metricPostgresqlCommits:    newMetricPostgresqlCommits(settings.PostgresqlCommits),
		metricPostgresqlDbSize:     newMetricPostgresqlDbSize(settings.PostgresqlDbSize),
		metricPostgresqlOperations: newMetricPostgresqlOperations(settings.PostgresqlOperations),
		metricPostgresqlRollbacks:  newMetricPostgresqlRollbacks(settings.PostgresqlRollbacks),
		metricPostgresqlRows:       newMetricPostgresqlRows(settings.PostgresqlRows),
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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/postgresqlreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricPostgresqlBackends.emit(ils.Metrics())
	mb.metricPostgresqlBlocksRead.emit(ils.Metrics())
	mb.metricPostgresqlCommits.emit(ils.Metrics())
	mb.metricPostgresqlDbSize.emit(ils.Metrics())
	mb.metricPostgresqlOperations.emit(ils.Metrics())
	mb.metricPostgresqlRollbacks.emit(ils.Metrics())
	mb.metricPostgresqlRows.emit(ils.Metrics())
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

// RecordPostgresqlBackendsDataPoint adds a data point to postgresql.backends metric.
func (mb *MetricsBuilder) RecordPostgresqlBackendsDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlBackends.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlBlocksReadDataPoint adds a data point to postgresql.blocks_read metric.
func (mb *MetricsBuilder) RecordPostgresqlBlocksReadDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue AttributeSource) {
	mb.metricPostgresqlBlocksRead.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, sourceAttributeValue.String())
}

// RecordPostgresqlCommitsDataPoint adds a data point to postgresql.commits metric.
func (mb *MetricsBuilder) RecordPostgresqlCommitsDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlCommits.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlDbSizeDataPoint adds a data point to postgresql.db_size metric.
func (mb *MetricsBuilder) RecordPostgresqlDbSizeDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlDbSize.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlOperationsDataPoint adds a data point to postgresql.operations metric.
func (mb *MetricsBuilder) RecordPostgresqlOperationsDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue AttributeOperation) {
	mb.metricPostgresqlOperations.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, operationAttributeValue.String())
}

// RecordPostgresqlRollbacksDataPoint adds a data point to postgresql.rollbacks metric.
func (mb *MetricsBuilder) RecordPostgresqlRollbacksDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string) {
	mb.metricPostgresqlRollbacks.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlRowsDataPoint adds a data point to postgresql.rows metric.
func (mb *MetricsBuilder) RecordPostgresqlRowsDataPoint(ts pcommon.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue AttributeState) {
	mb.metricPostgresqlRows.recordDataPoint(mb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, stateAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
