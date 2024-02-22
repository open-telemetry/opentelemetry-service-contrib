// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/cespare/xxhash/v2"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/model/value"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"

	prometheustranslator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"
)

const (
	sumStr        = "_sum"
	countStr      = "_count"
	bucketStr     = "_bucket"
	leStr         = "le"
	quantileStr   = "quantile"
	pInfStr       = "+Inf"
	createdSuffix = "_created"
	// maxExemplarRunes is the maximum number of UTF-8 exemplar characters
	// according to the prometheus specification
	// https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#exemplars
	maxExemplarRunes = 128
	// Trace and Span id keys are defined as part of the spec:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification%2Fmetrics%2Fdatamodel.md#exemplars-2
	traceIDKey       = "trace_id"
	spanIDKey        = "span_id"
	infoType         = "info"
	targetMetricName = "target_info"
)

type bucketBoundsData struct {
	ts    *prompb.TimeSeries
	bound float64
}

// byBucketBoundsData enables the usage of sort.Sort() with a slice of bucket bounds
type byBucketBoundsData []bucketBoundsData

func (m byBucketBoundsData) Len() int           { return len(m) }
func (m byBucketBoundsData) Less(i, j int) bool { return m[i].bound < m[j].bound }
func (m byBucketBoundsData) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// ByLabelName enables the usage of sort.Sort() with a slice of labels
type ByLabelName []prompb.Label

