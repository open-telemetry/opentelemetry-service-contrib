// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jaeger

import (
	"fmt"

	"github.com/jaegertracing/jaeger/model"
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"

	"github.com/open-telemetry/opentelemetry-collector/consumer/pdata"
	"github.com/open-telemetry/opentelemetry-collector/internal"
	"github.com/open-telemetry/opentelemetry-collector/translator/conventions"
	tracetranslator "github.com/open-telemetry/opentelemetry-collector/translator/trace"
)

// InternalTracesToJaegerProto translates internal trace data into the Jaeger Proto for GRPC.
// Returns slice of translated Jaeger batches and error if translation failed.
func InternalTracesToJaegerProto(td pdata.Traces) ([]*model.Batch, error) {
	resourceSpans := td.ResourceSpans()

	if resourceSpans.Len() == 0 {
		return nil, nil
	}

	batches := make([]*model.Batch, 0, resourceSpans.Len())

	for i := 0; i < resourceSpans.Len(); i++ {
		rs := resourceSpans.At(i)
		if rs.IsNil() {
			continue
		}

		batch, err := resourceSpansToJaegerProto(rs)
		if err != nil {
			return nil, err
		}
		if batch != nil {
			batches = append(batches, batch)
		}
	}

	return batches, nil
}

func resourceSpansToJaegerProto(rs pdata.ResourceSpans) (*model.Batch, error) {
	resource := rs.Resource()
	ilss := rs.InstrumentationLibrarySpans()

	if resource.IsNil() && ilss.Len() == 0 {
		return nil, nil
	}

	batch := &model.Batch{
		Process: resourceToJaegerProtoProcess(resource),
	}

	if ilss.Len() == 0 {
		return batch, nil
	}

	// Approximate the number of the spans as the number of the spans in the first
	// instrumentation library info.
	jSpans := make([]*model.Span, 0, ilss.At(0).Spans().Len())

	for i := 0; i < ilss.Len(); i++ {
		ils := ilss.At(i)
		if ils.IsNil() {
			continue
		}

		// TODO: Handle instrumentation library name and version.
		spans := ils.Spans()
		for j := 0; j < spans.Len(); j++ {
			span := spans.At(j)
			if span.IsNil() {
				continue
			}

			jSpan, err := spanToJaegerProto(span)
			if err != nil {
				return nil, err
			}
			if jSpan != nil {
				jSpans = append(jSpans, jSpan)
			}
		}
	}

	batch.Spans = jSpans

	return batch, nil
}

func resourceToJaegerProtoProcess(resource pdata.Resource) *model.Process {
	if resource.IsNil() {
		return nil
	}

	attrs := resource.Attributes()
	if attrs.Cap() == 0 {
		return nil
	}

	process := model.Process{}
	if serviceName, ok := attrs.Get(conventions.AttributeServiceName); ok {
		process.ServiceName = serviceName.StringVal()
	}
	process.Tags = resourceAttributesToJaegerProtoTags(attrs)
	return &process

}

func resourceAttributesToJaegerProtoTags(attrs pdata.AttributeMap) []model.KeyValue {
	if attrs.Cap() == 0 {
		return nil
	}

	tags := make([]model.KeyValue, 0, attrs.Cap())
	attrs.ForEach(func(key string, attr pdata.AttributeValue) {
		if key == conventions.AttributeServiceName {
			return
		}
		tags = append(tags, attributeToJaegerProtoTag(key, attr))
	})
	return tags
}

func attributesToJaegerProtoTags(attrs pdata.AttributeMap) []model.KeyValue {
	if attrs.Cap() == 0 {
		return nil
	}

	tags := make([]model.KeyValue, 0, attrs.Cap())
	attrs.ForEach(func(key string, attr pdata.AttributeValue) {
		tags = append(tags, attributeToJaegerProtoTag(key, attr))
	})
	return tags
}

func attributeToJaegerProtoTag(key string, attr pdata.AttributeValue) model.KeyValue {
	tag := model.KeyValue{Key: key}
	switch attr.Type() {
	case pdata.AttributeValueSTRING:
		// Jaeger-to-Internal maps binary tags to string attributes and encodes them as
		// base64 strings. Blindingly attempting to decode base64 seems too much.
		tag.VType = model.ValueType_STRING
		tag.VStr = attr.StringVal()
	case pdata.AttributeValueINT:
		tag.VType = model.ValueType_INT64
		tag.VInt64 = attr.IntVal()
	case pdata.AttributeValueBOOL:
		tag.VType = model.ValueType_BOOL
		tag.VBool = attr.BoolVal()
	case pdata.AttributeValueDOUBLE:
		tag.VType = model.ValueType_FLOAT64
		tag.VFloat64 = attr.DoubleVal()
	}
	return tag
}

func spanToJaegerProto(span pdata.Span) (*model.Span, error) {
	if span.IsNil() {
		return nil, nil
	}

	traceID, err := traceIDToJaegerProto(span.TraceID())
	if err != nil {
		return nil, err
	}

	spanID, err := spanIDToJaegerProto(span.SpanID())
	if err != nil {
		return nil, err
	}

	jReferences, err := makeJaegerProtoReferences(span.Links(), span.ParentSpanID(), traceID)
	if err != nil {
		return nil, fmt.Errorf("error converting span links to Jaeger references: %w", err)
	}

	tags := attributesToJaegerProtoTags(span.Attributes())
	tags = appendTagFromSpanKind(tags, span.Kind())
	tags = appendTagFromSpanStatus(tags, span.Status())
	startTime := internal.UnixNanoToTime(span.StartTime())

	return &model.Span{
		TraceID:       traceID,
		SpanID:        spanID,
		OperationName: span.Name(),
		References:    jReferences,
		StartTime:     startTime,
		Duration:      internal.UnixNanoToTime(span.EndTime()).Sub(startTime),
		Tags:          tags,
		Logs:          spanEventsToJaegerProtoLogs(span.Events()),
	}, nil
}

