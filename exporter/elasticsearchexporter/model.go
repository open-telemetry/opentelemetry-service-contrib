// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"math"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv "go.opentelemetry.io/collector/semconv/v1.22.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/exphistogram"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/objmodel"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/traceutil"
)

// resourceAttrsConversionMap contains conversions for resource-level attributes
// from their Semantic Conventions (SemConv) names to equivalent Elastic Common
// Schema (ECS) names.
// If the ECS field name is specified as an empty string (""), the converter will
// neither convert the SemConv key to the equivalent ECS name nor pass-through the
// SemConv key as-is to become the ECS name.
var resourceAttrsConversionMap = map[string]string{
	semconv.AttributeServiceInstanceID:      "service.node.name",
	semconv.AttributeDeploymentEnvironment:  "service.environment",
	semconv.AttributeTelemetrySDKName:       "",
	semconv.AttributeTelemetrySDKLanguage:   "",
	semconv.AttributeTelemetrySDKVersion:    "",
	semconv.AttributeTelemetryDistroName:    "",
	semconv.AttributeTelemetryDistroVersion: "",
	semconv.AttributeCloudPlatform:          "cloud.service.name",
	semconv.AttributeContainerImageTags:     "container.image.tag",
	semconv.AttributeHostName:               "host.hostname",
	semconv.AttributeHostArch:               "host.architecture",
	semconv.AttributeProcessExecutablePath:  "process.executable",
	semconv.AttributeProcessRuntimeName:     "service.runtime.name",
	semconv.AttributeProcessRuntimeVersion:  "service.runtime.version",
	semconv.AttributeOSName:                 "host.os.name",
	semconv.AttributeOSType:                 "host.os.platform",
	semconv.AttributeOSDescription:          "host.os.full",
	semconv.AttributeOSVersion:              "host.os.version",
	semconv.AttributeK8SDeploymentName:      "kubernetes.deployment.name",
	semconv.AttributeK8SNamespaceName:       "kubernetes.namespace",
	semconv.AttributeK8SNodeName:            "kubernetes.node.name",
	semconv.AttributeK8SPodName:             "kubernetes.pod.name",
	semconv.AttributeK8SPodUID:              "kubernetes.pod.uid",
}

// resourceAttrsToPreserve contains conventions that should be preserved in ECS mode.
// This can happen when an attribute needs to be mapped to an ECS equivalent but
// at the same time be preserved to its original form.
var resourceAttrsToPreserve = map[string]bool{
	semconv.AttributeHostName: true,
}

type mappingModel interface {
	encodeLog(pcommon.Resource, string, plog.LogRecord, pcommon.InstrumentationScope, string) ([]byte, error)
	encodeSpan(pcommon.Resource, string, ptrace.Span, pcommon.InstrumentationScope, string) ([]byte, error)
	upsertMetricDataPointValue(map[uint32]objmodel.Document, pcommon.Resource, string, pcommon.InstrumentationScope, string, pmetric.Metric, dataPoint, pcommon.Value) error
	encodeDocument(objmodel.Document) ([]byte, error)
}

// encodeModel tries to keep the event as close to the original open telemetry semantics as is.
// No fields will be mapped by default.
//
// Field deduplication and dedotting of attributes is supported by the encodeModel.
//
// See: https://github.com/open-telemetry/oteps/blob/master/text/logs/0097-log-data-model.md
type encodeModel struct {
	dedot bool
	mode  MappingMode
}

type dataPoint interface {
	Timestamp() pcommon.Timestamp
	StartTimestamp() pcommon.Timestamp
	Attributes() pcommon.Map
}

const (
	traceIDField   = "traceID"
	spanIDField    = "spanID"
	attributeField = "attribute"
)

func (m *encodeModel) encodeLog(resource pcommon.Resource, resourceSchemaURL string, record plog.LogRecord, scope pcommon.InstrumentationScope, scopeSchemaURL string) ([]byte, error) {
	var document objmodel.Document
	switch m.mode {
	case MappingECS:
		document = m.encodeLogECSMode(resource, record, scope)
	case MappingOTel:
		document = m.encodeLogOTelMode(resource, resourceSchemaURL, record, scope, scopeSchemaURL)
	default:
		document = m.encodeLogDefaultMode(resource, record, scope)
	}
	document.Dedup()

	var buf bytes.Buffer
	err := document.Serialize(&buf, m.dedot, m.mode == MappingOTel)
	return buf.Bytes(), err
}

