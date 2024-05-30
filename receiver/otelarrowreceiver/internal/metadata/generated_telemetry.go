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

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/otelarrowreceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/otelarrowreceiver")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	OtelArrowReceiverInFlightBytes    metric.Int64UpDownCounter
	OtelArrowReceiverInFlightItems    metric.Int64UpDownCounter
	OtelArrowReceiverInFlightRequests metric.Int64UpDownCounter
	level                             configtelemetry.Level
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
	var (
		err, errs error
		meter     metric.Meter
	)
	if builder.level >= configtelemetry.LevelBasic {
		meter = Meter(settings)
	} else {
		meter = noop.Meter{}
	}
	builder.OtelArrowReceiverInFlightBytes, err = meter.Int64UpDownCounter(
		"otel_arrow_receiver_in_flight_bytes",
		metric.WithDescription("Number of bytes in flight"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	builder.OtelArrowReceiverInFlightItems, err = meter.Int64UpDownCounter(
		"otel_arrow_receiver_in_flight_items",
		metric.WithDescription("Number of items in flight"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.OtelArrowReceiverInFlightRequests, err = meter.Int64UpDownCounter(
		"otel_arrow_receiver_in_flight_requests",
		metric.WithDescription("Number of requests in flight"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
