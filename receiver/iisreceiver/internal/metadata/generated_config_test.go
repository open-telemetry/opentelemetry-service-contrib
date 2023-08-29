// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					IisConnectionActive:       MetricConfig{Enabled: true},
					IisConnectionAnonymous:    MetricConfig{Enabled: true},
					IisConnectionAttemptCount: MetricConfig{Enabled: true},
					IisNetworkBlocked:         MetricConfig{Enabled: true},
					IisNetworkFileCount:       MetricConfig{Enabled: true},
					IisNetworkIo:              MetricConfig{Enabled: true},
					IisRequestCount:           MetricConfig{Enabled: true},
					IisRequestQueueAgeMax:     MetricConfig{Enabled: true},
					IisRequestQueueCount:      MetricConfig{Enabled: true},
					IisRequestRejected:        MetricConfig{Enabled: true},
					IisThreadActive:           MetricConfig{Enabled: true},
					IisUptime:                 MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					IisApplicationPool: ResourceAttributeConfig{Enabled: true},
					IisSite:            ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					IisConnectionActive:       MetricConfig{Enabled: false},
					IisConnectionAnonymous:    MetricConfig{Enabled: false},
					IisConnectionAttemptCount: MetricConfig{Enabled: false},
					IisNetworkBlocked:         MetricConfig{Enabled: false},
					IisNetworkFileCount:       MetricConfig{Enabled: false},
					IisNetworkIo:              MetricConfig{Enabled: false},
					IisRequestCount:           MetricConfig{Enabled: false},
					IisRequestQueueAgeMax:     MetricConfig{Enabled: false},
					IisRequestQueueCount:      MetricConfig{Enabled: false},
					IisRequestRejected:        MetricConfig{Enabled: false},
					IisThreadActive:           MetricConfig{Enabled: false},
					IisUptime:                 MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					IisApplicationPool: ResourceAttributeConfig{Enabled: false},
					IisSite:            ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				IisApplicationPool: ResourceAttributeConfig{Enabled: true},
				IisSite:            ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				IisApplicationPool: ResourceAttributeConfig{Enabled: false},
				IisSite:            ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
