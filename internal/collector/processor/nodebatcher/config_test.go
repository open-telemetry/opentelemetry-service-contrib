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

package nodebatcher

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-service/configv2"
	"github.com/open-telemetry/opentelemetry-service/configv2/configmodels"
	"github.com/open-telemetry/opentelemetry-service/processor"
)

var _ = configv2.RegisterTestFactories()

func TestLoadConfig(t *testing.T) {

	factory := processor.GetFactory(typeStr)

	config, err := configv2.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"))

	require.Nil(t, err)
	require.NotNil(t, config)

	p0 := config.Processors["batch"]
	assert.Equal(t, p0, factory.CreateDefaultConfig())

	p1 := config.Processors["batch/2"]

	timeout := time.Second * 10
	tickTime := time.Second * 5
	removeAfterTicks := 20
	sendBatchSize := 1000

	assert.Equal(t, p1,
		&ConfigV2{
			ProcessorSettings: configmodels.ProcessorSettings{
				TypeVal: "batch",
				NameVal: "batch/2",
			},
			Timeout:          &timeout,
			NumTickers:       10,
			RemoveAfterTicks: &removeAfterTicks,
			SendBatchSize:    &sendBatchSize,
			TickTime:         &tickTime,
		})
}
