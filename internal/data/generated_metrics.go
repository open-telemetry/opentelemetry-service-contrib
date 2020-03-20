// Copyright 2020 OpenTelemetry Authors
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

// Code generated by code_generator_main.go. DO NOT EDIT.
// To regenerate this file run "go run internal/data/internal/code_generator_main.go".

package data

import (
	otlpmetrics "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
)

// SummaryDataPointGeneratedSlice logically represents a slice of SummaryDataPoint.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSummaryDataPointGeneratedSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SummaryDataPointGeneratedSlice struct {
	orig *[]*otlpmetrics.SummaryDataPoint
}

// NewSummaryDataPointGeneratedSlice creates a SummaryDataPointGeneratedSlice with "len" empty elements.
//
// es := NewSummaryDataPointGeneratedSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewSummaryDataPointGeneratedSlice(len int) SummaryDataPointGeneratedSlice {
	// Slice for underlying orig.
	origs := make([]otlpmetrics.SummaryDataPoint, len)
	// Slice for wrappers.
	wrappers := make([]*otlpmetrics.SummaryDataPoint, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return SummaryDataPointGeneratedSlice{&wrappers}
}

func newSummaryDataPointGeneratedSlice(orig *[]*otlpmetrics.SummaryDataPoint) SummaryDataPointGeneratedSlice {
	return SummaryDataPointGeneratedSlice{orig}
}

// Len returns the number of elements in the slice.
func (es SummaryDataPointGeneratedSlice) Len() int {
	return len(*es.orig)
}

// Get returns the element associated with the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es SummaryDataPointGeneratedSlice) Get(ix int) SummaryDataPoint {
	return newSummaryDataPoint((*es.orig)[ix])
}

// Remove removes the element from the given index from the slice.
func (es SummaryDataPointGeneratedSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize resizes the slice. This operation is equivalent with slice[to:from].
func (es SummaryDataPointGeneratedSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// SummaryDataPoint is a single data point in a timeseries that describes the time-varying values of a Summary metric.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSummaryDataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SummaryDataPoint struct {
	// Wrap OTLP otlpmetrics.SummaryDataPoint.
	orig *otlpmetrics.SummaryDataPoint
}

// NewSummaryDataPoint creates a new SummaryDataPoint.
func NewSummaryDataPoint() SummaryDataPoint {
	return SummaryDataPoint{&otlpmetrics.SummaryDataPoint{}}
}

func newSummaryDataPoint(orig *otlpmetrics.SummaryDataPoint) SummaryDataPoint {
	return SummaryDataPoint{orig}
}

// StartTime returns the starttime associated with this SummaryDataPoint.
func (ms SummaryDataPoint) StartTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetStartTimeUnixnano())
}

// SetStartTime replaces the starttime associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetStartTime(v TimestampUnixNano) {
	ms.orig.StartTimeUnixnano = uint64(v)
}

// Timestamp returns the timestamp associated with this SummaryDataPoint.
func (ms SummaryDataPoint) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetTimestampUnixnano())
}

// SetTimestamp replaces the timestamp associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimestampUnixnano = uint64(v)
}

// Count returns the count associated with this SummaryDataPoint.
func (ms SummaryDataPoint) Count() uint64 {
	return ms.orig.GetCount()
}

// SetCount replaces the count associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetCount(v uint64) {
	ms.orig.Count = v
}

// Sum returns the sum associated with this SummaryDataPoint.
func (ms SummaryDataPoint) Sum() float64 {
	return ms.orig.GetSum()
}

// SetSum replaces the sum associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetSum(v float64) {
	ms.orig.Sum = v
}

// ValueAtPercentiles returns the PercentileValues associated with this SummaryDataPoint.
func (ms SummaryDataPoint) ValueAtPercentiles() SummaryValueAtPercentileSlice {
	return newSummaryValueAtPercentileSlice(&ms.orig.PercentileValues)
}

// SetValueAtPercentiles replaces the PercentileValues associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetValueAtPercentiles(v SummaryValueAtPercentileSlice) {
	ms.orig.PercentileValues = *v.orig
}

// LabelsMap returns the Labels associated with this SummaryDataPoint.
func (ms SummaryDataPoint) LabelsMap() StringMap {
	return newStringMap(&ms.orig.Labels)
}

// SetLabelsMap replaces the Labels associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetLabelsMap(v StringMap) {
	ms.orig.Labels = *v.orig
}

// SummaryValueAtPercentileSlice logically represents a slice of SummaryValueAtPercentile.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSummaryValueAtPercentileSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SummaryValueAtPercentileSlice struct {
	orig *[]*otlpmetrics.SummaryDataPoint_ValueAtPercentile
}

// NewSummaryValueAtPercentileSlice creates a SummaryValueAtPercentileSlice with "len" empty elements.
//
// es := NewSummaryValueAtPercentileSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewSummaryValueAtPercentileSlice(len int) SummaryValueAtPercentileSlice {
	// Slice for underlying orig.
	origs := make([]otlpmetrics.SummaryDataPoint_ValueAtPercentile, len)
	// Slice for wrappers.
	wrappers := make([]*otlpmetrics.SummaryDataPoint_ValueAtPercentile, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return SummaryValueAtPercentileSlice{&wrappers}
}

func newSummaryValueAtPercentileSlice(orig *[]*otlpmetrics.SummaryDataPoint_ValueAtPercentile) SummaryValueAtPercentileSlice {
	return SummaryValueAtPercentileSlice{orig}
}

// Len returns the number of elements in the slice.
func (es SummaryValueAtPercentileSlice) Len() int {
	return len(*es.orig)
}

// Get returns the element associated with the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es SummaryValueAtPercentileSlice) Get(ix int) SummaryValueAtPercentile {
	return newSummaryValueAtPercentile((*es.orig)[ix])
}

// Remove removes the element from the given index from the slice.
func (es SummaryValueAtPercentileSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize resizes the slice. This operation is equivalent with slice[to:from].
func (es SummaryValueAtPercentileSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// SummaryValueAtPercentile represents the value at a given percentile of a distribution..
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSummaryValueAtPercentile function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SummaryValueAtPercentile struct {
	// Wrap OTLP otlpmetrics.SummaryDataPoint_ValueAtPercentile.
	orig *otlpmetrics.SummaryDataPoint_ValueAtPercentile
}

// NewSummaryValueAtPercentile creates a new SummaryValueAtPercentile.
func NewSummaryValueAtPercentile() SummaryValueAtPercentile {
	return SummaryValueAtPercentile{&otlpmetrics.SummaryDataPoint_ValueAtPercentile{}}
}

func newSummaryValueAtPercentile(orig *otlpmetrics.SummaryDataPoint_ValueAtPercentile) SummaryValueAtPercentile {
	return SummaryValueAtPercentile{orig}
}

// Percentile returns the percentile associated with this SummaryValueAtPercentile.
func (ms SummaryValueAtPercentile) Percentile() float64 {
	return ms.orig.GetPercentile()
}

// SetPercentile replaces the percentile associated with this SummaryValueAtPercentile.
func (ms SummaryValueAtPercentile) SetPercentile(v float64) {
	ms.orig.Percentile = v
}

// Value returns the value associated with this SummaryValueAtPercentile.
func (ms SummaryValueAtPercentile) Value() float64 {
	return ms.orig.GetValue()
}

// SetValue replaces the value associated with this SummaryValueAtPercentile.
func (ms SummaryValueAtPercentile) SetValue(v float64) {
	ms.orig.Value = v
}