func (m *encodeModel) encodeLogDefaultMode(resource pcommon.Resource, record plog.LogRecord, scope pcommon.InstrumentationScope) objmodel.Document {
	var document objmodel.Document

	docTimeStamp := record.Timestamp()
	if docTimeStamp.AsTime().UnixNano() == 0 {
		docTimeStamp = record.ObservedTimestamp()
	}
	document.AddTimestamp("@timestamp", docTimeStamp) // We use @timestamp in order to ensure that we can index if the default data stream logs template is used.
	document.AddTraceID("TraceId", record.TraceID())
	document.AddSpanID("SpanId", record.SpanID())
	document.AddInt("TraceFlags", int64(record.Flags()))
	document.AddString("SeverityText", record.SeverityText())
	document.AddInt("SeverityNumber", int64(record.SeverityNumber()))
	document.AddAttribute("Body", record.Body())
	m.encodeAttributes(&document, record.Attributes())
	document.AddAttributes("Resource", resource.Attributes())
	document.AddAttributes("Scope", scopeToAttributes(scope))

	return document
}

func (m *encodeModel) encodeLogOTelMode(resource pcommon.Resource, resourceSchemaURL string, record plog.LogRecord, scope pcommon.InstrumentationScope, scopeSchemaURL string) objmodel.Document {
	var document objmodel.Document

	docTimeStamp := record.Timestamp()
	if docTimeStamp.AsTime().UnixNano() == 0 {
		docTimeStamp = record.ObservedTimestamp()
	}

	document.AddTimestamp("@timestamp", docTimeStamp)
	document.AddTimestamp("observed_timestamp", record.ObservedTimestamp())

	document.AddTraceID("trace_id", record.TraceID())
	document.AddSpanID("span_id", record.SpanID())
	document.AddString("severity_text", record.SeverityText())
	document.AddInt("severity_number", int64(record.SeverityNumber()))
	document.AddInt("dropped_attributes_count", int64(record.DroppedAttributesCount()))

	m.encodeAttributesOTelMode(&document, record.Attributes())
	m.encodeResourceOTelMode(&document, resource, resourceSchemaURL)
	m.encodeScopeOTelMode(&document, scope, scopeSchemaURL)

	// Body
	setOTelLogBody(&document, record.Body())

	return document
}

func setOTelLogBody(doc *objmodel.Document, body pcommon.Value) {
	switch body.Type() {
	case pcommon.ValueTypeMap:
		doc.AddAttribute("body_structured", body)
	case pcommon.ValueTypeSlice:
		slice := body.Slice()
		for i := 0; i < slice.Len(); i++ {
			switch slice.At(i).Type() {
			case pcommon.ValueTypeMap, pcommon.ValueTypeSlice:
				doc.AddAttribute("body_structured", body)
				return
			}
		}

		bodyTextVal := pcommon.NewValueSlice()
		bodyTextSlice := bodyTextVal.Slice()
		bodyTextSlice.EnsureCapacity(slice.Len())

		for i := 0; i < slice.Len(); i++ {
			elem := slice.At(i)
			bodyTextSlice.AppendEmpty().SetStr(elem.AsString())
		}
		doc.AddAttribute("body_text", bodyTextVal)
	default:
		doc.AddString("body_text", body.AsString())
	}
}

func (m *encodeModel) encodeLogECSMode(resource pcommon.Resource, record plog.LogRecord, scope pcommon.InstrumentationScope) objmodel.Document {
	var document objmodel.Document

	// First, try to map resource-level attributes to ECS fields.
	encodeAttributesECSMode(&document, resource.Attributes(), resourceAttrsConversionMap, resourceAttrsToPreserve)

	// Then, try to map scope-level attributes to ECS fields.
	scopeAttrsConversionMap := map[string]string{
		// None at the moment
	}
	encodeAttributesECSMode(&document, scope.Attributes(), scopeAttrsConversionMap, resourceAttrsToPreserve)

	// Finally, try to map record-level attributes to ECS fields.
	recordAttrsConversionMap := map[string]string{
		"event.name":                         "event.action",
		semconv.AttributeExceptionMessage:    "error.message",
		semconv.AttributeExceptionStacktrace: "error.stacktrace",
		semconv.AttributeExceptionType:       "error.type",
		semconv.AttributeExceptionEscaped:    "event.error.exception.handled",
	}
	encodeAttributesECSMode(&document, record.Attributes(), recordAttrsConversionMap, resourceAttrsToPreserve)

	// Handle special cases.
	encodeLogAgentNameECSMode(&document, resource)
	encodeLogAgentVersionECSMode(&document, resource)
	encodeLogHostOsTypeECSMode(&document, resource)
	encodeLogTimestampECSMode(&document, record)
	document.AddTraceID("trace.id", record.TraceID())
	document.AddSpanID("span.id", record.SpanID())
	if n := record.SeverityNumber(); n != plog.SeverityNumberUnspecified {
		document.AddInt("event.severity", int64(record.SeverityNumber()))
	}

	document.AddString("log.level", record.SeverityText())

	if record.Body().Type() == pcommon.ValueTypeStr {
		document.AddAttribute("message", record.Body())
	}

	return document
}

