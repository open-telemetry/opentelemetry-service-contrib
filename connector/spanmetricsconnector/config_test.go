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

package spanmetricsconnector

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/otelcol/otelcoltest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor/batchprocessor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
)

func TestLoadConfig(t *testing.T) {
	defaultMethod := "GET"
	testcases := []struct {
		configFile                  string
		wantLatencyHistogramBuckets []time.Duration
		wantDimensions              []Dimension
		wantDimensionsCacheSize     int
		wantAggregationTemporality  string
		wantMetricsFlushInterval    time.Duration
	}{
		{
			configFile:                 "config-2-pipelines.yaml",
			wantAggregationTemporality: cumulative,
			wantDimensionsCacheSize:    500,
			wantMetricsFlushInterval:   15 * time.Second, // Default.
		},
		{
			configFile: "config-full.yaml",
			wantLatencyHistogramBuckets: []time.Duration{
				100 * time.Microsecond,
				1 * time.Millisecond,
				2 * time.Millisecond,
				6 * time.Millisecond,
				10 * time.Millisecond,
				100 * time.Millisecond,
				250 * time.Millisecond,
			},
			wantDimensions: []Dimension{
				{"http.method", &defaultMethod},
				{"http.status_code", nil},
			},
			wantDimensionsCacheSize:    1500,
			wantAggregationTemporality: delta,
			wantMetricsFlushInterval:   30 * time.Second,
		},
	}

	require.NoError(t, featuregate.GlobalRegistry().Set("enableConnectors", true))
	defer func() {
		require.NoError(t, featuregate.GlobalRegistry().Set("enableConnectors", false))
	}()
	for _, tc := range testcases {
		t.Run(tc.configFile, func(t *testing.T) {
			// Prepare
			factories, err := otelcoltest.NopFactories()
			require.NoError(t, err)

			factories.Receivers["jaeger"] = jaegerreceiver.NewFactory()

			factories.Processors["batch"] = batchprocessor.NewFactory()

			factories.Exporters["prometheus"] = prometheusexporter.NewFactory()
			factories.Exporters["jaeger"] = jaegerexporter.NewFactory()

			factories.Connectors[typeStr] = NewFactory()

			// Test
			cfg, err := otelcoltest.LoadConfigAndValidate(filepath.Join("testdata", tc.configFile), factories)

			// Verify
			require.NoError(t, err)
			require.NotNil(t, cfg)
			assert.Equal(t,
				&Config{
					LatencyHistogramBuckets: tc.wantLatencyHistogramBuckets,
					Dimensions:              tc.wantDimensions,
					DimensionsCacheSize:     tc.wantDimensionsCacheSize,
					AggregationTemporality:  tc.wantAggregationTemporality,
					MetricsFlushInterval:    tc.wantMetricsFlushInterval,
				},
				cfg.Connectors[component.NewID(typeStr)],
			)
		})
	}
}

func TestGetAggregationTemporality(t *testing.T) {
	cfg := &Config{AggregationTemporality: delta}
	assert.Equal(t, pmetric.AggregationTemporalityDelta, cfg.GetAggregationTemporality())

	cfg = &Config{AggregationTemporality: cumulative}
	assert.Equal(t, pmetric.AggregationTemporalityCumulative, cfg.GetAggregationTemporality())

	cfg = &Config{}
	assert.Equal(t, pmetric.AggregationTemporalityCumulative, cfg.GetAggregationTemporality())
}
