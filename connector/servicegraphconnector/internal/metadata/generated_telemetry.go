// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                             metric.Meter
	mu                                sync.Mutex
	registrations                     []metric.Registration
	ConnectorServicegraphDroppedSpans metric.Int64Counter
	ConnectorServicegraphExpiredEdges metric.Int64Counter
	ConnectorServicegraphTotalEdges   metric.Int64Counter
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
	builder.ConnectorServicegraphDroppedSpans, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_dropped_spans",
		metric.WithDescription("Number of spans dropped when trying to add edges"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ConnectorServicegraphExpiredEdges, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_expired_edges",
		metric.WithDescription("Number of edges that expired before finding its matching span"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ConnectorServicegraphTotalEdges, err = builder.meter.Int64Counter(
		"otelcol_connector_servicegraph_total_edges",
		metric.WithDescription("Total number of unique edges"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
