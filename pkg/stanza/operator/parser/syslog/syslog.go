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

package syslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	sl "github.com/influxdata/go-syslog/v3"
	"github.com/influxdata/go-syslog/v3/nontransparent"
	"github.com/influxdata/go-syslog/v3/octetcounting"
	"github.com/influxdata/go-syslog/v3/rfc3164"
	"github.com/influxdata/go-syslog/v3/rfc5424"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const (
	operatorType = "syslog_parser"

	RFC3164 = "rfc3164"
	RFC5424 = "rfc5424"

	NULTrailer = "NUL"
	LFTrailer  = "LF"
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new syslog parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new syslog parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		ParserConfig: helper.NewParserConfig(operatorID, operatorType),
	}
}

// Config is the configuration of a syslog parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`
	BaseConfig          `mapstructure:",squash"`
}

// BaseConfig is the detailed configuration of a syslog parser.
type BaseConfig struct {
	Protocol                     string  `mapstructure:"protocol,omitempty"`
	Location                     string  `mapstructure:"location,omitempty"`
	EnableOctetCounting          bool    `mapstructure:"enable_octet_counting,omitempty"`
	NonTransparentFramingTrailer *string `mapstructure:"non_transparent_framing_trailer,omitempty"`
}

// Build will build a JSON parser operator.
func (c Config) Build(logger *zap.SugaredLogger) (operator.Operator, error) {
	if c.ParserConfig.TimeParser == nil {
		parseFromField := entry.NewAttributeField("timestamp")
		c.ParserConfig.TimeParser = &helper.TimeParser{
			ParseFrom:  &parseFromField,
			LayoutType: helper.NativeKey,
		}
	}

	parserOperator, err := c.ParserConfig.Build(logger)
	if err != nil {
		return nil, err
	}

	switch {
	case c.Protocol == "":
		return nil, fmt.Errorf("missing field 'protocol'")
	case c.Protocol == RFC3164 && (c.NonTransparentFramingTrailer != nil || c.EnableOctetCounting):
		return nil, errors.New("octet_counting and non_transparent_framing are only compatible with protocol rfc5424")
	case c.Protocol == RFC5424 && (c.NonTransparentFramingTrailer != nil && c.EnableOctetCounting):
		return nil, errors.New("only one of octet_counting or non_transparent_framing can be enabled")
	case c.Protocol == RFC5424 && c.NonTransparentFramingTrailer != nil:
		if err := validateNonTransparentFramingTrailer(c.NonTransparentFramingTrailer); err != nil {
			return nil, err
		}
	}

	if c.Location == "" {
		c.Location = "UTC"
	}

	location, err := time.LoadLocation(c.Location)
	if err != nil {
		return nil, fmt.Errorf("failed to load location %s: %w", c.Location, err)
	}

	return &Parser{
		ParserOperator:               parserOperator,
		protocol:                     c.Protocol,
		location:                     location,
		enableOctetCounting:          c.EnableOctetCounting,
		nonTransparentFramingTrailer: c.NonTransparentFramingTrailer,
	}, nil
}

// parseFunc a parseFunc determines how the raw input is to be parsed into a syslog message
type parseFunc func(input []byte) (sl.Message, error)

func (s *Parser) buildParseFunc() (parseFunc, error) {
	switch s.protocol {
	case RFC3164:
		return func(input []byte) (sl.Message, error) {
			machine := rfc3164.NewMachine(rfc3164.WithLocaleTimezone(s.location))
			return machine.Parse(input)
		}, nil
	case RFC5424:
		switch {
		// Octet Counting Parsing RFC6587
		case s.enableOctetCounting:
			return newOctetCountingParseFunc(), nil
		// Non-Transparent-Framing Parsing RFC6587
		case s.nonTransparentFramingTrailer != nil:
			return newNonTransparentFramingParseFunc(*s.nonTransparentFramingTrailer), nil
		// Raw RFC5424 parsing
		default:
			return func(input []byte) (sl.Message, error) {
				machine := rfc5424.NewMachine()
				return machine.Parse(input)
			}, nil
		}

	default:
		return nil, fmt.Errorf("invalid protocol %s", s.protocol)
	}
}

// Parser is an operator that parses syslog.
type Parser struct {
	helper.ParserOperator
	protocol                     string
	location                     *time.Location
	enableOctetCounting          bool
	nonTransparentFramingTrailer *string
}

// Process will parse an entry field as syslog.
func (s *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	return s.ParserOperator.ProcessWithCallback(ctx, entry, s.parse, postprocess)
}

// parse will parse a value as syslog.
func (s *Parser) parse(value interface{}) (interface{}, error) {
	bytes, err := toBytes(value)
	if err != nil {
		return nil, err
	}

	pFunc, err := s.buildParseFunc()
	if err != nil {
		return nil, err
	}

	slog, err := pFunc(bytes)
	if err != nil {
		return nil, err
	}

	switch message := slog.(type) {
	case *rfc3164.SyslogMessage:
		return s.parseRFC3164(message)
	case *rfc5424.SyslogMessage:
		return s.parseRFC5424(message)
	default:
		return nil, fmt.Errorf("parsed value was not rfc3164 or rfc5424 compliant")
	}
}

