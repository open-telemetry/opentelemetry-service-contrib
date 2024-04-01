// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package jsonarray // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jsonarray"

import (
	"strings"

	"github.com/valyala/fastjson"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

var jsonArrayParserFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"logs.jsonParserArray",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, allows usage of `json_array_parser`."),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/30321"),
)

var operatorType = component.MustNewType("json_array_parser")

func init() {
	if jsonArrayParserFeatureGate.IsEnabled() {
		operator.RegisterFactory(NewFactory())
	}
}

// NewFactory creates a new factory.
func NewFactory() operator.Factory {
	return operator.NewFactory(operatorType, newDefaultConfig, createOperator)
}

func newDefaultConfig(operatorID string) component.Config {
	return &Config{
		ParserConfig: helper.NewParserConfig(operatorID, operatorType.String()),
	}
}

func createOperator(set component.TelemetrySettings, cfg component.Config) (operator.Operator, error) {
	c := cfg.(*Config)
	parserOperator, err := helper.NewParser(set, c.ParserConfig)
	if err != nil {
		return nil, err
	}

	if c.Header != "" {
		return &Parser{
			ParserOperator: parserOperator,
			parse:          generateParseToMapFunc(new(fastjson.ParserPool), strings.Split(c.Header, headerDelimiter)),
		}, nil
	}

	return &Parser{
		ParserOperator: parserOperator,
		parse:          generateParseToArrayFunc(new(fastjson.ParserPool)),
	}, nil
}
