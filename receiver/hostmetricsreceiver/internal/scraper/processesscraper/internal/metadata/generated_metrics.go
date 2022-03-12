// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "hostmetricsreceiver/processes"

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
	SystemProcessesCount   MetricIntf
	SystemProcessesCreated MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"system.processes.count",
		"system.processes.created",
	}
}

var metricsByName = map[string]MetricIntf{
	"system.processes.count":   Metrics.SystemProcessesCount,
	"system.processes.created": Metrics.SystemProcessesCreated,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"system.processes.count",
		func(metric pdata.Metric) {
			metric.SetName("system.processes.count")
			metric.SetDescription("Total number of processes in each state.")
			metric.SetUnit("{processes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.processes.created",
		func(metric pdata.Metric) {
			metric.SetName("system.processes.created")
			metric.SetDescription("Total number of created processes.")
			metric.SetUnit("{processes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Status (Breakdown status of the processes.)
	Status string
}{
	"status",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeStatus are the possible values that the attribute "status" can have.
var AttributeStatus = struct {
	Blocked  string
	Daemon   string
	Detached string
	Idle     string
	Locked   string
	Orphan   string
	Paging   string
	Running  string
	Sleeping string
	Stopped  string
	System   string
	Unknown  string
	Zombies  string
}{
	"blocked",
	"daemon",
	"detached",
	"idle",
	"locked",
	"orphan",
	"paging",
	"running",
	"sleeping",
	"stopped",
	"system",
	"unknown",
	"zombies",
}