func (m *encodeModel) encodeDocument(document objmodel.Document) ([]byte, error) {
	document.Dedup()

	var buf bytes.Buffer
	err := document.Serialize(&buf, m.dedot, m.mode == MappingOTel)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// upsertMetricDataPointValue upserts a datapoint value to documents which is already hashed by resource and index
func (m *encodeModel) upsertMetricDataPointValue(documents map[uint32]objmodel.Document, resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, metric pmetric.Metric, dp dataPoint, value pcommon.Value) error {
	switch m.mode {
	case MappingOTel:
		return m.upsertMetricDataPointValueOTelMode(documents, resource, resourceSchemaURL, scope, scopeSchemaURL, metric, dp, value)
	case MappingECS:
		return m.upsertMetricDataPointValueECSMode(documents, resource, resourceSchemaURL, scope, scopeSchemaURL, metric, dp, value)
	default:
		// Defaults to ECS for backward compatibility
		return m.upsertMetricDataPointValueECSMode(documents, resource, resourceSchemaURL, scope, scopeSchemaURL, metric, dp, value)
	}
}

func (m *encodeModel) upsertMetricDataPointValueECSMode(documents map[uint32]objmodel.Document, resource pcommon.Resource, _ string, _ pcommon.InstrumentationScope, _ string, metric pmetric.Metric, dp dataPoint, value pcommon.Value) error {
	hash := metricECSHash(dp.Timestamp(), dp.Attributes())
	var (
		document objmodel.Document
		ok       bool
	)
	if document, ok = documents[hash]; !ok {
		encodeAttributesECSMode(&document, resource.Attributes(), resourceAttrsConversionMap, resourceAttrsToPreserve)
		document.AddTimestamp("@timestamp", dp.Timestamp())
		document.AddAttributes("", dp.Attributes())
	}

	document.AddAttribute(metric.Name(), value)

	documents[hash] = document
	return nil
}

func (m *encodeModel) upsertMetricDataPointValueOTelMode(documents map[uint32]objmodel.Document, resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, metric pmetric.Metric, dp dataPoint, value pcommon.Value) error {
	// documents is per-resource. Therefore, there is no need to hash resource attributes
	hash := metricOTelHash(dp, scope.Attributes(), metric.Unit())
	var (
		document objmodel.Document
		ok       bool
	)
	if document, ok = documents[hash]; !ok {
		document.AddTimestamp("@timestamp", dp.Timestamp())
		if dp.StartTimestamp() != 0 {
			document.AddTimestamp("start_timestamp", dp.StartTimestamp())
		}
		document.AddString("unit", metric.Unit())

		m.encodeAttributesOTelMode(&document, dp.Attributes())
		m.encodeResourceOTelMode(&document, resource, resourceSchemaURL)
		m.encodeScopeOTelMode(&document, scope, scopeSchemaURL)
	}

	switch value.Type() {
	case pcommon.ValueTypeMap:
		m := pcommon.NewMap()
		value.Map().CopyTo(m)
		document.Add("metrics."+metric.Name(), objmodel.UnflattenableObjectValue(m))
	default:
		document.Add("metrics."+metric.Name(), objmodel.ValueFromAttribute(value))
	}
	// TODO: support quantiles
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/34561

	document.AddDynamicTemplate("metrics."+metric.Name(), metricDpToDynamicTemplate(metric, dp))
	documents[hash] = document
	return nil
}

// metricDpToDynamicTemplate returns the name of dynamic template that applies to the metric and data point,
// so that the field is indexed into Elasticsearch with the correct mapping. The name should correspond to a
// dynamic template that is defined in ES mapping, e.g.
// https://github.com/elastic/elasticsearch/blob/8.15/x-pack/plugin/core/template-resources/src/main/resources/metrics%40mappings.json
func metricDpToDynamicTemplate(metric pmetric.Metric, dp dataPoint) string {
	switch metric.Type() {
	case pmetric.MetricTypeSum:
		switch dp.(pmetric.NumberDataPoint).ValueType() {
		case pmetric.NumberDataPointValueTypeDouble:
			if metric.Sum().IsMonotonic() {
				return "counter_double"
			}
			return "gauge_double"
		case pmetric.NumberDataPointValueTypeInt:
			if metric.Sum().IsMonotonic() {
				return "counter_long"
			}
			return "gauge_long"
		default:
			return "" // NumberDataPointValueTypeEmpty should already be discarded in numberToValue
		}
	case pmetric.MetricTypeGauge:
		switch dp.(pmetric.NumberDataPoint).ValueType() {
		case pmetric.NumberDataPointValueTypeDouble:
			return "gauge_double"
		case pmetric.NumberDataPointValueTypeInt:
			return "gauge_long"
		default:
			return "" // NumberDataPointValueTypeEmpty should already be discarded in numberToValue
		}
	case pmetric.MetricTypeHistogram, pmetric.MetricTypeExponentialHistogram:
		return "histogram"
	case pmetric.MetricTypeSummary:
		return "summary_metrics"
	}
	return ""
}

func summaryToValue(dp pmetric.SummaryDataPoint) pcommon.Value {
	// TODO: Add support for quantiles
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/34561
	vm := pcommon.NewValueMap()
	m := vm.Map()
	m.PutDouble("sum", dp.Sum())
	m.PutInt("value_count", int64(dp.Count()))
	return vm
}

// exponentialHistogramToValue converts an exponential histogram data point to T-digest.
func exponentialHistogramToValue(dp pmetric.ExponentialHistogramDataPoint) (pcommon.Value, error) {
	vm := pcommon.NewValueMap()
	m := vm.Map()
	counts := m.PutEmptySlice("counts")
	values := m.PutEmptySlice("values")

	scale := int(dp.Scale())

	offset := int(dp.Negative().Offset())
	bucketCounts := dp.Negative().BucketCounts()
	for i := bucketCounts.Len() - 1; i >= 0; i-- {
		count := bucketCounts.At(i)
		if count == 0 {
			continue
		}
		lb := -exphistogram.LowerBoundary(offset+i+1, scale)
		ub := -exphistogram.LowerBoundary(offset+i, scale)
		counts.AppendEmpty().SetInt(int64(count))
		values.AppendEmpty().SetDouble(lb + (ub-lb)/2)
	}

	if zeroCount := dp.ZeroCount(); zeroCount != 0 {
		counts.AppendEmpty().SetInt(int64(zeroCount))
		// The midpoint is only non-zero when positive offset and negative offset are not the same,
		// but the midpoint between negative and positive boundaries closest to zero will not be very meaningful anyway.
		// Using a zero here instead.
		values.AppendEmpty().SetDouble(0)
	}

	offset = int(dp.Positive().Offset())
	bucketCounts = dp.Positive().BucketCounts()
	for i := 0; i < bucketCounts.Len(); i++ {
		count := bucketCounts.At(i)
		if count == 0 {
			continue
		}
		lb := exphistogram.LowerBoundary(offset+i, scale)
		ub := exphistogram.LowerBoundary(offset+i+1, scale)
		counts.AppendEmpty().SetInt(int64(count))
		values.AppendEmpty().SetDouble(lb + (ub-lb)/2)
	}

	return vm, nil
}

func histogramToValue(dp pmetric.HistogramDataPoint) (pcommon.Value, error) {
	// Histogram conversion function is from
	// https://github.com/elastic/apm-data/blob/3b28495c3cbdc0902983134276eb114231730249/input/otlp/metrics.go#L277
	bucketCounts := dp.BucketCounts()
	explicitBounds := dp.ExplicitBounds()
	if bucketCounts.Len() != explicitBounds.Len()+1 || explicitBounds.Len() == 0 {
		return pcommon.Value{}, errors.New("invalid histogram data point")
	}

	vm := pcommon.NewValueMap()
	m := vm.Map()
	counts := m.PutEmptySlice("counts")
	values := m.PutEmptySlice("values")

	values.EnsureCapacity(bucketCounts.Len())
	counts.EnsureCapacity(bucketCounts.Len())
	for i := 0; i < bucketCounts.Len(); i++ {
		count := bucketCounts.At(i)
		if count == 0 {
			continue
		}

		var value float64
		switch i {
		// (-infinity, explicit_bounds[i]]
		case 0:
			value = explicitBounds.At(i)
			if value > 0 {
				value /= 2
			}

		// (explicit_bounds[i], +infinity)
		case bucketCounts.Len() - 1:
			value = explicitBounds.At(i - 1)

		// [explicit_bounds[i-1], explicit_bounds[i])
		default:
			// Use the midpoint between the boundaries.
			value = explicitBounds.At(i-1) + (explicitBounds.At(i)-explicitBounds.At(i-1))/2.0
		}

		counts.AppendEmpty().SetInt(int64(count))
		values.AppendEmpty().SetDouble(value)
	}

	return vm, nil
}

var errInvalidNumberDataPoint = errors.New("invalid number data point")

func numberToValue(dp pmetric.NumberDataPoint) (pcommon.Value, error) {
	switch dp.ValueType() {
	case pmetric.NumberDataPointValueTypeDouble:
		value := dp.DoubleValue()
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return pcommon.Value{}, errInvalidNumberDataPoint
		}
		return pcommon.NewValueDouble(value), nil
	case pmetric.NumberDataPointValueTypeInt:
		return pcommon.NewValueInt(dp.IntValue()), nil
	}
	return pcommon.Value{}, errInvalidNumberDataPoint
}