// parseRFC3164 will parse an RFC3164 syslog message.
func (s *Parser) parseRFC3164(syslogMessage *rfc3164.SyslogMessage) (map[string]interface{}, error) {
	value := map[string]interface{}{
		"timestamp": syslogMessage.Timestamp,
		"priority":  syslogMessage.Priority,
		"facility":  syslogMessage.Facility,
		"severity":  syslogMessage.Severity,
		"hostname":  syslogMessage.Hostname,
		"appname":   syslogMessage.Appname,
		"proc_id":   syslogMessage.ProcID,
		"msg_id":    syslogMessage.MsgID,
		"message":   syslogMessage.Message,
	}
	return s.toSafeMap(value)
}

// parseRFC5424 will parse an RFC5424 syslog message.
func (s *Parser) parseRFC5424(syslogMessage *rfc5424.SyslogMessage) (map[string]interface{}, error) {
	value := map[string]interface{}{
		"timestamp":       syslogMessage.Timestamp,
		"priority":        syslogMessage.Priority,
		"facility":        syslogMessage.Facility,
		"severity":        syslogMessage.Severity,
		"hostname":        syslogMessage.Hostname,
		"appname":         syslogMessage.Appname,
		"proc_id":         syslogMessage.ProcID,
		"msg_id":          syslogMessage.MsgID,
		"message":         syslogMessage.Message,
		"structured_data": syslogMessage.StructuredData,
		"version":         syslogMessage.Version,
	}
	return s.toSafeMap(value)
}

// toSafeMap will dereference any pointers on the supplied map.
func (s *Parser) toSafeMap(message map[string]interface{}) (map[string]interface{}, error) {
	for key, val := range message {
		switch v := val.(type) {
		case *string:
			if v == nil {
				delete(message, key)
				continue
			}
			message[key] = *v
		case *uint8:
			if v == nil {
				delete(message, key)
				continue
			}
			message[key] = int(*v)
		case uint16:
			message[key] = int(v)
		case *time.Time:
			if v == nil {
				delete(message, key)
				continue
			}
			message[key] = *v
		case *map[string]map[string]string:
			if v == nil {
				delete(message, key)
				continue
			}
			message[key] = *v
		default:
			return nil, fmt.Errorf("key %s has unknown field of type %T", key, v)
		}
	}

	return message, nil
}

func toBytes(value interface{}) ([]byte, error) {
	switch v := value.(type) {
	case string:
		return []byte(v), nil
	default:
		return nil, fmt.Errorf("unable to convert type '%T' to bytes", value)
	}
}

var severityMapping = [...]entry.Severity{
	0: entry.Fatal,
	1: entry.Error3,
	2: entry.Error2,
	3: entry.Error,
	4: entry.Warn,
	5: entry.Info2,
	6: entry.Info,
	7: entry.Debug,
}

var severityText = [...]string{
	0: "emerg",
	1: "alert",
	2: "crit",
	3: "err",
	4: "warning",
	5: "notice",
	6: "info",
	7: "debug",
}

var severityField = entry.NewAttributeField("severity")

func postprocess(e *entry.Entry) error {
	sev, ok := severityField.Delete(e)
	if !ok {
		return fmt.Errorf("severity field does not exist")
	}

	sevInt, ok := sev.(int)
	if !ok {
		return fmt.Errorf("severity field is not an int")
	}

	if sevInt < 0 || sevInt > 7 {
		return fmt.Errorf("invalid severity '%d'", sevInt)
	}

	e.Severity = severityMapping[sevInt]
	e.SeverityText = severityText[sevInt]

	_, ok = entry.NewAttributeField("timestamp").Delete(e)
	if !ok {
		return fmt.Errorf("failed to cleanup timestamp")
	}

	return nil
}

func validateNonTransparentFramingTrailer(trailer *string) error {
	if trailer == nil {
		return nil
	}

	if *trailer != NULTrailer && *trailer != LFTrailer {
		return fmt.Errorf("invalid non_transparent_framing_trailer '%s'. Must be either 'LF' or 'NUL'", *trailer)
	}

	return nil
}

func newOctetCountingParseFunc() parseFunc {
	return func(input []byte) (message sl.Message, err error) {
		listener := func(res *sl.Result) {
			message = res.Message
			err = res.Error
		}
		parser := octetcounting.NewParser(sl.WithBestEffort(), sl.WithListener(listener))
		reader := bytes.NewReader(input)
		parser.Parse(reader)
		return
	}
}

func newNonTransparentFramingParseFunc(trailer string) parseFunc {
	return func(input []byte) (message sl.Message, err error) {
		listener := func(res *sl.Result) {
			message = res.Message
			err = res.Error
		}

		trailerType := nontransparent.LF
		if trailer == NULTrailer {
			trailerType = nontransparent.NUL
		}

		parser := nontransparent.NewParser(sl.WithBestEffort(), nontransparent.WithTrailer(trailerType), sl.WithListener(listener))
		reader := bytes.NewReader(input)
		parser.Parse(reader)
		return
	}
}
