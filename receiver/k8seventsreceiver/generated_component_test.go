// Code generated by mdatagen. DO NOT EDIT.

package k8seventsreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

var typ = component.MustNewType("k8s_events")

func TestComponentFactoryType(t *testing.T) {
	require.Equal(t, typ, NewFactory().Type())
}

func TestComponentConfigStruct(t *testing.T) {
	require.NoError(t, componenttest.CheckConfigStruct(NewFactory().CreateDefaultConfig()))
}

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	tests := []struct {
		createFn func(ctx context.Context, set receiver.Settings, cfg component.Config) (component.Component, error)
		name     string
	}{

		{
			name: "logs",
			createFn: func(ctx context.Context, set receiver.Settings, cfg component.Config) (component.Component, error) {
				return factory.CreateLogs(ctx, set, cfg, consumertest.NewNop())
			},
		},
	}

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, sub.Unmarshal(&cfg))

	for _, tt := range tests {
		t.Run(tt.name+"-shutdown", func(t *testing.T) {
			c, err := tt.createFn(context.Background(), receivertest.NewNopSettings(), cfg)
			require.NoError(t, err)
			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})
	}
}