func (m *encodeModel) encodeResourceOTelMode(document *objmodel.Document, resource pcommon.Resource, resourceSchemaURL string) {
	resourceMapVal := pcommon.NewValueMap()
	resourceMap := resourceMapVal.Map()
	if resourceSchemaURL != "" {
		resourceMap.PutStr("schema_url", resourceSchemaURL)
	}
	resourceMap.PutInt("dropped_attributes_count", int64(resource.DroppedAttributesCount()))
	resourceAttrMap := resourceMap.PutEmptyMap("attributes")
	resource.Attributes().CopyTo(resourceAttrMap)
	resourceAttrMap.RemoveIf(func(key string, _ pcommon.Value) bool {
		switch key {
		case dataStreamType, dataStreamDataset, dataStreamNamespace:
			return true
		}
		return false
	})

	document.Add("resource", objmodel.ValueFromAttribute(resourceMapVal))
}

func (m *encodeModel) encodeScopeOTelMode(document *objmodel.Document, scope pcommon.InstrumentationScope, scopeSchemaURL string) {
	scopeMapVal := pcommon.NewValueMap()
	scopeMap := scopeMapVal.Map()
	if scope.Name() != "" {
		scopeMap.PutStr("name", scope.Name())
	}
	if scope.Version() != "" {
		scopeMap.PutStr("version", scope.Version())
	}
	if scopeSchemaURL != "" {
		scopeMap.PutStr("schema_url", scopeSchemaURL)
	}
	scopeMap.PutInt("dropped_attributes_count", int64(scope.DroppedAttributesCount()))
	scopeAttrMap := scopeMap.PutEmptyMap("attributes")
	scope.Attributes().CopyTo(scopeAttrMap)
	scopeAttrMap.RemoveIf(func(key string, _ pcommon.Value) bool {
		switch key {
		case dataStreamType, dataStreamDataset, dataStreamNamespace:
			return true
		}
		return false
	})
	document.Add("scope", objmodel.ValueFromAttribute(scopeMapVal))
}

