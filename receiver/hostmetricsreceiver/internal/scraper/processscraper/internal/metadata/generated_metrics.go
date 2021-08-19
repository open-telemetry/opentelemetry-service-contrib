// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "process"

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
	ProcessCPUTime             MetricIntf
	ProcessDiskIo              MetricIntf
	ProcessMemoryPhysicalUsage MetricIntf
	ProcessMemoryVirtualUsage  MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"process.cpu.time",
		"process.disk.io",
		"process.memory.physical_usage",
		"process.memory.virtual_usage",
	}
}

var metricsByName = map[string]MetricIntf{
	"process.cpu.time":              Metrics.ProcessCPUTime,
	"process.disk.io":               Metrics.ProcessDiskIo,
	"process.memory.physical_usage": Metrics.ProcessMemoryPhysicalUsage,
	"process.memory.virtual_usage":  Metrics.ProcessMemoryVirtualUsage,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"process.cpu.time",
		func(metric pdata.Metric) {
			metric.SetName("process.cpu.time")
			metric.SetDescription("Total CPU seconds broken down by different states.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.disk.io",
		func(metric pdata.Metric) {
			metric.SetName("process.disk.io")
			metric.SetDescription("Disk bytes transferred.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.memory.physical_usage",
		func(metric pdata.Metric) {
			metric.SetName("process.memory.physical_usage")
			metric.SetDescription("The amount of physical memory in use.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.memory.virtual_usage",
		func(metric pdata.Metric) {
			metric.SetName("process.memory.virtual_usage")
			metric.SetDescription("Virtual memory size.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Direction (Direction of flow of bytes (read or write).)
	Direction string
	// State (Breakdown of CPU usage by type.)
	State string
}{
	"direction",
	"state",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelDirection are the possible values that the label "direction" can have.
var LabelDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}

// LabelState are the possible values that the label "state" can have.
var LabelState = struct {
	System string
	User   string
	Wait   string
}{
	"system",
	"user",
	"wait",
}
