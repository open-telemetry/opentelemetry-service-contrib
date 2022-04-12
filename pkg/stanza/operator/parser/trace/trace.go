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

package trace // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/trace"

import (
	"context"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

func init() {
	operator.Register("trace_parser", func() operator.Builder { return NewParserConfig("") })
}

// NewParserConfig creates a new trace parser config with default values
func NewParserConfig(operatorID string) *ParserConfig {
	return &ParserConfig{
		TransformerConfig: helper.NewTransformerConfig(operatorID, "trace_parser"),
		TraceParser:       helper.NewTraceParser(),
	}
}

// ParserConfig is the configuration of a trace parser operator.
type ParserConfig struct {
	helper.TransformerConfig `mapstructure:",squash"           yaml:",inline"`
	helper.TraceParser       `mapstructure:",omitempty,squash" yaml:",omitempty,inline"`
}

// Build will build a trace parser operator.
func (c ParserConfig) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	transformerOperator, err := c.TransformerConfig.Build(logger)
	if err != nil {
		return nil, err
	}

	if err := c.TraceParser.Validate(); err != nil {
		return nil, err
	}

	return &Parser{
		TransformerOperator: transformerOperator,
		TraceParser:         c.TraceParser,
	}, nil
}

// Parser is an operator that parses traces from fields to an entry.
type Parser struct {
	helper.TransformerOperator
	helper.TraceParser
}

// Process will parse traces from an entry.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	return p.ProcessWith(ctx, entry, p.Parse)
}
