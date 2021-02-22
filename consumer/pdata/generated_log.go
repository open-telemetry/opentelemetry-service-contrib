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

// Code generated by "cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run cmd/pdatagen/main.go".

package pdata

import (
	"go.opentelemetry.io/collector/internal/data"
	otlplogs "go.opentelemetry.io/collector/internal/data/protogen/logs/v1"
)

// ResourceLogsSlice logically represents a slice of ResourceLogs.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogsSlice struct {
	// orig points to the slice otlplogs.ResourceLogs field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlplogs.ResourceLogs
}

func newResourceLogsSlice(orig *[]*otlplogs.ResourceLogs) ResourceLogsSlice {
	return ResourceLogsSlice{orig}
}

// NewResourceLogsSlice creates a ResourceLogsSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewResourceLogsSlice() ResourceLogsSlice {
	orig := []*otlplogs.ResourceLogs(nil)
	return ResourceLogsSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceLogsSlice()".
func (es ResourceLogsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es ResourceLogsSlice) At(ix int) ResourceLogs {
	return newResourceLogs((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceLogsSlice) MoveAndAppendTo(dest ResourceLogsSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es ResourceLogsSlice) CopyTo(dest ResourceLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newResourceLogs((*es.orig)[i]).CopyTo(newResourceLogs((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.ResourceLogs, srcLen)
	wrappers := make([]*otlplogs.ResourceLogs, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newResourceLogs((*es.orig)[i]).CopyTo(newResourceLogs(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new ResourceLogsSlice can be initialized:
// es := NewResourceLogsSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es ResourceLogsSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlplogs.ResourceLogs, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlplogs.ResourceLogs, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the ResourceLogsSlice by one and set the
// given ResourceLogs at that new position.  The original ResourceLogs
// could still be referenced so do not reuse it after passing it to this
// method.
func (es ResourceLogsSlice) Append(e ResourceLogs) {
	*es.orig = append(*es.orig, e.orig)
}

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogs struct {
	orig *otlplogs.ResourceLogs
}

func newResourceLogs(orig *otlplogs.ResourceLogs) ResourceLogs {
	return ResourceLogs{orig: orig}
}

// NewResourceLogs creates a new empty ResourceLogs.
//
// This must be used only in testing code since no "Set" method available.
func NewResourceLogs() ResourceLogs {
	return newResourceLogs(&otlplogs.ResourceLogs{})
}

// Resource returns the resource associated with this ResourceLogs.
func (ms ResourceLogs) Resource() Resource {
	return newResource(&(*ms.orig).Resource)
}

// InstrumentationLibraryLogs returns the InstrumentationLibraryLogs associated with this ResourceLogs.
func (ms ResourceLogs) InstrumentationLibraryLogs() InstrumentationLibraryLogsSlice {
	return newInstrumentationLibraryLogsSlice(&(*ms.orig).InstrumentationLibraryLogs)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms ResourceLogs) CopyTo(dest ResourceLogs) {
	ms.Resource().CopyTo(dest.Resource())
	ms.InstrumentationLibraryLogs().CopyTo(dest.InstrumentationLibraryLogs())
}

// InstrumentationLibraryLogsSlice logically represents a slice of InstrumentationLibraryLogs.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibraryLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibraryLogsSlice struct {
	// orig points to the slice otlplogs.InstrumentationLibraryLogs field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlplogs.InstrumentationLibraryLogs
}

func newInstrumentationLibraryLogsSlice(orig *[]*otlplogs.InstrumentationLibraryLogs) InstrumentationLibraryLogsSlice {
	return InstrumentationLibraryLogsSlice{orig}
}

// NewInstrumentationLibraryLogsSlice creates a InstrumentationLibraryLogsSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewInstrumentationLibraryLogsSlice() InstrumentationLibraryLogsSlice {
	orig := []*otlplogs.InstrumentationLibraryLogs(nil)
	return InstrumentationLibraryLogsSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewInstrumentationLibraryLogsSlice()".
func (es InstrumentationLibraryLogsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es InstrumentationLibraryLogsSlice) At(ix int) InstrumentationLibraryLogs {
	return newInstrumentationLibraryLogs((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es InstrumentationLibraryLogsSlice) MoveAndAppendTo(dest InstrumentationLibraryLogsSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es InstrumentationLibraryLogsSlice) CopyTo(dest InstrumentationLibraryLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newInstrumentationLibraryLogs((*es.orig)[i]).CopyTo(newInstrumentationLibraryLogs((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.InstrumentationLibraryLogs, srcLen)
	wrappers := make([]*otlplogs.InstrumentationLibraryLogs, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newInstrumentationLibraryLogs((*es.orig)[i]).CopyTo(newInstrumentationLibraryLogs(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new InstrumentationLibraryLogsSlice can be initialized:
// es := NewInstrumentationLibraryLogsSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es InstrumentationLibraryLogsSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlplogs.InstrumentationLibraryLogs, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlplogs.InstrumentationLibraryLogs, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the InstrumentationLibraryLogsSlice by one and set the
// given InstrumentationLibraryLogs at that new position.  The original InstrumentationLibraryLogs
// could still be referenced so do not reuse it after passing it to this
// method.
func (es InstrumentationLibraryLogsSlice) Append(e InstrumentationLibraryLogs) {
	*es.orig = append(*es.orig, e.orig)
}

// InstrumentationLibraryLogs is a collection of logs from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibraryLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibraryLogs struct {
	orig *otlplogs.InstrumentationLibraryLogs
}

func newInstrumentationLibraryLogs(orig *otlplogs.InstrumentationLibraryLogs) InstrumentationLibraryLogs {
	return InstrumentationLibraryLogs{orig: orig}
}

// NewInstrumentationLibraryLogs creates a new empty InstrumentationLibraryLogs.
//
// This must be used only in testing code since no "Set" method available.
func NewInstrumentationLibraryLogs() InstrumentationLibraryLogs {
	return newInstrumentationLibraryLogs(&otlplogs.InstrumentationLibraryLogs{})
}

// InstrumentationLibrary returns the instrumentationlibrary associated with this InstrumentationLibraryLogs.
func (ms InstrumentationLibraryLogs) InstrumentationLibrary() InstrumentationLibrary {
	return newInstrumentationLibrary(&(*ms.orig).InstrumentationLibrary)
}

// Logs returns the Logs associated with this InstrumentationLibraryLogs.
func (ms InstrumentationLibraryLogs) Logs() LogSlice {
	return newLogSlice(&(*ms.orig).Logs)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms InstrumentationLibraryLogs) CopyTo(dest InstrumentationLibraryLogs) {
	ms.InstrumentationLibrary().CopyTo(dest.InstrumentationLibrary())
	ms.Logs().CopyTo(dest.Logs())
}

// LogSlice logically represents a slice of LogRecord.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogSlice struct {
	// orig points to the slice otlplogs.LogRecord field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlplogs.LogRecord
}

func newLogSlice(orig *[]*otlplogs.LogRecord) LogSlice {
	return LogSlice{orig}
}

// NewLogSlice creates a LogSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewLogSlice() LogSlice {
	orig := []*otlplogs.LogRecord(nil)
	return LogSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewLogSlice()".
func (es LogSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es LogSlice) At(ix int) LogRecord {
	return newLogRecord((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es LogSlice) MoveAndAppendTo(dest LogSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es LogSlice) CopyTo(dest LogSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newLogRecord((*es.orig)[i]).CopyTo(newLogRecord((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.LogRecord, srcLen)
	wrappers := make([]*otlplogs.LogRecord, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newLogRecord((*es.orig)[i]).CopyTo(newLogRecord(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new LogSlice can be initialized:
// es := NewLogSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es LogSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlplogs.LogRecord, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlplogs.LogRecord, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the LogSlice by one and set the
// given LogRecord at that new position.  The original LogRecord
// could still be referenced so do not reuse it after passing it to this
// method.
func (es LogSlice) Append(e LogRecord) {
	*es.orig = append(*es.orig, e.orig)
}

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecord struct {
	orig *otlplogs.LogRecord
}

func newLogRecord(orig *otlplogs.LogRecord) LogRecord {
	return LogRecord{orig: orig}
}

// NewLogRecord creates a new empty LogRecord.
//
// This must be used only in testing code since no "Set" method available.
func NewLogRecord() LogRecord {
	return newLogRecord(&otlplogs.LogRecord{})
}

// Timestamp returns the timestamp associated with this LogRecord.
func (ms LogRecord) Timestamp() TimestampUnixNano {
	return TimestampUnixNano((*ms.orig).TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this LogRecord.
func (ms LogRecord) SetTimestamp(v TimestampUnixNano) {
	(*ms.orig).TimeUnixNano = uint64(v)
}

// TraceID returns the traceid associated with this LogRecord.
func (ms LogRecord) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this LogRecord.
func (ms LogRecord) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this LogRecord.
func (ms LogRecord) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this LogRecord.
func (ms LogRecord) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = data.SpanID(v)
}

// Flags returns the flags associated with this LogRecord.
func (ms LogRecord) Flags() uint32 {
	return uint32((*ms.orig).Flags)
}

// SetFlags replaces the flags associated with this LogRecord.
func (ms LogRecord) SetFlags(v uint32) {
	(*ms.orig).Flags = uint32(v)
}

// SeverityText returns the severitytext associated with this LogRecord.
func (ms LogRecord) SeverityText() string {
	return (*ms.orig).SeverityText
}

// SetSeverityText replaces the severitytext associated with this LogRecord.
func (ms LogRecord) SetSeverityText(v string) {
	(*ms.orig).SeverityText = v
}

// SeverityNumber returns the severitynumber associated with this LogRecord.
func (ms LogRecord) SeverityNumber() SeverityNumber {
	return SeverityNumber((*ms.orig).SeverityNumber)
}

// SetSeverityNumber replaces the severitynumber associated with this LogRecord.
func (ms LogRecord) SetSeverityNumber(v SeverityNumber) {
	(*ms.orig).SeverityNumber = otlplogs.SeverityNumber(v)
}

// Name returns the name associated with this LogRecord.
func (ms LogRecord) Name() string {
	return (*ms.orig).Name
}

// SetName replaces the name associated with this LogRecord.
func (ms LogRecord) SetName(v string) {
	(*ms.orig).Name = v
}

// Body returns the body associated with this LogRecord.
func (ms LogRecord) Body() AttributeValue {
	return newAttributeValue(&(*ms.orig).Body)
}

// Attributes returns the Attributes associated with this LogRecord.
func (ms LogRecord) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this LogRecord.
func (ms LogRecord) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this LogRecord.
func (ms LogRecord) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms LogRecord) CopyTo(dest LogRecord) {
	dest.SetTimestamp(ms.Timestamp())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetFlags(ms.Flags())
	dest.SetSeverityText(ms.SeverityText())
	dest.SetSeverityNumber(ms.SeverityNumber())
	dest.SetName(ms.Name())
	ms.Body().CopyTo(dest.Body())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}
