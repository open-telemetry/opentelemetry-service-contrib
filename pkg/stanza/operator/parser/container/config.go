// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package container // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/container"

import (
	"fmt"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/errors"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/recombine"
)

const operatorType = "container"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new JSON parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new JSON parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		ParserConfig:            helper.NewParserConfig(operatorID, operatorType),
		Format:                  "",
		AddMetadataFromFilePath: true,
	}
}

// Config is the configuration of a Container parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Format                  string `mapstructure:"format"`
	AddMetadataFromFilePath bool   `mapstructure:"add_metadata_from_filepath"`
}

// Build will build a Container parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	parserOperator, err := c.ParserConfig.Build(set)
	if err != nil {
		return nil, err
	}

	cLogEmitter := helper.NewLogEmitter(set.Logger.Sugar())
	recombineParser, err := createRecombine(set, cLogEmitter)
	if err != nil {
		return nil, fmt.Errorf("failed to create internal recombine config: %w", err)
	}

	wg := sync.WaitGroup{}

	if c.Format != "" {
		switch c.Format {
		case dockerFormat, crioFormat, containerdFormat:
		default:
			return &Parser{}, errors.NewError(
				"operator config has an invalid `format` field.",
				"ensure that the `format` field is set to one of `docker`, `crio`, `containerd`.",
				"format", c.OnError,
			)
		}
	}

	p := &Parser{
		ParserOperator:          parserOperator,
		recombineParser:         recombineParser,
		json:                    jsoniter.ConfigFastest,
		format:                  c.Format,
		addMetadataFromFilepath: c.AddMetadataFromFilePath,
		crioLogEmitter:          cLogEmitter,
		criConsumers:            &wg,
	}
	return p, nil
}

// createRecombine creates an internal recombine operator which outputs to an async helper.LogEmitter
// the equivalent recombine config:
//
//	combine_field: body
//	combine_with: ""
//	is_last_entry: attributes.logtag == 'F'
//	max_log_size: 102400
//	source_identifier: attributes["log.file.path"]
//	type: recombine
func createRecombine(set component.TelemetrySettings, cLogEmitter *helper.LogEmitter) (operator.Operator, error) {
	recombineParserCfg := createRecombineConfig()
	recombineParser, err := recombineParserCfg.Build(set)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve internal recombine config: %w", err)
	}

	// set the LogEmmiter as the output of the recombine parser
	recombineParser.SetOutputIDs([]string{cLogEmitter.OperatorID})
	if err := recombineParser.SetOutputs([]operator.Operator{cLogEmitter}); err != nil {
		return nil, fmt.Errorf("failed to set outputs of internal recombine")
	}

	return recombineParser, nil
}

func createRecombineConfig() *recombine.Config {
	recombineParserCfg := recombine.NewConfigWithID(recombineInternalID)
	recombineParserCfg.IsLastEntry = "attributes.logtag == 'F'"
	recombineParserCfg.CombineField = entry.NewBodyField()
	recombineParserCfg.CombineWith = ""
	recombineParserCfg.SourceIdentifier = entry.NewAttributeField("log.file.path")
	recombineParserCfg.MaxLogSize = 102400
	return recombineParserCfg
}
