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

// Int64DataPoint is a single data point in a timeseries that describes the time-varying values of a int64 metric.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInt64DataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Int64DataPoint struct {
	// Wrap OTLP otlpmetrics.Int64DataPoint.
	orig *otlpmetrics.Int64DataPoint
}

// NewInt64DataPoint creates a new Int64DataPoint.
func NewInt64DataPoint() Int64DataPoint {
	return Int64DataPoint{&otlpmetrics.Int64DataPoint{}}
}

func newInt64DataPoint(orig *otlpmetrics.Int64DataPoint) Int64DataPoint {
	return Int64DataPoint{orig}
}

// LabelsMap returns the Labels associated with this Int64DataPoint.
func (ms Int64DataPoint) LabelsMap() StringMap {
	return newStringMap(&ms.orig.Labels)
}

// SetLabelsMap replaces the Labels associated with this Int64DataPoint.
func (ms Int64DataPoint) SetLabelsMap(v StringMap) {
	ms.orig.Labels = *v.orig
}

// StartTime returns the starttime associated with this Int64DataPoint.
func (ms Int64DataPoint) StartTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetStartTimeUnixnano())
}

// SetStartTime replaces the starttime associated with this Int64DataPoint.
func (ms Int64DataPoint) SetStartTime(v TimestampUnixNano) {
	ms.orig.StartTimeUnixnano = uint64(v)
}

// Timestamp returns the timestamp associated with this Int64DataPoint.
func (ms Int64DataPoint) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetTimestampUnixnano())
}

// SetTimestamp replaces the timestamp associated with this Int64DataPoint.
func (ms Int64DataPoint) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimestampUnixnano = uint64(v)
}

// Value returns the value associated with this Int64DataPoint.
func (ms Int64DataPoint) Value() int64 {
	return ms.orig.GetValue()
}

// SetValue replaces the value associated with this Int64DataPoint.
func (ms Int64DataPoint) SetValue(v int64) {
	ms.orig.Value = v
}

// DoubleDataPoint is a single data point in a timeseries that describes the time-varying value of a double metric.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewDoubleDataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type DoubleDataPoint struct {
	// Wrap OTLP otlpmetrics.DoubleDataPoint.
	orig *otlpmetrics.DoubleDataPoint
}

// NewDoubleDataPoint creates a new DoubleDataPoint.
func NewDoubleDataPoint() DoubleDataPoint {
	return DoubleDataPoint{&otlpmetrics.DoubleDataPoint{}}
}

func newDoubleDataPoint(orig *otlpmetrics.DoubleDataPoint) DoubleDataPoint {
	return DoubleDataPoint{orig}
}

// LabelsMap returns the Labels associated with this DoubleDataPoint.
func (ms DoubleDataPoint) LabelsMap() StringMap {
	return newStringMap(&ms.orig.Labels)
}

// SetLabelsMap replaces the Labels associated with this DoubleDataPoint.
func (ms DoubleDataPoint) SetLabelsMap(v StringMap) {
	ms.orig.Labels = *v.orig
}

// StartTime returns the starttime associated with this DoubleDataPoint.
func (ms DoubleDataPoint) StartTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetStartTimeUnixnano())
}

// SetStartTime replaces the starttime associated with this DoubleDataPoint.
func (ms DoubleDataPoint) SetStartTime(v TimestampUnixNano) {
	ms.orig.StartTimeUnixnano = uint64(v)
}

// Timestamp returns the timestamp associated with this DoubleDataPoint.
func (ms DoubleDataPoint) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetTimestampUnixnano())
}

// SetTimestamp replaces the timestamp associated with this DoubleDataPoint.
func (ms DoubleDataPoint) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimestampUnixnano = uint64(v)
}

// Value returns the value associated with this DoubleDataPoint.
func (ms DoubleDataPoint) Value() float64 {
	return ms.orig.GetValue()
}

// SetValue replaces the value associated with this DoubleDataPoint.
func (ms DoubleDataPoint) SetValue(v float64) {
	ms.orig.Value = v
}

