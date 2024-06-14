// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
	embeddedmetric "go.opentelemetry.io/otel/metric/embedded"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/trace"
	embeddedtrace "go.opentelemetry.io/otel/trace/embedded"
	nooptrace "go.opentelemetry.io/otel/trace/noop"

	"go.opentelemetry.io/collector/component"
)

type mockMeter struct {
	noopmetric.Meter
	name string
}
type mockMeterProvider struct {
	embeddedmetric.MeterProvider
}

func (m mockMeterProvider) Meter(name string, opts ...metric.MeterOption) metric.Meter {
	return mockMeter{name: name}
}

type mockTracer struct {
	nooptrace.Tracer
	name string
}

type mockTracerProvider struct {
	embeddedtrace.TracerProvider
}

func (m mockTracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return mockTracer{name: name}
}

func TestProviders(t *testing.T) {
	set := component.TelemetrySettings{
		MeterProvider:  mockMeterProvider{},
		TracerProvider: mockTracerProvider{},
	}

	meter := Meter(set)
	if m, ok := meter.(mockMeter); ok {
		require.Equal(t, "otelcol/tailsampling", m.name)
	} else {
		require.Fail(t, "returned Meter not mockMeter")
	}

	tracer := Tracer(set)
	if m, ok := tracer.(mockTracer); ok {
		require.Equal(t, "otelcol/tailsampling", m.name)
	} else {
		require.Fail(t, "returned Meter not mockTracer")
	}
}

func TestNewTelemetryBuilder(t *testing.T) {
	set := component.TelemetrySettings{
		MeterProvider:  mockMeterProvider{},
		TracerProvider: mockTracerProvider{},
	}
	applied := false
	_, err := NewTelemetryBuilder(set, func(b *TelemetryBuilder) {
		applied = true
	})
	require.NoError(t, err)
	require.True(t, applied)
}
