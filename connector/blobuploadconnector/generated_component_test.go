// Code generated by mdatagen. DO NOT EDIT.

package blobuploadconnector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/connectortest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pipeline"
)

func TestComponentFactoryType(t *testing.T) {
	require.Equal(t, "blobuploadconnector", NewFactory().Type().String())
}

func TestComponentConfigStruct(t *testing.T) {
	require.NoError(t, componenttest.CheckConfigStruct(NewFactory().CreateDefaultConfig()))
}

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	tests := []struct {
		name     string
		createFn func(ctx context.Context, set connector.Settings, cfg component.Config) (component.Component, error)
	}{

		{
			name: "logs_to_logs",
			createFn: func(ctx context.Context, set connector.Settings, cfg component.Config) (component.Component, error) {
				router := connector.NewLogsRouter(map[pipeline.ID]consumer.Logs{pipeline.NewID(pipeline.SignalLogs): consumertest.NewNop()})
				return factory.CreateLogsToLogs(ctx, set, cfg, router)
			},
		},

		{
			name: "traces_to_traces",
			createFn: func(ctx context.Context, set connector.Settings, cfg component.Config) (component.Component, error) {
				router := connector.NewTracesRouter(map[pipeline.ID]consumer.Traces{pipeline.NewID(pipeline.SignalTraces): consumertest.NewNop()})
				return factory.CreateTracesToTraces(ctx, set, cfg, router)
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
			c, err := tt.createFn(context.Background(), connectortest.NewNopSettings(), cfg)
			require.NoError(t, err)
			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})
		t.Run(tt.name+"-lifecycle", func(t *testing.T) {
			firstConnector, err := tt.createFn(context.Background(), connectortest.NewNopSettings(), cfg)
			require.NoError(t, err)
			host := componenttest.NewNopHost()
			require.NoError(t, err)
			require.NoError(t, firstConnector.Start(context.Background(), host))
			require.NoError(t, firstConnector.Shutdown(context.Background()))
			secondConnector, err := tt.createFn(context.Background(), connectortest.NewNopSettings(), cfg)
			require.NoError(t, err)
			require.NoError(t, secondConnector.Start(context.Background(), host))
			require.NoError(t, secondConnector.Shutdown(context.Background()))
		})
	}
}
