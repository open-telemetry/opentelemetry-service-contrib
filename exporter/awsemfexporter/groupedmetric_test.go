// Copyright 2020, OpenTelemetry Authors
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

package awsemfexporter

import (
	"testing"
	"time"

	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"
	"go.opentelemetry.io/collector/translator/internaldata"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestAddToGroupedMetric(t *testing.T) {
	namespace := "namespace"
	instrumentationLibName := "cloudwatch-otel"
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	logger := zap.NewNop()

	metadata := CWMetricMetadata{
		Namespace:                  namespace,
		Timestamp:                  timestamp,
		LogGroup:                   logGroup,
		LogStream:                  logStreamName,
		InstrumentationLibraryName: instrumentationLibName,
	}

	testCases := []struct {
		testName string
		metric   *metricspb.Metric
		expected map[string]*MetricInfo
	}{
		{
			"Int gauge",
			generateTestIntGauge("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: int64(1),
					Unit:  "Count",
				},
			},
		},
		{
			"Double gauge",
			generateTestDoubleGauge("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: 0.1,
					Unit:  "Count",
				},
			},
		},
		{
			"Int sum",
			generateTestIntSum("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: int64(0),
					Unit:  "Count",
				},
			},
		},
		{
			"Double sum",
			generateTestDoubleSum("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: float64(0),
					Unit:  "Count",
				},
			},
		},
		{
			"Double histogram",
			generateTestDoubleHistogram("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: &CWMetricStats{
						Min:   0,
						Max:   10,
						Count: 18,
						Sum:   35.0,
					},
					Unit: "Seconds",
				},
			},
		},
		{
			"Summary",
			generateTestSummary("foo"),
			map[string]*MetricInfo{
				"foo": {
					Value: &CWMetricStats{
						Min:   1,
						Max:   5,
						Count: 5,
						Sum:   15,
					},
					Unit: "Seconds",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			groupedMetrics := make(map[string]*GroupedMetric)
			oc := consumerdata.MetricsData{
				Node: &commonpb.Node{},
				Resource: &resourcepb.Resource{
					Labels: map[string]string{
						conventions.AttributeServiceName:      "myServiceName",
						conventions.AttributeServiceNamespace: "myServiceNS",
					},
				},
				Metrics: []*metricspb.Metric{tc.metric},
			}

			// Retrieve *pdata.Metric
			rm := internaldata.OCToMetrics(oc)
			rms := rm.ResourceMetrics()
			assert.Equal(t, 1, rms.Len())
			ilms := rms.At(0).InstrumentationLibraryMetrics()
			assert.Equal(t, 1, ilms.Len())
			metrics := ilms.At(0).Metrics()
			assert.Equal(t, 1, metrics.Len())
			metric := metrics.At(0)

			addToGroupedMetric(&metric, groupedMetrics, metadata, zap.NewNop())
			expectedLabels := map[string]string{
				(OTellibDimensionKey): instrumentationLibName,
				"label1":              "value1",
			}

			for _, v := range groupedMetrics {
				assert.Equal(t, len(tc.expected), len(v.Metrics))
				assert.Equal(t, tc.expected, v.Metrics)
				assert.Equal(t, 2, len(v.Labels))
				assert.Equal(t, metadata, v.Metadata)
				assert.Equal(t, expectedLabels, v.Labels)
			}
		})
	}

	t.Run("Add multiple different metrics", func(t *testing.T) {
		groupedMetrics := make(map[string]*GroupedMetric)
		oc := consumerdata.MetricsData{
			Node: &commonpb.Node{},
			Resource: &resourcepb.Resource{
				Labels: map[string]string{
					conventions.AttributeServiceName:      "myServiceName",
					conventions.AttributeServiceNamespace: "myServiceNS",
				},
			},
			Metrics: []*metricspb.Metric{
				generateTestIntGauge("int-gauge"),
				generateTestDoubleGauge("double-gauge"),
				generateTestIntSum("int-sum"),
				generateTestDoubleSum("double-sum"),
				generateTestDoubleHistogram("double-histogram"),
				generateTestSummary("summary"),
			},
		}
		rm := internaldata.OCToMetrics(oc)
		rms := rm.ResourceMetrics()
		ilms := rms.At(0).InstrumentationLibraryMetrics()
		metrics := ilms.At(0).Metrics()
		assert.Equal(t, 6, metrics.Len())

		for i := 0; i < metrics.Len(); i++ {
			metric := metrics.At(i)
			addToGroupedMetric(&metric, groupedMetrics, metadata, logger)
		}

		assert.Equal(t, 1, len(groupedMetrics))
		for _, group := range groupedMetrics {
			assert.Equal(t, 6, len(group.Metrics))
			for metricName, metricInfo := range group.Metrics {
				if metricName == "double-histogram" || metricName == "summary" {
					assert.Equal(t, "Seconds", metricInfo.Unit)
				} else {
					assert.Equal(t, "Count", metricInfo.Unit)
				}
			}
			expectedLabels := map[string]string{
				(OTellibDimensionKey): "cloudwatch-otel",
				"label1":              "value1",
			}
			assert.Equal(t, expectedLabels, group.Labels)
			assert.Equal(t, metadata, group.Metadata)
		}
	})

	t.Run("Add same metric but different log group", func(t *testing.T) {
		groupedMetrics := make(map[string]*GroupedMetric)
		oc := consumerdata.MetricsData{
			Metrics: []*metricspb.Metric{
				generateTestIntGauge("int-gauge"),
			},
		}
		rm := internaldata.OCToMetrics(oc)
		ilms := rm.ResourceMetrics().At(0).InstrumentationLibraryMetrics()
		metric := ilms.At(0).Metrics().At(0)

		metricMetadata1 := CWMetricMetadata{
			Namespace:                  namespace,
			Timestamp:                  timestamp,
			LogGroup:                   "log-group-1",
			LogStream:                  logStreamName,
			InstrumentationLibraryName: instrumentationLibName,
		}
		addToGroupedMetric(&metric, groupedMetrics, metricMetadata1, logger)

		metricMetadata2 := CWMetricMetadata{
			Namespace:                  namespace,
			Timestamp:                  timestamp,
			LogGroup:                   "log-group-2",
			LogStream:                  logStreamName,
			InstrumentationLibraryName: instrumentationLibName,
		}
		addToGroupedMetric(&metric, groupedMetrics, metricMetadata2, logger)

		assert.Equal(t, 2, len(groupedMetrics))
		seenLogGroup1 := false
		seenLogGroup2 := false
		for _, group := range groupedMetrics {
			assert.Equal(t, 1, len(group.Metrics))
			expectedMetrics := map[string]*MetricInfo{
				"int-gauge": {
					Value: int64(1),
					Unit:  "Count",
				},
			}
			assert.Equal(t, expectedMetrics, group.Metrics)
			expectedLabels := map[string]string{
				(OTellibDimensionKey): "cloudwatch-otel",
				"label1":              "value1",
			}
			assert.Equal(t, expectedLabels, group.Labels)

			if group.Metadata.LogGroup == "log-group-2" {
				seenLogGroup2 = true
			} else if group.Metadata.LogGroup == "log-group-1" {
				seenLogGroup1 = true
			}
		}
		assert.True(t, seenLogGroup1)
		assert.True(t, seenLogGroup2)
	})

	t.Run("Duplicate metric names", func(t *testing.T) {
		groupedMetrics := make(map[string]*GroupedMetric)
		oc := consumerdata.MetricsData{
			Resource: &resourcepb.Resource{
				Labels: map[string]string{
					conventions.AttributeServiceName:      "myServiceName",
					conventions.AttributeServiceNamespace: "myServiceNS",
				},
			},
			Metrics: []*metricspb.Metric{
				generateTestIntGauge("foo"),
				generateTestDoubleGauge("foo"),
			},
		}
		rm := internaldata.OCToMetrics(oc)
		rms := rm.ResourceMetrics()
		ilms := rms.At(0).InstrumentationLibraryMetrics()
		metrics := ilms.At(0).Metrics()
		assert.Equal(t, 2, metrics.Len())

		obs, logs := observer.New(zap.WarnLevel)
		obsLogger := zap.New(obs)

		for i := 0; i < metrics.Len(); i++ {
			metric := metrics.At(i)
			addToGroupedMetric(&metric, groupedMetrics, metadata, obsLogger)
		}
		assert.Equal(t, 1, len(groupedMetrics))

		labels := map[string]string{
			(OTellibDimensionKey): instrumentationLibName,
			"label1":              "value1",
		}
		// Test output warning logs
		expectedLogs := []observer.LoggedEntry{
			{
				Entry: zapcore.Entry{Level: zap.WarnLevel, Message: "Duplicate metric found"},
				Context: []zapcore.Field{
					zap.String("Name", "foo"),
					zap.Any("Labels", labels),
				},
			},
		}
		assert.Equal(t, 1, logs.Len())
		assert.Equal(t, expectedLogs, logs.AllUntimed())
	})

	t.Run("Unhandled metric type", func(t *testing.T) {
		groupedMetrics := make(map[string]*GroupedMetric)
		md := pdata.NewMetrics()
		rms := md.ResourceMetrics()
		rms.Resize(1)
		rms.At(0).InstrumentationLibraryMetrics().Resize(1)
		rms.At(0).InstrumentationLibraryMetrics().At(0).Metrics().Resize(1)
		metric := rms.At(0).InstrumentationLibraryMetrics().At(0).Metrics().At(0)
		metric.SetName("foo")
		metric.SetUnit("Count")
		metric.SetDataType(pdata.MetricDataTypeIntHistogram)

		obs, logs := observer.New(zap.WarnLevel)
		obsLogger := zap.New(obs)
		addToGroupedMetric(&metric, groupedMetrics, metadata, obsLogger)
		assert.Equal(t, 0, len(groupedMetrics))

		// Test output warning logs
		expectedLogs := []observer.LoggedEntry{
			{
				Entry: zapcore.Entry{Level: zap.WarnLevel, Message: "Unhandled metric data type."},
				Context: []zapcore.Field{
					zap.String("DataType", "IntHistogram"),
					zap.String("Name", "foo"),
					zap.String("Unit", "Count"),
				},
			},
		}
		assert.Equal(t, 1, logs.Len())
		assert.Equal(t, expectedLogs, logs.AllUntimed())
	})

	t.Run("Nil metric", func(t *testing.T) {
		groupedMetrics := make(map[string]*GroupedMetric)
		addToGroupedMetric(nil, groupedMetrics, metadata, logger)
		assert.Equal(t, 0, len(groupedMetrics))
	})
}

func BenchmarkAddToGroupedMetric(b *testing.B) {
	oc := consumerdata.MetricsData{
		Metrics: []*metricspb.Metric{
			generateTestIntGauge("int-gauge"),
			generateTestDoubleGauge("double-gauge"),
			generateTestIntSum("int-sum"),
			generateTestDoubleSum("double-sum"),
			generateTestDoubleHistogram("double-histogram"),
			generateTestSummary("summary"),
		},
	}
	rms := internaldata.OCToMetrics(oc).ResourceMetrics()
	metrics := rms.At(0).InstrumentationLibraryMetrics().At(0).Metrics()
	numMetrics := metrics.Len()

	metadata := CWMetricMetadata{
		Namespace:                  "Namespace",
		Timestamp:                  int64(1596151098037),
		LogGroup:                   "log-group",
		LogStream:                  "log-stream",
		InstrumentationLibraryName: "cloudwatch-otel",
	}

	logger := zap.NewNop()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		groupedMetrics := make(map[string]*GroupedMetric)
		for i := 0; i < numMetrics; i++ {
			metric := metrics.At(i)
			addToGroupedMetric(&metric, groupedMetrics, metadata, logger)
		}
	}
}
