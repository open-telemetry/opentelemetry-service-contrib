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
	"testing"

	"github.com/stretchr/testify/assert"

	otlptrace "go.opentelemetry.io/collector/model/internal/data/protogen/trace/v1"
)

func TestResourceSpansSlice(t *testing.T) {
	es := NewResourceSpansSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newResourceSpansSlice(&[]*otlptrace.ResourceSpans{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newResourceSpans(&otlptrace.ResourceSpans{})
	testVal := generateTestResourceSpans()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestResourceSpans(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestResourceSpansSlice_CopyTo(t *testing.T) {
	dest := NewResourceSpansSlice()
	// Test CopyTo to empty
	NewResourceSpansSlice().CopyTo(dest)
	assert.EqualValues(t, NewResourceSpansSlice(), dest)

	// Test CopyTo larger slice
	generateTestResourceSpansSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceSpansSlice(), dest)

	// Test CopyTo same size slice
	generateTestResourceSpansSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceSpansSlice(), dest)
}

func TestResourceSpansSlice_EnsureCapacity(t *testing.T) {
	es := generateTestResourceSpansSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.ResourceSpans]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.ResourceSpans]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.ResourceSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.ResourceSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestResourceSpansSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestResourceSpansSlice()
	dest := NewResourceSpansSlice()
	src := generateTestResourceSpansSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceSpansSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceSpansSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestResourceSpansSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceSpansSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewResourceSpansSlice()
	emptySlice.RemoveIf(func(el ResourceSpans) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestResourceSpansSlice()
	pos := 0
	filtered.RemoveIf(func(el ResourceSpans) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestResourceSpans_CopyTo(t *testing.T) {
	ms := NewResourceSpans()
	generateTestResourceSpans().CopyTo(ms)
	assert.EqualValues(t, generateTestResourceSpans(), ms)
}

func TestResourceSpans_Resource(t *testing.T) {
	ms := NewResourceSpans()
	fillTestResource(ms.Resource())
	assert.EqualValues(t, generateTestResource(), ms.Resource())
}

func TestResourceSpans_SchemaUrl(t *testing.T) {
	ms := NewResourceSpans()
	assert.EqualValues(t, "", ms.SchemaUrl())
	testValSchemaUrl := "https://opentelemetry.io/schemas/1.5.0"
	ms.SetSchemaUrl(testValSchemaUrl)
	assert.EqualValues(t, testValSchemaUrl, ms.SchemaUrl())
}

func TestResourceSpans_InstrumentationLibrarySpans(t *testing.T) {
	ms := NewResourceSpans()
	assert.EqualValues(t, NewInstrumentationLibrarySpansSlice(), ms.InstrumentationLibrarySpans())
	fillTestInstrumentationLibrarySpansSlice(ms.InstrumentationLibrarySpans())
	testValInstrumentationLibrarySpans := generateTestInstrumentationLibrarySpansSlice()
	assert.EqualValues(t, testValInstrumentationLibrarySpans, ms.InstrumentationLibrarySpans())
}

func TestInstrumentationLibrarySpansSlice(t *testing.T) {
	es := NewInstrumentationLibrarySpansSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newInstrumentationLibrarySpansSlice(&[]*otlptrace.InstrumentationLibrarySpans{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newInstrumentationLibrarySpans(&otlptrace.InstrumentationLibrarySpans{})
	testVal := generateTestInstrumentationLibrarySpans()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestInstrumentationLibrarySpans(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestInstrumentationLibrarySpansSlice_CopyTo(t *testing.T) {
	dest := NewInstrumentationLibrarySpansSlice()
	// Test CopyTo to empty
	NewInstrumentationLibrarySpansSlice().CopyTo(dest)
	assert.EqualValues(t, NewInstrumentationLibrarySpansSlice(), dest)

	// Test CopyTo larger slice
	generateTestInstrumentationLibrarySpansSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibrarySpansSlice(), dest)

	// Test CopyTo same size slice
	generateTestInstrumentationLibrarySpansSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibrarySpansSlice(), dest)
}

func TestInstrumentationLibrarySpansSlice_EnsureCapacity(t *testing.T) {
	es := generateTestInstrumentationLibrarySpansSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.InstrumentationLibrarySpans]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.InstrumentationLibrarySpans]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.InstrumentationLibrarySpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.InstrumentationLibrarySpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestInstrumentationLibrarySpansSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestInstrumentationLibrarySpansSlice()
	dest := NewInstrumentationLibrarySpansSlice()
	src := generateTestInstrumentationLibrarySpansSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibrarySpansSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestInstrumentationLibrarySpansSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestInstrumentationLibrarySpansSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestInstrumentationLibrarySpansSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewInstrumentationLibrarySpansSlice()
	emptySlice.RemoveIf(func(el InstrumentationLibrarySpans) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestInstrumentationLibrarySpansSlice()
	pos := 0
	filtered.RemoveIf(func(el InstrumentationLibrarySpans) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestInstrumentationLibrarySpans_CopyTo(t *testing.T) {
	ms := NewInstrumentationLibrarySpans()
	generateTestInstrumentationLibrarySpans().CopyTo(ms)
	assert.EqualValues(t, generateTestInstrumentationLibrarySpans(), ms)
}

func TestInstrumentationLibrarySpans_InstrumentationLibrary(t *testing.T) {
	ms := NewInstrumentationLibrarySpans()
	fillTestInstrumentationLibrary(ms.InstrumentationLibrary())
	assert.EqualValues(t, generateTestInstrumentationLibrary(), ms.InstrumentationLibrary())
}

func TestInstrumentationLibrarySpans_SchemaUrl(t *testing.T) {
	ms := NewInstrumentationLibrarySpans()
	assert.EqualValues(t, "", ms.SchemaUrl())
	testValSchemaUrl := "https://opentelemetry.io/schemas/1.5.0"
	ms.SetSchemaUrl(testValSchemaUrl)
	assert.EqualValues(t, testValSchemaUrl, ms.SchemaUrl())
}

func TestInstrumentationLibrarySpans_Spans(t *testing.T) {
	ms := NewInstrumentationLibrarySpans()
	assert.EqualValues(t, NewSpanSlice(), ms.Spans())
	fillTestSpanSlice(ms.Spans())
	testValSpans := generateTestSpanSlice()
	assert.EqualValues(t, testValSpans, ms.Spans())
}

func TestSpanSlice(t *testing.T) {
	es := NewSpanSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newSpanSlice(&[]*otlptrace.Span{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpan(&otlptrace.Span{})
	testVal := generateTestSpan()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestSpan(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestSpanSlice_CopyTo(t *testing.T) {
	dest := NewSpanSlice()
	// Test CopyTo to empty
	NewSpanSlice().CopyTo(dest)
	assert.EqualValues(t, NewSpanSlice(), dest)

	// Test CopyTo larger slice
	generateTestSpanSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanSlice(), dest)

	// Test CopyTo same size slice
	generateTestSpanSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanSlice(), dest)
}

func TestSpanSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSpanSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.Span]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestSpanSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSpanSlice()
	dest := NewSpanSlice()
	src := generateTestSpanSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSpanSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanSlice()
	emptySlice.RemoveIf(func(el Span) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSpanSlice()
	pos := 0
	filtered.RemoveIf(func(el Span) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpan_CopyTo(t *testing.T) {
	ms := NewSpan()
	generateTestSpan().CopyTo(ms)
	assert.EqualValues(t, generateTestSpan(), ms)
}

func TestSpan_TraceID(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewTraceID([16]byte{}), ms.TraceID())
	testValTraceID := NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	ms.SetTraceID(testValTraceID)
	assert.EqualValues(t, testValTraceID, ms.TraceID())
}

func TestSpan_SpanID(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewSpanID([8]byte{}), ms.SpanID())
	testValSpanID := NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
	ms.SetSpanID(testValSpanID)
	assert.EqualValues(t, testValSpanID, ms.SpanID())
}

func TestSpan_TraceState(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, TraceState(""), ms.TraceState())
	testValTraceState := TraceState("congo=congos")
	ms.SetTraceState(testValTraceState)
	assert.EqualValues(t, testValTraceState, ms.TraceState())
}

func TestSpan_ParentSpanID(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewSpanID([8]byte{}), ms.ParentSpanID())
	testValParentSpanID := NewSpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	ms.SetParentSpanID(testValParentSpanID)
	assert.EqualValues(t, testValParentSpanID, ms.ParentSpanID())
}

func TestSpan_Name(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())
}

func TestSpan_Kind(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, SpanKindUnspecified, ms.Kind())
	testValKind := SpanKindServer
	ms.SetKind(testValKind)
	assert.EqualValues(t, testValKind, ms.Kind())
}

func TestSpan_StartTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, Timestamp(0), ms.StartTimestamp())
	testValStartTimestamp := Timestamp(1234567890)
	ms.SetStartTimestamp(testValStartTimestamp)
	assert.EqualValues(t, testValStartTimestamp, ms.StartTimestamp())
}

func TestSpan_EndTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, Timestamp(0), ms.EndTimestamp())
	testValEndTimestamp := Timestamp(1234567890)
	ms.SetEndTimestamp(testValEndTimestamp)
	assert.EqualValues(t, testValEndTimestamp, ms.EndTimestamp())
}