func (m *encodeModel) encodeAttributesOTelMode(document *objmodel.Document, attributeMap pcommon.Map) {
	attributeMap.RemoveIf(func(key string, val pcommon.Value) bool {
		switch key {
		case dataStreamType, dataStreamDataset, dataStreamNamespace:
			// At this point the data_stream attributes are expected to be in the record attributes,
			// updated by the router.
			// Move them to the top of the document and remove them from the record
			document.AddAttribute(key, val)
			return true
		}
		return false
	})
	document.AddAttributes("attributes", attributeMap)
}

func (m *encodeModel) encodeSpan(resource pcommon.Resource, resourceSchemaURL string, span ptrace.Span, scope pcommon.InstrumentationScope, scopeSchemaURL string) ([]byte, error) {
	var document objmodel.Document
	switch m.mode {
	case MappingOTel:
		document = m.encodeSpanOTelMode(resource, resourceSchemaURL, span, scope, scopeSchemaURL)
	default:
		document = m.encodeSpanDefaultMode(resource, span, scope)
	}
	document.Dedup()
	var buf bytes.Buffer
	err := document.Serialize(&buf, m.dedot, m.mode == MappingOTel)
	return buf.Bytes(), err
}

func (m *encodeModel) encodeSpanOTelMode(resource pcommon.Resource, resourceSchemaURL string, span ptrace.Span, scope pcommon.InstrumentationScope, scopeSchemaURL string) objmodel.Document {
	var document objmodel.Document
	document.AddTimestamp("@timestamp", span.StartTimestamp())
	document.AddTraceID("trace_id", span.TraceID())
	document.AddSpanID("span_id", span.SpanID())
	document.AddString("trace_state", span.TraceState().AsRaw())
	document.AddSpanID("parent_span_id", span.ParentSpanID())
	document.AddString("name", span.Name())
	document.AddString("kind", span.Kind().String())
	document.AddInt("duration", int64(span.EndTimestamp()-span.StartTimestamp()))

	m.encodeAttributesOTelMode(&document, span.Attributes())

	document.AddInt("dropped_attributes_count", int64(span.DroppedAttributesCount()))
	document.AddInt("dropped_events_count", int64(span.DroppedEventsCount()))

	links := pcommon.NewValueSlice()
	linkSlice := links.SetEmptySlice()
	spanLinks := span.Links()
	for i := 0; i < spanLinks.Len(); i++ {
		linkMap := linkSlice.AppendEmpty().SetEmptyMap()
		spanLink := spanLinks.At(i)
		linkMap.PutStr("trace_id", spanLink.TraceID().String())
		linkMap.PutStr("span_id", spanLink.SpanID().String())
		linkMap.PutStr("trace_state", spanLink.TraceState().AsRaw())
		mAttr := linkMap.PutEmptyMap("attributes")
		spanLink.Attributes().CopyTo(mAttr)
		linkMap.PutInt("dropped_attributes_count", int64(spanLink.DroppedAttributesCount()))
	}
	document.AddAttribute("links", links)

	document.AddInt("dropped_links_count", int64(span.DroppedLinksCount()))
	document.AddString("status.message", span.Status().Message())
	document.AddString("status.code", span.Status().Code().String())

	m.encodeResourceOTelMode(&document, resource, resourceSchemaURL)
	m.encodeScopeOTelMode(&document, scope, scopeSchemaURL)

	// TODO: add span events to log data streams

	return document
}

