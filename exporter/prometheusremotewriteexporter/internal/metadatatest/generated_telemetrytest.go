// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

// Deprecated: [v0.119.0] Use componenttest.Telemetry
type Telemetry struct {
	*componenttest.Telemetry
}

// Deprecated: [v0.119.0] Use componenttest.NewTelemetry
func SetupTelemetry(opts ...componenttest.TelemetryOption) Telemetry {
	return Telemetry{Telemetry: componenttest.NewTelemetry(opts...)}
}

// Deprecated: [v0.119.0] Use metadatatest.NewSettings
func (tt *Telemetry) NewSettings() exporter.Settings {
	return NewSettings(tt.Telemetry)
}

func NewSettings(tt *componenttest.Telemetry) exporter.Settings {
	set := exportertest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("prometheusremotewrite"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

// Deprecated: [v0.119.0] Use metadatatest.AssertEqual*
func (tt *Telemetry) AssertMetrics(t *testing.T, expected []metricdata.Metrics, opts ...metricdatatest.Option) {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.Reader.Collect(context.Background(), &md))
	// ensure all required metrics are present
	for _, want := range expected {
		got := getMetricFromResource(want.Name, md)
		metricdatatest.AssertEqual(t, want, got, opts...)
	}

	// ensure no additional metrics are emitted
	require.Equal(t, len(expected), lenMetrics(md))
}

func AssertEqualExporterPrometheusremotewriteConsumers(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_exporter_prometheusremotewrite_consumers",
		Description: "Number of configured workers to use to fan out the outgoing requests",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_exporter_prometheusremotewrite_consumers")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualExporterPrometheusremotewriteFailedTranslations(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_exporter_prometheusremotewrite_failed_translations",
		Description: "Number of translation operations that failed to translate metrics from Otel to Prometheus",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_exporter_prometheusremotewrite_failed_translations")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualExporterPrometheusremotewriteSentBatchCount(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_exporter_prometheusremotewrite_sent_batch_count",
		Description: "Number of remote write request batches sent to the remote write endpoint",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_exporter_prometheusremotewrite_sent_batch_count")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualExporterPrometheusremotewriteTranslatedTimeSeries(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_exporter_prometheusremotewrite_translated_time_series",
		Description: "Number of Prometheus time series that were translated from OTel metrics",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_exporter_prometheusremotewrite_translated_time_series")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func getMetricFromResource(name string, got metricdata.ResourceMetrics) metricdata.Metrics {
	for _, sm := range got.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m
			}
		}
	}

	return metricdata.Metrics{}
}

func lenMetrics(got metricdata.ResourceMetrics) int {
	metricsCount := 0
	for _, sm := range got.ScopeMetrics {
		metricsCount += len(sm.Metrics)
	}

	return metricsCount
}
