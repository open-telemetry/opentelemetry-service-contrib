// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for aerospikereceiver metrics.
type MetricsSettings struct {
	AerospikeNamespaceDiskAvailable    MetricSettings `mapstructure:"aerospike.namespace.disk.available"`
	AerospikeNamespaceMemoryFree       MetricSettings `mapstructure:"aerospike.namespace.memory.free"`
	AerospikeNamespaceMemoryUsage      MetricSettings `mapstructure:"aerospike.namespace.memory.usage"`
	AerospikeNamespaceScanCount        MetricSettings `mapstructure:"aerospike.namespace.scan.count"`
	AerospikeNamespaceTransactionCount MetricSettings `mapstructure:"aerospike.namespace.transaction.count"`
	AerospikeNodeConnectionCount       MetricSettings `mapstructure:"aerospike.node.connection.count"`
	AerospikeNodeConnectionOpen        MetricSettings `mapstructure:"aerospike.node.connection.open"`
	AerospikeNodeMemoryFree            MetricSettings `mapstructure:"aerospike.node.memory.free"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		AerospikeNamespaceDiskAvailable: MetricSettings{
			Enabled: true,
		},
		AerospikeNamespaceMemoryFree: MetricSettings{
			Enabled: true,
		},
		AerospikeNamespaceMemoryUsage: MetricSettings{
			Enabled: true,
		},
		AerospikeNamespaceScanCount: MetricSettings{
			Enabled: true,
		},
		AerospikeNamespaceTransactionCount: MetricSettings{
			Enabled: true,
		},
		AerospikeNodeConnectionCount: MetricSettings{
			Enabled: true,
		},
		AerospikeNodeConnectionOpen: MetricSettings{
			Enabled: true,
		},
		AerospikeNodeMemoryFree: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeConnectionOp specifies the a value connection_op attribute.
type AttributeConnectionOp int

const (
	_ AttributeConnectionOp = iota
	AttributeConnectionOpClose
	AttributeConnectionOpOpen
)

// String returns the string representation of the AttributeConnectionOp.
func (av AttributeConnectionOp) String() string {
	switch av {
	case AttributeConnectionOpClose:
		return "close"
	case AttributeConnectionOpOpen:
		return "open"
	}
	return ""
}

// MapAttributeConnectionOp is a helper map of string to AttributeConnectionOp attribute value.
var MapAttributeConnectionOp = map[string]AttributeConnectionOp{
	"close": AttributeConnectionOpClose,
	"open":  AttributeConnectionOpOpen,
}

// AttributeConnectionType specifies the a value connection_type attribute.
type AttributeConnectionType int

const (
	_ AttributeConnectionType = iota
	AttributeConnectionTypeClient
	AttributeConnectionTypeFabric
	AttributeConnectionTypeHeartbeat
)

// String returns the string representation of the AttributeConnectionType.
func (av AttributeConnectionType) String() string {
	switch av {
	case AttributeConnectionTypeClient:
		return "client"
	case AttributeConnectionTypeFabric:
		return "fabric"
	case AttributeConnectionTypeHeartbeat:
		return "heartbeat"
	}
	return ""
}

// MapAttributeConnectionType is a helper map of string to AttributeConnectionType attribute value.
var MapAttributeConnectionType = map[string]AttributeConnectionType{
	"client":    AttributeConnectionTypeClient,
	"fabric":    AttributeConnectionTypeFabric,
	"heartbeat": AttributeConnectionTypeHeartbeat,
}

// AttributeNamespaceComponent specifies the a value namespace_component attribute.
type AttributeNamespaceComponent int

const (
	_ AttributeNamespaceComponent = iota
	AttributeNamespaceComponentData
	AttributeNamespaceComponentIndex
	AttributeNamespaceComponentSetIndex
	AttributeNamespaceComponentSecondaryIndex
)

// String returns the string representation of the AttributeNamespaceComponent.
func (av AttributeNamespaceComponent) String() string {
	switch av {
	case AttributeNamespaceComponentData:
		return "data"
	case AttributeNamespaceComponentIndex:
		return "index"
	case AttributeNamespaceComponentSetIndex:
		return "set_index"
	case AttributeNamespaceComponentSecondaryIndex:
		return "secondary_index"
	}
	return ""
}

// MapAttributeNamespaceComponent is a helper map of string to AttributeNamespaceComponent attribute value.
var MapAttributeNamespaceComponent = map[string]AttributeNamespaceComponent{
	"data":            AttributeNamespaceComponentData,
	"index":           AttributeNamespaceComponentIndex,
	"set_index":       AttributeNamespaceComponentSetIndex,
	"secondary_index": AttributeNamespaceComponentSecondaryIndex,
}

// AttributeScanResult specifies the a value scan_result attribute.
type AttributeScanResult int

const (
	_ AttributeScanResult = iota
	AttributeScanResultAbort
	AttributeScanResultComplete
	AttributeScanResultError
)

// String returns the string representation of the AttributeScanResult.
func (av AttributeScanResult) String() string {
	switch av {
	case AttributeScanResultAbort:
		return "abort"
	case AttributeScanResultComplete:
		return "complete"
	case AttributeScanResultError:
		return "error"
	}
	return ""
}

// MapAttributeScanResult is a helper map of string to AttributeScanResult attribute value.
var MapAttributeScanResult = map[string]AttributeScanResult{
	"abort":    AttributeScanResultAbort,
	"complete": AttributeScanResultComplete,
	"error":    AttributeScanResultError,
}

// AttributeScanType specifies the a value scan_type attribute.
type AttributeScanType int

const (
	_ AttributeScanType = iota
	AttributeScanTypeAggregation
	AttributeScanTypeBasic
	AttributeScanTypeOpsBackground
	AttributeScanTypeUdfBackground
)

// String returns the string representation of the AttributeScanType.
func (av AttributeScanType) String() string {
	switch av {
	case AttributeScanTypeAggregation:
		return "aggregation"
	case AttributeScanTypeBasic:
		return "basic"
	case AttributeScanTypeOpsBackground:
		return "ops_background"
	case AttributeScanTypeUdfBackground:
		return "udf_background"
	}
	return ""
}

// MapAttributeScanType is a helper map of string to AttributeScanType attribute value.
var MapAttributeScanType = map[string]AttributeScanType{
	"aggregation":    AttributeScanTypeAggregation,
	"basic":          AttributeScanTypeBasic,
	"ops_background": AttributeScanTypeOpsBackground,
	"udf_background": AttributeScanTypeUdfBackground,
}

// AttributeTransactionResult specifies the a value transaction_result attribute.
type AttributeTransactionResult int

const (
	_ AttributeTransactionResult = iota
	AttributeTransactionResultError
	AttributeTransactionResultFilteredOut
	AttributeTransactionResultNotFound
	AttributeTransactionResultSuccess
	AttributeTransactionResultTimeout
)

// String returns the string representation of the AttributeTransactionResult.
func (av AttributeTransactionResult) String() string {
	switch av {
	case AttributeTransactionResultError:
		return "error"
	case AttributeTransactionResultFilteredOut:
		return "filtered_out"
	case AttributeTransactionResultNotFound:
		return "not_found"
	case AttributeTransactionResultSuccess:
		return "success"
	case AttributeTransactionResultTimeout:
		return "timeout"
	}
	return ""
}

// MapAttributeTransactionResult is a helper map of string to AttributeTransactionResult attribute value.
var MapAttributeTransactionResult = map[string]AttributeTransactionResult{
	"error":        AttributeTransactionResultError,
	"filtered_out": AttributeTransactionResultFilteredOut,
	"not_found":    AttributeTransactionResultNotFound,
	"success":      AttributeTransactionResultSuccess,
	"timeout":      AttributeTransactionResultTimeout,
}

// AttributeTransactionType specifies the a value transaction_type attribute.
type AttributeTransactionType int

const (
	_ AttributeTransactionType = iota
	AttributeTransactionTypeDelete
	AttributeTransactionTypeRead
	AttributeTransactionTypeUdf
	AttributeTransactionTypeWrite
)

// String returns the string representation of the AttributeTransactionType.
func (av AttributeTransactionType) String() string {
	switch av {
	case AttributeTransactionTypeDelete:
		return "delete"
	case AttributeTransactionTypeRead:
		return "read"
	case AttributeTransactionTypeUdf:
		return "udf"
	case AttributeTransactionTypeWrite:
		return "write"
	}
	return ""
}

// MapAttributeTransactionType is a helper map of string to AttributeTransactionType attribute value.
var MapAttributeTransactionType = map[string]AttributeTransactionType{
	"delete": AttributeTransactionTypeDelete,
	"read":   AttributeTransactionTypeRead,
	"udf":    AttributeTransactionTypeUdf,
	"write":  AttributeTransactionTypeWrite,
}

type metricAerospikeNamespaceDiskAvailable struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.namespace.disk.available metric with initial data.
func (m *metricAerospikeNamespaceDiskAvailable) init() {
	m.data.SetName("aerospike.namespace.disk.available")
	m.data.SetDescription("Minimum percentage of contiguous disk space free to the namespace across all devices")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricAerospikeNamespaceDiskAvailable) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNamespaceDiskAvailable) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNamespaceDiskAvailable) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNamespaceDiskAvailable(settings MetricSettings) metricAerospikeNamespaceDiskAvailable {
	m := metricAerospikeNamespaceDiskAvailable{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNamespaceMemoryFree struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.namespace.memory.free metric with initial data.
func (m *metricAerospikeNamespaceMemoryFree) init() {
	m.data.SetName("aerospike.namespace.memory.free")
	m.data.SetDescription("Percentage of the namespace's memory which is still free")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricAerospikeNamespaceMemoryFree) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNamespaceMemoryFree) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNamespaceMemoryFree) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNamespaceMemoryFree(settings MetricSettings) metricAerospikeNamespaceMemoryFree {
	m := metricAerospikeNamespaceMemoryFree{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNamespaceMemoryUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.namespace.memory.usage metric with initial data.
func (m *metricAerospikeNamespaceMemoryUsage) init() {
	m.data.SetName("aerospike.namespace.memory.usage")
	m.data.SetDescription("Memory currently used by each component of the namespace")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricAerospikeNamespaceMemoryUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, namespaceComponentAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("component", pcommon.NewValueString(namespaceComponentAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNamespaceMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNamespaceMemoryUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNamespaceMemoryUsage(settings MetricSettings) metricAerospikeNamespaceMemoryUsage {
	m := metricAerospikeNamespaceMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNamespaceScanCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.namespace.scan.count metric with initial data.
func (m *metricAerospikeNamespaceScanCount) init() {
	m.data.SetName("aerospike.namespace.scan.count")
	m.data.SetDescription("Number of scan operations performed on the namespace")
	m.data.SetUnit("{scans}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricAerospikeNamespaceScanCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, scanTypeAttributeValue string, scanResultAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("type", pcommon.NewValueString(scanTypeAttributeValue))
	dp.Attributes().Insert("result", pcommon.NewValueString(scanResultAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNamespaceScanCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNamespaceScanCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNamespaceScanCount(settings MetricSettings) metricAerospikeNamespaceScanCount {
	m := metricAerospikeNamespaceScanCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNamespaceTransactionCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.namespace.transaction.count metric with initial data.
func (m *metricAerospikeNamespaceTransactionCount) init() {
	m.data.SetName("aerospike.namespace.transaction.count")
	m.data.SetDescription("Number of transactions performed on the namespace")
	m.data.SetUnit("{transactions}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricAerospikeNamespaceTransactionCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, transactionTypeAttributeValue string, transactionResultAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("type", pcommon.NewValueString(transactionTypeAttributeValue))
	dp.Attributes().Insert("result", pcommon.NewValueString(transactionResultAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNamespaceTransactionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNamespaceTransactionCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNamespaceTransactionCount(settings MetricSettings) metricAerospikeNamespaceTransactionCount {
	m := metricAerospikeNamespaceTransactionCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNodeConnectionCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.node.connection.count metric with initial data.
func (m *metricAerospikeNodeConnectionCount) init() {
	m.data.SetName("aerospike.node.connection.count")
	m.data.SetDescription("Number of connections opened and closed to the node")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricAerospikeNodeConnectionCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, connectionTypeAttributeValue string, connectionOpAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("type", pcommon.NewValueString(connectionTypeAttributeValue))
	dp.Attributes().Insert("operation", pcommon.NewValueString(connectionOpAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNodeConnectionCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNodeConnectionCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNodeConnectionCount(settings MetricSettings) metricAerospikeNodeConnectionCount {
	m := metricAerospikeNodeConnectionCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNodeConnectionOpen struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.node.connection.open metric with initial data.
func (m *metricAerospikeNodeConnectionOpen) init() {
	m.data.SetName("aerospike.node.connection.open")
	m.data.SetDescription("Current number of open connections to the node")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricAerospikeNodeConnectionOpen) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, connectionTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert("type", pcommon.NewValueString(connectionTypeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNodeConnectionOpen) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNodeConnectionOpen) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNodeConnectionOpen(settings MetricSettings) metricAerospikeNodeConnectionOpen {
	m := metricAerospikeNodeConnectionOpen{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricAerospikeNodeMemoryFree struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills aerospike.node.memory.free metric with initial data.
func (m *metricAerospikeNodeMemoryFree) init() {
	m.data.SetName("aerospike.node.memory.free")
	m.data.SetDescription("Percentage of the node's memory which is still free")
	m.data.SetUnit("%")
	m.data.SetDataType(pmetric.MetricDataTypeGauge)
}

func (m *metricAerospikeNodeMemoryFree) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricAerospikeNodeMemoryFree) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricAerospikeNodeMemoryFree) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricAerospikeNodeMemoryFree(settings MetricSettings) metricAerospikeNodeMemoryFree {
	m := metricAerospikeNodeMemoryFree{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                                pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                          int                 // maximum observed number of metrics per resource.
	resourceCapacity                         int                 // maximum observed number of resource attributes.
	metricsBuffer                            pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                                component.BuildInfo // contains version information
	metricAerospikeNamespaceDiskAvailable    metricAerospikeNamespaceDiskAvailable
	metricAerospikeNamespaceMemoryFree       metricAerospikeNamespaceMemoryFree
	metricAerospikeNamespaceMemoryUsage      metricAerospikeNamespaceMemoryUsage
	metricAerospikeNamespaceScanCount        metricAerospikeNamespaceScanCount
	metricAerospikeNamespaceTransactionCount metricAerospikeNamespaceTransactionCount
	metricAerospikeNodeConnectionCount       metricAerospikeNodeConnectionCount
	metricAerospikeNodeConnectionOpen        metricAerospikeNodeConnectionOpen
	metricAerospikeNodeMemoryFree            metricAerospikeNodeMemoryFree
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
		startTime:                                pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                            pmetric.NewMetrics(),
		buildInfo:                                buildInfo,
		metricAerospikeNamespaceDiskAvailable:    newMetricAerospikeNamespaceDiskAvailable(settings.AerospikeNamespaceDiskAvailable),
		metricAerospikeNamespaceMemoryFree:       newMetricAerospikeNamespaceMemoryFree(settings.AerospikeNamespaceMemoryFree),
		metricAerospikeNamespaceMemoryUsage:      newMetricAerospikeNamespaceMemoryUsage(settings.AerospikeNamespaceMemoryUsage),
		metricAerospikeNamespaceScanCount:        newMetricAerospikeNamespaceScanCount(settings.AerospikeNamespaceScanCount),
		metricAerospikeNamespaceTransactionCount: newMetricAerospikeNamespaceTransactionCount(settings.AerospikeNamespaceTransactionCount),
		metricAerospikeNodeConnectionCount:       newMetricAerospikeNodeConnectionCount(settings.AerospikeNodeConnectionCount),
		metricAerospikeNodeConnectionOpen:        newMetricAerospikeNodeConnectionOpen(settings.AerospikeNodeConnectionOpen),
		metricAerospikeNodeMemoryFree:            newMetricAerospikeNodeMemoryFree(settings.AerospikeNodeMemoryFree),
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

// WithAerospikeNamespace sets provided value as "aerospike.namespace" attribute for current resource.
func WithAerospikeNamespace(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("aerospike.namespace", val)
	}
}

// WithAerospikeNodeName sets provided value as "aerospike.node.name" attribute for current resource.
func WithAerospikeNodeName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().UpsertString("aerospike.node.name", val)
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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/aerospikereceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricAerospikeNamespaceDiskAvailable.emit(ils.Metrics())
	mb.metricAerospikeNamespaceMemoryFree.emit(ils.Metrics())
	mb.metricAerospikeNamespaceMemoryUsage.emit(ils.Metrics())
	mb.metricAerospikeNamespaceScanCount.emit(ils.Metrics())
	mb.metricAerospikeNamespaceTransactionCount.emit(ils.Metrics())
	mb.metricAerospikeNodeConnectionCount.emit(ils.Metrics())
	mb.metricAerospikeNodeConnectionOpen.emit(ils.Metrics())
	mb.metricAerospikeNodeMemoryFree.emit(ils.Metrics())
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

// RecordAerospikeNamespaceDiskAvailableDataPoint adds a data point to aerospike.namespace.disk.available metric.
func (mb *MetricsBuilder) RecordAerospikeNamespaceDiskAvailableDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNamespaceDiskAvailable, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNamespaceDiskAvailable.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordAerospikeNamespaceMemoryFreeDataPoint adds a data point to aerospike.namespace.memory.free metric.
func (mb *MetricsBuilder) RecordAerospikeNamespaceMemoryFreeDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNamespaceMemoryFree, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNamespaceMemoryFree.recordDataPoint(mb.startTime, ts, val)
	return nil
}

// RecordAerospikeNamespaceMemoryUsageDataPoint adds a data point to aerospike.namespace.memory.usage metric.
func (mb *MetricsBuilder) RecordAerospikeNamespaceMemoryUsageDataPoint(ts pcommon.Timestamp, inputVal string, namespaceComponentAttributeValue AttributeNamespaceComponent) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNamespaceMemoryUsage, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNamespaceMemoryUsage.recordDataPoint(mb.startTime, ts, val, namespaceComponentAttributeValue.String())
	return nil
}

// RecordAerospikeNamespaceScanCountDataPoint adds a data point to aerospike.namespace.scan.count metric.
func (mb *MetricsBuilder) RecordAerospikeNamespaceScanCountDataPoint(ts pcommon.Timestamp, inputVal string, scanTypeAttributeValue AttributeScanType, scanResultAttributeValue AttributeScanResult) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNamespaceScanCount, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNamespaceScanCount.recordDataPoint(mb.startTime, ts, val, scanTypeAttributeValue.String(), scanResultAttributeValue.String())
	return nil
}

// RecordAerospikeNamespaceTransactionCountDataPoint adds a data point to aerospike.namespace.transaction.count metric.
func (mb *MetricsBuilder) RecordAerospikeNamespaceTransactionCountDataPoint(ts pcommon.Timestamp, inputVal string, transactionTypeAttributeValue AttributeTransactionType, transactionResultAttributeValue AttributeTransactionResult) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNamespaceTransactionCount, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNamespaceTransactionCount.recordDataPoint(mb.startTime, ts, val, transactionTypeAttributeValue.String(), transactionResultAttributeValue.String())
	return nil
}

// RecordAerospikeNodeConnectionCountDataPoint adds a data point to aerospike.node.connection.count metric.
func (mb *MetricsBuilder) RecordAerospikeNodeConnectionCountDataPoint(ts pcommon.Timestamp, inputVal string, connectionTypeAttributeValue AttributeConnectionType, connectionOpAttributeValue AttributeConnectionOp) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNodeConnectionCount, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNodeConnectionCount.recordDataPoint(mb.startTime, ts, val, connectionTypeAttributeValue.String(), connectionOpAttributeValue.String())
	return nil
}

// RecordAerospikeNodeConnectionOpenDataPoint adds a data point to aerospike.node.connection.open metric.
func (mb *MetricsBuilder) RecordAerospikeNodeConnectionOpenDataPoint(ts pcommon.Timestamp, inputVal string, connectionTypeAttributeValue AttributeConnectionType) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNodeConnectionOpen, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNodeConnectionOpen.recordDataPoint(mb.startTime, ts, val, connectionTypeAttributeValue.String())
	return nil
}

// RecordAerospikeNodeMemoryFreeDataPoint adds a data point to aerospike.node.memory.free metric.
func (mb *MetricsBuilder) RecordAerospikeNodeMemoryFreeDataPoint(ts pcommon.Timestamp, inputVal string) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for AerospikeNodeMemoryFree, value was %s: %w", inputVal, err)
	}
	mb.metricAerospikeNodeMemoryFree.recordDataPoint(mb.startTime, ts, val)
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
