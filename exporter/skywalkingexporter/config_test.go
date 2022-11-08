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

package skywalkingexporter

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	defaultCfg := createDefaultConfig().(*Config)
	defaultCfg.Endpoint = "1.2.3.4:11800"

	tests := []struct {
		id       component.ID
		expected component.ExporterConfig
	}{
		{
			id:       component.NewIDWithName(typeStr, ""),
			expected: defaultCfg,
		},
		{
			id: component.NewIDWithName(typeStr, "2"),
			expected: &Config{
				ExporterSettings: config.NewExporterSettings(component.NewID(typeStr)),
				RetrySettings: exporterhelper.RetrySettings{
					Enabled:         true,
					InitialInterval: 10 * time.Second,
					MaxInterval:     1 * time.Minute,
					MaxElapsedTime:  10 * time.Minute,
				},
				QueueSettings: exporterhelper.QueueSettings{
					Enabled:      true,
					NumConsumers: 2,
					QueueSize:    10,
				},
				TimeoutSettings: exporterhelper.TimeoutSettings{
					Timeout: 10 * time.Second,
				},
				GRPCClientSettings: configgrpc.GRPCClientSettings{
					Headers: map[string]string{
						"can you have a . here?": "F0000000-0000-0000-0000-000000000000",
						"header1":                "234",
						"another":                "somevalue",
					},
					Endpoint:    "1.2.3.4:11800",
					Compression: "gzip",
					TLSSetting: configtls.TLSClientSetting{
						TLSSetting: configtls.TLSSetting{
							CAFile: "/var/lib/mycert.pem",
						},
						Insecure: false,
					},
					Keepalive: &configgrpc.KeepaliveClientConfig{
						Time:                20,
						PermitWithoutStream: true,
						Timeout:             30,
					},
					WriteBufferSize: 512 * 1024,
					BalancerName:    "round_robin",
				},
				NumStreams: 233,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalExporterConfig(sub, cfg))

			assert.NoError(t, cfg.Validate())
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestValidate(t *testing.T) {
	c1 := &Config{
		ExporterSettings: config.NewExporterSettings(component.NewID(typeStr)),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Endpoint: "",
		},
		NumStreams: 3,
	}
	err := c1.Validate()
	assert.Error(t, err)
	c2 := &Config{
		ExporterSettings: config.NewExporterSettings(component.NewID(typeStr)),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Endpoint: "",
		},
		NumStreams: 0,
	}
	err2 := c2.Validate()
	assert.Error(t, err2)
}