func TestSpan_Attributes(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewAttributeMap(), ms.Attributes())
	fillTestAttributeMap(ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	assert.EqualValues(t, testValAttributes, ms.Attributes())
}

func TestSpan_DroppedAttributesCount(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())
}

func TestSpan_Events(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewSpanEventSlice(), ms.Events())
	fillTestSpanEventSlice(ms.Events())
	testValEvents := generateTestSpanEventSlice()
	assert.EqualValues(t, testValEvents, ms.Events())
}

func TestSpan_DroppedEventsCount(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, uint32(0), ms.DroppedEventsCount())
	testValDroppedEventsCount := uint32(17)
	ms.SetDroppedEventsCount(testValDroppedEventsCount)
	assert.EqualValues(t, testValDroppedEventsCount, ms.DroppedEventsCount())
}

func TestSpan_Links(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, NewSpanLinkSlice(), ms.Links())
	fillTestSpanLinkSlice(ms.Links())
	testValLinks := generateTestSpanLinkSlice()
	assert.EqualValues(t, testValLinks, ms.Links())
}

func TestSpan_DroppedLinksCount(t *testing.T) {
	ms := NewSpan()
	assert.EqualValues(t, uint32(0), ms.DroppedLinksCount())
	testValDroppedLinksCount := uint32(17)
	ms.SetDroppedLinksCount(testValDroppedLinksCount)
	assert.EqualValues(t, testValDroppedLinksCount, ms.DroppedLinksCount())
}

