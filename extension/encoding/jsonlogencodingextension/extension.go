// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jsonlogencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/jsonlogencodingextension"

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.LogsMarshalerExtension   = (*jsonLogExtension)(nil)
	_ encoding.LogsUnmarshalerExtension = (*jsonLogExtension)(nil)
)

type jsonLogExtension struct {
	config     component.Config
}

func (e *jsonLogExtension) MarshalLogs(ld plog.Logs) ([]byte, error) {
	if !e.config.(*Config).RawLog {
		return e.LogProcessor(ld)
	}
	logRecord := ld.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Body()
	var raw map[string]any
	switch logRecord.Type() {
	case pcommon.ValueTypeMap:
		raw = logRecord.Map().AsRaw()
	default:
		return nil, fmt.Errorf("Marshal: Expected 'Map' found '%v'", logRecord.Type().String())
	}
	buf, err := jsoniter.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (e *jsonLogExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	p := plog.NewLogs()

	// get json logs from the buffer
	jsonVal := map[string]any{}
	if err := jsoniter.Unmarshal(buf, &jsonVal); err != nil {
		return p, err
	}

	// create a new log record
	logRecords := p.ResourceLogs().AppendEmpty().ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
	logRecords.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))

	// Set the unmarshaled jsonVal as the body of the log record
	if err := logRecords.Body().SetEmptyMap().FromRaw(jsonVal); err != nil {
		return p, err
	}
	return p, nil
}

func (e *jsonLogExtension) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (e *jsonLogExtension) Shutdown(_ context.Context) error {
	return nil
}

func (e *jsonLogExtension) LogProcessor(ld plog.Logs) ([]byte, error){
	prettyLogs := []prettyLogBody{}

	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)
		resourceAttrs := attrsValue(rl.Resource().Attributes())

		sls := rl.ScopeLogs()
		for j := 0; j < sls.Len(); j++ {
			sl := sls.At(j)
			logs := sl.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				log := logs.At(k)
				logEvent := prettyLogBody{
					Body: log.Body().AsRaw(),
					Resource: resourceAttrs,
					Attributes: attrsValue(log.Attributes()),
				}
				prettyLogs = append(prettyLogs, logEvent)
			}
		}
	}

	return jsoniter.Marshal(prettyLogs)
}


type prettyLogBody struct {
	Body                   any             `json:"body,omitempty"`
	Attributes             map[string]any  `json:"attributes,omitempty"`
	Resource               map[string]any  `json:"resource,omitempty"`
}

func attrsValue(attrs pcommon.Map) map[string]any {
	if attrs.Len() == 0 {
		return nil
	}
	out := make(map[string]any, attrs.Len())
	attrs.Range(func(k string, v pcommon.Value) bool {
		out[k] = v.AsRaw()
		return true
	})
	return out
}
