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

package metricstransformprocessor

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadingFullConfig(t *testing.T) {
	tests := []struct {
		configFile string
		filterName string
		expCfg     *Config
	}{
		{
			configFile: "config_full.yaml",
			filterName: "metricstransform",
			expCfg: &Config{
				ProcessorSettings: configmodels.ProcessorSettings{
					TypeVal: "metricstransform",
					NameVal: "metricstransform",
				},
				Transforms: []Transform{
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "name",
							MatchType: "",
						},
						Action:  "update",
						NewName: "new_name",
					},
				},
			},
		},
		{
			configFile: "config_full.yaml",
			filterName: "metricstransform/multiple",
			expCfg: &Config{
				ProcessorSettings: configmodels.ProcessorSettings{
					TypeVal: "metricstransform",
					NameVal: "metricstransform/multiple",
				},
				Transforms: []Transform{
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "name1",
							MatchType: "strict",
						},
						Action:  "insert",
						NewName: "new_name",
						Operations: []Operation{
							{
								Action:   "add_label",
								NewLabel: "my_label",
								NewValue: "my_value",
							},
						},
					},
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "name2",
							MatchType: "",
						},
						Action: "update",
						Operations: []Operation{
							{
								Action:   "update_label",
								Label:    "label",
								NewLabel: "new_label_key",
								ValueActions: []ValueAction{
									{Value: "label1", NewValue: "new_label1"},
								},
							},
							{
								Action:          "aggregate_labels",
								LabelSet:        []string{"new_label1", "label2"},
								AggregationType: "sum",
							},
							{
								Action:           "aggregate_label_values",
								Label:            "new_label1",
								AggregationType:  "sum",
								AggregatedValues: []string{"value1", "value2"},
								NewValue:         "new_value",
							},
						},
					},
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "name3",
							MatchType: "strict",
						},
						Action: "update",
						Operations: []Operation{
							{
								Action:     "delete_label_value",
								Label:      "my_label",
								LabelValue: "delete_me",
							},
						},
					},
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "^regexp (?P<my_label>.*)$",
							MatchType: "regexp",
						},
						Action:       "combine",
						NewName:      "combined_metric_name",
						SubmatchCase: "lower",
					},
					{
						MetricIncludeFilter: FilterConfig{
							Include:   "name2",
							MatchType: "strict",
						},
						Action:              "group",
						GroupResourceLabels: map[string]string{"metric_group": "2"},
					},
				},
			},
		},
		{
			configFile: "config_deprecated.yaml",
			filterName: "metricstransform",
			expCfg: &Config{
				ProcessorSettings: configmodels.ProcessorSettings{
					NameVal: "metricstransform",
					TypeVal: typeStr,
				},
				Transforms: []Transform{
					{
						MetricName: "old_name",
						Action:     Update,
						NewName:    "new_name",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.filterName, func(t *testing.T) {

			factories, err := componenttest.NopFactories()
			assert.NoError(t, err)

			factory := NewFactory()
			factories.Processors[configmodels.Type(typeStr)] = factory
			config, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", test.configFile), factories)
			assert.NoError(t, err)
			require.NotNil(t, config)

			cfg := config.Processors[test.filterName]
			assert.Equal(t, test.expCfg, cfg)
		})
	}
}