func (m *encodeModel) encodeSpanDefaultMode(resource pcommon.Resource, span ptrace.Span, scope pcommon.InstrumentationScope) objmodel.Document {
	var document objmodel.Document
	document.AddTimestamp("@timestamp", span.StartTimestamp()) // We use @timestamp in order to ensure that we can index if the default data stream logs template is used.
	document.AddTimestamp("EndTimestamp", span.EndTimestamp())
	document.AddTraceID("TraceId", span.TraceID())
	document.AddSpanID("SpanId", span.SpanID())
	document.AddSpanID("ParentSpanId", span.ParentSpanID())
	document.AddString("Name", span.Name())
	document.AddString("Kind", traceutil.SpanKindStr(span.Kind()))
	document.AddInt("TraceStatus", int64(span.Status().Code()))
	document.AddString("TraceStatusDescription", span.Status().Message())
	document.AddString("Link", spanLinksToString(span.Links()))
	m.encodeAttributes(&document, span.Attributes())
	document.AddAttributes("Resource", resource.Attributes())
	m.encodeEvents(&document, span.Events())
	document.AddInt("Duration", durationAsMicroseconds(span.StartTimestamp().AsTime(), span.EndTimestamp().AsTime())) // unit is microseconds
	document.AddAttributes("Scope", scopeToAttributes(scope))
	return document
}

func (m *encodeModel) encodeAttributes(document *objmodel.Document, attributes pcommon.Map) {
	key := "Attributes"
	if m.mode == MappingRaw {
		key = ""
	}
	document.AddAttributes(key, attributes)
}

