// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processortest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

type Telemetry struct {
	componenttest.Telemetry
}

func SetupTelemetry(opts ...componenttest.TelemetryOption) Telemetry {
	return Telemetry{Telemetry: componenttest.NewTelemetry(opts...)}
}

func (tt *Telemetry) NewSettings() processor.Settings {
	set := processortest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("groupbytrace"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

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

func AssertEqualProcessorGroupbytraceConfNumTraces(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_conf_num_traces",
		Description: "Maximum number of traces to hold in the internal storage",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_conf_num_traces")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceEventLatency(t *testing.T, tt componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_event_latency",
		Description: "How long the queue events are taking to be processed",
		Unit:        "ms",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_event_latency")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceIncompleteReleases(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_incomplete_releases",
		Description: "Releases that are suspected to have been incomplete",
		Unit:        "{releases}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_incomplete_releases")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceNumEventsInQueue(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_num_events_in_queue",
		Description: "Number of events currently in the queue",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_num_events_in_queue")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceNumTracesInMemory(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_num_traces_in_memory",
		Description: "Number of traces currently in the in-memory storage",
		Unit:        "1",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_num_traces_in_memory")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceSpansReleased(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_spans_released",
		Description: "Spans released to the next consumer",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_spans_released")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceTracesEvicted(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_traces_evicted",
		Description: "Traces evicted from the internal buffer",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_traces_evicted")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorGroupbytraceTracesReleased(t *testing.T, tt componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_groupbytrace_traces_released",
		Description: "Traces released to the next consumer",
		Unit:        "1",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got := getMetric(t, tt, "otelcol_processor_groupbytrace_traces_released")
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func getMetric(t *testing.T, tt componenttest.Telemetry, name string) metricdata.Metrics {
	var md metricdata.ResourceMetrics
	require.NoError(t, tt.Reader.Collect(context.Background(), &md))
	return getMetricFromResource(name, md)
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