func TestSpan_Status(t *testing.T) {
	ms := NewSpan()
	fillTestSpanStatus(ms.Status())
	assert.EqualValues(t, generateTestSpanStatus(), ms.Status())
}

func TestSpanEventSlice(t *testing.T) {
	es := NewSpanEventSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newSpanEventSlice(&[]*otlptrace.Span_Event{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpanEvent(&otlptrace.Span_Event{})
	testVal := generateTestSpanEvent()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestSpanEvent(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestSpanEventSlice_CopyTo(t *testing.T) {
	dest := NewSpanEventSlice()
	// Test CopyTo to empty
	NewSpanEventSlice().CopyTo(dest)
	assert.EqualValues(t, NewSpanEventSlice(), dest)

	// Test CopyTo larger slice
	generateTestSpanEventSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanEventSlice(), dest)

	// Test CopyTo same size slice
	generateTestSpanEventSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanEventSlice(), dest)
}

func TestSpanEventSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSpanEventSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span_Event]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span_Event]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span_Event]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.Span_Event]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestSpanEventSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSpanEventSlice()
	dest := NewSpanEventSlice()
	src := generateTestSpanEventSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanEventSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanEventSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSpanEventSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanEventSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanEventSlice()
	emptySlice.RemoveIf(func(el SpanEvent) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSpanEventSlice()
	pos := 0
	filtered.RemoveIf(func(el SpanEvent) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanEvent_CopyTo(t *testing.T) {
	ms := NewSpanEvent()
	generateTestSpanEvent().CopyTo(ms)
	assert.EqualValues(t, generateTestSpanEvent(), ms)
}

func TestSpanEvent_Timestamp(t *testing.T) {
	ms := NewSpanEvent()
	assert.EqualValues(t, Timestamp(0), ms.Timestamp())
	testValTimestamp := Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())
}

func TestSpanEvent_Name(t *testing.T) {
	ms := NewSpanEvent()
	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())
}