func (a ByLabelName) Len() int           { return len(a) }
func (a ByLabelName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByLabelName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// timeSeriesSignature returns a hashed label set signature.
// The label slice should not contain duplicate label names; this method sorts the slice by label name before creating
// the signature.
// The algorithm is the same as in Prometheus' Labels.Hash method.
func timeSeriesSignature(labels []prompb.Label) uint64 {
	sort.Sort(ByLabelName(labels))

	// Use xxhash.Sum64(b) for fast path as it's faster.
	b := make([]byte, 0, 1024)

	for i, v := range labels {
		if len(b)+len(v.Name)+len(v.Value)+2 >= cap(b) {
			// If labels entry is 1KB+ do not allocate whole entry.
			h := xxhash.New()
			_, _ = h.Write(b)
			for _, v := range labels[i:] {
				_, _ = h.WriteString(v.Name)
				_, _ = h.Write(seps)
				_, _ = h.WriteString(v.Value)
				_, _ = h.Write(seps)
			}
			return h.Sum64()
		}

		b = append(b, v.Name...)
		b = append(b, seps[0])
		b = append(b, v.Value...)
		b = append(b, seps[0])
	}
	return xxhash.Sum64(b)
}

var seps = []byte{'\xff'}

// createAttributes creates a slice of Prometheus Labels with OTLP attributes and pairs of string values.
// Unpaired string values are ignored. String pairs overwrite OTLP labels if collisions happen, and overwrites are
// logged. Resulting label names are sanitized.
func createAttributes(resource pcommon.Resource, attributes pcommon.Map, externalLabels map[string]string, extras ...string) []prompb.Label {
	resourceAttrs := resource.Attributes()
	serviceName, haveServiceName := resourceAttrs.Get(conventions.AttributeServiceName)
	instance, haveInstanceID := resourceAttrs.Get(conventions.AttributeServiceInstanceID)

	// Calculate the maximum possible number of labels we could return so we can preallocate l
	maxLabelCount := attributes.Len() + len(externalLabels) + len(extras)/2
	if haveServiceName {
		maxLabelCount++
	}
	if haveInstanceID {
		maxLabelCount++
	}

	// map ensures no duplicate label name
	l := make(map[string]string, maxLabelCount)

	// Ensure attributes are sorted by key for consistent merging of keys which
	// collide when sanitized.
	labels := make([]prompb.Label, 0, maxLabelCount)
	attributes.Range(func(key string, value pcommon.Value) bool {
		labels = append(labels, prompb.Label{Name: key, Value: value.AsString()})
		return true
	})
	sort.Stable(ByLabelName(labels))

	for _, label := range labels {
		var finalKey = prometheustranslator.NormalizeLabel(label.Name)
		if existingValue, alreadyExists := l[finalKey]; alreadyExists {
			l[finalKey] = existingValue + ";" + label.Value
		} else {
			l[finalKey] = label.Value
		}
	}

	// Map service.name + service.namespace to job
	if haveServiceName {
		val := serviceName.AsString()
		if serviceNamespace, ok := resourceAttrs.Get(conventions.AttributeServiceNamespace); ok {
			val = fmt.Sprintf("%s/%s", serviceNamespace.AsString(), val)
		}
		l[model.JobLabel] = val
	}
	// Map service.instance.id to instance
	if haveInstanceID {
		l[model.InstanceLabel] = instance.AsString()
	}
	for key, value := range externalLabels {
		// External labels have already been sanitized
		if _, alreadyExists := l[key]; alreadyExists {
			// Skip external labels if they are overridden by metric attributes
			continue
		}
		l[key] = value
	}

	for i := 0; i < len(extras); i += 2 {
		if i+1 >= len(extras) {
			break
		}
		_, found := l[extras[i]]
		if found {
			log.Println("label " + extras[i] + " is overwritten. Check if Prometheus reserved labels are used.")
		}
		// internal labels should be maintained
		name := extras[i]
		if !(len(name) > 4 && name[:2] == "__" && name[len(name)-2:] == "__") {
			name = prometheustranslator.NormalizeLabel(name)
		}
		l[name] = extras[i+1]
	}

	labels = labels[:0]
	for k, v := range l {
		labels = append(labels, prompb.Label{Name: k, Value: v})
	}

	return labels
}

// isValidAggregationTemporality checks whether an OTel metric has a valid
// aggregation temporality for conversion to a Prometheus metric.
func isValidAggregationTemporality(metric pmetric.Metric) bool {
	//exhaustive:enforce
	switch metric.Type() {
	case pmetric.MetricTypeGauge, pmetric.MetricTypeSummary:
		return true
	case pmetric.MetricTypeSum:
		return metric.Sum().AggregationTemporality() == pmetric.AggregationTemporalityCumulative
	case pmetric.MetricTypeHistogram:
		return metric.Histogram().AggregationTemporality() == pmetric.AggregationTemporalityCumulative
	case pmetric.MetricTypeExponentialHistogram:
		return metric.ExponentialHistogram().AggregationTemporality() == pmetric.AggregationTemporalityCumulative
	}
	return false
}

func (c *PrometheusConverter) AddHistogramDataPoints(dataPoints pmetric.HistogramDataPointSlice,
	resource pcommon.Resource, settings Settings, baseName string) error {
	createLabels := func(nameSuffix string, baseLabels []prompb.Label, extras ...string) []prompb.Label {
		extraLabelCount := len(extras) / 2
		labels := make([]prompb.Label, len(baseLabels), len(baseLabels)+extraLabelCount+1) // +1 for name
		copy(labels, baseLabels)

		for extrasIdx := 0; extrasIdx < extraLabelCount; extrasIdx++ {
			labels = append(labels, prompb.Label{Name: extras[extrasIdx], Value: extras[extrasIdx+1]})
		}

		// sum, count, and buckets of the histogram should append suffix to baseName
		labels = append(labels, prompb.Label{Name: model.MetricNameLabel, Value: baseName + nameSuffix})

		return labels
	}

	for x := 0; x < dataPoints.Len(); x++ {
		pt := dataPoints.At(x)
		timestamp := convertTimeStamp(pt.Timestamp())
		baseLabels := createAttributes(resource, pt.Attributes(), settings.ExternalLabels)

		// If the sum is unset, it indicates the _sum metric point should be
		// omitted
		if pt.HasSum() {
			// treat sum as a sample in an individual TimeSeries
			sum := &prompb.Sample{
				Value:     pt.Sum(),
				Timestamp: timestamp,
			}
			if pt.Flags().NoRecordedValue() {
				sum.Value = math.Float64frombits(value.StaleNaN)
			}

			sumlabels := createLabels(sumStr, baseLabels)
			c.AddSample(sum, sumlabels)

		}

		// treat count as a sample in an individual TimeSeries
		count := &prompb.Sample{
			Value:     float64(pt.Count()),
			Timestamp: timestamp,
		}
		if pt.Flags().NoRecordedValue() {
			count.Value = math.Float64frombits(value.StaleNaN)
		}

		countlabels := createLabels(countStr, baseLabels)
		c.AddSample(count, countlabels)

		// cumulative count for conversion to cumulative histogram
		var cumulativeCount uint64

		var bucketBounds []bucketBoundsData

		// process each bound, based on histograms proto definition, # of buckets = # of explicit bounds + 1
		for i := 0; i < pt.ExplicitBounds().Len() && i < pt.BucketCounts().Len(); i++ {
			bound := pt.ExplicitBounds().At(i)
			cumulativeCount += pt.BucketCounts().At(i)
			bucket := &prompb.Sample{
				Value:     float64(cumulativeCount),
				Timestamp: timestamp,
			}
			if pt.Flags().NoRecordedValue() {
				bucket.Value = math.Float64frombits(value.StaleNaN)
			}
			boundStr := strconv.FormatFloat(bound, 'f', -1, 64)
			labels := createLabels(bucketStr, baseLabels, leStr, boundStr)
			ts := c.AddSample(bucket, labels)

			bucketBounds = append(bucketBounds, bucketBoundsData{ts: ts, bound: bound})
		}
		// add le=+Inf bucket
		infBucket := &prompb.Sample{
			Timestamp: timestamp,
		}
		if pt.Flags().NoRecordedValue() {
			infBucket.Value = math.Float64frombits(value.StaleNaN)
		} else {
			infBucket.Value = float64(pt.Count())
		}
		infLabels := createLabels(bucketStr, baseLabels, leStr, pInfStr)
		ts := c.AddSample(infBucket, infLabels)

		bucketBounds = append(bucketBounds, bucketBoundsData{ts: ts, bound: math.Inf(1)})
		c.AddExemplars(pt, bucketBounds)

		startTimestamp := pt.StartTimestamp()
		if settings.ExportCreatedMetric && startTimestamp != 0 {
			labels := createLabels(createdSuffix, baseLabels)
			c.AddMetricIfNeeded(labels, startTimestamp, pt.Timestamp())
		}
	}

	return nil
}

type exemplarType interface {
	pmetric.ExponentialHistogramDataPoint | pmetric.HistogramDataPoint | pmetric.NumberDataPoint
	Exemplars() pmetric.ExemplarSlice
}

func getPromExemplars[T exemplarType](pt T) []prompb.Exemplar {
	promExemplars := make([]prompb.Exemplar, 0, pt.Exemplars().Len())
	for i := 0; i < pt.Exemplars().Len(); i++ {
		exemplar := pt.Exemplars().At(i)
		exemplarRunes := 0

		promExemplar := prompb.Exemplar{
			Value:     exemplar.DoubleValue(),
			Timestamp: timestamp.FromTime(exemplar.Timestamp().AsTime()),
		}
		if traceID := exemplar.TraceID(); !traceID.IsEmpty() {
			val := hex.EncodeToString(traceID[:])
			exemplarRunes += utf8.RuneCountInString(traceIDKey) + utf8.RuneCountInString(val)
			promLabel := prompb.Label{
				Name:  traceIDKey,
				Value: val,
			}
			promExemplar.Labels = append(promExemplar.Labels, promLabel)
		}
		if spanID := exemplar.SpanID(); !spanID.IsEmpty() {
			val := hex.EncodeToString(spanID[:])
			exemplarRunes += utf8.RuneCountInString(spanIDKey) + utf8.RuneCountInString(val)
			promLabel := prompb.Label{
				Name:  spanIDKey,
				Value: val,
			}
			promExemplar.Labels = append(promExemplar.Labels, promLabel)
		}

		attrs := exemplar.FilteredAttributes()
		labelsFromAttributes := make([]prompb.Label, 0, attrs.Len())
		attrs.Range(func(key string, value pcommon.Value) bool {
			val := value.AsString()
			exemplarRunes += utf8.RuneCountInString(key) + utf8.RuneCountInString(val)
			promLabel := prompb.Label{
				Name:  key,
				Value: val,
			}

			labelsFromAttributes = append(labelsFromAttributes, promLabel)

			return true
		})
		if exemplarRunes <= maxExemplarRunes {
			// only append filtered attributes if it does not cause exemplar
			// labels to exceed the max number of runes
			promExemplar.Labels = append(promExemplar.Labels, labelsFromAttributes...)
		}

		promExemplars = append(promExemplars, promExemplar)
	}

	return promExemplars
}

// mostRecentTimestampInMetric returns the latest timestamp in a batch of metrics
func mostRecentTimestampInMetric(metric pmetric.Metric) pcommon.Timestamp {
	var ts pcommon.Timestamp
	// handle individual metric based on type
	//exhaustive:enforce
	switch metric.Type() {
	case pmetric.MetricTypeGauge:
		dataPoints := metric.Gauge().DataPoints()
		for x := 0; x < dataPoints.Len(); x++ {
			ts = maxTimestamp(ts, dataPoints.At(x).Timestamp())
		}
	case pmetric.MetricTypeSum:
		dataPoints := metric.Sum().DataPoints()
		for x := 0; x < dataPoints.Len(); x++ {
			ts = maxTimestamp(ts, dataPoints.At(x).Timestamp())
		}
	case pmetric.MetricTypeHistogram:
		dataPoints := metric.Histogram().DataPoints()
		for x := 0; x < dataPoints.Len(); x++ {
			ts = maxTimestamp(ts, dataPoints.At(x).Timestamp())
		}
	case pmetric.MetricTypeExponentialHistogram:
		dataPoints := metric.ExponentialHistogram().DataPoints()
		for x := 0; x < dataPoints.Len(); x++ {
			ts = maxTimestamp(ts, dataPoints.At(x).Timestamp())
		}
	case pmetric.MetricTypeSummary:
		dataPoints := metric.Summary().DataPoints()
		for x := 0; x < dataPoints.Len(); x++ {
			ts = maxTimestamp(ts, dataPoints.At(x).Timestamp())
		}
	}
	return ts
}

func maxTimestamp(a, b pcommon.Timestamp) pcommon.Timestamp {
	if a > b {
		return a
	}
	return b
}

func (c *PrometheusConverter) AddSummaryDataPoints(dataPoints pmetric.SummaryDataPointSlice, resource pcommon.Resource,
	settings Settings, baseName string) error {
	createLabels := func(name string, baseLabels []prompb.Label, extras ...string) []prompb.Label {
		extraLabelCount := len(extras) / 2
		labels := make([]prompb.Label, len(baseLabels), len(baseLabels)+extraLabelCount+1) // +1 for name
		copy(labels, baseLabels)

		for extrasIdx := 0; extrasIdx < extraLabelCount; extrasIdx++ {
			labels = append(labels, prompb.Label{Name: extras[extrasIdx], Value: extras[extrasIdx+1]})
		}

		labels = append(labels, prompb.Label{Name: model.MetricNameLabel, Value: name})

		return labels
	}

	for x := 0; x < dataPoints.Len(); x++ {
		pt := dataPoints.At(x)
		timestamp := convertTimeStamp(pt.Timestamp())
		baseLabels := createAttributes(resource, pt.Attributes(), settings.ExternalLabels)

		// treat sum as a sample in an individual TimeSeries
		sum := &prompb.Sample{
			Value:     pt.Sum(),
			Timestamp: timestamp,
		}
		if pt.Flags().NoRecordedValue() {
			sum.Value = math.Float64frombits(value.StaleNaN)
		}
		// sum and count of the summary should append suffix to baseName
		sumlabels := createLabels(baseName+sumStr, baseLabels)
		c.AddSample(sum, sumlabels)

		// treat count as a sample in an individual TimeSeries
		count := &prompb.Sample{
			Value:     float64(pt.Count()),
			Timestamp: timestamp,
		}
		if pt.Flags().NoRecordedValue() {
			count.Value = math.Float64frombits(value.StaleNaN)
		}
		countlabels := createLabels(baseName+countStr, baseLabels)
		c.AddSample(count, countlabels)

		// process each percentile/quantile
		for i := 0; i < pt.QuantileValues().Len(); i++ {
			qt := pt.QuantileValues().At(i)
			quantile := &prompb.Sample{
				Value:     qt.Value(),
				Timestamp: timestamp,
			}
			if pt.Flags().NoRecordedValue() {
				quantile.Value = math.Float64frombits(value.StaleNaN)
			}
			percentileStr := strconv.FormatFloat(qt.Quantile(), 'f', -1, 64)
			qtlabels := createLabels(baseName, baseLabels, quantileStr, percentileStr)
			c.AddSample(quantile, qtlabels)
		}

		startTimestamp := pt.StartTimestamp()
		if settings.ExportCreatedMetric && startTimestamp != 0 {
			createdLabels := createLabels(baseName+createdSuffix, baseLabels)
			c.AddMetricIfNeeded(createdLabels, startTimestamp, pt.Timestamp())
		}
	}

	return nil
}

func (c *PrometheusConverter) AddMetricIfNeeded(lbls []prompb.Label, startTimestamp pcommon.Timestamp, timestamp pcommon.Timestamp) {
	h := timeSeriesSignature(lbls)
	ts := c.unique[h]
	if ts != nil {
		if isSameMetric(ts, lbls) {
			// We already have this metric
			return
		}

		// Look for a matching conflict
		for _, ts := range c.conflicts[h] {
			if isSameMetric(ts, lbls) {
				// We already have this metric
				return
			}
		}

		// New conflict
		c.conflicts[h] = append(c.conflicts[h], &prompb.TimeSeries{
			Labels: lbls,
			Samples: []prompb.Sample{
				{ // convert ns to ms
					Value:     float64(convertTimeStamp(startTimestamp)),
					Timestamp: convertTimeStamp(timestamp),
				},
			},
		})
		return
	}

	// This metric is new
	c.unique[h] = &prompb.TimeSeries{
		Labels: lbls,
		Samples: []prompb.Sample{
			{ // convert ns to ms
				Value:     float64(convertTimeStamp(startTimestamp)),
				Timestamp: convertTimeStamp(timestamp),
			},
		},
	}
}

// addResourceTargetInfo converts the resource to the target info metric
func addResourceTargetInfo(resource pcommon.Resource, settings Settings, timestamp pcommon.Timestamp, converter MetricsConverter) {
	if settings.DisableTargetInfo {
		return
	}

	attributes := resource.Attributes()
	serviceName, haveServiceName := attributes.Get(conventions.AttributeServiceName)
	serviceNamespace, haveServiceNamespace := attributes.Get(conventions.AttributeServiceNamespace)
	instance, haveInstanceID := attributes.Get(conventions.AttributeServiceInstanceID)

	nonIdentifyingAttrs := attributes.Len()
	if haveServiceName {
		nonIdentifyingAttrs--
	}
	if haveInstanceID {
		nonIdentifyingAttrs--
	}
	if haveServiceNamespace {
		nonIdentifyingAttrs--
	}
	if nonIdentifyingAttrs == 0 {
		// If we only have job + instance, then target_info isn't useful, so don't add it.
		return
	}

	name := targetMetricName
	if len(settings.Namespace) > 0 {
		name = settings.Namespace + "_" + name
	}

	// Calculate the maximum possible number of labels we could return so we can preallocate l
	maxLabelCount := attributes.Len() + len(settings.ExternalLabels) + 1
	// map ensures no duplicate label names
	ls := make(map[string]string, maxLabelCount)

	// Pre-allocate the labels.
	labels := make([]prompb.Label, 0, maxLabelCount)
	attributes.Range(func(key string, value pcommon.Value) bool {
		// Ignore resource attributes used for job + instance
		if key != conventions.AttributeServiceName && key != conventions.AttributeServiceNamespace && key != conventions.AttributeServiceInstanceID {
			labels = append(labels, prompb.Label{Name: key, Value: value.AsString()})
		}
		return true
	})
	// Ensure attributes are sorted by key for consistent merging of keys which
	// collide when sanitized.
	sort.Stable(ByLabelName(labels))

	for _, label := range labels {
		finalKey := prometheustranslator.NormalizeLabel(label.Name)
		if existingValue, alreadyExists := ls[finalKey]; alreadyExists {
			ls[finalKey] = existingValue + ";" + label.Value
		} else {
			ls[finalKey] = label.Value
		}
	}

	// Map service.name + service.namespace to job
	if haveServiceName {
		val := serviceName.AsString()
		if haveServiceNamespace {
			val = fmt.Sprintf("%s/%s", serviceNamespace.AsString(), val)
		}
		ls[model.JobLabel] = val
	}
	// Map service.instance.id to instance
	if haveInstanceID {
		ls[model.InstanceLabel] = instance.AsString()
	}
	for key, value := range settings.ExternalLabels {
		// External labels have already been sanitized
		if _, alreadyExists := ls[key]; alreadyExists {
			// Skip external labels if they are overridden by metric attributes
			continue
		}
		ls[key] = value
	}

	if _, found := ls[model.MetricNameLabel]; found {
		log.Println("label " + model.MetricNameLabel + " is overwritten. Check if Prometheus reserved labels are used.")
	}
	ls[model.MetricNameLabel] = name

	labels = labels[:0]
	for k, v := range ls {
		labels = append(labels, prompb.Label{Name: k, Value: v})
	}

	sample := &prompb.Sample{
		Value: float64(1),
		// convert ns to ms
		Timestamp: convertTimeStamp(timestamp),
	}
	converter.AddSample(sample, labels)
}

// convertTimeStamp converts OTLP timestamp in ns to timestamp in ms
func convertTimeStamp(timestamp pcommon.Timestamp) int64 {
	return timestamp.AsTime().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
