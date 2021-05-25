// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sampling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

// TestStringAttributeCfg is replicated with StringAttributeCfg
type TestStringAttributeCfg struct {
	Key                  string
	Values               []string
	EnabledRegexMatching bool
	CacheMaxSize         int
}

func TestStringTagFilter(t *testing.T) {

	var empty = map[string]pdata.AttributeValue{}

	cases := []struct {
		Desc      string
		Trace     *TraceData
		filterCfg *TestStringAttributeCfg
		Decision  Decision
	}{
		{
			Desc:      "nonmatching node attribute key",
			Trace:     newTraceStringAttrs(map[string]pdata.AttributeValue{"non_matching": pdata.NewAttributeValueString("value")}, "", ""),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  NotSampled,
		},
		{
			Desc:      "nonmatching node attribute value",
			Trace:     newTraceStringAttrs(map[string]pdata.AttributeValue{"example": pdata.NewAttributeValueString("non_matching")}, "", ""),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  NotSampled,
		},
		{
			Desc:      "matching node attribute",
			Trace:     newTraceStringAttrs(map[string]pdata.AttributeValue{"example": pdata.NewAttributeValueString("value")}, "", ""),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  Sampled,
		},
		{
			Desc:      "nonmatching span attribute key",
			Trace:     newTraceStringAttrs(empty, "nonmatching", "value"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  NotSampled,
		},
		{
			Desc:      "nonmatching span attribute value",
			Trace:     newTraceStringAttrs(empty, "example", "nonmatching"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  NotSampled,
		},
		{
			Desc:      "matching span attribute",
			Trace:     newTraceStringAttrs(empty, "example", "value"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: false, CacheMaxSize: defaultCacheSize},
			Decision:  Sampled,
		},
		{
			Desc:      "matching span attribute with regex",
			Trace:     newTraceStringAttrs(empty, "example", "grpc.health.v1.HealthCheck"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"v[0-9]+.HealthCheck$"}, EnabledRegexMatching: true, CacheMaxSize: defaultCacheSize},
			Decision:  Sampled,
		},
		{
			Desc:      "nonmatching span attribute with regex",
			Trace:     newTraceStringAttrs(empty, "example", "grpc.health.v1.HealthCheck"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"v[a-z]+.HealthCheck$"}, EnabledRegexMatching: true, CacheMaxSize: defaultCacheSize},
			Decision:  NotSampled,
		},
		{
			Desc:      "matching span attribute with regex without CacheSize provided in config",
			Trace:     newTraceStringAttrs(empty, "example", "grpc.health.v1.HealthCheck"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"v[0-9]+.HealthCheck$"}, EnabledRegexMatching: true},
			Decision:  Sampled,
		},
		{
			Desc:      "matching plain text node attribute in regex",
			Trace:     newTraceStringAttrs(map[string]pdata.AttributeValue{"example": pdata.NewAttributeValueString("value")}, "", ""),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{"value"}, EnabledRegexMatching: true, CacheMaxSize: defaultCacheSize},
			Decision:  Sampled,
		},
		{
			Desc:      "nonmatching span attribute on empty filter list",
			Trace:     newTraceStringAttrs(empty, "example", "grpc.health.v1.HealthCheck"),
			filterCfg: &TestStringAttributeCfg{Key: "example", Values: []string{}, EnabledRegexMatching: true},
			Decision:  NotSampled,
		},
	}

	for _, c := range cases {
		t.Run(c.Desc, func(t *testing.T) {
			filter := NewStringAttributeFilter(zap.NewNop(), c.filterCfg.Key, c.filterCfg.Values, c.filterCfg.EnabledRegexMatching, c.filterCfg.CacheMaxSize)
			decision, err := filter.Evaluate(pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}), c.Trace)
			assert.NoError(t, err)
			assert.Equal(t, decision, c.Decision)
		})
	}
}

func BenchmarkStringTagFilterEvaluatePlainText(b *testing.B) {
	trace := newTraceStringAttrs(map[string]pdata.AttributeValue{"example": pdata.NewAttributeValueString("value")}, "", "")
	filter := NewStringAttributeFilter(zap.NewNop(), "example", []string{"value"}, false, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Evaluate(pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}), trace)
	}
}

func BenchmarkStringTagFilterEvaluateRegex(b *testing.B) {
	trace := newTraceStringAttrs(map[string]pdata.AttributeValue{"example": pdata.NewAttributeValueString("grpc.health.v1.HealthCheck")}, "", "")
	filter := NewStringAttributeFilter(zap.NewNop(), "example", []string{"v[0-9]+.HealthCheck$"}, true, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Evaluate(pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}), trace)
	}
}

func newTraceStringAttrs(nodeAttrs map[string]pdata.AttributeValue, spanAttrKey string, spanAttrValue string) *TraceData {
	var traceBatches []pdata.Traces
	traces := pdata.NewTraces()
	rs := traces.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().InitFromMap(nodeAttrs)
	ils := rs.InstrumentationLibrarySpans().AppendEmpty()
	span := ils.Spans().AppendEmpty()
	span.SetTraceID(pdata.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}))
	span.SetSpanID(pdata.NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	attributes := make(map[string]pdata.AttributeValue)
	attributes[spanAttrKey] = pdata.NewAttributeValueString(spanAttrValue)
	span.Attributes().InitFromMap(attributes)
	traceBatches = append(traceBatches, traces)
	return &TraceData{
		ReceivedBatches: traceBatches,
	}
}

func TestOnLateArrivingSpans_StringAttribute(t *testing.T) {
	filter := NewStringAttributeFilter(zap.NewNop(), "example", []string{"value"}, false, defaultCacheSize)
	err := filter.OnLateArrivingSpans(NotSampled, nil)
	assert.Nil(t, err)
}
