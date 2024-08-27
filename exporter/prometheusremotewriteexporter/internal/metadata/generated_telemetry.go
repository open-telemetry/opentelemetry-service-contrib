// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// Deprecated: [v0.108.0] use LeveledMeter instead.
func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                             metric.Meter
	ExporterPrometheusremotewriteFailedTranslations   metric.Int64Counter
	ExporterPrometheusremotewriteTranslatedTimeSeries metric.Int64Counter
	level                                             configtelemetry.Level
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// WithLevel sets the current telemetry level for the component.
func WithLevel(lvl configtelemetry.Level) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.level = lvl
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{level: configtelemetry.LevelBasic}
	for _, op := range options {
		op(&builder)
	}
	var err, errs error
	if builder.level >= configtelemetry.LevelBasic {
		builder.meter = Meter(settings)
	} else {
		builder.meter = noop.Meter{}
	}
	builder.ExporterPrometheusremotewriteFailedTranslations, err = builder.meter.Int64Counter(
		"otelcol_exporter_prometheusremotewrite_failed_translations",
		metric.WithDescription("Number of translation operations that failed to translate metrics from Otel to Prometheus"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterPrometheusremotewriteTranslatedTimeSeries, err = builder.meter.Int64Counter(
		"otelcol_exporter_prometheusremotewrite_translated_time_series",
		metric.WithDescription("Number of Prometheus time series that were translated from OTel metrics"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
