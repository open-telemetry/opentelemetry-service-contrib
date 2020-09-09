// Copyright The OpenTelemetry Authors
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

package httpforwarder

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/config/configtls"
	"go.uber.org/zap"
)

func TestFactory(t *testing.T) {
	f := NewFactory()
	expectType := "http_forwarder"
	require.Equal(t, configmodels.Type(expectType), f.Type())

	cfg := f.CreateDefaultConfig().(*Config)
	require.Equal(t, expectType, cfg.Name())
	require.Equal(t, configmodels.Type(expectType), cfg.Type())
	require.Equal(t, ":6060", cfg.Endpoint)
	require.Equal(t, 10*time.Second, cfg.Upstream.Timeout)

	tests := []struct {
		name           string
		config         *Config
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:           "Default config",
			config:         cfg,
			wantErr:        true,
			wantErrMessage: "'upstream.endpoint' config option cannot be empty",
		},
		{
			name:    "Invalid config",
			config:  &Config{Upstream: confighttp.HTTPClientSettings{Endpoint: "123.456.7.89:9090"}},
			wantErr: true,
		},
		{
			name: "Invalid config - HTTP Client creation fails",
			config: &Config{
				Upstream: confighttp.HTTPClientSettings{
					Endpoint: "localhost:9090",
					TLSSetting: configtls.TLSClientSetting{
						TLSSetting: configtls.TLSSetting{
							CAFile: "/non/existent",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "Valid config",
			config: &Config{Upstream: confighttp.HTTPClientSettings{Endpoint: "localhost:9090"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e, err := f.CreateExtension(
				context.Background(),
				component.ExtensionCreateParams{
					Logger: zap.NewNop(),
				},
				test.config,
			)
			if test.wantErr {
				if test.wantErrMessage != "" {
					require.EqualError(t, err, test.wantErrMessage)
				}
				require.Error(t, err)
				require.Nil(t, e)
			} else {
				require.NoError(t, err)
				require.NotNil(t, e)

				ctx := context.Background()
				require.NoError(t, e.Start(ctx, componenttest.NewNopHost()))
				require.NoError(t, e.Shutdown(ctx))
			}
		})
	}
}
