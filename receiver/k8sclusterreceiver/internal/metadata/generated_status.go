// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var (
	Type = component.MustNewType("k8s_cluster")
)

const (
	MetricsStability = component.StabilityLevelBeta
	LogsStability    = component.StabilityLevelDevelopment
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/k8sclusterreceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/k8sclusterreceiver")
}
