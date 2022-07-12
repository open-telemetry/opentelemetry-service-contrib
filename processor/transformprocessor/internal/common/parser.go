// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"encoding/hex"
	"errors"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/multierr"
)

// ParsedQuery represents a parsed query. It is the entry point into the query DSL.
// nolint:govet
type ParsedQuery struct {
	Invocation Invocation `parser:"@@" mapstructure:"invocation"`
	Condition  *Condition `parser:"( 'where' @@ )?" mapstructure:"condition"`
}

// Condition represents an optional boolean condition on the RHS of a query.
// nolint:govet
type Condition struct {
	Left  Value  `parser:"@@" mapstructure:"left"`
	Op    string `parser:"@('==' | '!=')" mapstructure:"op"`
	Right Value  `parser:"@@" mapstructure:"right"`
}

// Invocation represents a function call.
// nolint:govet
type Invocation struct {
	Function  string  `parser:"@Ident" mapstructure:"function"`
	Arguments []Value `parser:"'(' ( @@ ( ',' @@ )* )? ')'" mapstructure:"arguments"`
}

// Value represents a part of a parsed query which is resolved to a value of some sort. This can be a telemetry path
// expression, function call, or literal.
// nolint:govet
type Value struct {
	Invocation *Invocation `parser:"( @@" mapstructure:"invocation"`
	Bytes      *Bytes      `parser:"| @Bytes" mapstructure:"bytes"`
	String     *string     `parser:"| @String" mapstructure:"string"`
	Float      *float64    `parser:"| @Float" mapstructure:"float"`
	Int        *int64      `parser:"| @Int" mapstructure:"int"`
	Bool       *Boolean    `parser:"| @('true' | 'false')" mapstructure:"bool"`
	IsNil      *IsNil      `parser:"| @'nil'" mapstructure:"is_nil"`
	Path       *Path       `parser:"| @@ )" mapstructure:"path"`
}

// Path represents a telemetry path expression.
// nolint:govet
type Path struct {
	Fields []Field `parser:"@@ ( '.' @@ )*" mapstructure:"fields"`
}

// Field is an item within a Path.
// nolint:govet
type Field struct {
	Name   string  `parser:"@Ident" mapstructure:"name"`
	MapKey *string `parser:"( '[' @String ']' )?" mapstructure:"map_key"`
}

// Query holds a top level Query for processing telemetry data. A Query is a combination of a function
// invocation and the condition to match telemetry for invoking the function.
type Query struct {
	Function  ExprFunc
	Condition condFunc
}

// Bytes type for capturing byte arrays
type Bytes []byte

func (b *Bytes) Capture(values []string) error {
	rawStr := values[0][2:]
	bytes, err := hex.DecodeString(rawStr)
	if err != nil {
		return err
	}
	*b = bytes
	return nil
}

// Boolean Type for capturing booleans, see:
// https://github.com/alecthomas/participle#capturing-boolean-value
type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type IsNil bool

func (n *IsNil) Capture(_ []string) error {
	*n = true
	return nil
}

func ParseQueries(statements []string, functions map[string]interface{}, pathParser PathExpressionParser, enumParser EnumParser) ([]Query, error) {
	parsedQueries := make([]ParsedQuery, 0)
	var errors error

	for _, statement := range statements {
		parsed, err := parseQuery(statement)
		if err != nil {
			errors = multierr.Append(errors, err)
		} else {
			parsedQueries = append(parsedQueries, *parsed)
		}
	}

	queries, err := InterpretQueries(parsedQueries, functions, pathParser, enumParser)
	errors = multierr.Append(errors, err)
	if errors != nil {
		return nil, errors
	}
	return queries, nil
}

func InterpretQueries(parsedQueries []ParsedQuery, functions map[string]interface{}, pathParser PathExpressionParser, enumParser EnumParser) ([]Query, error) {
	queries := make([]Query, 0)
	var errors error

	for _, parsedQuery := range parsedQueries {
		query, err := generateQuery(parsedQuery, functions, pathParser, enumParser)
		if err != nil {
			errors = multierr.Append(errors, err)
		} else {
			queries = append(queries, query)
		}
	}

	if errors != nil {
		return nil, errors
	}
	return queries, nil
}

func generateQuery(parsedQuery ParsedQuery, functions map[string]interface{}, pathParser PathExpressionParser, enumParser EnumParser) (Query, error) {
	var errors error
	function, err := NewFunctionCall(parsedQuery.Invocation, functions, pathParser, enumParser)
	if err != nil {
		errors = multierr.Append(errors, err)
	}
	condition, err := newConditionEvaluator(parsedQuery.Condition, functions, pathParser, enumParser)
	if err != nil {
		errors = multierr.Append(errors, err)
	}
	if errors != nil {
		return Query{}, errors
	}
	return Query{
		Function:  function,
		Condition: condition,
	}, nil
}

var parser = newParser()

func parseQuery(raw string) (*ParsedQuery, error) {
	parsed := &ParsedQuery{}
	err := parser.ParseString("", raw, parsed)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

// newParser returns a parser that can be used to read a string into a ParsedQuery. An error will be returned if the string
// is not formatted for the DSL.
func newParser() *participle.Parser {
	lex := lexer.MustSimple([]lexer.SimpleRule{
		{Name: `Ident`, Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
		{Name: `Bytes`, Pattern: `0x[a-fA-F0-9]+`},
		{Name: `Float`, Pattern: `[-+]?\d*\.\d+([eE][-+]?\d+)?`},
		{Name: `Int`, Pattern: `[-+]?\d+`},
		{Name: `String`, Pattern: `"(\\"|[^"])*"`},
		{Name: `Operators`, Pattern: `==|!=|[,.()\[\]]`},
		{Name: "whitespace", Pattern: `\s+`},
	})
	parser, err := participle.Build(&ParsedQuery{},
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.Elide("whitespace"),
	)
	if err != nil {
		panic("Unable to initialize parser, this is a programming error in the transformprocesor")
	}
	return parser
}

func ParseSpanID(spanIDStr string) (pcommon.SpanID, error) {
	id, err := hex.DecodeString(spanIDStr)
	if err != nil {
		return pcommon.SpanID{}, err
	}
	if len(id) != 8 {
		return pcommon.SpanID{}, errors.New("span ids must be 8 bytes")
	}
	var idArr [8]byte
	copy(idArr[:8], id)
	return pcommon.NewSpanID(idArr), nil
}

func ParseTraceID(traceIDStr string) (pcommon.TraceID, error) {
	id, err := hex.DecodeString(traceIDStr)
	if err != nil {
		return pcommon.TraceID{}, err
	}
	if len(id) != 16 {
		return pcommon.TraceID{}, errors.New("traces ids must be 16 bytes")
	}
	var idArr [16]byte
	copy(idArr[:16], id)
	return pcommon.NewTraceID(idArr), nil
}