func traceIDToJaegerProto(traceID pdata.TraceID) (model.TraceID, error) {
	traceIDHigh, traceIDLow, err := tracetranslator.BytesToUInt64TraceID([]byte(traceID))
	if err != nil {
		return model.TraceID{}, err
	}
	if traceIDLow == 0 && traceIDHigh == 0 {
		return model.TraceID{}, errZeroTraceID
	}
	return model.TraceID{
		Low:  traceIDLow,
		High: traceIDHigh,
	}, nil
}

func spanIDToJaegerProto(spanID pdata.SpanID) (model.SpanID, error) {
	uSpanID, err := tracetranslator.BytesToUInt64SpanID([]byte(spanID))
	if err != nil {
		return model.SpanID(0), err
	}
	if uSpanID == 0 {
		return model.SpanID(0), errZeroSpanID
	}
	return model.SpanID(uSpanID), nil
}

// makeJaegerProtoReferences constructs jaeger span references based on parent span ID and span links
func makeJaegerProtoReferences(
	links pdata.SpanLinkSlice,
	parentSpanID pdata.SpanID,
	traceID model.TraceID,
) ([]model.SpanRef, error) {
	parentSpanIDSet := len(parentSpanID.Bytes()) != 0
	if !parentSpanIDSet && links.Len() == 0 {
		return nil, nil
	}

	refsCount := links.Len()
	if parentSpanIDSet {
		refsCount++
	}

	refs := make([]model.SpanRef, 0, refsCount)

	// Put parent span ID at the first place because usually backends look for it
	// as the first CHILD_OF item in the model.SpanRef slice.
	if parentSpanIDSet {
		jParentSpanID, err := spanIDToJaegerProto(parentSpanID)
		if err != nil {
			return nil, fmt.Errorf("OC incorrect parent span ID: %v", err)
		}

		refs = append(refs, model.SpanRef{
			TraceID: traceID,
			SpanID:  jParentSpanID,
			RefType: model.SpanRefType_CHILD_OF,
		})
	}

	for i := 0; i < links.Len(); i++ {
		link := links.At(i)
		if link.IsNil() {
			continue
		}

		traceID, err := traceIDToJaegerProto(link.TraceID())
		if err != nil {
			continue // skip invalid link
		}

		spanID, err := spanIDToJaegerProto(link.SpanID())
		if err != nil {
			continue // skip invalid link
		}

		refs = append(refs, model.SpanRef{
			TraceID: traceID,
			SpanID:  spanID,

			// Since Jaeger RefType is not captured in internal data,
			// use SpanRefType_FOLLOWS_FROM by default.
			// SpanRefType_CHILD_OF supposed to be set only from parentSpanID.
			RefType: model.SpanRefType_FOLLOWS_FROM,
		})
	}

	return refs, nil
}

func spanEventsToJaegerProtoLogs(events pdata.SpanEventSlice) []model.Log {
	if events.Len() == 0 {
		return nil
	}

	logs := make([]model.Log, 0, events.Len())
	for i := 0; i < events.Len(); i++ {
		event := events.At(i)
		if event.IsNil() {
			continue
		}

		logs = append(logs, model.Log{
			Timestamp: internal.UnixNanoToTime(event.Timestamp()),
			Fields:    attributesToJaegerProtoTags(event.Attributes()),
		})
	}

	return logs
}

func appendTagFromSpanKind(tags []model.KeyValue, spanKind pdata.SpanKind) []model.KeyValue {
	tag := model.KeyValue{
		Key:   tracetranslator.TagSpanKind,
		VType: model.ValueType_STRING,
	}

	switch spanKind {
	case pdata.SpanKindCLIENT:
		tag.VStr = string(tracetranslator.OpenTracingSpanKindClient)
	case pdata.SpanKindSERVER:
		tag.VStr = string(tracetranslator.OpenTracingSpanKindServer)
	case pdata.SpanKindPRODUCER:
		tag.VStr = string(tracetranslator.OpenTracingSpanKindProducer)
	case pdata.SpanKindCONSUMER:
		tag.VStr = string(tracetranslator.OpenTracingSpanKindConsumer)
	default:
		return tags
	}

	if tags == nil {
		return []model.KeyValue{tag}
	}

	return append(tags, tag)
}

func appendTagFromSpanStatus(tags []model.KeyValue, status pdata.SpanStatus) []model.KeyValue {
	if status.IsNil() {
		return tags
	}

	tags = append(tags, model.KeyValue{
		Key:    tracetranslator.TagStatusCode,
		VInt64: int64(status.Code()),
		VType:  model.ValueType_INT64,
	})

	if status.Code() != pdata.StatusCode(otlptrace.Status_Ok) {
		tags = append(tags, model.KeyValue{
			Key:   tracetranslator.TagError,
			VBool: true,
			VType: model.ValueType_BOOL,
		})
	}

	if status.Message() != "" {
		tags = append(tags, model.KeyValue{
			Key:   tracetranslator.TagStatusMsg,
			VStr:  status.Message(),
			VType: model.ValueType_STRING,
		})
	}

	return tags
}
