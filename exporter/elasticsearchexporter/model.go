// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"bytes"
	"encoding/json"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv "go.opentelemetry.io/collector/semconv/v1.22.0"

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
	"k8s.namespace.name":                    "kubernetes.namespace",
	"k8s.node.name":                         "kubernetes.node.name",
	"k8s.pod.name":                          "kubernetes.pod.name",
	"k8s.pod.uid":                           "kubernetes.pod.uid",
}

// resourceAttrsConversionMap contains conversions for log record attributes
// from their Semantic Conventions (SemConv) names to equivalent Elastic Common
// Schema (ECS) names.
var recordAttrsConversionMap = map[string]string{
	"event.name":                         "event.action",
	semconv.AttributeExceptionMessage:    "error.message",
	semconv.AttributeExceptionStacktrace: "error.stacktrace",
	semconv.AttributeExceptionType:       "error.type",
	semconv.AttributeExceptionEscaped:    "event.error.exception.handled",
}

type mappingModel interface {
	encodeLog(pcommon.Resource, plog.LogRecord, pcommon.InstrumentationScope) ([]byte, error)
	encodeSpan(pcommon.Resource, ptrace.Span, pcommon.InstrumentationScope) ([]byte, error)
}

// encodeModel tries to keep the event as close to the original open telemetry semantics as is.
// No fields will be mapped by default.
//
// Field deduplication and dedotting of attributes is supported by the encodeModel.
//
// See: https://github.com/open-telemetry/oteps/blob/master/text/logs/0097-log-data-model.md
type encodeModel struct {
	dedup bool
	dedot bool
	mode  MappingMode
}

const (
	traceIDField   = "traceID"
	spanIDField    = "spanID"
	attributeField = "attribute"
)

func (m *encodeModel) encodeLog(resource pcommon.Resource, record plog.LogRecord, scope pcommon.InstrumentationScope) ([]byte, error) {
	var document objmodel.Document
	switch m.mode {
	case MappingECS:
		document = m.encodeLogECSMode(resource, record, scope)
	default:
		document = m.encodeLogDefaultMode(resource, record, scope)
	}

	var buf bytes.Buffer
	if m.dedup {
		document.Dedup()
	} else if m.dedot {
		document.Sort()
	}
	err := document.Serialize(&buf, m.dedot)
	return buf.Bytes(), err
}

func (m *encodeModel) encodeLogDefaultMode(resource pcommon.Resource, record plog.LogRecord, scope pcommon.InstrumentationScope) objmodel.Document {
	docTs := record.Timestamp()
	if docTs.AsTime().UnixNano() == 0 {
		docTs = record.ObservedTimestamp()
	}

	// TODO (lahsivjar): Should be a function?
	recordAttrKey := "Attributes"
	if m.mode == MappingRaw {
		recordAttrKey = ""
	}

	var document objmodel.Document
	document.AddMultiple(
		// We use @timestamp in order to ensure that we can index if the
		// default data stream logs template is used.
		objmodel.NewKV("@timestamp", objmodel.TimestampValue(docTs)),
		objmodel.NewKV("TraceId", objmodel.TraceIDValue(record.TraceID())),
		objmodel.NewKV("SpanId", objmodel.SpanIDValue(record.SpanID())),
		objmodel.NewKV("TraceFlags", objmodel.IntValue(int64(record.Flags()))),
		objmodel.NewKV("SeverityText", objmodel.NonEmptyStringValue(record.SeverityText())),
		objmodel.NewKV("SeverityNumber", objmodel.IntValue(int64(record.SeverityNumber()))),
		objmodel.NewKV("Body", objmodel.RawValue(record.Body())),
		// Add scope name and version as additional scope attributes.
		// Empty values are also allowed to be added.
		objmodel.NewKV("Scope.name", objmodel.StringValue(scope.Name())),
		objmodel.NewKV("Scope.version", objmodel.StringValue(scope.Version())),
		objmodel.NewKV("Resource", objmodel.RawMapValue(resource.Attributes())),
		objmodel.NewKV("Scope", objmodel.RawMapValue(scope.Attributes())),
		objmodel.NewKV(recordAttrKey, objmodel.RawMapValue(record.Attributes())),
	)

	return document

}

func (m *encodeModel) encodeLogECSMode(resource pcommon.Resource, record plog.LogRecord, scope pcommon.InstrumentationScope) objmodel.Document {
	var document objmodel.Document

	// Appoximate capacity to reduce allocations
	resourceAttrs := resource.Attributes()
	scopeAttrs := scope.Attributes()
	recordAttrs := record.Attributes()
	document.EnsureCapacity(9 + // constant number of fields added
		resourceAttrs.Len() +
		scopeAttrs.Len() +
		recordAttrs.Len(),
	)

	// TODO (lahsivjar): move to objmodel with a new function, something like ...WithKeyMutator(...)
	// First, try to map resource-level attributes to ECS fields.
	encodeLogAttributesECSMode(&document, resourceAttrs, resourceAttrsConversionMap)

	// Then, try to map scope-level attributes to ECS fields.
	scopeAttrsConversionMap := map[string]string{
		// None at the moment
	}
	encodeLogAttributesECSMode(&document, scopeAttrs, scopeAttrsConversionMap)

	// Finally, try to map record-level attributes to ECS fields.
	encodeLogAttributesECSMode(&document, recordAttrs, recordAttrsConversionMap)

	severity := objmodel.NilValue
	if n := record.SeverityNumber(); n != plog.SeverityNumberUnspecified {
		severity = objmodel.IntValue(int64(record.SeverityNumber()))
	}

	// Handle special cases.
	document.AddMultiple(
		objmodel.NewKV("@timestamp", getTimestampForECS(record)),
		objmodel.NewKV("agent.name", getAgentNameForECS(resource)),
		objmodel.NewKV("agent.version", getAgentVersionForECS(resource)),
		objmodel.NewKV("host.os.type", getHostOsTypeForECS(resource)),
		objmodel.NewKV("trace.id", objmodel.TraceIDValue(record.TraceID())),
		objmodel.NewKV("span.id", objmodel.SpanIDValue(record.SpanID())),
		objmodel.NewKV("log.level", objmodel.NonEmptyStringValue(record.SeverityText())),
		objmodel.NewKV("event.severity", severity),
		objmodel.NewKV("message", objmodel.RawValue(record.Body())),
	)

	return document
}

