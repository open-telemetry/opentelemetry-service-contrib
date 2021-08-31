// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opencensus

import (
	"time"

	occommon "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	agentmetricspb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/metrics/v1"
	ocmetrics "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	ocresource "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.5.0"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/occonventions"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/testdata"
)

func generateOCTestDataNoMetrics() *agentmetricspb.ExportMetricsServiceRequest {
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
	}
}

func generateOCTestDataNoPoints() *agentmetricspb.ExportMetricsServiceRequest {
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestGaugeDoubleMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_GAUGE_DOUBLE,
				},
			},
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestGaugeIntMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_GAUGE_INT64,
				},
			},
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestSumDoubleMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_CUMULATIVE_DOUBLE,
				},
			},
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestSumIntMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_CUMULATIVE_INT64,
				},
			},
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestDoubleHistogramMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_CUMULATIVE_DISTRIBUTION,
				},
			},
			{
				MetricDescriptor: &ocmetrics.MetricDescriptor{
					Name:        testdata.TestDoubleSummaryMetricName,
					Description: "",
					Unit:        "1",
					Type:        ocmetrics.MetricDescriptor_SUMMARY,
				},
			},
		},
	}
}

func generateOCTestDataNoLabels() *agentmetricspb.ExportMetricsServiceRequest {
	m := generateOCTestMetricCumulativeInt()
	m.MetricDescriptor.LabelKeys = nil
	m.Timeseries[0].LabelValues = nil
	m.Timeseries[1].LabelValues = nil
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{m},
	}
}

func generateOCTestDataMetricsOneMetric() *agentmetricspb.ExportMetricsServiceRequest {
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{generateOCTestMetricCumulativeInt()},
	}
}

func generateOCTestDataMetricsOneMetricOneNil() *agentmetricspb.ExportMetricsServiceRequest {
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{generateOCTestMetricCumulativeInt(), nil},
	}
}

func generateOCTestDataMetricsOneMetricOneNilTimeseries() *agentmetricspb.ExportMetricsServiceRequest {
	m := generateOCTestMetricCumulativeInt()
	m.Timeseries = append(m.Timeseries, nil)
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{m},
	}
}

func generateOCTestDataMetricsOneMetricOneNilPoint() *agentmetricspb.ExportMetricsServiceRequest {
	m := generateOCTestMetricCumulativeInt()
	m.Timeseries[0].Points = append(m.Timeseries[0].Points, nil)
	return &agentmetricspb.ExportMetricsServiceRequest{
		Node: &occommon.Node{},
		Resource: &ocresource.Resource{
			Labels: map[string]string{"resource-attr": "resource-attr-val-1"},
		},
		Metrics: []*ocmetrics.Metric{m},
	}
}

func generateOCTestMetricGaugeInt() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name:        testdata.TestGaugeIntMetricName,
			Description: "",
			Unit:        "1",
			Type:        ocmetrics.MetricDescriptor_GAUGE_INT64,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_Int64Value{
							Int64Value: 123,
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						HasValue: false,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_Int64Value{
							Int64Value: 456,
						},
					},
				},
			},
		},
	}
}

func generateOCTestMetricGaugeDouble() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name: testdata.TestGaugeDoubleMetricName,
			Unit: "1",
			Type: ocmetrics.MetricDescriptor_GAUGE_DOUBLE,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
				{Key: testdata.TestLabelKey3},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
					{
						// key3
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DoubleValue{
							DoubleValue: 1.23,
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
					{
						// key3
						Value:    testdata.TestLabelValue3,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DoubleValue{
							DoubleValue: 4.56,
						},
					},
				},
			},
		},
	}
}

func generateOCTestMetricCumulativeInt() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name:        testdata.TestSumIntMetricName,
			Description: "",
			Unit:        "1",
			Type:        ocmetrics.MetricDescriptor_CUMULATIVE_INT64,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_Int64Value{
							Int64Value: 123,
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						HasValue: false,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_Int64Value{
							Int64Value: 456,
						},
					},
				},
			},
		},
	}
}

func generateOCTestMetricCumulativeDouble() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name: testdata.TestSumDoubleMetricName,
			Unit: "1",
			Type: ocmetrics.MetricDescriptor_CUMULATIVE_DOUBLE,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
				{Key: testdata.TestLabelKey3},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
					{
						// key3
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DoubleValue{
							DoubleValue: 1.23,
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
					{
						// key3
						Value:    testdata.TestLabelValue3,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DoubleValue{
							DoubleValue: 4.56,
						},
					},
				},
			},
		},
	}
}

