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
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/pdata"
)

// Type is the component type name.
const Type configmodels.Type = "hostmetricsreceiver"

type metricIntf interface {
	Name() string
	New() pdata.Metric
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name    string
	newFunc func() pdata.Metric
}

func (m *metricImpl) Name() string {
	return m.name
}

func (m *metricImpl) New() pdata.Metric {
	return m.newFunc()
}

type metricStruct struct {
	SystemCPUTime     metricIntf
	SystemMemoryUsage metricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"system.cpu.time",
		"system.memory.usage",
	}
}

var metricsByName = map[string]metricIntf{
	"system.cpu.time":     Metrics.SystemCPUTime,
	"system.memory.usage": Metrics.SystemMemoryUsage,
}

func (m *metricStruct) ByName(n string) metricIntf {
	return metricsByName[n]
}

func (m *metricStruct) FactoriesByName() map[string]func() pdata.Metric {
	return map[string]func() pdata.Metric{
		Metrics.SystemCPUTime.Name():     Metrics.SystemCPUTime.New,
		Metrics.SystemMemoryUsage.Name(): Metrics.SystemMemoryUsage.New,
	}
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"system.cpu.time",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("system.cpu.time")
			metric.SetDescription("Total CPU seconds broken down by different states.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeDoubleSum)
			data := metric.DoubleSum()
			data.SetIsMonotonic(true)
			data.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

			return metric
		},
	},
	&metricImpl{
		"system.memory.usage",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("system.memory.usage")
			metric.SetDescription("Bytes of memory in use.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			data := metric.IntSum()
			data.SetIsMonotonic(false)
			data.SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

			return metric
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Cpu (CPU number starting at 0.)
	Cpu string
	// CPUState (Breakdown of CPU usage by type.)
	CPUState string
	// MemState (Breakdown of memory usage by type.)
	MemState string
}{
	"cpu",
	"state",
	"state",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelCPUState are the possible values that the label "cpu.state" can have.
var LabelCPUState = struct {
	Idle      string
	Interrupt string
	Nice      string
	Softirq   string
	Steal     string
	System    string
	User      string
	Wait      string
}{
	"idle",
	"interrupt",
	"nice",
	"softirq",
	"steal",
	"system",
	"user",
	"wait",
}

// LabelMemState are the possible values that the label "mem.state" can have.
var LabelMemState = struct {
	Buffered          string
	Cached            string
	Inactive          string
	Free              string
	SlabReclaimable   string
	SlabUnreclaimable string
	Used              string
}{
	"buffered",
	"cached",
	"inactive",
	"free",
	"slab_reclaimable",
	"slab_unreclaimable",
	"used",
}