func (m *encodeModel) encodeSpan(resource pcommon.Resource, span ptrace.Span, scope pcommon.InstrumentationScope) ([]byte, error) {
	// TODO (lahsivjar): Should be a function?
	spanAttrKey := "Attributes"
	eventsKey := "Events"
	if m.mode == MappingRaw {
		spanAttrKey = ""
		eventsKey = ""
	}

	var document objmodel.Document
	document.AddMultiple(
		objmodel.NewKV("@timestamp", objmodel.TimestampValue(span.StartTimestamp())),
		objmodel.NewKV("EndTimestamp", objmodel.TimestampValue(span.EndTimestamp())),
		objmodel.NewKV("TraceId", objmodel.TraceIDValue(span.TraceID())),
		objmodel.NewKV("SpanId", objmodel.SpanIDValue(span.SpanID())),
		objmodel.NewKV("ParentSpanId", objmodel.SpanIDValue(span.ParentSpanID())),
		objmodel.NewKV("Name", objmodel.NonEmptyStringValue(span.Name())),
		objmodel.NewKV("Kind", objmodel.NonEmptyStringValue(traceutil.SpanKindStr(span.Kind()))),
		objmodel.NewKV("TraceStatus", objmodel.IntValue(int64(span.Status().Code()))),
		objmodel.NewKV("TraceStatusDescription", objmodel.NonEmptyStringValue(span.Status().Message())),
		objmodel.NewKV("Link", objmodel.NonEmptyStringValue(spanLinksToString(span.Links()))),
		objmodel.NewKV("Duration", objmodel.IntValue(durationAsMicroseconds(span.StartTimestamp(), span.EndTimestamp()))),
		// Add scope name and version as additional scope attributes
		// Empty values are also allowed to be added.
		objmodel.NewKV("Scope.name", objmodel.StringValue(scope.Name())),
		objmodel.NewKV("Scope.version", objmodel.StringValue(scope.Version())),
		objmodel.NewKV("Resource", objmodel.RawMapValue(resource.Attributes())),
		objmodel.NewKV("Scope", objmodel.RawMapValue(scope.Attributes())),
		objmodel.NewKV(spanAttrKey, objmodel.RawMapValue(span.Attributes())),
		objmodel.NewKV(eventsKey, objmodel.RawSpans(span.Events())),
	)

	if m.dedup {
		document.Dedup()
	} else if m.dedot {
		document.Sort()
	}

	var buf bytes.Buffer
	err := document.Serialize(&buf, m.dedot)
	return buf.Bytes(), err
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

// durationAsMicroseconds calculate span duration through end - start
// nanoseconds and converts time.Time to microseconds, which is the format
// the Duration field is stored in the Span.
func durationAsMicroseconds(start, end pcommon.Timestamp) int64 {
	return (end.AsTime().UnixNano() - start.AsTime().UnixNano()) / 1000
}

func encodeLogAttributesECSMode(document *objmodel.Document, attrs pcommon.Map, conversionMap map[string]string) {
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
			return true
		}

		// Otherwise, add key at top level with attribute name as-is.
		document.AddAttribute(k, v)
		return true
	})
}

func getAgentNameForECS(resource pcommon.Resource) objmodel.Value {
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

	return objmodel.NonEmptyStringValue(agentName)
}

func getAgentVersionForECS(resource pcommon.Resource) objmodel.Value {
	attrs := resource.Attributes()

	if telemetryDistroVersion, exists := attrs.Get(semconv.AttributeTelemetryDistroVersion); exists {
		return objmodel.NonEmptyStringValue(telemetryDistroVersion.Str())
	}

	if telemetrySdkVersion, exists := attrs.Get(semconv.AttributeTelemetrySDKVersion); exists {
		return objmodel.NonEmptyStringValue(telemetrySdkVersion.Str())
	}
	return objmodel.NilValue
}

func getHostOsTypeForECS(resource pcommon.Resource) objmodel.Value {
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
		return objmodel.NilValue
	}
	return objmodel.NonEmptyStringValue(ecsHostOsType)
}

func getTimestampForECS(record plog.LogRecord) objmodel.Value {
	if record.Timestamp() != 0 {
		return objmodel.TimestampValue(record.Timestamp())
	}

	return objmodel.TimestampValue(record.ObservedTimestamp())
}
