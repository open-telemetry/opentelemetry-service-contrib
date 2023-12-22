// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	Type             = "alibabacloud_logservice"
	TracesStability  = component.StabilityLevelUnmaintained
	MetricsStability = component.StabilityLevelUnmaintained
	LogsStability    = component.StabilityLevelUnmaintained
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/alibabacloudlogservice")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/alibabacloudlogservice")
}
