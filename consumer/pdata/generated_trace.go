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
	otlptrace "go.opentelemetry.io/collector/internal/data/protogen/trace/v1"
)

// ResourceSpansSlice logically represents a slice of ResourceSpans.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpansSlice struct {
	// orig points to the slice otlptrace.ResourceSpans field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.ResourceSpans
}

func newResourceSpansSlice(orig *[]*otlptrace.ResourceSpans) ResourceSpansSlice {
	return ResourceSpansSlice{orig}
}

// NewResourceSpansSlice creates a ResourceSpansSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewResourceSpansSlice() ResourceSpansSlice {
	orig := []*otlptrace.ResourceSpans(nil)
	return ResourceSpansSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceSpansSlice()".
func (es ResourceSpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es ResourceSpansSlice) At(ix int) ResourceSpans {
	return newResourceSpans((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceSpansSlice) MoveAndAppendTo(dest ResourceSpansSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es ResourceSpansSlice) CopyTo(dest ResourceSpansSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newResourceSpans((*es.orig)[i]).CopyTo(newResourceSpans((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlptrace.ResourceSpans, srcLen)
	wrappers := make([]*otlptrace.ResourceSpans, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newResourceSpans((*es.orig)[i]).CopyTo(newResourceSpans(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new ResourceSpansSlice can be initialized:
// es := NewResourceSpansSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es ResourceSpansSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlptrace.ResourceSpans, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlptrace.ResourceSpans, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the ResourceSpansSlice by one and set the
// given ResourceSpans at that new position.  The original ResourceSpans
// could still be referenced so do not reuse it after passing it to this
// method.
func (es ResourceSpansSlice) Append(e ResourceSpans) {
	*es.orig = append(*es.orig, e.orig)
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpans struct {
	orig *otlptrace.ResourceSpans
}

func newResourceSpans(orig *otlptrace.ResourceSpans) ResourceSpans {
	return ResourceSpans{orig: orig}
}

// NewResourceSpans creates a new empty ResourceSpans.
//
// This must be used only in testing code since no "Set" method available.
func NewResourceSpans() ResourceSpans {
	return newResourceSpans(&otlptrace.ResourceSpans{})
}

// Resource returns the resource associated with this ResourceSpans.
func (ms ResourceSpans) Resource() Resource {
	return newResource(&(*ms.orig).Resource)
}

// InstrumentationLibrarySpans returns the InstrumentationLibrarySpans associated with this ResourceSpans.
func (ms ResourceSpans) InstrumentationLibrarySpans() InstrumentationLibrarySpansSlice {
	return newInstrumentationLibrarySpansSlice(&(*ms.orig).InstrumentationLibrarySpans)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms ResourceSpans) CopyTo(dest ResourceSpans) {
	ms.Resource().CopyTo(dest.Resource())
	ms.InstrumentationLibrarySpans().CopyTo(dest.InstrumentationLibrarySpans())
}

// InstrumentationLibrarySpansSlice logically represents a slice of InstrumentationLibrarySpans.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpansSlice struct {
	// orig points to the slice otlptrace.InstrumentationLibrarySpans field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.InstrumentationLibrarySpans
}

func newInstrumentationLibrarySpansSlice(orig *[]*otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpansSlice {
	return InstrumentationLibrarySpansSlice{orig}
}

// NewInstrumentationLibrarySpansSlice creates a InstrumentationLibrarySpansSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewInstrumentationLibrarySpansSlice() InstrumentationLibrarySpansSlice {
	orig := []*otlptrace.InstrumentationLibrarySpans(nil)
	return InstrumentationLibrarySpansSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewInstrumentationLibrarySpansSlice()".
func (es InstrumentationLibrarySpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es InstrumentationLibrarySpansSlice) At(ix int) InstrumentationLibrarySpans {
	return newInstrumentationLibrarySpans((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es InstrumentationLibrarySpansSlice) MoveAndAppendTo(dest InstrumentationLibrarySpansSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es InstrumentationLibrarySpansSlice) CopyTo(dest InstrumentationLibrarySpansSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newInstrumentationLibrarySpans((*es.orig)[i]).CopyTo(newInstrumentationLibrarySpans((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlptrace.InstrumentationLibrarySpans, srcLen)
	wrappers := make([]*otlptrace.InstrumentationLibrarySpans, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newInstrumentationLibrarySpans((*es.orig)[i]).CopyTo(newInstrumentationLibrarySpans(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new InstrumentationLibrarySpansSlice can be initialized:
// es := NewInstrumentationLibrarySpansSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es InstrumentationLibrarySpansSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlptrace.InstrumentationLibrarySpans, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlptrace.InstrumentationLibrarySpans, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the InstrumentationLibrarySpansSlice by one and set the
// given InstrumentationLibrarySpans at that new position.  The original InstrumentationLibrarySpans
// could still be referenced so do not reuse it after passing it to this
// method.
func (es InstrumentationLibrarySpansSlice) Append(e InstrumentationLibrarySpans) {
	*es.orig = append(*es.orig, e.orig)
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpans struct {
	orig *otlptrace.InstrumentationLibrarySpans
}

func newInstrumentationLibrarySpans(orig *otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpans {
	return InstrumentationLibrarySpans{orig: orig}
}

// NewInstrumentationLibrarySpans creates a new empty InstrumentationLibrarySpans.
//
// This must be used only in testing code since no "Set" method available.
func NewInstrumentationLibrarySpans() InstrumentationLibrarySpans {
	return newInstrumentationLibrarySpans(&otlptrace.InstrumentationLibrarySpans{})
}

// InstrumentationLibrary returns the instrumentationlibrary associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) InstrumentationLibrary() InstrumentationLibrary {
	return newInstrumentationLibrary(&(*ms.orig).InstrumentationLibrary)
}

// Spans returns the Spans associated with this InstrumentationLibrarySpans.
func (ms InstrumentationLibrarySpans) Spans() SpanSlice {
	return newSpanSlice(&(*ms.orig).Spans)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms InstrumentationLibrarySpans) CopyTo(dest InstrumentationLibrarySpans) {
	ms.InstrumentationLibrary().CopyTo(dest.InstrumentationLibrary())
	ms.Spans().CopyTo(dest.Spans())
}

// SpanSlice logically represents a slice of Span.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanSlice struct {
	// orig points to the slice otlptrace.Span field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span
}

func newSpanSlice(orig *[]*otlptrace.Span) SpanSlice {
	return SpanSlice{orig}
}

// NewSpanSlice creates a SpanSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanSlice() SpanSlice {
	orig := []*otlptrace.Span(nil)
	return SpanSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanSlice()".
func (es SpanSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanSlice) At(ix int) Span {
	return newSpan((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es SpanSlice) MoveAndAppendTo(dest SpanSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es SpanSlice) CopyTo(dest SpanSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newSpan((*es.orig)[i]).CopyTo(newSpan((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlptrace.Span, srcLen)
	wrappers := make([]*otlptrace.Span, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newSpan((*es.orig)[i]).CopyTo(newSpan(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new SpanSlice can be initialized:
// es := NewSpanSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlptrace.Span, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlptrace.Span, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the SpanSlice by one and set the
// given Span at that new position.  The original Span
// could still be referenced so do not reuse it after passing it to this
// method.
func (es SpanSlice) Append(e Span) {
	*es.orig = append(*es.orig, e.orig)
}

// Span represents a single operation within a trace.
// See Span definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto#L37
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpan function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Span struct {
	orig *otlptrace.Span
}

func newSpan(orig *otlptrace.Span) Span {
	return Span{orig: orig}
}

// NewSpan creates a new empty Span.
//
// This must be used only in testing code since no "Set" method available.
func NewSpan() Span {
	return newSpan(&otlptrace.Span{})
}

// TraceID returns the traceid associated with this Span.
func (ms Span) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this Span.
func (ms Span) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this Span.
func (ms Span) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this Span.
func (ms Span) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = data.SpanID(v)
}

// TraceState returns the tracestate associated with this Span.
func (ms Span) TraceState() TraceState {
	return TraceState((*ms.orig).TraceState)
}

// SetTraceState replaces the tracestate associated with this Span.
func (ms Span) SetTraceState(v TraceState) {
	(*ms.orig).TraceState = string(v)
}

// ParentSpanID returns the parentspanid associated with this Span.
func (ms Span) ParentSpanID() SpanID {
	return SpanID((*ms.orig).ParentSpanId)
}

// SetParentSpanID replaces the parentspanid associated with this Span.
func (ms Span) SetParentSpanID(v SpanID) {
	(*ms.orig).ParentSpanId = data.SpanID(v)
}

// Name returns the name associated with this Span.
func (ms Span) Name() string {
	return (*ms.orig).Name
}

// SetName replaces the name associated with this Span.
func (ms Span) SetName(v string) {
	(*ms.orig).Name = v
}

// Kind returns the kind associated with this Span.
func (ms Span) Kind() SpanKind {
	return SpanKind((*ms.orig).Kind)
}

// SetKind replaces the kind associated with this Span.
func (ms Span) SetKind(v SpanKind) {
	(*ms.orig).Kind = otlptrace.Span_SpanKind(v)
}

// StartTimestamp returns the starttimestamp associated with this Span.
func (ms Span) StartTimestamp() Timestamp {
	return Timestamp((*ms.orig).StartTimeUnixNano)
}

// SetStartTimestamp replaces the starttimestamp associated with this Span.
func (ms Span) SetStartTimestamp(v Timestamp) {
	(*ms.orig).StartTimeUnixNano = uint64(v)
}

// EndTimestamp returns the endtimestamp associated with this Span.
func (ms Span) EndTimestamp() Timestamp {
	return Timestamp((*ms.orig).EndTimeUnixNano)
}

// SetEndTimestamp replaces the endtimestamp associated with this Span.
func (ms Span) SetEndTimestamp(v Timestamp) {
	(*ms.orig).EndTimeUnixNano = uint64(v)
}

// Attributes returns the Attributes associated with this Span.
func (ms Span) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this Span.
func (ms Span) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this Span.
func (ms Span) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// Events returns the Events associated with this Span.
func (ms Span) Events() SpanEventSlice {
	return newSpanEventSlice(&(*ms.orig).Events)
}

// DroppedEventsCount returns the droppedeventscount associated with this Span.
func (ms Span) DroppedEventsCount() uint32 {
	return (*ms.orig).DroppedEventsCount
}

// SetDroppedEventsCount replaces the droppedeventscount associated with this Span.
func (ms Span) SetDroppedEventsCount(v uint32) {
	(*ms.orig).DroppedEventsCount = v
}

// Links returns the Links associated with this Span.
func (ms Span) Links() SpanLinkSlice {
	return newSpanLinkSlice(&(*ms.orig).Links)
}

// DroppedLinksCount returns the droppedlinkscount associated with this Span.
func (ms Span) DroppedLinksCount() uint32 {
	return (*ms.orig).DroppedLinksCount
}

// SetDroppedLinksCount replaces the droppedlinkscount associated with this Span.
func (ms Span) SetDroppedLinksCount(v uint32) {
	(*ms.orig).DroppedLinksCount = v
}

// Status returns the status associated with this Span.
func (ms Span) Status() SpanStatus {
	return newSpanStatus(&(*ms.orig).Status)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms Span) CopyTo(dest Span) {
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetTraceState(ms.TraceState())
	dest.SetParentSpanID(ms.ParentSpanID())
	dest.SetName(ms.Name())
	dest.SetKind(ms.Kind())
	dest.SetStartTimestamp(ms.StartTimestamp())
	dest.SetEndTimestamp(ms.EndTimestamp())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
	ms.Events().CopyTo(dest.Events())
	dest.SetDroppedEventsCount(ms.DroppedEventsCount())
	ms.Links().CopyTo(dest.Links())
	dest.SetDroppedLinksCount(ms.DroppedLinksCount())
	ms.Status().CopyTo(dest.Status())
}

// SpanEventSlice logically represents a slice of SpanEvent.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEventSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEventSlice struct {
	// orig points to the slice otlptrace.Span_Event field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span_Event
}

func newSpanEventSlice(orig *[]*otlptrace.Span_Event) SpanEventSlice {
	return SpanEventSlice{orig}
}

// NewSpanEventSlice creates a SpanEventSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanEventSlice() SpanEventSlice {
	orig := []*otlptrace.Span_Event(nil)
	return SpanEventSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanEventSlice()".
func (es SpanEventSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanEventSlice) At(ix int) SpanEvent {
	return newSpanEvent((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es SpanEventSlice) MoveAndAppendTo(dest SpanEventSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es SpanEventSlice) CopyTo(dest SpanEventSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newSpanEvent((*es.orig)[i]).CopyTo(newSpanEvent((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlptrace.Span_Event, srcLen)
	wrappers := make([]*otlptrace.Span_Event, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newSpanEvent((*es.orig)[i]).CopyTo(newSpanEvent(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new SpanEventSlice can be initialized:
// es := NewSpanEventSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanEventSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlptrace.Span_Event, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlptrace.Span_Event, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the SpanEventSlice by one and set the
// given SpanEvent at that new position.  The original SpanEvent
// could still be referenced so do not reuse it after passing it to this
// method.
func (es SpanEventSlice) Append(e SpanEvent) {
	*es.orig = append(*es.orig, e.orig)
}

// SpanEvent is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs. See OTLP for event definition.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEvent function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEvent struct {
	orig *otlptrace.Span_Event
}

func newSpanEvent(orig *otlptrace.Span_Event) SpanEvent {
	return SpanEvent{orig: orig}
}

// NewSpanEvent creates a new empty SpanEvent.
//
// This must be used only in testing code since no "Set" method available.
func NewSpanEvent() SpanEvent {
	return newSpanEvent(&otlptrace.Span_Event{})
}

// Timestamp returns the timestamp associated with this SpanEvent.
func (ms SpanEvent) Timestamp() Timestamp {
	return Timestamp((*ms.orig).TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this SpanEvent.
func (ms SpanEvent) SetTimestamp(v Timestamp) {
	(*ms.orig).TimeUnixNano = uint64(v)
}

// Name returns the name associated with this SpanEvent.
func (ms SpanEvent) Name() string {
	return (*ms.orig).Name
}

// SetName replaces the name associated with this SpanEvent.
func (ms SpanEvent) SetName(v string) {
	(*ms.orig).Name = v
}

// Attributes returns the Attributes associated with this SpanEvent.
func (ms SpanEvent) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanEvent.
func (ms SpanEvent) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanEvent.
func (ms SpanEvent) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms SpanEvent) CopyTo(dest SpanEvent) {
	dest.SetTimestamp(ms.Timestamp())
	dest.SetName(ms.Name())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}

// SpanLinkSlice logically represents a slice of SpanLink.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLinkSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLinkSlice struct {
	// orig points to the slice otlptrace.Span_Link field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span_Link
}

func newSpanLinkSlice(orig *[]*otlptrace.Span_Link) SpanLinkSlice {
	return SpanLinkSlice{orig}
}

// NewSpanLinkSlice creates a SpanLinkSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanLinkSlice() SpanLinkSlice {
	orig := []*otlptrace.Span_Link(nil)
	return SpanLinkSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanLinkSlice()".
func (es SpanLinkSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanLinkSlice) At(ix int) SpanLink {
	return newSpanLink((*es.orig)[ix])
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es SpanLinkSlice) MoveAndAppendTo(dest SpanLinkSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// CopyTo copies all elements from the current slice to the dest.
func (es SpanLinkSlice) CopyTo(dest SpanLinkSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newSpanLink((*es.orig)[i]).CopyTo(newSpanLink((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlptrace.Span_Link, srcLen)
	wrappers := make([]*otlptrace.Span_Link, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newSpanLink((*es.orig)[i]).CopyTo(newSpanLink(wrappers[i]))
	}
	*dest.orig = wrappers
}

// Resize is an operation that resizes the slice:
// 1. If the newLen <= len then equivalent with slice[0:newLen:cap].
// 2. If the newLen > len then (newLen - cap) empty elements will be appended to the slice.
//
// Here is how a new SpanLinkSlice can be initialized:
// es := NewSpanLinkSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanLinkSlice) Resize(newLen int) {
	oldLen := len(*es.orig)
	oldCap := cap(*es.orig)
	if newLen <= oldLen {
		*es.orig = (*es.orig)[:newLen:oldCap]
		return
	}

	if newLen > oldCap {
		newOrig := make([]*otlptrace.Span_Link, oldLen, newLen)
		copy(newOrig, *es.orig)
		*es.orig = newOrig
	}

	// Add extra empty elements to the array.
	extraOrigs := make([]otlptrace.Span_Link, newLen-oldLen)
	for i := range extraOrigs {
		*es.orig = append(*es.orig, &extraOrigs[i])
	}
}

// Append will increase the length of the SpanLinkSlice by one and set the
// given SpanLink at that new position.  The original SpanLink
// could still be referenced so do not reuse it after passing it to this
// method.
func (es SpanLinkSlice) Append(e SpanLink) {
	*es.orig = append(*es.orig, e.orig)
}

// SpanLink is a pointer from the current span to another span in the same trace or in a
// different trace. See OTLP for link definition.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLink struct {
	orig *otlptrace.Span_Link
}

func newSpanLink(orig *otlptrace.Span_Link) SpanLink {
	return SpanLink{orig: orig}
}

// NewSpanLink creates a new empty SpanLink.
//
// This must be used only in testing code since no "Set" method available.
func NewSpanLink() SpanLink {
	return newSpanLink(&otlptrace.Span_Link{})
}

// TraceID returns the traceid associated with this SpanLink.
func (ms SpanLink) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this SpanLink.
func (ms SpanLink) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this SpanLink.
func (ms SpanLink) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this SpanLink.
func (ms SpanLink) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = data.SpanID(v)
}

// TraceState returns the tracestate associated with this SpanLink.
func (ms SpanLink) TraceState() TraceState {
	return TraceState((*ms.orig).TraceState)
}

// SetTraceState replaces the tracestate associated with this SpanLink.
func (ms SpanLink) SetTraceState(v TraceState) {
	(*ms.orig).TraceState = string(v)
}

// Attributes returns the Attributes associated with this SpanLink.
func (ms SpanLink) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanLink.
func (ms SpanLink) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanLink.
func (ms SpanLink) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms SpanLink) CopyTo(dest SpanLink) {
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetTraceState(ms.TraceState())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}

// SpanStatus is an optional final status for this span. Semantically when Status wasn't set
// it is means span ended without errors and assume Status.Ok (code = 0).
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanStatus function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanStatus struct {
	orig *otlptrace.Status
}

func newSpanStatus(orig *otlptrace.Status) SpanStatus {
	return SpanStatus{orig: orig}
}

// NewSpanStatus creates a new empty SpanStatus.
//
// This must be used only in testing code since no "Set" method available.
func NewSpanStatus() SpanStatus {
	return newSpanStatus(&otlptrace.Status{})
}

// Code returns the code associated with this SpanStatus.
func (ms SpanStatus) Code() StatusCode {
	return StatusCode((*ms.orig).Code)
}

// Message returns the message associated with this SpanStatus.
func (ms SpanStatus) Message() string {
	return (*ms.orig).Message
}

// SetMessage replaces the message associated with this SpanStatus.
func (ms SpanStatus) SetMessage(v string) {
	(*ms.orig).Message = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms SpanStatus) CopyTo(dest SpanStatus) {
	dest.SetCode(ms.Code())
	dest.SetMessage(ms.Message())
}
