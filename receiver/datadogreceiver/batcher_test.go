// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogreceiver

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestMetricBatcher(t *testing.T) {
	tests := []struct {
		name   string
		series SeriesList
		expect func(t *testing.T, result pmetric.Metrics)
	}{
		{
			name: "Same metric, same tags, different hosts",
			series: SeriesList{
				Series: []datadogV1.Series{
					{
						Metric: "TestCount1",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:tag1", "service:test1", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
					{
						Metric: "TestCount1",
						Host:   strPtr("Host2"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:tag1", "service:test1", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
				},
			},
			expect: func(t *testing.T, result pmetric.Metrics) {
				// Different hosts should result in different ResourceMetrics
				require.Equal(t, 2, result.ResourceMetrics().Len())
				resource1 := result.ResourceMetrics().At(0)
				resource2 := result.ResourceMetrics().At(1)
				v, exists := resource1.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())
				v, exists = resource2.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host2", v.AsString())

				require.Equal(t, 1, resource1.ScopeMetrics().Len())
				require.Equal(t, 1, resource2.ScopeMetrics().Len())

				require.Equal(t, 1, resource1.ScopeMetrics().At(0).Metrics().Len())
				require.Equal(t, 1, resource2.ScopeMetrics().At(0).Metrics().Len())

				require.Equal(t, "TestCount1", resource1.ScopeMetrics().At(0).Metrics().At(0).Name())
				require.Equal(t, "TestCount1", resource2.ScopeMetrics().At(0).Metrics().At(0).Name())
			},
		},
		{
			name: "Same host, different metric names",
			series: SeriesList{
				Series: []datadogV1.Series{
					{
						Metric: "TestCount1",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:tag1", "service:test1", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
					{
						Metric: "TestCount2",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:tag1", "service:test1", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
				},
			},
			expect: func(t *testing.T, result pmetric.Metrics) {
				// The different metrics will fall under the same ResourceMetric and ScopeMetric
				// and there will be separate metrics under the ScopeMetric.Metrics()
				require.Equal(t, 1, result.ResourceMetrics().Len())
				resource := result.ResourceMetrics().At(0)

				v, exists := resource.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())

				require.Equal(t, 1, resource.ScopeMetrics().Len())

				require.Equal(t, 2, resource.ScopeMetrics().At(0).Metrics().Len())
				require.Equal(t, "TestCount1", resource.ScopeMetrics().At(0).Metrics().At(0).Name())
				require.Equal(t, "TestCount2", resource.ScopeMetrics().At(0).Metrics().At(1).Name())
			},
		},
		{
			name: "Same host, same metric name, single tag diff",
			series: SeriesList{
				Series: []datadogV1.Series{
					{
						Metric: "TestCount1",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:dev", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
					{
						Metric: "TestCount1",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:prod", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
				},
			},
			expect: func(t *testing.T, result pmetric.Metrics) {
				// Differences in attribute values should result in different resourceMetrics
				require.Equal(t, 2, result.ResourceMetrics().Len())
				resource1 := result.ResourceMetrics().At(0)
				resource2 := result.ResourceMetrics().At(1)
				v, exists := resource1.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())
				v, exists = resource2.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())
				v, exists = resource1.Resource().Attributes().Get("deployment.environment")
				require.True(t, exists)
				require.Equal(t, "dev", v.AsString())
				v, exists = resource2.Resource().Attributes().Get("deployment.environment")
				require.True(t, exists)
				require.Equal(t, "prod", v.AsString())

				require.Equal(t, 1, resource1.ScopeMetrics().Len())
				require.Equal(t, 1, resource1.ScopeMetrics().Len())

				require.Equal(t, 1, resource1.ScopeMetrics().At(0).Metrics().Len())
				require.Equal(t, 1, resource2.ScopeMetrics().At(0).Metrics().Len())

				require.Equal(t, 1, resource1.ScopeMetrics().At(0).Metrics().Len())
				require.Equal(t, 1, resource2.ScopeMetrics().At(0).Metrics().Len())

				require.Equal(t, "TestCount1", resource1.ScopeMetrics().At(0).Metrics().At(0).Name())
				require.Equal(t, "TestCount1", resource2.ScopeMetrics().At(0).Metrics().At(0).Name())
			},
		},
		{
			name: "Same host, same metric name, same tags, different type",
			series: SeriesList{
				Series: []datadogV1.Series{
					{
						Metric: "TestMetric",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:dev", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
					{
						Metric: "TestMetric",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeGauge),
						Tags:   []string{"env:dev", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
				},
			},
			expect: func(t *testing.T, result pmetric.Metrics) {
				// The different metrics will fall under the same ResourceMetric and ScopeMetric
				// and there will be separate metrics under the ScopeMetric.Metrics() due to the different
				// data types
				require.Equal(t, 1, result.ResourceMetrics().Len())
				resource := result.ResourceMetrics().At(0)

				v, exists := resource.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())

				require.Equal(t, 1, resource.ScopeMetrics().Len())

				require.Equal(t, 2, resource.ScopeMetrics().At(0).Metrics().Len())

				require.Equal(t, "TestMetric", resource.ScopeMetrics().At(0).Metrics().At(0).Name())
				require.Equal(t, "TestMetric", resource.ScopeMetrics().At(0).Metrics().At(1).Name())

				require.Equal(t, pmetric.MetricTypeSum, resource.ScopeMetrics().At(0).Metrics().At(0).Type())
				require.Equal(t, pmetric.MetricTypeGauge, resource.ScopeMetrics().At(0).Metrics().At(1).Type())
			},
		},
		{
			name: "Same host, same metric name, same tags, diff datapoints",
			series: SeriesList{
				Series: []datadogV1.Series{
					{
						Metric: "TestMetric",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:dev", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629071, 0.5,
							},
						}),
					},
					{
						Metric: "TestMetric",
						Host:   strPtr("Host1"),
						Type:   strPtr(TypeCount),
						Tags:   []string{"env:dev", "version:tag1"},
						Points: testPointsToDatadogPoints([]testPoint{
							{
								1636629081, 1.0, // Different timestamp and value
							},
						}),
					},
				},
			},
			expect: func(t *testing.T, result pmetric.Metrics) {
				// Same host, tags, and metric name but two different datapoints
				// should result in a single resourceMetric, scopeMetric, and metric
				// but two different datapoints under that metric
				require.Equal(t, 1, result.ResourceMetrics().Len())
				resource := result.ResourceMetrics().At(0)

				v, exists := resource.Resource().Attributes().Get("host.name")
				require.True(t, exists)
				require.Equal(t, "Host1", v.AsString())

				require.Equal(t, 1, resource.ScopeMetrics().Len())

				require.Equal(t, 1, resource.ScopeMetrics().At(0).Metrics().Len())

				require.Equal(t, "TestMetric", resource.ScopeMetrics().At(0).Metrics().At(0).Name())

				require.Equal(t, pmetric.MetricTypeSum, resource.ScopeMetrics().At(0).Metrics().At(0).Type())
				require.Equal(t, 2, resource.ScopeMetrics().At(0).Metrics().At(0).Sum().DataPoints().Len())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := newMetricsTranslator()
			mt.buildInfo = component.BuildInfo{
				Command:     "otelcol",
				Description: "OpenTelemetry Collector",
				Version:     "latest",
			}
			result := mt.translateMetricsV1(tt.series)

			tt.expect(t, result)
		})
	}
}
