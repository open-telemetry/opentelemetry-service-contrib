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

package mesosreceiver

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc        string
		endpoint    string
		errExpected bool
		errText     string
	}{
		{
			desc:        "default_endpoint",
			endpoint:    "http://remote-docker-ageorgiades:5050/metrics/snapshot",
			errExpected: false,
		},
		{
			desc:        "custom_host",
			endpoint:    "http://133.713:371:337:5050/metrics/snapshot",
			errExpected: false,
		},
		{
			desc:        "custom_port",
			endpoint:    "http://remote-docker-ageorgiades:7942/metrics/snapshot",
			errExpected: false,
		},
		{
			desc:        "custom_path",
			endpoint:    "http://remote-docker-ageorgiades:5050/hey/thatsprettygood",
			errExpected: false,
		},
		{
			desc:        "empty_path",
			endpoint:    "",
			errExpected: true,
			errText:     "missing hostname: ''",
		},
		{
			desc:        "missing_hostname",
			endpoint:    "http://:5050/metrics/snapshot",
			errExpected: true,
			errText:     "missing hostname: 'http://:5050/metrics/snapshot'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cfg := NewFactory().CreateDefaultConfig().(*Config)
			cfg.Endpoint = tc.endpoint
			err := component.ValidateConfig(cfg)
			if tc.errExpected {
				require.EqualError(t, err, tc.errText)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("tmp", "config.yaml"))
	require.NoError(t, err)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub(component.NewIDWithName(typeStr, "").String())
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))

	expected := factory.CreateDefaultConfig().(*Config)
	expected.Endpoint = "http://remote-docker-ageorgiades:5050/metrics/snapshot"
	expected.CollectionInterval = 10 * time.Second

	require.Equal(t, expected, cfg)
}
