// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

import (
	"testing"
	// client is an autogenerated mock type for the client type
	"github.com/stretchr/testify/require"
)

func TestGetMetricScalarOIDs(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no scalar OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: "1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricScalarOIDs()
				require.ElementsMatch(t, []string{}, actual)
			},
		},
		{
			desc: "Returns all scalar OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
								{
									OID: ".2",
								},
							},
						},
						"m2": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".3",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricScalarOIDs()
				require.ElementsMatch(t, []string{".1", ".2", ".3"}, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetMetricColumnOIDs(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no column OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricColumnOIDs()
				require.ElementsMatch(t, []string{}, actual)
			},
		},
		{
			desc: "Returns all column OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
								},
								{
									OID: ".2",
								},
							},
						},
						"m2": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".3",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricColumnOIDs()
				require.ElementsMatch(t, []string{".1", ".2", ".3"}, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetAttributeColumnOIDs(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no attributes",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeColumnOIDs()
				require.ElementsMatch(t, []string{}, actual)
			},
		},
		{
			desc: "Returns all column OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
						"a2": {
							OID: ".3",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeColumnOIDs()
				require.ElementsMatch(t, []string{".2", ".3"}, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetResourceAttributeColumnOIDs(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no resource attributes",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeColumnOIDs()
				require.ElementsMatch(t, []string{}, actual)
			},
		},
		{
			desc: "Returns all column OIDs",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							OID: ".3",
						},
						"ra2": {
							OID: ".4",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeColumnOIDs()
				require.ElementsMatch(t, []string{".3", ".4"}, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetMetricName(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no metric matches OID",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricName(".2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns metric name it exists by matching scalar OID",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricName(".1")
				require.Equal(t, "m1", actual)
			},
		},
		{
			desc: "Returns metric name it exists by matching column OID",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricName(".1")
				require.Equal(t, "m1", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetMetricConfig(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns nil when no metric config exists",
			testFunc: func(t *testing.T) {
				var m1 *MetricConfig
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricConfig("m2")
				require.Equal(t, m1, actual)
			},
		},
		{
			desc: "Returns metric config",
			testFunc: func(t *testing.T) {
				m1 := &MetricConfig{
					ScalarOIDs: []ScalarOID{
						{
							OID: ".1",
						},
					},
				}
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": m1,
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricConfig("m1")
				require.Equal(t, m1, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetAttributeConfigValue(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no attribute config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							Value: "a1v",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigValue("a2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns empty string when attribute config does not have value",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigValue("a1")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns value for attribute config",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							Value: "a1v",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigValue("a1")
				require.Equal(t, "a1v", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetAttributeConfigIndexedValuePrefix(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no attribute config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigIndexedValuePrefix("a2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns empty string when attribute config does not have indexed value prefix",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigIndexedValuePrefix("a1")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns indexed value prefix for attribute config",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigIndexedValuePrefix("a1")
				require.Equal(t, "prefix", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetAttributeConfigOID(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no attribute config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigOID("a2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns empty string when attribute config does not have OID",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigOID("a1")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns indexed value prefix for attribute config",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getAttributeConfigOID("a1")
				require.Equal(t, ".2", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetResourceAttributeConfigIndexedValuePrefix(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no resource attribute config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigIndexedValuePrefix("ra2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns empty string when resource attribute config does not have indexed value prefix",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigIndexedValuePrefix("ra1")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns indexed value prefix for resource attribute config",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigIndexedValuePrefix("ra1")
				require.Equal(t, "prefix", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetResourceAttributeConfigOID(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty string when no resource attribute config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigOID("ra2")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns empty string when resource attribute config does not have OID",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigOID("ra1")
				require.Equal(t, "", actual)
			},
		},
		{
			desc: "Returns indexed value prefix for resource attribute config",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							OID: ".2",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeConfigOID("ra1")
				require.Equal(t, ".2", actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetMetricConfigAttributes(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no metric config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
									Attributes: []Attribute{
										{
											Name: "a1",
										},
									},
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricConfigAttributes(".2")
				require.ElementsMatch(t, []Attribute{}, actual)
			},
		},
		{
			desc: "Returns empty slice when metric config has no attributes",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID: ".1",
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricConfigAttributes(".1")
				require.ElementsMatch(t, []Attribute{}, actual)
			},
		},
		{
			desc: "Returns metric config attributes",
			testFunc: func(t *testing.T) {
				attributes := []Attribute{
					{
						Name:  "a2",
						Value: "v2",
					},
					{
						Name:  "a3",
						Value: "v3",
					},
				}
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ScalarOIDs: []ScalarOID{
								{
									OID:        ".1",
									Attributes: attributes,
								},
							},
						},
					},
					Attributes: map[string]*AttributeConfig{
						"a1": {
							IndexedValuePrefix: "prefix",
						},
						"a2": {
							Enum: []string{"v2"},
						},
						"a3": {
							Enum: []string{"v3"},
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getMetricConfigAttributes(".1")
				require.ElementsMatch(t, attributes, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}

func TestGetResourceAttributeNames(t *testing.T) {
	testCases := []struct {
		desc     string
		testFunc func(*testing.T)
	}{
		{
			desc: "Returns empty slice when no metric config exists",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
									ResourceAttributes: []string{
										"ra1",
									},
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeNames(".2")
				require.ElementsMatch(t, []Attribute{}, actual)
			},
		},
		{
			desc: "Returns empty slice when metric config has no resource attributes",
			testFunc: func(t *testing.T) {
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID: ".1",
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeNames(".1")
				require.ElementsMatch(t, []Attribute{}, actual)
			},
		},
		{
			desc: "Returns metric config resource attributes",
			testFunc: func(t *testing.T) {
				attributes := []string{"ra2", "ra3"}
				cfg := Config{
					Metrics: map[string]*MetricConfig{
						"m1": {
							ColumnOIDs: []ColumnOID{
								{
									OID:                ".1",
									ResourceAttributes: attributes,
								},
							},
						},
					},
					ResourceAttributes: map[string]*ResourceAttributeConfig{
						"ra1": {
							IndexedValuePrefix: "prefix",
						},
						"ra2": {
							OID: ".2",
						},
						"ra3": {
							OID: ".3",
						},
					},
				}
				helper := newConfigHelper(&cfg)
				actual := helper.getResourceAttributeNames(".1")
				require.ElementsMatch(t, attributes, actual)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, tc.testFunc)
	}
}
