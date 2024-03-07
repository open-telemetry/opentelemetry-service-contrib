// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"sort"

	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// PrometheusConverter converts from OTel write format to Prometheus write format.
type PrometheusConverter struct {
	unique    map[uint64]*prompb.TimeSeries
	conflicts map[uint64][]*prompb.TimeSeries
}

func NewPrometheusConverter() *PrometheusConverter {
	return &PrometheusConverter{
		unique:    map[uint64]*prompb.TimeSeries{},
		conflicts: map[uint64][]*prompb.TimeSeries{},
	}
}

// FromMetrics converts pmetric.Metrics to Prometheus remote write format.
func (c *PrometheusConverter) FromMetrics(md pmetric.Metrics, settings Settings) error {
	return FromMetrics(md, settings, c)
}

// TimeSeries returns a slice of the prompb.TimeSeries that were converted from OTel format.
func (c *PrometheusConverter) TimeSeries() []prompb.TimeSeries {
	conflicts := 0
	for _, ts := range c.conflicts {
		conflicts += len(ts)
	}
	allTS := make([]prompb.TimeSeries, 0, len(c.unique)+conflicts)
	for _, ts := range c.unique {
		allTS = append(allTS, *ts)
	}
	for _, cTS := range c.conflicts {
		for _, ts := range cTS {
			allTS = append(allTS, *ts)
		}
	}

	return allTS
}

func isSameMetric(ts *prompb.TimeSeries, lbls []prompb.Label) bool {
	if len(ts.Labels) != len(lbls) {
		return false
	}
	for i, l := range ts.Labels {
		if l.Name != ts.Labels[i].Name || l.Value != ts.Labels[i].Value {
			return false
		}
	}
	return true
}

// AddExemplars adds exemplars for the dataPoint. For each exemplar, if it can find a bucket bound corresponding to its value,
// the exemplar is added to the bucket bound's time series, provided that the time series' has samples.
func (c *PrometheusConverter) AddExemplars(dataPoint pmetric.HistogramDataPoint, bucketBounds []bucketBoundsData) {
	if len(bucketBounds) == 0 {
		return
	}

	exemplars := getPromExemplars(dataPoint)
	if len(exemplars) == 0 {
		return
	}

	sort.Sort(byBucketBoundsData(bucketBounds))
	for _, exemplar := range exemplars {
		for _, bound := range bucketBounds {
			if len(bound.ts.Samples) > 0 && exemplar.Value <= bound.bound {
				bound.ts.Exemplars = append(bound.ts.Exemplars, exemplar)
				break
			}
		}
	}
}

// AddSample finds a TimeSeries in tsMap that corresponds to the label set labels, and add sample to the TimeSeries; it
// creates a new TimeSeries in the map if not found and returns the time series signature.
// tsMap will be unmodified if either labels or sample is nil, but can still be modified if the exemplar is nil.
func (c *PrometheusConverter) AddSample(sample *prompb.Sample, lbls []prompb.Label) *prompb.TimeSeries {
	if sample == nil || len(lbls) == 0 {
		// This shouldn't happen
		return nil
	}

	ts, _ := c.getOrCreateTimeSeries(lbls)
	ts.Samples = append(ts.Samples, *sample)
	return ts
}