func TestSpanEvent_Attributes(t *testing.T) {
	ms := NewSpanEvent()
	assert.EqualValues(t, NewAttributeMap(), ms.Attributes())
	fillTestAttributeMap(ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	assert.EqualValues(t, testValAttributes, ms.Attributes())
}

func TestSpanEvent_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanEvent()
	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())
}

func TestSpanLinkSlice(t *testing.T) {
	es := NewSpanLinkSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newSpanLinkSlice(&[]*otlptrace.Span_Link{})
	assert.EqualValues(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpanLink(&otlptrace.Span_Link{})
	testVal := generateTestSpanLink()
	assert.EqualValues(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.EqualValues(t, emptyVal, el)
		fillTestSpanLink(el)
		assert.EqualValues(t, testVal, el)
	}
}

func TestSpanLinkSlice_CopyTo(t *testing.T) {
	dest := NewSpanLinkSlice()
	// Test CopyTo to empty
	NewSpanLinkSlice().CopyTo(dest)
	assert.EqualValues(t, NewSpanLinkSlice(), dest)

	// Test CopyTo larger slice
	generateTestSpanLinkSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanLinkSlice(), dest)

	// Test CopyTo same size slice
	generateTestSpanLinkSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestSpanLinkSlice(), dest)
}

func TestSpanLinkSlice_EnsureCapacity(t *testing.T) {
	es := generateTestSpanLinkSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span_Link]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span_Link]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span_Link]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.Span_Link]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
}

func TestSpanLinkSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestSpanLinkSlice()
	dest := NewSpanLinkSlice()
	src := generateTestSpanLinkSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanLinkSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestSpanLinkSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestSpanLinkSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanLinkSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanLinkSlice()
	emptySlice.RemoveIf(func(el SpanLink) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestSpanLinkSlice()
	pos := 0
	filtered.RemoveIf(func(el SpanLink) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanLink_CopyTo(t *testing.T) {
	ms := NewSpanLink()
	generateTestSpanLink().CopyTo(ms)
	assert.EqualValues(t, generateTestSpanLink(), ms)
}

func TestSpanLink_TraceID(t *testing.T) {
	ms := NewSpanLink()
	assert.EqualValues(t, NewTraceID([16]byte{}), ms.TraceID())
	testValTraceID := NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	ms.SetTraceID(testValTraceID)
	assert.EqualValues(t, testValTraceID, ms.TraceID())
}

func TestSpanLink_SpanID(t *testing.T) {
	ms := NewSpanLink()
	assert.EqualValues(t, NewSpanID([8]byte{}), ms.SpanID())
	testValSpanID := NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
	ms.SetSpanID(testValSpanID)
	assert.EqualValues(t, testValSpanID, ms.SpanID())
}

func TestSpanLink_TraceState(t *testing.T) {
	ms := NewSpanLink()
	assert.EqualValues(t, TraceState(""), ms.TraceState())
	testValTraceState := TraceState("congo=congos")
	ms.SetTraceState(testValTraceState)
	assert.EqualValues(t, testValTraceState, ms.TraceState())
}

func TestSpanLink_Attributes(t *testing.T) {
	ms := NewSpanLink()
	assert.EqualValues(t, NewAttributeMap(), ms.Attributes())
	fillTestAttributeMap(ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	assert.EqualValues(t, testValAttributes, ms.Attributes())
}

func TestSpanLink_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanLink()
	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())
}

func TestSpanStatus_CopyTo(t *testing.T) {
	ms := NewSpanStatus()
	generateTestSpanStatus().CopyTo(ms)
	assert.EqualValues(t, generateTestSpanStatus(), ms)
}

func TestSpanStatus_Code(t *testing.T) {
	ms := NewSpanStatus()
	assert.EqualValues(t, StatusCode(0), ms.Code())
	testValCode := StatusCode(1)
	ms.SetCode(testValCode)
	assert.EqualValues(t, testValCode, ms.Code())
}

func TestSpanStatus_Message(t *testing.T) {
	ms := NewSpanStatus()
	assert.EqualValues(t, "", ms.Message())
	testValMessage := "cancelled"
	ms.SetMessage(testValMessage)
	assert.EqualValues(t, testValMessage, ms.Message())
}

func generateTestResourceSpansSlice() ResourceSpansSlice {
	tv := NewResourceSpansSlice()
	fillTestResourceSpansSlice(tv)
	return tv
}

func fillTestResourceSpansSlice(tv ResourceSpansSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestResourceSpans(tv.AppendEmpty())
	}
}

