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

package attributesprocessor

import (
	"testing"

	tracepb "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/config/configerror"
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
	"github.com/open-telemetry/opentelemetry-service/exporter/exportertest"
)

func TestFactory_Type(t *testing.T) {
	factory := Factory{}
	assert.Equal(t, factory.Type(), typeStr)
}

func TestFactory_CreateDefaultConfig(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	assert.Equal(t, cfg, &Config{
		ProcessorSettings: configmodels.ProcessorSettings{
			NameVal: typeStr,
			TypeVal: typeStr,
		},
	})
}

func TestFactory_CreateTraceProcessor(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	oCfg := cfg.(*Config)
	oCfg.Actions = []ActionKeyValue{
		{Key: "a key", Action: DELETE},
	}

	tp, err := factory.CreateTraceProcessor(zap.NewNop(), exportertest.NewNopTraceExporter(), cfg)
	assert.NotNil(t, tp)
	assert.Nil(t, err)
}

func TestFactory_CreateMetricsProcessor(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()

	mp, err := factory.CreateMetricsProcessor(zap.NewNop(), nil, cfg)
	require.Nil(t, mp)
	assert.Equal(t, err, configerror.ErrDataTypeIsNotSupported)
}

func TestFactory_attributeValue(t *testing.T) {
	val, err := attributeValue(123)
	assert.Equal(t, &tracepb.AttributeValue{
		Value: &tracepb.AttributeValue_IntValue{IntValue: cast.ToInt64(123)}}, val)
	assert.Nil(t, err)

	val, err = attributeValue(234.129312)
	assert.Equal(t, &tracepb.AttributeValue{
		Value: &tracepb.AttributeValue_DoubleValue{DoubleValue: cast.ToFloat64(234.129312)}}, val)
	assert.Nil(t, err)

	val, err = attributeValue(true)
	assert.Equal(t, &tracepb.AttributeValue{
		Value: &tracepb.AttributeValue_BoolValue{BoolValue: true},
	}, val)
	assert.Nil(t, err)

	val, err = attributeValue("bob the builder")
	assert.Equal(t, &tracepb.AttributeValue{
		Value: &tracepb.AttributeValue_StringValue{StringValue: &tracepb.TruncatableString{Value: "bob the builder"}},
	}, val)
	assert.Nil(t, err)

	val, err = attributeValue(nil)
	assert.Nil(t, val)
	assert.Equal(t, "error unsupported value type \"<nil>\"", err.Error())

	val, err = attributeValue(Factory{})
	assert.Nil(t, val)
	assert.Equal(t, "error unsupported value type \"attributesprocessor.Factory\"", err.Error())
}

func TestFactory_validateAttributesConfiguration(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	oCfg := cfg.(*Config)

	oCfg.Actions = []ActionKeyValue{
		{Key: "one", Action: "Delete"},
		{Key: "two", Value: 123, Action: "INSERT"},
		{Key: "three", FromAttribute: "two", Action: "upDaTE"},
		{Key: "five", FromAttribute: "two", Action: "upsert"},
	}
	output, err := buildAttributesConfiguration(*oCfg)
	assert.Equal(t, []attributeAction{
		{Key: "one", Action: DELETE},
		{Key: "two", Action: INSERT, AttributeValue: &tracepb.AttributeValue{
			Value: &tracepb.AttributeValue_IntValue{IntValue: cast.ToInt64(123)},
		}},
		{Key: "three", FromAttribute: "two", Action: UPDATE},
		{Key: "five", FromAttribute: "two", Action: UPSERT},
	}, output)
	assert.NoError(t, err)

}

func TestFactory_validateAttributesConfiguration_InvalidConfig(t *testing.T) {
	testcase := []struct {
		name        string
		actionLists []ActionKeyValue
		errorString string
	}{
		{
			name:        "empty action lists",
			actionLists: []ActionKeyValue{},
			errorString: "error creating \"attributes\" processor due to missing required field \"actions\" of processor \"attributes/error\"",
		},
		{
			name: "missing key",
			actionLists: []ActionKeyValue{
				{Key: "one", Action: DELETE},
				{Key: "", Value: 123, Action: UPSERT},
			},
			errorString: "error creating \"attributes\" processor due to missing required field \"key\" at the 1-th actions of processor \"attributes/error\"",
		},
		{
			name: "invalid action",
			actionLists: []ActionKeyValue{
				{Key: "invalid", Action: "invalid"},
			},
			errorString: "error creating \"attributes\" processor due to unsupported action \"invalid\" at the 0-th actions of processor \"attributes/error\"",
		},
		{
			name: "missing value or from attribute",
			actionLists: []ActionKeyValue{
				{Key: "MissingValueFromAttributes", Action: INSERT},
			},
			errorString: "error creating \"attributes\" processor. Either field \"value\" or \"from_attribute\" setting must be specified for 0-th action of processor \"attributes/error\"",
		},
		{
			name: "both set value and from attribute",
			actionLists: []ActionKeyValue{
				{Key: "BothSet", Value: 123, FromAttribute: "aa", Action: UPSERT},
			},
			errorString: "error creating \"attributes\" processor due to both fields \"value\" and \"from_attribute\" being set at the 0-th actions of processor \"attributes/error\"",
		},
	}
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	oCfg := cfg.(*Config)
	oCfg.NameVal = "attributes/error"
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			oCfg.Actions = tc.actionLists
			output, err := buildAttributesConfiguration(*oCfg)
			assert.Nil(t, output)
			assert.Equal(t, tc.errorString, err.Error())
		})
	}
}
