// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

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
				AwsLogGroupNames:  ResourceAttributeConfig{Enabled: true},
				AwsLogStreamNames: ResourceAttributeConfig{Enabled: true},
				CloudPlatform:     ResourceAttributeConfig{Enabled: true},
				CloudProvider:     ResourceAttributeConfig{Enabled: true},
				CloudRegion:       ResourceAttributeConfig{Enabled: true},
				FaasInstance:      ResourceAttributeConfig{Enabled: true},
				FaasMaxMemory:     ResourceAttributeConfig{Enabled: true},
				FaasName:          ResourceAttributeConfig{Enabled: true},
				FaasVersion:       ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				AwsLogGroupNames:  ResourceAttributeConfig{Enabled: false},
				AwsLogStreamNames: ResourceAttributeConfig{Enabled: false},
				CloudPlatform:     ResourceAttributeConfig{Enabled: false},
				CloudProvider:     ResourceAttributeConfig{Enabled: false},
				CloudRegion:       ResourceAttributeConfig{Enabled: false},
				FaasInstance:      ResourceAttributeConfig{Enabled: false},
				FaasMaxMemory:     ResourceAttributeConfig{Enabled: false},
				FaasName:          ResourceAttributeConfig{Enabled: false},
				FaasVersion:       ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
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
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
