// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
		metadata.WithGrafanacloudHostCountCallback(func() int64 { return 1 }),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)
	tb.GrafanacloudDatapointCount.Add(context.Background(), 1)
	tb.GrafanacloudFlushCount.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_grafanacloud_datapoint_count",
			Description: "Number of datapoints sent to Grafana Cloud",
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
			Name:        "otelcol_grafanacloud_flush_count",
			Description: "Number of metrics flushes",
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
			Name:        "otelcol_grafanacloud_host_count",
			Description: "Number of unique hosts",
			Unit:        "1",
			Data: metricdata.Gauge[int64]{
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}
