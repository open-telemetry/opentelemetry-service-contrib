// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logzioexporter

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadConfig(tester *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(tester, err)

	factory := NewFactory()
	factories.Exporters[config.Type(typeStr)] = factory
	cfg, err := configtest.LoadConfigFile(
		tester, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(tester, err)
	require.NotNil(tester, cfg)

	assert.Equal(tester, 2, len(cfg.Exporters))

	config := cfg.Exporters["logzio/2"].(*Config)
	assert.Equal(tester, &Config{
		ExporterSettings: config.ExporterSettings{TypeVal: typeStr, NameVal: "logzio/2"},
		TracesToken:      "logzioTESTtoken",
		Region:           "eu",
		CustomEndpoint:   "https://some-url.com:8888",
	}, config)
}
