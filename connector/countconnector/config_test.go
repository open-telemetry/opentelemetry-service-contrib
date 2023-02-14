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

package countconnector

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	testCases := []struct {
		name   string
		expect *Config
	}{
		{
			name: "",
			expect: &Config{
				Spans: map[string]MetricInfo{
					defaultMetricNameSpans: {
						Description: defaultMetricDescSpans,
					},
				},
				SpanEvents: map[string]MetricInfo{
					defaultMetricNameSpanEvents: {
						Description: defaultMetricDescSpanEvents,
					},
				},
				Metrics: map[string]MetricInfo{
					defaultMetricNameMetrics: {
						Description: defaultMetricDescMetrics,
					},
				},
				DataPoints: map[string]MetricInfo{
					defaultMetricNameDataPoints: {
						Description: defaultMetricDescDataPoints,
					},
				},
				Logs: map[string]MetricInfo{
					defaultMetricNameLogRecords: {
						Description: defaultMetricDescLogRecords,
					},
				},
			},
		},
		{
			name: "custom_description",
			expect: &Config{
				Spans: map[string]MetricInfo{
					defaultMetricNameSpans: {
						Description: "My description for default span count metric.",
					},
				},
				SpanEvents: map[string]MetricInfo{
					defaultMetricNameSpanEvents: {
						Description: "My description for default span event count metric.",
					},
				},
				Metrics: map[string]MetricInfo{
					defaultMetricNameMetrics: {
						Description: "My description for default metric count metric.",
					},
				},
				DataPoints: map[string]MetricInfo{
					defaultMetricNameDataPoints: {
						Description: "My description for default datapoint count metric.",
					},
				},
				Logs: map[string]MetricInfo{
					defaultMetricNameLogRecords: {
						Description: "My description for default log count metric.",
					},
				},
			},
		},
		{
			name: "custom_metric",
			expect: &Config{
				Spans: map[string]MetricInfo{
					"my.span.count": {
						Description: "My span count.",
					},
				},
				SpanEvents: map[string]MetricInfo{
					"my.spanevent.count": {
						Description: "My span event count.",
					},
				},
				Metrics: map[string]MetricInfo{
					"my.metric.count": {
						Description: "My metric count.",
					},
				},
				DataPoints: map[string]MetricInfo{
					"my.datapoint.count": {
						Description: "My data point count.",
					},
				},
				Logs: map[string]MetricInfo{
					"my.logrecord.count": {
						Description: "My log record count.",
					},
				},
			},
		},
		{
			name: "condition",
			expect: &Config{
				Spans: map[string]MetricInfo{
					"my.span.count": {
						Description: "My span count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-s") == true`},
					},
				},
				SpanEvents: map[string]MetricInfo{
					"my.spanevent.count": {
						Description: "My span event count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-e") == true`},
					},
				},
				Metrics: map[string]MetricInfo{
					"my.metric.count": {
						Description: "My metric count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-m") == true`},
					},
				},
				DataPoints: map[string]MetricInfo{
					"my.datapoint.count": {
						Description: "My data point count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-d") == true`},
					},
				},
				Logs: map[string]MetricInfo{
					"my.logrecord.count": {
						Description: "My log record count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-l") == true`},
					},
				},
			},
		},
		{
			name: "multiple_condition",
			expect: &Config{
				Spans: map[string]MetricInfo{
					"my.span.count": {
						Description: "My span count.",
						Conditions: []string{
							`IsMatch(resource.attributes["host.name"], "pod-s") == true`,
							`IsMatch(resource.attributes["foo"], "bar-s") == true`,
						},
					},
				},
				SpanEvents: map[string]MetricInfo{
					"my.spanevent.count": {
						Description: "My span event count.",
						Conditions: []string{
							`IsMatch(resource.attributes["host.name"], "pod-e") == true`,
							`IsMatch(resource.attributes["foo"], "bar-e") == true`,
						},
					},
				},
				Metrics: map[string]MetricInfo{
					"my.metric.count": {
						Description: "My metric count.",
						Conditions: []string{
							`IsMatch(resource.attributes["host.name"], "pod-m") == true`,
							`IsMatch(resource.attributes["foo"], "bar-m") == true`,
						},
					},
				},
				DataPoints: map[string]MetricInfo{
					"my.datapoint.count": {
						Description: "My data point count.",
						Conditions: []string{
							`IsMatch(resource.attributes["host.name"], "pod-d") == true`,
							`IsMatch(resource.attributes["foo"], "bar-d") == true`,
						},
					},
				},
				Logs: map[string]MetricInfo{
					"my.logrecord.count": {
						Description: "My log record count.",
						Conditions: []string{
							`IsMatch(resource.attributes["host.name"], "pod-l") == true`,
							`IsMatch(resource.attributes["foo"], "bar-l") == true`,
						},
					},
				},
			},
		},
		{
			name: "multiple_metrics",
			expect: &Config{
				Spans: map[string]MetricInfo{
					"my.span.count": {
						Description: "My span count."},
					"limited.span.count": {
						Description: "Limited span count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-s") == true`},
					},
				},
				SpanEvents: map[string]MetricInfo{
					"my.spanevent.count": {
						Description: "My span event count."},
					"limited.spanevent.count": {
						Description: "Limited span event count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-e") == true`},
					},
				},
				Metrics: map[string]MetricInfo{
					"my.metric.count": {
						Description: "My metric count."},
					"limited.metric.count": {
						Description: "Limited metric count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-m") == true`},
					},
				},
				DataPoints: map[string]MetricInfo{
					"my.datapoint.count": {
						Description: "My data point count.",
					},
					"limited.datapoint.count": {
						Description: "Limited data point count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-d") == true`},
					},
				},
				Logs: map[string]MetricInfo{
					"my.logrecord.count": {
						Description: "My log record count.",
					},
					"limited.logrecord.count": {
						Description: "Limited log record count.",
						Conditions:  []string{`IsMatch(resource.attributes["host.name"], "pod-l") == true`},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
			require.NoError(t, err)

			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(component.NewIDWithName(typeStr, tc.name).String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			assert.Equal(t, tc.expect, cfg)
		})
	}
}