// HistogramDataPoint is a single data point in a timeseries that describes the time-varying values of a Histogram.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewHistogramDataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type HistogramDataPoint struct {
	// Wrap OTLP otlpmetrics.HistogramDataPoint.
	orig *otlpmetrics.HistogramDataPoint
}

// NewHistogramDataPoint creates a new HistogramDataPoint.
func NewHistogramDataPoint() HistogramDataPoint {
	return HistogramDataPoint{&otlpmetrics.HistogramDataPoint{}}
}

func newHistogramDataPoint(orig *otlpmetrics.HistogramDataPoint) HistogramDataPoint {
	return HistogramDataPoint{orig}
}

// LabelsMap returns the Labels associated with this HistogramDataPoint.
func (ms HistogramDataPoint) LabelsMap() StringMap {
	return newStringMap(&ms.orig.Labels)
}

// SetLabelsMap replaces the Labels associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetLabelsMap(v StringMap) {
	ms.orig.Labels = *v.orig
}

// StartTime returns the starttime associated with this HistogramDataPoint.
func (ms HistogramDataPoint) StartTime() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetStartTimeUnixnano())
}

// SetStartTime replaces the starttime associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetStartTime(v TimestampUnixNano) {
	ms.orig.StartTimeUnixnano = uint64(v)
}

// Timestamp returns the timestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetTimestampUnixnano())
}

// SetTimestamp replaces the timestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimestampUnixnano = uint64(v)
}

// Count returns the count associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Count() uint64 {
	return ms.orig.GetCount()
}

// SetCount replaces the count associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetCount(v uint64) {
	ms.orig.Count = v
}

// Sum returns the sum associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Sum() float64 {
	return ms.orig.GetSum()
}

// SetSum replaces the sum associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetSum(v float64) {
	ms.orig.Sum = v
}

// Buckets returns the Buckets associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Buckets() HistogramBucketSlice {
	return newHistogramBucketSlice(&ms.orig.Buckets)
}

// SetBuckets replaces the Buckets associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetBuckets(v HistogramBucketSlice) {
	ms.orig.Buckets = *v.orig
}

// ExplicitBounds returns the explicitbounds associated with this HistogramDataPoint.
func (ms HistogramDataPoint) ExplicitBounds() []float64 {
	return ms.orig.GetExplicitBounds()
}

// SetExplicitBounds replaces the explicitbounds associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetExplicitBounds(v []float64) {
	ms.orig.ExplicitBounds = v
}

// HistogramBucketSlice logically represents a slice of HistogramBucket.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewHistogramBucketSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type HistogramBucketSlice struct {
	orig *[]*otlpmetrics.HistogramDataPoint_Bucket
}

// NewHistogramBucketSlice creates a HistogramBucketSlice with "len" empty elements.
//
// es := NewHistogramBucketSlice(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     // Here should set all the values for e.
// }
func NewHistogramBucketSlice(len int) HistogramBucketSlice {
	// Slice for underlying orig.
	origs := make([]otlpmetrics.HistogramDataPoint_Bucket, len)
	// Slice for wrappers.
	wrappers := make([]*otlpmetrics.HistogramDataPoint_Bucket, len)
	for i := range origs {
		wrappers[i] = &origs[i]
	}
	return HistogramBucketSlice{&wrappers}
}

func newHistogramBucketSlice(orig *[]*otlpmetrics.HistogramDataPoint_Bucket) HistogramBucketSlice {
	return HistogramBucketSlice{orig}
}

// Len returns the number of elements in the slice.
func (es HistogramBucketSlice) Len() int {
	return len(*es.orig)
}

// Get returns the element associated with the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.Get(i)
//     ... // Do something with the element
// }
func (es HistogramBucketSlice) Get(ix int) HistogramBucket {
	return newHistogramBucket((*es.orig)[ix])
}

// Remove removes the element from the given index from the slice.
func (es HistogramBucketSlice) Remove(ix int) {
	(*es.orig)[ix] = (*es.orig)[len(*es.orig)-1]
	*es.orig = (*es.orig)[:len(*es.orig)-1]
}

