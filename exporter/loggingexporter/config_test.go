// Copyright 2019, OpenTelemetry Authors
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

package loggingexporter

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-service/config"
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
	"github.com/open-telemetry/opentelemetry-service/exporter"
)

var _ = config.RegisterTestFactories()

func TestLoadConfig(t *testing.T) {
	factory := exporter.GetFactory(typeStr)

	cfg, err := config.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"))

	require.NoError(t, err)
	require.NotNil(t, cfg)

	e0 := cfg.Exporters["logging"]
	assert.Equal(t, e0, factory.CreateDefaultConfig())

	e1 := cfg.Exporters["logging/2"]
	assert.Equal(t, e1,
		&Config{
			ExporterSettings: configmodels.ExporterSettings{
				NameVal: "logging/2",
				TypeVal: "logging",
				Enabled: true,
			},
		})
}
