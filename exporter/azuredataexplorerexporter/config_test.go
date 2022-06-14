// Copyright OpenTelemetry Authors
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

package azuredataexplorerexporter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/service/servicetest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(t, err)

	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, _ := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)

	require.NotNil(t, cfg)

	// There are 3 stanzas in the exporter.
	assert.Equal(t, len(cfg.Exporters), 3)

	exporter := cfg.Exporters[config.NewComponentID(typeStr)]

	// Is a valid configuration that has no errors
	assert.NoError(t, configtest.CheckConfigStruct(exporter))
	assert.Equal(
		t,
		&Config{
			ClusterName:    "https://CLUSTER.kusto.windows.net",
			ClientId:       "f80da32c-108c-415c-a19e-643f461a677a",
			ClientSecret:   "17cc3f47-e95e-4045-af6c-ec2eea163cc6",
			TenantId:       "21ff9e36-fbaa-43c8-98ba-00431ea10bc3",
			Database:       "oteldb",
			RawMetricTable: "RawMetrics",
			IngestionType:  managedingesttype,
		},
		exporter)

	// The second one has a validation error
	exporter = cfg.Exporters[config.NewComponentIDWithName(typeStr, "2")].(*Config)
	err = cfg.Validate()
	assert.EqualError(t, err, `exporter "azuredataexplorer/2" has invalid configuration: mandatory configurations "cluster_name" ,"client_id" , "client_secret" and "tenant_id" are missing or empty `)
}