// Resize resizes the slice. This operation is equivalent with slice[to:from].
func (es HistogramBucketSlice) Resize(from, to int) {
	*es.orig = (*es.orig)[from:to]
}

// HistogramBucket contains values for a histogram bucket.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewHistogramBucket function to create new instances.
// Important: zero-initialized instance is not valid for use.
type HistogramBucket struct {
	// Wrap OTLP otlpmetrics.HistogramDataPoint_Bucket.
	orig *otlpmetrics.HistogramDataPoint_Bucket
}

// NewHistogramBucket creates a new HistogramBucket.
func NewHistogramBucket() HistogramBucket {
	return HistogramBucket{&otlpmetrics.HistogramDataPoint_Bucket{}}
}

func newHistogramBucket(orig *otlpmetrics.HistogramDataPoint_Bucket) HistogramBucket {
	return HistogramBucket{orig}
}

// Count returns the count associated with this HistogramBucket.
func (ms HistogramBucket) Count() uint64 {
	return ms.orig.GetCount()
}

// SetCount replaces the count associated with this HistogramBucket.
func (ms HistogramBucket) SetCount(v uint64) {
	ms.orig.Count = v
}

// Exemplar returns the exemplar associated with this HistogramBucket.
func (ms HistogramBucket) Exemplar() HistogramBucketExemplar {
	if ms.orig.Exemplar == nil {
		// No Exemplar available, initialize one to make all operations on HistogramBucketExemplar available.
		ms.orig.Exemplar = &otlpmetrics.HistogramDataPoint_Bucket_Exemplar{}
	}
	return newHistogramBucketExemplar(ms.orig.Exemplar)
}

// SetExemplar replaces the exemplar associated with this HistogramBucket.
func (ms HistogramBucket) SetExemplar(v HistogramBucketExemplar) {
	ms.orig.Exemplar = v.orig
}

// HistogramBucketExemplar are example points that may be used to annotate aggregated Histogram values.
// They are metadata that gives information about a particular value added to a Histogram bucket.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewHistogramBucketExemplar function to create new instances.
// Important: zero-initialized instance is not valid for use.
type HistogramBucketExemplar struct {
	// Wrap OTLP otlpmetrics.HistogramDataPoint_Bucket_Exemplar.
	orig *otlpmetrics.HistogramDataPoint_Bucket_Exemplar
}

// NewHistogramBucketExemplar creates a new HistogramBucketExemplar.
func NewHistogramBucketExemplar() HistogramBucketExemplar {
	return HistogramBucketExemplar{&otlpmetrics.HistogramDataPoint_Bucket_Exemplar{}}
}

func newHistogramBucketExemplar(orig *otlpmetrics.HistogramDataPoint_Bucket_Exemplar) HistogramBucketExemplar {
	return HistogramBucketExemplar{orig}
}

// Timestamp returns the timestamp associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) Timestamp() TimestampUnixNano {
	return TimestampUnixNano(ms.orig.GetTimestampUnixnano())
}

// SetTimestamp replaces the timestamp associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) SetTimestamp(v TimestampUnixNano) {
	ms.orig.TimestampUnixnano = uint64(v)
}

// Value returns the value associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) Value() float64 {
	return ms.orig.GetValue()
}

// SetValue replaces the value associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) SetValue(v float64) {
	ms.orig.Value = v
}

// Attachments returns the Attachments associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) Attachments() StringMap {
	return newStringMap(&ms.orig.Attachments)
}

// SetAttachments replaces the Attachments associated with this HistogramBucketExemplar.
func (ms HistogramBucketExemplar) SetAttachments(v StringMap) {
	ms.orig.Attachments = *v.orig
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

// LabelsMap returns the Labels associated with this SummaryDataPoint.
func (ms SummaryDataPoint) LabelsMap() StringMap {
	return newStringMap(&ms.orig.Labels)
}

// SetLabelsMap replaces the Labels associated with this SummaryDataPoint.
func (ms SummaryDataPoint) SetLabelsMap(v StringMap) {
	ms.orig.Labels = *v.orig
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

// SummaryValueAtPercentile represents the value at a given percentile of a distribution.
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

