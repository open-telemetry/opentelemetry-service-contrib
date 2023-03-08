// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/otel/attribute"
)

var (
	errEmptyMetricName  = errors.New("empty metric name")
	errEmptyMetricValue = errors.New("empty metric value")
)

type (
	MetricType   string // From the statsd line e.g., "c", "g", "h"
	TypeName     string // How humans describe the MetricTypes ("counter", "gauge")
	ObserverType string // How the server will aggregate histogram and timings ("gauge", "summary")
)

const (
	tagMetricType = "metric_type"

	CounterType   MetricType = "c"
	GaugeType     MetricType = "g"
	HistogramType MetricType = "h"
	TimingType    MetricType = "ms"

	CounterTypeName   TypeName = "counter"
	GaugeTypeName     TypeName = "gauge"
	HistogramTypeName TypeName = "histogram"
	TimingTypeName    TypeName = "timing"
	TimingAltTypeName TypeName = "timer"

	GaugeObserver     ObserverType = "gauge"
	SummaryObserver   ObserverType = "summary"
	HistogramObserver ObserverType = "histogram"
	DisableObserver   ObserverType = "disabled"

	DefaultObserverType = DisableObserver
)

type TimerHistogramMapping struct {
	StatsdType   TypeName        `mapstructure:"statsd_type"`
	ObserverType ObserverType    `mapstructure:"observer_type"`
	Histogram    HistogramConfig `mapstructure:"histogram"`
}

type HistogramConfig struct {
	MaxSize int32 `mapstructure:"max_size"`
}

type ObserverCategory struct {
	method          ObserverType
	histogramConfig structure.Config
}

var defaultObserverCategory = ObserverCategory{
	method: DefaultObserverType,
}

// StatsDParser supports the Parse method for parsing StatsD messages with Tags.
type StatsDParser struct {
	gauges                 map[statsDMetricDescription]pmetric.ScopeMetrics
	counters               map[statsDMetricDescription]pmetric.ScopeMetrics
	summaries              map[statsDMetricDescription]summaryMetric
	histograms             map[statsDMetricDescription]histogramMetric
	timersAndDistributions []pmetric.ScopeMetrics
	enableMetricType       bool
	isMonotonicCounter     bool
	timerEvents            ObserverCategory
	histogramEvents        ObserverCategory
	lastIntervalTime       time.Time
}

type sampleValue struct {
	value float64
	count float64
}

type summaryMetric struct {
	points  []float64
	weights []float64
}

type histogramStructure = structure.Histogram[float64]

type histogramMetric struct {
	agg *histogramStructure
}

type statsDMetric struct {
	description statsDMetricDescription
	asFloat     float64
	addition    bool
	unit        string
	sampleRate  float64
}

type statsDMetricDescription struct {
	name       string
	metricType MetricType
	attrs      attribute.Set
}

func (t MetricType) FullName() TypeName {
	switch t {
	case GaugeType:
		return GaugeTypeName
	case CounterType:
		return CounterTypeName
	case TimingType:
		return TimingTypeName
	case HistogramType:
		return HistogramTypeName
	}
	return TypeName(fmt.Sprintf("unknown(%s)", t))
}

func (p *StatsDParser) resetState(when time.Time) {
	p.lastIntervalTime = when
	p.gauges = make(map[statsDMetricDescription]pmetric.ScopeMetrics)
	p.counters = make(map[statsDMetricDescription]pmetric.ScopeMetrics)
	p.timersAndDistributions = nil
	p.summaries = make(map[statsDMetricDescription]summaryMetric)
	p.histograms = make(map[statsDMetricDescription]histogramMetric)
}

func (p *StatsDParser) Initialize(enableMetricType bool, isMonotonicCounter bool, sendTimerHistogram []TimerHistogramMapping) error {
	p.resetState(timeNowFunc())

	p.histogramEvents = defaultObserverCategory
	p.timerEvents = defaultObserverCategory
	p.enableMetricType = enableMetricType
	p.isMonotonicCounter = isMonotonicCounter
	// Note: validation occurs in ("../".Config).validate()
	for _, eachMap := range sendTimerHistogram {
		switch eachMap.StatsdType {
		case HistogramTypeName:
			p.histogramEvents.method = eachMap.ObserverType
			p.histogramEvents.histogramConfig = expoHistogramConfig(eachMap.Histogram)
		case TimingTypeName, TimingAltTypeName:
			p.timerEvents.method = eachMap.ObserverType
			p.timerEvents.histogramConfig = expoHistogramConfig(eachMap.Histogram)
		}
	}
	return nil
}

func expoHistogramConfig(opts HistogramConfig) structure.Config {
	var r []structure.Option
	if opts.MaxSize >= structure.MinSize {
		r = append(r, structure.WithMaxSize(opts.MaxSize))
	}
	return structure.NewConfig(r...)
}

// GetMetrics gets the metrics preparing for flushing and reset the state.
func (p *StatsDParser) GetMetrics() pmetric.Metrics {
	metrics := pmetric.NewMetrics()
	rm := metrics.ResourceMetrics().AppendEmpty()

	for _, metric := range p.gauges {
		metric.CopyTo(rm.ScopeMetrics().AppendEmpty())
	}

	for _, metric := range p.timersAndDistributions {
		metric.CopyTo(rm.ScopeMetrics().AppendEmpty())
	}

	now := timeNowFunc()

	for _, metric := range p.counters {
		setTimestampsForCounterMetric(metric, p.lastIntervalTime, now)
		metric.CopyTo(rm.ScopeMetrics().AppendEmpty())
	}

	for desc, summaryMetric := range p.summaries {
		buildSummaryMetric(
			desc,
			summaryMetric,
			p.lastIntervalTime,
			now,
			statsDDefaultPercentiles,
			rm.ScopeMetrics().AppendEmpty(),
		)
	}

	for desc, histogramMetric := range p.histograms {
		buildHistogramMetric(
			desc,
			histogramMetric,
			p.lastIntervalTime,
			now,
			rm.ScopeMetrics().AppendEmpty(),
		)
	}
	p.resetState(now)
	return metrics
}

