// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "memcachedreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	MemcachedBytes              MetricIntf
	MemcachedCommands           MetricIntf
	MemcachedConnectionsCurrent MetricIntf
	MemcachedConnectionsTotal   MetricIntf
	MemcachedCPUUsage           MetricIntf
	MemcachedCurrentItems       MetricIntf
	MemcachedEvictions          MetricIntf
	MemcachedNetwork            MetricIntf
	MemcachedOperationHitRatio  MetricIntf
	MemcachedOperations         MetricIntf
	MemcachedThreads            MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"memcached.bytes",
		"memcached.commands",
		"memcached.connections.current",
		"memcached.connections.total",
		"memcached.cpu.usage",
		"memcached.current_items",
		"memcached.evictions",
		"memcached.network",
		"memcached.operation_hit_ratio",
		"memcached.operations",
		"memcached.threads",
	}
}

var metricsByName = map[string]MetricIntf{
	"memcached.bytes":               Metrics.MemcachedBytes,
	"memcached.commands":            Metrics.MemcachedCommands,
	"memcached.connections.current": Metrics.MemcachedConnectionsCurrent,
	"memcached.connections.total":   Metrics.MemcachedConnectionsTotal,
	"memcached.cpu.usage":           Metrics.MemcachedCPUUsage,
	"memcached.current_items":       Metrics.MemcachedCurrentItems,
	"memcached.evictions":           Metrics.MemcachedEvictions,
	"memcached.network":             Metrics.MemcachedNetwork,
	"memcached.operation_hit_ratio": Metrics.MemcachedOperationHitRatio,
	"memcached.operations":          Metrics.MemcachedOperations,
	"memcached.threads":             Metrics.MemcachedThreads,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"memcached.bytes",
		func(metric pdata.Metric) {
			metric.SetName("memcached.bytes")
			metric.SetDescription("Current number of bytes used by this server to store items.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"memcached.commands",
		func(metric pdata.Metric) {
			metric.SetName("memcached.commands")
			metric.SetDescription("Commands executed.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.connections.current",
		func(metric pdata.Metric) {
			metric.SetName("memcached.connections.current")
			metric.SetDescription("The current number of open connections.")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.connections.total",
		func(metric pdata.Metric) {
			metric.SetName("memcached.connections.total")
			metric.SetDescription("Total number of connections opened since the server started running.")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.cpu.usage",
		func(metric pdata.Metric) {
			metric.SetName("memcached.cpu.usage")
			metric.SetDescription("Accumulated user and system time.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.current_items",
		func(metric pdata.Metric) {
			metric.SetName("memcached.current_items")
			metric.SetDescription("Number of items currently stored in the cache.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.evictions",
		func(metric pdata.Metric) {
			metric.SetName("memcached.evictions")
			metric.SetDescription("Cache item evictions.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.network",
		func(metric pdata.Metric) {
			metric.SetName("memcached.network")
			metric.SetDescription("Bytes transferred over the network.")
			metric.SetUnit("by")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.operation_hit_ratio",
		func(metric pdata.Metric) {
			metric.SetName("memcached.operation_hit_ratio")
			metric.SetDescription("Hit ratio for operations, expressed as a percentage value between 0.0 and 100.0.")
			metric.SetUnit("%")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"memcached.operations",
		func(metric pdata.Metric) {
			metric.SetName("memcached.operations")
			metric.SetDescription("Operation counts.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"memcached.threads",
		func(metric pdata.Metric) {
			metric.SetName("memcached.threads")
			metric.SetDescription("Number of threads used by the memcached instance.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Command (The type of command.)
	Command string
	// Direction (Direction of data flow.)
	Direction string
	// Operation (The type of operation.)
	Operation string
	// State (The type of CPU usage.)
	State string
	// Type (Result of cache request.)
	Type string
}{
	"command",
	"direction",
	"operation",
	"state",
	"type",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeCommand are the possible values that the attribute "command" can have.
var AttributeCommand = struct {
	Get   string
	Set   string
	Flush string
	Touch string
}{
	"get",
	"set",
	"flush",
	"touch",
}

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	Sent     string
	Received string
}{
	"sent",
	"received",
}

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Increment string
	Decrement string
	Get       string
}{
	"increment",
	"decrement",
	"get",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	System string
	User   string
}{
	"system",
	"user",
}

// AttributeType are the possible values that the attribute "type" can have.
var AttributeType = struct {
	Hit  string
	Miss string
}{
	"hit",
	"miss",
}
