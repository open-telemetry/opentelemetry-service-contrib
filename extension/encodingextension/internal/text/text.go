// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package text // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encodingextension/internal/text"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/textutils"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/codec"
)

type textLogCodec struct {
	enc *textutils.Encoding
}

func NewLogCodec(textEncoding string) codec.Log {
	encCfg := textutils.NewEncodingConfig()
	encCfg.Encoding = textEncoding
	enc, err := encCfg.Build()
	if err != nil {
		panic(err)
	}
	return &textLogCodec{
		enc: &enc,
	}
}

func (r *textLogCodec) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	p := plog.NewLogs()
	decoded, err := r.enc.Decode(buf)
	if err != nil {
		return p, err
	}

	l := p.ResourceLogs().AppendEmpty().ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
	l.SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	l.Body().SetStr(string(decoded))
	return p, nil
}

func (r *textLogCodec) MarshalLogs(ld plog.Logs) ([]byte, error) {
	marshaler := &plog.JSONMarshaler{}
	return marshaler.MarshalLogs(ld)
}
