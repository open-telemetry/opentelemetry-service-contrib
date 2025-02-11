// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

func NewSettings(tt *componenttest.Telemetry) receiver.Settings {
	set := receivertest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("kafka"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

func AssertEqualKafkaReceiverCurrentOffset(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_current_offset",
		Description: "Current message offset",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_current_offset")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverMessages(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_messages",
		Description: "Number of received messages",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_messages")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverOffsetLag(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_offset_lag",
		Description: "Current offset lag",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_offset_lag")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverPartitionClose(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_partition_close",
		Description: "Number of finished partitions",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_partition_close")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverPartitionStart(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_partition_start",
		Description: "Number of started partitions",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_partition_start")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverUnmarshalFailedLogRecords(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_unmarshal_failed_log_records",
		Description: "Number of log records failed to be unmarshaled",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_unmarshal_failed_log_records")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverUnmarshalFailedMetricPoints(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_unmarshal_failed_metric_points",
		Description: "Number of metric points failed to be unmarshaled",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_unmarshal_failed_metric_points")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualKafkaReceiverUnmarshalFailedSpans(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_kafka_receiver_unmarshal_failed_spans",
		Description: "Number of spans failed to be unmarshaled",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_kafka_receiver_unmarshal_failed_spans")
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
