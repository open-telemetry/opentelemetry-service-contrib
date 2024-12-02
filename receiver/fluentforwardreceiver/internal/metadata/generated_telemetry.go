// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

// Deprecated: [v0.114.0] use Meter instead.
func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                   metric.Meter
	FluentClosedConnections metric.Int64UpDownCounter
	FluentEventsParsed      metric.Int64UpDownCounter
	FluentOpenedConnections metric.Int64UpDownCounter
	FluentParseFailures     metric.Int64UpDownCounter
	FluentRecordsGenerated  metric.Int64UpDownCounter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.FluentClosedConnections, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64UpDownCounter(
		"otelcol_fluent_closed_connections",
		metric.WithDescription("Number of connections closed to the fluentforward receiver"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentEventsParsed, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64UpDownCounter(
		"otelcol_fluent_events_parsed",
		metric.WithDescription("Number of Fluent events parsed successfully"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentOpenedConnections, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64UpDownCounter(
		"otelcol_fluent_opened_connections",
		metric.WithDescription("Number of connections opened to the fluentforward receiver"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentParseFailures, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64UpDownCounter(
		"otelcol_fluent_parse_failures",
		metric.WithDescription("Number of times Fluent messages failed to be decoded"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.FluentRecordsGenerated, err = getLeveledMeter(builder.meter, configtelemetry.LevelBasic, settings.MetricsLevel).Int64UpDownCounter(
		"otelcol_fluent_records_generated",
		metric.WithDescription("Number of log records generated from Fluent forward input"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}

func getLeveledMeter(meter metric.Meter, cfgLevel, srvLevel configtelemetry.Level) metric.Meter {
	if cfgLevel <= srvLevel {
		return meter
	}
	return noop.Meter{}
}
