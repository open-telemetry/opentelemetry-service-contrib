// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"encoding/json"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	metricsHistogramTableSuffix = "_histogram"
	metricsHistogramDDL         = `
CREATE TABLE IF NOT EXISTS %s` + metricsHistogramTableSuffix + `
(
    service_name            VARCHAR(200),
    timestamp               DATETIME(6),
    metric_name             VARCHAR(200),
    metric_description      STRING,
    metric_unit             STRING,
    attributes              VARIANT,
    start_time              DATETIME(6),
    count                   BIGINT,
    sum                     DOUBLE,
    bucket_counts           ARRAY<BIGINT>,
    explicit_bounds         ARRAY<DOUBLE>,
    exemplars               ARRAY<STRUCT<filtered_attributes:MAP<STRING,STRING>, timestamp:DATETIME(6), value:DOUBLE, span_id:STRING, trace_id:STRING>>,
    min                     DOUBLE,
    max                     DOUBLE,
    aggregation_temporality STRING,
    resource_attributes     VARIANT,
    scope_name              STRING,
    scope_version           STRING,

    INDEX idx_service_name(service_name) USING INVERTED,
    INDEX idx_timestamp(timestamp) USING INVERTED,
    INDEX idx_metric_name(metric_name) USING INVERTED,
    INDEX idx_metric_description(metric_description) USING INVERTED,
    INDEX idx_metric_unit(metric_unit) USING INVERTED,
    INDEX idx_attributes(attributes) USING INVERTED,
    INDEX idx_start_time(start_time) USING INVERTED,
    INDEX idx_count(count) USING INVERTED,
    INDEX idx_aggregation_temporality(aggregation_temporality) USING INVERTED,
    INDEX idx_resource_attributes(resource_attributes) USING INVERTED,
    INDEX idx_scope_name(scope_name) USING INVERTED,
    INDEX idx_scope_version(scope_version) USING INVERTED
)
ENGINE = OLAP
DUPLICATE KEY(service_name, timestamp)
PARTITION BY RANGE(timestamp) ()
DISTRIBUTED BY HASH(metric_name) BUCKETS AUTO
%s;
`
)

// dMetricHistogram Histogram Metric to Doris
type dMetricHistogram struct {
	*dMetric               `json:",inline"`
	Timestamp              string         `json:"timestamp"`
	Attributes             map[string]any `json:"attributes"`
	StartTime              string         `json:"start_time"`
	Count                  int64          `json:"count"`
	Sum                    float64        `json:"sum"`
	BucketCounts           []int64        `json:"bucket_counts"`
	ExplicitBounds         []float64      `json:"explicit_bounds"`
	Exemplars              []*dExemplar   `json:"exemplars"`
	Min                    float64        `json:"min"`
	Max                    float64        `json:"max"`
	AggregationTemporality string         `json:"aggregation_temporality"`
}

type metricModelHistogram struct {
	data []*dMetricHistogram
}

func (m *metricModelHistogram) metricType() pmetric.MetricType {
	return pmetric.MetricTypeHistogram
}

func (m *metricModelHistogram) tableSuffix() string {
	return metricsHistogramTableSuffix
}

func (m *metricModelHistogram) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	if pm.Type() != pmetric.MetricTypeHistogram {
		return fmt.Errorf("metric type is not Histogram: %v", pm.Type().String())
	}

	dataPoints := pm.Histogram().DataPoints()
	for i := 0; i < dataPoints.Len(); i++ {
		dp := dataPoints.At(i)

		exemplars := dp.Exemplars()
		newExeplars := make([]*dExemplar, 0, exemplars.Len())
		for j := 0; j < exemplars.Len(); j++ {
			exemplar := exemplars.At(j)

			newExeplar := &dExemplar{
				FilteredAttributes: exemplar.FilteredAttributes().AsRaw(),
				Timestamp:          e.formatTime(exemplar.Timestamp().AsTime()),
				Value:              e.getExemplarValue(exemplar),
				SpanID:             exemplar.SpanID().String(),
				TraceID:            exemplar.TraceID().String(),
			}

			newExeplars = append(newExeplars, newExeplar)
		}

		bucketCounts := dp.BucketCounts()
		newBucketCounts := make([]int64, 0, bucketCounts.Len())
		for j := 0; j < bucketCounts.Len(); j++ {
			newBucketCounts = append(newBucketCounts, int64(bucketCounts.At(j)))
		}

		explicitBounds := dp.ExplicitBounds()
		newExplicitBounds := make([]float64, 0, explicitBounds.Len())
		for j := 0; j < explicitBounds.Len(); j++ {
			newExplicitBounds = append(newExplicitBounds, explicitBounds.At(j))
		}

		metric := &dMetricHistogram{
			dMetric:                dm,
			Timestamp:              e.formatTime(dp.Timestamp().AsTime()),
			Attributes:             dp.Attributes().AsRaw(),
			StartTime:              e.formatTime(dp.StartTimestamp().AsTime()),
			Count:                  int64(dp.Count()),
			Sum:                    dp.Sum(),
			BucketCounts:           newBucketCounts,
			ExplicitBounds:         newExplicitBounds,
			Exemplars:              newExeplars,
			Min:                    dp.Min(),
			Max:                    dp.Max(),
			AggregationTemporality: pm.Histogram().AggregationTemporality().String(),
		}
		m.data = append(m.data, metric)
	}

	return nil
}

func (m *metricModelHistogram) raw() any {
	return m.data
}

func (m *metricModelHistogram) size() int {
	return len(m.data)
}

func (m *metricModelHistogram) bytes() ([]byte, error) {
	return json.Marshal(m.data)
}
