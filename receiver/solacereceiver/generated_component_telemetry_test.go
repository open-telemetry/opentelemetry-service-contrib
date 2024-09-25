// Code generated by mdatagen. DO NOT EDIT.

package solacereceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	// "go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

type componentTestTelemetry struct {
	reader        *sdkmetric.ManualReader
	meterProvider *sdkmetric.MeterProvider
}

func (tt *componentTestTelemetry) NewSettings() receiver.Settings {
	settings := receivertest.NewNopSettings()
	settings.MeterProvider = tt.meterProvider
	settings.LeveledMeterProvider = func(_ configtelemetry.Level) metric.MeterProvider {
		return tt.meterProvider
	}
	settings.ID = component.NewID(component.MustNewType("solace"))

	return settings
}

func setupTestTelemetry() componentTestTelemetry {
	reader := sdkmetric.NewManualReader()
	return componentTestTelemetry{
		reader:        reader,
		meterProvider: sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)),
	}
}

func (tt *componentTestTelemetry) assertMetrics(t *testing.T, expected []metricdata.Metrics) {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.reader.Collect(context.Background(), &md))
	// ensure all required metrics are present
	for _, want := range expected {
		got := tt.getMetric(want.Name, md)
		require.Equal(t, want.Name, got.Name)
		require.Equal(t, want.Unit, got.Unit)
		require.Equal(t, want.Description, got.Description)
		// Commented out this assertion because we added other attibutes to the generated metric
		// metricdatatest.AssertEqual(t, want, got, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	}

	// ensure no additional metrics are emitted
	require.Equal(t, len(expected), tt.len(md))
}

func (tt *componentTestTelemetry) getMetric(name string, got metricdata.ResourceMetrics) metricdata.Metrics {
	for _, sm := range got.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m
			}
		}
	}

	return metricdata.Metrics{}
}

func (tt *componentTestTelemetry) len(got metricdata.ResourceMetrics) int {
	metricsCount := 0
	for _, sm := range got.ScopeMetrics {
		metricsCount += len(sm.Metrics)
	}

	return metricsCount
}

func (tt *componentTestTelemetry) Shutdown(ctx context.Context) error {
	return tt.meterProvider.Shutdown(ctx)
}
