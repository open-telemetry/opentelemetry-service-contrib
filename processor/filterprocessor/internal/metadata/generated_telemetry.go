// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                             metric.Meter
	mu                                sync.Mutex
	registrations                     []metric.Registration
	ProcessorFilterDatapointsFiltered metric.Int64Counter
	ProcessorFilterLogsFiltered       metric.Int64Counter
	ProcessorFilterSpansFiltered      metric.Int64Counter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// Shutdown unregister all registered callbacks for async instruments.
func (builder *TelemetryBuilder) Shutdown() {
	builder.mu.Lock()
	defer builder.mu.Unlock()
	for _, reg := range builder.registrations {
		reg.Unregister()
	}
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
	builder.ProcessorFilterDatapointsFiltered, err = builder.meter.Int64Counter(
		"otelcol_processor_filter_datapoints.filtered",
		metric.WithDescription("Number of metric data points dropped by the filter processor"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorFilterLogsFiltered, err = builder.meter.Int64Counter(
		"otelcol_processor_filter_logs.filtered",
		metric.WithDescription("Number of logs dropped by the filter processor"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorFilterSpansFiltered, err = builder.meter.Int64Counter(
		"otelcol_processor_filter_spans.filtered",
		metric.WithDescription("Number of spans dropped by the filter processor"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