func generateOCTestMetricDoubleHistogram() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name:        testdata.TestDoubleHistogramMetricName,
			Description: "",
			Unit:        "1",
			Type:        ocmetrics.MetricDescriptor_CUMULATIVE_DISTRIBUTION,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
				{Key: testdata.TestLabelKey3},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
					{
						// key3
						Value:    testdata.TestLabelValue3,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DistributionValue{
							DistributionValue: &ocmetrics.DistributionValue{
								Count: 1,
								Sum:   15,
							},
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						HasValue: false,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
					{
						// key3
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_DistributionValue{
							DistributionValue: &ocmetrics.DistributionValue{
								Count: 1,
								Sum:   15,
								BucketOptions: &ocmetrics.DistributionValue_BucketOptions{
									Type: &ocmetrics.DistributionValue_BucketOptions_Explicit_{
										Explicit: &ocmetrics.DistributionValue_BucketOptions_Explicit{
											Bounds: []float64{1},
										},
									},
								},
								Buckets: []*ocmetrics.DistributionValue_Bucket{
									{
										Count: 0,
									},
									{
										Count: 1,
										Exemplar: &ocmetrics.DistributionValue_Exemplar{
											Timestamp:   timestamppb.New(testdata.TestMetricExemplarTime),
											Value:       15,
											Attachments: map[string]string{testdata.TestAttachmentKey: testdata.TestAttachmentValue},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func generateOCTestMetricDoubleSummary() *ocmetrics.Metric {
	return &ocmetrics.Metric{
		MetricDescriptor: &ocmetrics.MetricDescriptor{
			Name:        testdata.TestDoubleSummaryMetricName,
			Description: "",
			Unit:        "1",
			Type:        ocmetrics.MetricDescriptor_SUMMARY,
			LabelKeys: []*ocmetrics.LabelKey{
				{Key: testdata.TestLabelKey1},
				{Key: testdata.TestLabelKey2},
				{Key: testdata.TestLabelKey3},
			},
		},
		Timeseries: []*ocmetrics.TimeSeries{
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						Value:    testdata.TestLabelValue1,
						HasValue: true,
					},
					{
						// key2
						HasValue: false,
					},
					{
						// key3
						Value:    testdata.TestLabelValue3,
						HasValue: true,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_SummaryValue{
							SummaryValue: &ocmetrics.SummaryValue{
								Count: &wrapperspb.Int64Value{
									Value: 1,
								},
								Sum: &wrapperspb.DoubleValue{
									Value: 15,
								},
								Snapshot: &ocmetrics.SummaryValue_Snapshot{
									PercentileValues: nil,
								},
							},
						},
					},
				},
			},
			{
				StartTimestamp: timestamppb.New(testdata.TestMetricStartTime),
				LabelValues: []*ocmetrics.LabelValue{
					{
						// key1
						HasValue: false,
					},
					{
						// key2
						Value:    testdata.TestLabelValue2,
						HasValue: true,
					},
					{
						// key3
						HasValue: false,
					},
				},
				Points: []*ocmetrics.Point{
					{
						Timestamp: timestamppb.New(testdata.TestMetricTime),
						Value: &ocmetrics.Point_SummaryValue{
							SummaryValue: &ocmetrics.SummaryValue{
								Count: &wrapperspb.Int64Value{
									Value: 1,
								},
								Sum: &wrapperspb.DoubleValue{
									Value: 15,
								},
								Snapshot: &ocmetrics.SummaryValue_Snapshot{
									PercentileValues: []*ocmetrics.SummaryValue_Snapshot_ValueAtPercentile{
										{
											Percentile: 1,
											Value:      15,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func generateResourceWithOcNodeAndResource() pdata.Resource {
	resource := pdata.NewResource()
	resource.Attributes().InitFromMap(map[string]pdata.AttributeValue{
		occonventions.AttributeProcessStartTime:   pdata.NewAttributeValueString("2020-02-11T20:26:00Z"),
		conventions.AttributeHostName:             pdata.NewAttributeValueString("host1"),
		conventions.AttributeProcessPID:           pdata.NewAttributeValueInt(123),
		conventions.AttributeTelemetrySDKVersion:  pdata.NewAttributeValueString("v2.0.1"),
		occonventions.AttributeExporterVersion:    pdata.NewAttributeValueString("v1.2.0"),
		conventions.AttributeTelemetrySDKLanguage: pdata.NewAttributeValueString("cpp"),
		occonventions.AttributeResourceType:       pdata.NewAttributeValueString("good-resource"),
		"node-str-attr":                           pdata.NewAttributeValueString("node-str-attr-val"),
		"resource-str-attr":                       pdata.NewAttributeValueString("resource-str-attr-val"),
		"resource-int-attr":                       pdata.NewAttributeValueInt(123),
	})
	return resource
}

func generateOcNode() *occommon.Node {
	ts := timestamppb.New(time.Date(2020, 2, 11, 20, 26, 0, 0, time.UTC))

	return &occommon.Node{
		Identifier: &occommon.ProcessIdentifier{
			HostName:       "host1",
			Pid:            123,
			StartTimestamp: ts,
		},
		LibraryInfo: &occommon.LibraryInfo{
			Language:           occommon.LibraryInfo_CPP,
			ExporterVersion:    "v1.2.0",
			CoreLibraryVersion: "v2.0.1",
		},
		Attributes: map[string]string{
			"node-str-attr": "node-str-attr-val",
		},
	}
}

func generateOcResource() *ocresource.Resource {
	return &ocresource.Resource{
		Type: "good-resource",
		Labels: map[string]string{
			"resource-str-attr": "resource-str-attr-val",
			"resource-int-attr": "123",
		},
	}
}
