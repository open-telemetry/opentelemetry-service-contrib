// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package json_array // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jarray"

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

	"github.com/valyala/fastjson"
)

const operatorType = "json_array_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new jarray parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new jarray parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		ParserConfig: helper.NewParserConfig(operatorID, operatorType),
	}
}

// Config is the configuration of a jarray parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`
}

// Build will build a jarray parser operator.
func (c Config) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	parserOperator, err := c.ParserConfig.Build(logger)
	if err != nil {
		return nil, err
	}

	return &Parser{
		ParserOperator: parserOperator,
		pool:           new(fastjson.ParserPool),
	}, nil
}

// Parser is an operator that parses jarray in an entry.
type Parser struct {
	helper.ParserOperator
	pool *fastjson.ParserPool
}

// Process will parse an entry for jarray.
func (r *Parser) Process(ctx context.Context, e *entry.Entry) error {
	return r.ParserOperator.ProcessWith(ctx, e, r.parse)
}

func (r *Parser) parse(value any) (any, error) {
	jArrayLine, err := valueAsString(value)
	if err != nil {
		return nil, err
	}

	p := r.pool.Get()
	v, err := p.Parse(jArrayLine)
	r.pool.Put(p)
	if err != nil {
		return nil, errors.New("failed to parse entry")
	}

	jArray := v.GetArray() // a is a []*Value slice
	parsedValues := make([]any, len(jArray))
	for i := range jArray {
		switch jArray[i].Type() {
		case fastjson.TypeNumber:
			parsedValues[i] = jArray[i].GetInt64()
		case fastjson.TypeString:
			parsedValues[i] = string(jArray[i].GetStringBytes())
		case fastjson.TypeTrue:
			parsedValues[i] = true
		case fastjson.TypeFalse:
			parsedValues[i] = false
		case fastjson.TypeNull:
			parsedValues[i] = nil
		case fastjson.TypeObject:
			// Nested objects handled as a string since this parser doesn't support nested headers
			parsedValues[i] = jArray[i].String()
		default:
			return nil, errors.New("failed to parse entry: " + string(jArray[i].MarshalTo(nil)))
		}
	}

	return parsedValues, nil
}

// valueAsString interprets the given value as a string.
func valueAsString(value any) (string, error) {
	var s string
	switch t := value.(type) {
	case string:
		s += t
	case []byte:
		s += string(t)
	default:
		return s, fmt.Errorf("type '%T' cannot be parsed as jarray", value)
	}

	return s, nil
}
