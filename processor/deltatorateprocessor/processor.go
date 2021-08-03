// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deltatorateprocessor

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

type deltaToRateProcessor struct {
	metrics  map[string]bool
	timeUnit StringTimeUnit
	logger   *zap.Logger
}

func newDeltaToRateProcessor(config *Config, logger *zap.Logger) *deltaToRateProcessor {
	inputMetricSet := make(map[string]bool, len(config.Metrics))
	for _, name := range config.Metrics {
		inputMetricSet[name] = true
	}

	return &deltaToRateProcessor{
		metrics:  inputMetricSet,
		timeUnit: config.TimeUnit,
		logger:   logger,
	}
}

// Start is invoked during service startup.
func (dtrp *deltaToRateProcessor) Start(context.Context, component.Host) error {
	return nil
}

// processMetrics implements the ProcessMetricsFunc type.
func (dtrp *deltaToRateProcessor) processMetrics(_ context.Context, md pdata.Metrics) (pdata.Metrics, error) {
	resourceMetricsSlice := md.ResourceMetrics()

	for i := 0; i < resourceMetricsSlice.Len(); i++ {
		rm := resourceMetricsSlice.At(i)
		ilms := rm.InstrumentationLibraryMetrics()
		for i := 0; i < ilms.Len(); i++ {
			ilm := ilms.At(i)
			metricSlice := ilm.Metrics()
			for j := 0; j < metricSlice.Len(); j++ {
				metric := metricSlice.At(j)
				_, ok := dtrp.metrics[metric.Name()]
				if ok {
					newDoubleDataPointSlice := pdata.NewNumberDataPointSlice()
					if metric.DataType() == pdata.MetricDataTypeSum && metric.Sum().AggregationTemporality() == pdata.AggregationTemporalityDelta {
						dataPoints := metric.Sum().DataPoints()

						for i := 0; i < dataPoints.Len(); i++ {
							fromDataPoint := dataPoints.At(i)
							newDp := newDoubleDataPointSlice.AppendEmpty()
							fromDataPoint.CopyTo(newDp)

							durationNanos := fromDataPoint.Timestamp().AsTime().Sub(fromDataPoint.StartTimestamp().AsTime())
							rate := calculateRate(fromDataPoint.Value(), durationNanos, dtrp.timeUnit)
							newDp.SetValue(rate)
						}

						metric.SetDataType(pdata.MetricDataTypeGauge)
						for d := 0; d < newDoubleDataPointSlice.Len(); d++ {
							dp := metric.Gauge().DataPoints().AppendEmpty()
							newDoubleDataPointSlice.At(d).CopyTo(dp)
						}
					}
				}
			}
		}
	}
	return md, nil
}

// Shutdown is invoked during service shutdown.
func (dtrp *deltaToRateProcessor) Shutdown(context.Context) error {
	return nil
}

func calculateRate(value float64, durationNanos time.Duration, timeUnit StringTimeUnit) float64 {
	duration := durationNanos.Seconds()

	switch timeUnit {
	case nanosecond:
		duration = float64(durationNanos.Nanoseconds())
	case millisecond:
		duration = float64(durationNanos.Milliseconds())
	case minute:
		duration = durationNanos.Minutes()
	}
	if duration > 0 {
		rate := value / duration
		return rate
	}
	return 0
}