var timeNowFunc = time.Now

func (p *StatsDParser) observerCategoryFor(t MetricType) ObserverCategory {
	switch t {
	case HistogramType:
		return p.histogramEvents
	case TimingType:
		return p.timerEvents
	}
	return defaultObserverCategory
}

// Aggregate for each metric line.
func (p *StatsDParser) Aggregate(line string) error {
	parsedMetric, err := parseMessageToMetric(line, p.enableMetricType)
	if err != nil {
		return err
	}
	switch parsedMetric.description.metricType {
	case GaugeType:
		_, ok := p.gauges[parsedMetric.description]
		if !ok {
			p.gauges[parsedMetric.description] = buildGaugeMetric(parsedMetric, timeNowFunc())
		} else {
			if parsedMetric.addition {
				point := p.gauges[parsedMetric.description].Metrics().At(0).Gauge().DataPoints().At(0)
				point.SetDoubleValue(point.DoubleValue() + parsedMetric.gaugeValue())
			} else {
				p.gauges[parsedMetric.description] = buildGaugeMetric(parsedMetric, timeNowFunc())
			}
		}

	case CounterType:
		_, ok := p.counters[parsedMetric.description]
		if !ok {
			p.counters[parsedMetric.description] = buildCounterMetric(parsedMetric, p.isMonotonicCounter)
		} else {
			point := p.counters[parsedMetric.description].Metrics().At(0).Sum().DataPoints().At(0)
			point.SetIntValue(point.IntValue() + parsedMetric.counterValue())
		}

	case TimingType, HistogramType:
		category := p.observerCategoryFor(parsedMetric.description.metricType)
		switch category.method {
		case GaugeObserver:
			p.timersAndDistributions = append(p.timersAndDistributions, buildGaugeMetric(parsedMetric, timeNowFunc()))
		case SummaryObserver:
			raw := parsedMetric.sampleValue()
			if existing, ok := p.summaries[parsedMetric.description]; !ok {
				p.summaries[parsedMetric.description] = summaryMetric{
					points:  []float64{raw.value},
					weights: []float64{raw.count},
				}
			} else {
				p.summaries[parsedMetric.description] = summaryMetric{
					points:  append(existing.points, raw.value),
					weights: append(existing.weights, raw.count),
				}
			}
		case HistogramObserver:
			raw := parsedMetric.sampleValue()
			var agg *histogramStructure
			if existing, ok := p.histograms[parsedMetric.description]; ok {
				agg = existing.agg
			} else {
				agg = new(histogramStructure)
				agg.Init(category.histogramConfig)

				p.histograms[parsedMetric.description] = histogramMetric{
					agg: agg,
				}
			}
			agg.UpdateByIncr(
				raw.value,
				uint64(raw.count), // Note! Rounding float64 to uint64 here.
			)

		case DisableObserver:
			// No action.
		}
	}

	return nil
}

func parseMessageToMetric(line string, enableMetricType bool) (statsDMetric, error) {
	result := statsDMetric{}

	parts := strings.Split(line, "|")
	if len(parts) < 2 {
		return result, fmt.Errorf("invalid message format: %s", line)
	}

	separatorIndex := strings.IndexByte(parts[0], ':')
	if separatorIndex < 0 {
		return result, fmt.Errorf("invalid <name>:<value> format: %s", parts[0])
	}

	result.description.name = parts[0][0:separatorIndex]
	if result.description.name == "" {
		return result, errEmptyMetricName
	}
	valueStr := parts[0][separatorIndex+1:]
	if valueStr == "" {
		return result, errEmptyMetricValue
	}
	if strings.HasPrefix(valueStr, "-") || strings.HasPrefix(valueStr, "+") {
		result.addition = true
	}

	inType := MetricType(parts[1])
	switch inType {
	case CounterType, GaugeType, HistogramType, TimingType:
		result.description.metricType = inType
	default:
		return result, fmt.Errorf("unsupported metric type: %s", inType)
	}

	additionalParts := parts[2:]

	var kvs []attribute.KeyValue

	for _, part := range additionalParts {
		switch {
		case strings.HasPrefix(part, "@"):
			sampleRateStr := strings.TrimPrefix(part, "@")

			f, err := strconv.ParseFloat(sampleRateStr, 64)
			if err != nil {
				return result, fmt.Errorf("parse sample rate: %s", sampleRateStr)
			}

			result.sampleRate = f
		case strings.HasPrefix(part, "#"):
			tagsStr := strings.TrimPrefix(part, "#")

			tagSets := strings.Split(tagsStr, ",")

			for _, tagSet := range tagSets {
				tagParts := strings.SplitN(tagSet, ":", 2)
				if len(tagParts) != 2 {
					return result, fmt.Errorf("invalid tag format: %s", tagParts)
				}
				k := tagParts[0]
				v := tagParts[1]
				kvs = append(kvs, attribute.String(k, v))
			}
		default:
			return result, fmt.Errorf("unrecognized message part: %s", part)
		}
	}
	var err error
	result.asFloat, err = strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return result, fmt.Errorf("parse metric value string: %s", valueStr)
	}

	// add metric_type dimension for all metrics
	if enableMetricType {
		metricType := string(result.description.metricType.FullName())

		kvs = append(kvs, attribute.String(tagMetricType, metricType))
	}

	if len(kvs) != 0 {
		result.description.attrs = attribute.NewSet(kvs...)
	}

	return result, nil
}
