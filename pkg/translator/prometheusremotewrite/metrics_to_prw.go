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

// AddExemplars finds a bucket bound that corresponds to the exemplars value and add the exemplar to the specific sig;
// we only add exemplars if samples are presents
// tsMap is unmodified if either of its parameters is nil and samples are nil.
func (c *PrometheusConverter) AddExemplars(dataPoint pmetric.HistogramDataPoint, boundsData []bucketBoundsData) {
	if len(boundsData) == 0 {
		return
	}

	exemplars := getPromExemplars(dataPoint)
	if len(exemplars) == 0 {
		return
	}
	sort.Sort(byBucketBoundsData(boundsData))

	for _, exemplar := range exemplars {
		for _, bound := range boundsData {
			if len(bound.ts.Samples) > 0 && exemplar.Value <= bound.bound {
				bound.ts.Exemplars = append(bound.ts.Exemplars, exemplar)
				return
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

	h := timeSeriesSignature(lbls)
	ts := c.unique[h]
	if ts != nil {
		if !isSameMetric(ts, lbls) {
			// Collision, try to find it
			ts = nil
			for _, cTS := range c.conflicts[h] {
				if isSameMetric(cTS, lbls) {
					ts = cTS
					break
				}
			}
			if ts == nil {
				// New collision
				ts = &prompb.TimeSeries{
					Labels: lbls,
				}
				c.conflicts[h] = append(c.conflicts[h], ts)
			}
		}

		ts.Samples = append(ts.Samples, *sample)
		return ts
	}

	// The metric is new
	ts = &prompb.TimeSeries{
		Labels:  lbls,
		Samples: []prompb.Sample{*sample},
	}
	c.unique[h] = ts
	return ts
}
