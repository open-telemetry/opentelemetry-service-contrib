// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.6.1"
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

type metricPostgresqlBackends struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.backends metric with initial data.
func (m *metricPostgresqlBackends) init() {
	m.data.SetName("postgresql.backends")
	m.data.SetDescription("The number of backends.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBackends) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBackends) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBackends) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBackends(settings MetricSettings) metricPostgresqlBackends {
	m := metricPostgresqlBackends{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlBlocksRead struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.blocks_read metric with initial data.
func (m *metricPostgresqlBlocksRead) init() {
	m.data.SetName("postgresql.blocks_read")
	m.data.SetDescription("The number of blocks read.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlBlocksRead) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewAttributeValueString(tableAttributeValue))
	dp.Attributes().Insert(A.Source, pdata.NewAttributeValueString(sourceAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlBlocksRead) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlBlocksRead) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlBlocksRead(settings MetricSettings) metricPostgresqlBlocksRead {
	m := metricPostgresqlBlocksRead{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlCommits struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.commits metric with initial data.
func (m *metricPostgresqlCommits) init() {
	m.data.SetName("postgresql.commits")
	m.data.SetDescription("The number of commits.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlCommits) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlCommits) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlCommits) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlCommits(settings MetricSettings) metricPostgresqlCommits {
	m := metricPostgresqlCommits{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlDbSize struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.db_size metric with initial data.
func (m *metricPostgresqlDbSize) init() {
	m.data.SetName("postgresql.db_size")
	m.data.SetDescription("The database disk usage.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlDbSize) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlDbSize) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlDbSize) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlDbSize(settings MetricSettings) metricPostgresqlDbSize {
	m := metricPostgresqlDbSize{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.operations metric with initial data.
func (m *metricPostgresqlOperations) init() {
	m.data.SetName("postgresql.operations")
	m.data.SetDescription("The number of db row operations.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewAttributeValueString(tableAttributeValue))
	dp.Attributes().Insert(A.Operation, pdata.NewAttributeValueString(operationAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlOperations(settings MetricSettings) metricPostgresqlOperations {
	m := metricPostgresqlOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRollbacks struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rollbacks metric with initial data.
func (m *metricPostgresqlRollbacks) init() {
	m.data.SetName("postgresql.rollbacks")
	m.data.SetDescription("The number of rollbacks.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRollbacks) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRollbacks) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRollbacks) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRollbacks(settings MetricSettings) metricPostgresqlRollbacks {
	m := metricPostgresqlRollbacks{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPostgresqlRows struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills postgresql.rows metric with initial data.
func (m *metricPostgresqlRows) init() {
	m.data.SetName("postgresql.rows")
	m.data.SetDescription("The number of rows in the database.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPostgresqlRows) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Database, pdata.NewAttributeValueString(databaseAttributeValue))
	dp.Attributes().Insert(A.Table, pdata.NewAttributeValueString(tableAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewAttributeValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPostgresqlRows) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPostgresqlRows) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPostgresqlRows(settings MetricSettings) metricPostgresqlRows {
	m := metricPostgresqlRows{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime        pdata.Timestamp
	settings         MetricsSettings
	resourceBuilders []*ResourceBuilder
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime: pdata.NewTimestampFromTime(time.Now()),
		settings:  settings,
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// emitOption applies changes to pdata.Metrics emitted.
type emitOption func(*pdata.Metrics)

// WithResource copies the pdata.Resource into the emitted pdata.Metrics.
func WithResource(resource pdata.Resource) emitOption {
	return func(md *pdata.Metrics) {
		resource.CopyTo(md.ResourceMetrics().At(0).Resource())
	}
}

// WithCapacity calls EnsureCapacity on the pdata.Metrics.
func WithCapacity(capacity int) emitOption {
	return func(md *pdata.Metrics) {
		if capacity > 0 {
			md.ResourceMetrics().At(0).InstrumentationLibraryMetrics().At(0).Metrics().EnsureCapacity(capacity)
		}
	}
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(options ...emitOption) pdata.Metrics {
	md := mb.newMetricData()

	for _, op := range options {
		op(&md)
	}

	for _, rb := range mb.resourceBuilders {
		rm := md.ResourceMetrics().AppendEmpty()
		ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
		ilm.InstrumentationLibrary().SetName("otelcol/postgresqlreceiver")
		rm.SetSchemaUrl(conventions.SchemaURL)
		rb.emit(rm)
	}
	return md
}

type ResourceBuilder struct {
	resource                   pdata.Resource
	startTime                  pdata.Timestamp
	metricPostgresqlBackends   metricPostgresqlBackends
	metricPostgresqlBlocksRead metricPostgresqlBlocksRead
	metricPostgresqlCommits    metricPostgresqlCommits
	metricPostgresqlDbSize     metricPostgresqlDbSize
	metricPostgresqlOperations metricPostgresqlOperations
	metricPostgresqlRollbacks  metricPostgresqlRollbacks
	metricPostgresqlRows       metricPostgresqlRows
}

func (rb *ResourceBuilder) Attributes() pdata.AttributeMap {
	return rb.resource.Attributes()
}

func (rb *ResourceBuilder) emit(rm pdata.ResourceMetrics) {
	rb.resource.CopyTo(rm.Resource())

	metrics := rm.InstrumentationLibraryMetrics().At(0).Metrics()

	rb.metricPostgresqlBackends.emit(metrics)
	rb.metricPostgresqlBlocksRead.emit(metrics)
	rb.metricPostgresqlCommits.emit(metrics)
	rb.metricPostgresqlDbSize.emit(metrics)
	rb.metricPostgresqlOperations.emit(metrics)
	rb.metricPostgresqlRollbacks.emit(metrics)
	rb.metricPostgresqlRows.emit(metrics)
}

func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	rb := ResourceBuilder{
		metricPostgresqlBackends:   newMetricPostgresqlBackends(mb.settings.PostgresqlBackends),
		metricPostgresqlBlocksRead: newMetricPostgresqlBlocksRead(mb.settings.PostgresqlBlocksRead),
		metricPostgresqlCommits:    newMetricPostgresqlCommits(mb.settings.PostgresqlCommits),
		metricPostgresqlDbSize:     newMetricPostgresqlDbSize(mb.settings.PostgresqlDbSize),
		metricPostgresqlOperations: newMetricPostgresqlOperations(mb.settings.PostgresqlOperations),
		metricPostgresqlRollbacks:  newMetricPostgresqlRollbacks(mb.settings.PostgresqlRollbacks),
		metricPostgresqlRows:       newMetricPostgresqlRows(mb.settings.PostgresqlRows),
		resource:                   pdata.NewResource(),
	}
	mb.resourceBuilders = append(mb.resourceBuilders, &rb)
	return &rb
}

// RecordPostgresqlBackendsDataPoint adds a data point to postgresql.backends metric.
func (rb *ResourceBuilder) RecordPostgresqlBackendsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	rb.metricPostgresqlBackends.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlBlocksReadDataPoint adds a data point to postgresql.blocks_read metric.
func (rb *ResourceBuilder) RecordPostgresqlBlocksReadDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, sourceAttributeValue string) {
	rb.metricPostgresqlBlocksRead.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, sourceAttributeValue)
}

// RecordPostgresqlCommitsDataPoint adds a data point to postgresql.commits metric.
func (rb *ResourceBuilder) RecordPostgresqlCommitsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	rb.metricPostgresqlCommits.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlDbSizeDataPoint adds a data point to postgresql.db_size metric.
func (rb *ResourceBuilder) RecordPostgresqlDbSizeDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	rb.metricPostgresqlDbSize.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlOperationsDataPoint adds a data point to postgresql.operations metric.
func (rb *ResourceBuilder) RecordPostgresqlOperationsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, operationAttributeValue string) {
	rb.metricPostgresqlOperations.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, operationAttributeValue)
}

// RecordPostgresqlRollbacksDataPoint adds a data point to postgresql.rollbacks metric.
func (rb *ResourceBuilder) RecordPostgresqlRollbacksDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string) {
	rb.metricPostgresqlRollbacks.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue)
}

// RecordPostgresqlRowsDataPoint adds a data point to postgresql.rows metric.
func (rb *ResourceBuilder) RecordPostgresqlRowsDataPoint(ts pdata.Timestamp, val int64, databaseAttributeValue string, tableAttributeValue string, stateAttributeValue string) {
	rb.metricPostgresqlRows.recordDataPoint(rb.startTime, ts, val, databaseAttributeValue, tableAttributeValue, stateAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// newMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) newMetricData() pdata.Metrics {
	md := pdata.NewMetrics()
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Database (The name of the database.)
	Database string
	// Operation (The database operation.)
	Operation string
	// Source (The block read source type.)
	Source string
	// State (The tuple (row) state.)
	State string
	// Table (The schema name followed by the table name.)
	Table string
}{
	"database",
	"operation",
	"source",
	"state",
	"table",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Ins    string
	Upd    string
	Del    string
	HotUpd string
}{
	"ins",
	"upd",
	"del",
	"hot_upd",
}

// AttributeSource are the possible values that the attribute "source" can have.
var AttributeSource = struct {
	HeapRead  string
	HeapHit   string
	IdxRead   string
	IdxHit    string
	ToastRead string
	ToastHit  string
	TidxRead  string
	TidxHit   string
}{
	"heap_read",
	"heap_hit",
	"idx_read",
	"idx_hit",
	"toast_read",
	"toast_hit",
	"tidx_read",
	"tidx_hit",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Dead string
	Live string
}{
	"dead",
	"live",
}
