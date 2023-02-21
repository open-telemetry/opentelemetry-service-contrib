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

package ptracetest // import "github.com/asserts/opentelemetry-collector-contrib/pkg/pdatatest/ptracetest"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/asserts/opentelemetry-collector-contrib/pkg/pdatatest/internal"
	"github.com/asserts/opentelemetry-collector-contrib/pkg/pdatautil"
)

// CompareTracesOption can be used to mutate expected and/or actual traces before comparing.
type CompareTracesOption interface {
	applyOnTraces(expected, actual ptrace.Traces)
}

type compareTracesOptionFunc func(expected, actual ptrace.Traces)

func (f compareTracesOptionFunc) applyOnTraces(expected, actual ptrace.Traces) {
	f(expected, actual)
}

// IgnoreResourceAttributeValue is a CompareTracesOption that removes a resource attribute
// from all resources.
func IgnoreResourceAttributeValue(attributeName string) CompareTracesOption {
	return compareTracesOptionFunc(func(expected, actual ptrace.Traces) {
		maskTracesResourceAttributeValue(expected, attributeName)
		maskTracesResourceAttributeValue(actual, attributeName)
	})
}

func maskTracesResourceAttributeValue(traces ptrace.Traces, attributeName string) {
	rss := traces.ResourceSpans()
	for i := 0; i < rss.Len(); i++ {
		internal.MaskResourceAttributeValue(rss.At(i).Resource(), attributeName)
	}
}

// IgnoreResourceSpansOrder is a CompareTracesOption that ignores the order of resource traces/metrics/logs.
func IgnoreResourceSpansOrder() CompareTracesOption {
	return compareTracesOptionFunc(func(expected, actual ptrace.Traces) {
		sortResourceSpansSlice(expected.ResourceSpans())
		sortResourceSpansSlice(actual.ResourceSpans())
	})
}

func sortResourceSpansSlice(rms ptrace.ResourceSpansSlice) {
	rms.Sort(func(a, b ptrace.ResourceSpans) bool {
		if a.SchemaUrl() != b.SchemaUrl() {
			return a.SchemaUrl() < b.SchemaUrl()
		}
		aAttrs := pdatautil.MapHash(a.Resource().Attributes())
		bAttrs := pdatautil.MapHash(b.Resource().Attributes())
		return bytes.Compare(aAttrs[:], bAttrs[:]) < 0
	})
}

// IgnoreScopeSpansOrder is a CompareTracesOption that ignores the order of instrumentation scope traces/metrics/logs.
func IgnoreScopeSpansOrder() CompareTracesOption {
	return compareTracesOptionFunc(func(expected, actual ptrace.Traces) {
		sortScopeSpansSlices(expected)
		sortScopeSpansSlices(actual)
	})
}

func sortScopeSpansSlices(ts ptrace.Traces) {
	for i := 0; i < ts.ResourceSpans().Len(); i++ {
		ts.ResourceSpans().At(i).ScopeSpans().Sort(func(a, b ptrace.ScopeSpans) bool {
			if a.SchemaUrl() != b.SchemaUrl() {
				return a.SchemaUrl() < b.SchemaUrl()
			}
			if a.Scope().Name() != b.Scope().Name() {
				return a.Scope().Name() < b.Scope().Name()
			}
			return a.Scope().Version() < b.Scope().Version()
		})
	}
}

// IgnoreSpansOrder is a CompareTracesOption that ignores the order of spans.
func IgnoreSpansOrder() CompareTracesOption {
	return compareTracesOptionFunc(func(expected, actual ptrace.Traces) {
		sortSpanSlices(expected)
		sortSpanSlices(actual)
	})
}

func sortSpanSlices(ts ptrace.Traces) {
	for i := 0; i < ts.ResourceSpans().Len(); i++ {
		for j := 0; j < ts.ResourceSpans().At(i).ScopeSpans().Len(); j++ {
			ts.ResourceSpans().At(i).ScopeSpans().At(j).Spans().Sort(func(a, b ptrace.Span) bool {
				if a.Kind() != b.Kind() {
					return a.Kind() < b.Kind()
				}
				if a.Name() != b.Name() {
					return a.Name() < b.Name()
				}
				at := a.TraceID()
				bt := b.TraceID()
				if !bytes.Equal(at[:], bt[:]) {
					return bytes.Compare(at[:], bt[:]) < 0
				}
				as := a.SpanID()
				bs := b.SpanID()
				if !bytes.Equal(as[:], bs[:]) {
					return bytes.Compare(as[:], bs[:]) < 0
				}
				aps := a.ParentSpanID()
				bps := b.ParentSpanID()
				if !bytes.Equal(aps[:], bps[:]) {
					return bytes.Compare(aps[:], bps[:]) < 0
				}
				aAttrs := pdatautil.MapHash(a.Attributes())
				bAttrs := pdatautil.MapHash(b.Attributes())
				if !bytes.Equal(aAttrs[:], bAttrs[:]) {
					return bytes.Compare(aAttrs[:], bAttrs[:]) < 0
				}
				if a.StartTimestamp() != b.StartTimestamp() {
					return a.StartTimestamp() < b.StartTimestamp()
				}
				return a.EndTimestamp() < b.EndTimestamp()
			})
		}
	}
}
