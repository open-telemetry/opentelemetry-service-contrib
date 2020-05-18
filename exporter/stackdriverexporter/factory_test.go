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

package stackdriverexporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config/configcheck"
	"go.uber.org/zap"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configcheck.ValidateConfig(cfg))
}

func TestCreateExporter(t *testing.T) {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		t.Skip("Default credentials not set, skip creating Stackdriver exporter")
	}
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	eCfg := cfg.(*Config)
	eCfg.ProjectID = "test"

	te, err := factory.CreateTraceExporter(zap.NewNop(), eCfg)
	assert.Nil(t, err)
	assert.NotNil(t, te, "failed to create trace exporter")

	me, err := factory.CreateMetricsExporter(zap.NewNop(), eCfg)
	assert.Nil(t, err)
	assert.NotNil(t, me, "failed to create metrics exporter")
}
