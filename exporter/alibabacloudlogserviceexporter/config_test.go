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

package alibabacloudlogserviceexporter

import (
	"context"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
	"go.uber.org/zap"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(t, err)

	factory := NewFactory()
	factories.Exporters[config.Type(typeStr)] = factory
	cfg, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	e0 := cfg.Exporters[typeStr]

	// Endpoint doesn't have a default value so set it directly.
	defaultCfg := factory.CreateDefaultConfig().(*Config)
	defaultCfg.Endpoint = "cn-hangzhou.log.aliyuncs.com"
	assert.Equal(t, defaultCfg, e0)

	expectedName := typeStr + "/2"

	e1 := cfg.Exporters[expectedName]
	expectedCfg := Config{
		ExporterSettings: &config.ExporterSettings{
			TypeVal: config.Type(typeStr),
			NameVal: expectedName,
		},
		Endpoint:        "cn-hangzhou.log.aliyuncs.com",
		Project:         "demo-project",
		Logstore:        "demo-logstore",
		AccessKeyID:     "test-id",
		AccessKeySecret: "test-secret",
	}
	assert.Equal(t, &expectedCfg, e1)

	params := component.ExporterCreateParams{Logger: zap.NewNop()}

	// missing params
	te, err := factory.CreateTracesExporter(context.Background(), params, e0)
	require.Error(t, err)
	require.Nil(t, te)
	me, err := factory.CreateMetricsExporter(context.Background(), params, e0)
	require.Error(t, err)
	require.Nil(t, me)
	le, err := factory.CreateLogsExporter(context.Background(), params, e0)
	require.Error(t, err)
	require.Nil(t, le)

	te, err = factory.CreateTracesExporter(context.Background(), params, e1)
	require.NoError(t, err)
	require.NotNil(t, te)
	me, err = factory.CreateMetricsExporter(context.Background(), params, e1)
	require.NoError(t, err)
	require.NotNil(t, me)
	le, err = factory.CreateLogsExporter(context.Background(), params, e1)
	require.NoError(t, err)
	require.NotNil(t, le)
}
