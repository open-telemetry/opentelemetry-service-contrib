// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)
	tb.ExporterPrometheusremotewriteFailedTranslations.Add(context.Background(), 1)
	tb.ExporterPrometheusremotewriteTranslatedTimeSeries.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_exporter_prometheusremotewrite_failed_translations",
			Description: "Number of translation operations that failed to translate metrics from Otel to Prometheus",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_exporter_prometheusremotewrite_translated_time_series",
			Description: "Number of Prometheus time series that were translated from OTel metrics",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}