func (m *encodeModel) encodeEvents(document *objmodel.Document, events ptrace.SpanEventSlice) {
	key := "Events"
	if m.mode == MappingRaw {
		key = ""
	}
	document.AddEvents(key, events)
}

func spanLinksToString(spanLinkSlice ptrace.SpanLinkSlice) string {
	linkArray := make([]map[string]any, 0, spanLinkSlice.Len())
	for i := 0; i < spanLinkSlice.Len(); i++ {
		spanLink := spanLinkSlice.At(i)
		link := map[string]any{}
		link[spanIDField] = traceutil.SpanIDToHexOrEmptyString(spanLink.SpanID())
		link[traceIDField] = traceutil.TraceIDToHexOrEmptyString(spanLink.TraceID())
		link[attributeField] = spanLink.Attributes().AsRaw()
		linkArray = append(linkArray, link)
	}
	linkArrayBytes, _ := json.Marshal(&linkArray)
	return string(linkArrayBytes)
}

// durationAsMicroseconds calculate span duration through end - start nanoseconds and converts time.Time to microseconds,
// which is the format the Duration field is stored in the Span.
func durationAsMicroseconds(start, end time.Time) int64 {
	return (end.UnixNano() - start.UnixNano()) / 1000
}

func scopeToAttributes(scope pcommon.InstrumentationScope) pcommon.Map {
	attrs := pcommon.NewMap()
	attrs.PutStr("name", scope.Name())
	attrs.PutStr("version", scope.Version())
	for k, v := range scope.Attributes().AsRaw() {
		attrs.PutStr(k, v.(string))
	}
	return attrs
}

func encodeAttributesECSMode(document *objmodel.Document, attrs pcommon.Map, conversionMap map[string]string, preserveMap map[string]bool) {
	if len(conversionMap) == 0 {
		// No conversions to be done; add all attributes at top level of
		// document.
		document.AddAttributes("", attrs)
		return
	}

	attrs.Range(func(k string, v pcommon.Value) bool {
		// If ECS key is found for current k in conversion map, use it.
		if ecsKey, exists := conversionMap[k]; exists {
			if ecsKey == "" {
				// Skip the conversion for this k.
				return true
			}

			document.AddAttribute(ecsKey, v)
			if preserve := preserveMap[k]; preserve {
				document.AddAttribute(k, v)
			}
			return true
		}

		// Otherwise, add key at top level with attribute name as-is.
		document.AddAttribute(k, v)
		return true
	})
}

func encodeLogAgentNameECSMode(document *objmodel.Document, resource pcommon.Resource) {
	// Parse out telemetry SDK name, language, and distro name from resource
	// attributes, setting defaults as needed.
	telemetrySdkName := "otlp"
	var telemetrySdkLanguage, telemetryDistroName string

	attrs := resource.Attributes()
	if v, exists := attrs.Get(semconv.AttributeTelemetrySDKName); exists {
		telemetrySdkName = v.Str()
	}
	if v, exists := attrs.Get(semconv.AttributeTelemetrySDKLanguage); exists {
		telemetrySdkLanguage = v.Str()
	}
	if v, exists := attrs.Get(semconv.AttributeTelemetryDistroName); exists {
		telemetryDistroName = v.Str()
		if telemetrySdkLanguage == "" {
			telemetrySdkLanguage = "unknown"
		}
	}

	// Construct agent name from telemetry SDK name, language, and distro name.
	agentName := telemetrySdkName
	if telemetryDistroName != "" {
		agentName = fmt.Sprintf("%s/%s/%s", agentName, telemetrySdkLanguage, telemetryDistroName)
	} else if telemetrySdkLanguage != "" {
		agentName = fmt.Sprintf("%s/%s", agentName, telemetrySdkLanguage)
	}

	// Set agent name in document.
	document.AddString("agent.name", agentName)
}

func encodeLogAgentVersionECSMode(document *objmodel.Document, resource pcommon.Resource) {
	attrs := resource.Attributes()

	if telemetryDistroVersion, exists := attrs.Get(semconv.AttributeTelemetryDistroVersion); exists {
		document.AddString("agent.version", telemetryDistroVersion.Str())
		return
	}

	if telemetrySdkVersion, exists := attrs.Get(semconv.AttributeTelemetrySDKVersion); exists {
		document.AddString("agent.version", telemetrySdkVersion.Str())
		return
	}
}