func generateTestResourceSpans() ResourceSpans {
	tv := NewResourceSpans()
	fillTestResourceSpans(tv)
	return tv
}

func fillTestResourceSpans(tv ResourceSpans) {
	fillTestResource(tv.Resource())
	tv.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	fillTestInstrumentationLibrarySpansSlice(tv.InstrumentationLibrarySpans())
}

func generateTestInstrumentationLibrarySpansSlice() InstrumentationLibrarySpansSlice {
	tv := NewInstrumentationLibrarySpansSlice()
	fillTestInstrumentationLibrarySpansSlice(tv)
	return tv
}

func fillTestInstrumentationLibrarySpansSlice(tv InstrumentationLibrarySpansSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestInstrumentationLibrarySpans(tv.AppendEmpty())
	}
}

func generateTestInstrumentationLibrarySpans() InstrumentationLibrarySpans {
	tv := NewInstrumentationLibrarySpans()
	fillTestInstrumentationLibrarySpans(tv)
	return tv
}

func fillTestInstrumentationLibrarySpans(tv InstrumentationLibrarySpans) {
	fillTestInstrumentationLibrary(tv.InstrumentationLibrary())
	tv.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	fillTestSpanSlice(tv.Spans())
}

func generateTestSpanSlice() SpanSlice {
	tv := NewSpanSlice()
	fillTestSpanSlice(tv)
	return tv
}

func fillTestSpanSlice(tv SpanSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestSpan(tv.AppendEmpty())
	}
}

func generateTestSpan() Span {
	tv := NewSpan()
	fillTestSpan(tv)
	return tv
}

func fillTestSpan(tv Span) {
	tv.SetTraceID(NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	tv.SetSpanID(NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	tv.SetTraceState(TraceState("congo=congos"))
	tv.SetParentSpanID(NewSpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	tv.SetName("test_name")
	tv.SetKind(SpanKindServer)
	tv.SetStartTimestamp(Timestamp(1234567890))
	tv.SetEndTimestamp(Timestamp(1234567890))
	fillTestAttributeMap(tv.Attributes())
	tv.SetDroppedAttributesCount(uint32(17))
	fillTestSpanEventSlice(tv.Events())
	tv.SetDroppedEventsCount(uint32(17))
	fillTestSpanLinkSlice(tv.Links())
	tv.SetDroppedLinksCount(uint32(17))
	fillTestSpanStatus(tv.Status())
}

func generateTestSpanEventSlice() SpanEventSlice {
	tv := NewSpanEventSlice()
	fillTestSpanEventSlice(tv)
	return tv
}

func fillTestSpanEventSlice(tv SpanEventSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestSpanEvent(tv.AppendEmpty())
	}
}

func generateTestSpanEvent() SpanEvent {
	tv := NewSpanEvent()
	fillTestSpanEvent(tv)
	return tv
}

func fillTestSpanEvent(tv SpanEvent) {
	tv.SetTimestamp(Timestamp(1234567890))
	tv.SetName("test_name")
	fillTestAttributeMap(tv.Attributes())
	tv.SetDroppedAttributesCount(uint32(17))
}

func generateTestSpanLinkSlice() SpanLinkSlice {
	tv := NewSpanLinkSlice()
	fillTestSpanLinkSlice(tv)
	return tv
}

func fillTestSpanLinkSlice(tv SpanLinkSlice) {
	l := 7
	tv.EnsureCapacity(l)
	for i := 0; i < l; i++ {
		fillTestSpanLink(tv.AppendEmpty())
	}
}

func generateTestSpanLink() SpanLink {
	tv := NewSpanLink()
	fillTestSpanLink(tv)
	return tv
}

func fillTestSpanLink(tv SpanLink) {
	tv.SetTraceID(NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	tv.SetSpanID(NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	tv.SetTraceState(TraceState("congo=congos"))
	fillTestAttributeMap(tv.Attributes())
	tv.SetDroppedAttributesCount(uint32(17))
}

func generateTestSpanStatus() SpanStatus {
	tv := NewSpanStatus()
	fillTestSpanStatus(tv)
	return tv
}

func fillTestSpanStatus(tv SpanStatus) {
	tv.SetCode(StatusCode(1))
	tv.SetMessage("cancelled")
}