func encodeLogHostOsTypeECSMode(document *objmodel.Document, resource pcommon.Resource) {
	// https://www.elastic.co/guide/en/ecs/current/ecs-os.html#field-os-type:
	//
	// "One of these following values should be used (lowercase): linux, macos, unix, windows.
	// If the OS you’re dealing with is not in the list, the field should not be populated."

	var ecsHostOsType string
	if semConvOsType, exists := resource.Attributes().Get(semconv.AttributeOSType); exists {
		switch semConvOsType.Str() {
		case "windows", "linux":
			ecsHostOsType = semConvOsType.Str()
		case "darwin":
			ecsHostOsType = "macos"
		case "aix", "hpux", "solaris":
			ecsHostOsType = "unix"
		}
	}

	if semConvOsName, exists := resource.Attributes().Get(semconv.AttributeOSName); exists {
		switch semConvOsName.Str() {
		case "Android":
			ecsHostOsType = "android"
		case "iOS":
			ecsHostOsType = "ios"
		}
	}

	if ecsHostOsType == "" {
		return
	}
	document.AddString("host.os.type", ecsHostOsType)
}

func encodeLogTimestampECSMode(document *objmodel.Document, record plog.LogRecord) {
	if record.Timestamp() != 0 {
		document.AddTimestamp("@timestamp", record.Timestamp())
		return
	}

	document.AddTimestamp("@timestamp", record.ObservedTimestamp())
}

// TODO use https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/internal/exp/metrics/identity
func metricECSHash(timestamp pcommon.Timestamp, attributes pcommon.Map) uint32 {
	hasher := fnv.New32a()

	timestampBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestampBuf, uint64(timestamp))
	hasher.Write(timestampBuf)

	mapHashExcludeDataStreamAttr(hasher, attributes)

	return hasher.Sum32()
}

func metricOTelHash(dp dataPoint, scopeAttrs pcommon.Map, unit string) uint32 {
	hasher := fnv.New32a()

	timestampBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestampBuf, uint64(dp.Timestamp()))
	hasher.Write(timestampBuf)

	binary.LittleEndian.PutUint64(timestampBuf, uint64(dp.StartTimestamp()))
	hasher.Write(timestampBuf)

	hasher.Write([]byte(unit))

	mapHashExcludeDataStreamAttr(hasher, scopeAttrs)
	mapHashExcludeDataStreamAttr(hasher, dp.Attributes())

	return hasher.Sum32()
}

// mapHashExcludeDataStreamAttr is mapHash but ignoring DS attributes.
// It is useful for cases where index is already considered during routing and no need to be considered in hashing.
func mapHashExcludeDataStreamAttr(hasher hash.Hash, m pcommon.Map) {
	m.Range(func(k string, v pcommon.Value) bool {
		switch k {
		case dataStreamType, dataStreamDataset, dataStreamNamespace:
			return true
		}
		hasher.Write([]byte(k))
		valueHash(hasher, v)

		return true
	})
}

func mapHash(hasher hash.Hash, m pcommon.Map) {
	m.Range(func(k string, v pcommon.Value) bool {
		hasher.Write([]byte(k))
		valueHash(hasher, v)

		return true
	})
}

func valueHash(h hash.Hash, v pcommon.Value) {
	switch v.Type() {
	case pcommon.ValueTypeEmpty:
		h.Write([]byte{0})
	case pcommon.ValueTypeStr:
		h.Write([]byte(v.Str()))
	case pcommon.ValueTypeBool:
		if v.Bool() {
			h.Write([]byte{1})
		} else {
			h.Write([]byte{0})
		}
	case pcommon.ValueTypeDouble:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, math.Float64bits(v.Double()))
		h.Write(buf)
	case pcommon.ValueTypeInt:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(v.Int()))
		h.Write(buf)
	case pcommon.ValueTypeBytes:
		h.Write(v.Bytes().AsRaw())
	case pcommon.ValueTypeMap:
		mapHash(h, v.Map())
	case pcommon.ValueTypeSlice:
		sliceHash(h, v.Slice())
	}
}

func sliceHash(h hash.Hash, s pcommon.Slice) {
	for i := 0; i < s.Len(); i++ {
		valueHash(h, s.At(i))
	}
}
