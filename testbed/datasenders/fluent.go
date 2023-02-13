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

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

const (
	fluentDatafileVar = "FLUENT_DATA_SENDER_DATA_FILE"
	fluentPortVar     = "FLUENT_DATA_SENDER_RECEIVER_PORT"
)

// fluentLogsForwarder forwards logs to fluent forwader
type fluentLogsForwarder struct {
	consumer.Logs
	testbed.DataSenderBase
	fluentLogger *fluent.Fluent
	dataFile     *os.File
}

// Ensure fluentLogsForwarder implements LogDataSender.
var _ testbed.LogDataSender = (*fluentLogsForwarder)(nil)

func NewFluentLogsForwarder(t *testing.T, port int) *fluentLogsForwarder {
	var err error
	portOverride := os.Getenv(fluentPortVar)
	if portOverride != "" {
		port, err = strconv.Atoi(portOverride)
		require.NoError(t, err)
	}

	f := &fluentLogsForwarder{DataSenderBase: testbed.DataSenderBase{Port: port}}
	f.Logs, _ = consumer.NewLogs(f.consumeLogs)

	// When FLUENT_DATA_SENDER_DATA_FILE is set, the data sender, writes to a
	// file. This enables users to optionally run the e2e test against a real
	// fluentd/fluentbit agent rather than using the fluent writer the data sender
	// uses by default. In case, one is looking to point a real fluentd/fluentbit agent
	// to the e2e test, they can do so by configuring the fluent agent to read from the
	// file FLUENT_DATA_SENDER_DATA_FILE and forward data to FLUENT_DATA_SENDER_RECEIVER_PORT
	// on 127.0.0.1.
	if dataFileName := os.Getenv(fluentDatafileVar); dataFileName != "" {
		f.dataFile, err = os.OpenFile(dataFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		require.NoError(t, err)
	} else {
		logger, err := fluent.New(fluent.Config{FluentPort: port, Async: true})
		require.NoError(t, err)
		f.fluentLogger = logger
	}
	return f
}

func (f *fluentLogsForwarder) Start() error {
	return nil
}

func (f *fluentLogsForwarder) Stop() error {
	return f.fluentLogger.Close()
}

func (f *fluentLogsForwarder) consumeLogs(_ context.Context, logs plog.Logs) error {
	for i := 0; i < logs.ResourceLogs().Len(); i++ {
		for j := 0; j < logs.ResourceLogs().At(i).ScopeLogs().Len(); j++ {
			ills := logs.ResourceLogs().At(i).ScopeLogs().At(j)
			for k := 0; k < ills.LogRecords().Len(); k++ {
				if f.dataFile == nil {
					if err := f.fluentLogger.Post("", f.convertLogToMap(ills.LogRecords().At(k))); err != nil {
						return err
					}
				} else {
					if _, err := f.dataFile.Write(append(f.convertLogToJSON(ills.LogRecords().At(k)), '\n')); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (f *fluentLogsForwarder) convertLogToMap(lr plog.LogRecord) map[string]string {
	out := map[string]string{}

	if lr.Body().Type() == pcommon.ValueTypeStr {
		out["log"] = lr.Body().Str()
	}

	lr.Attributes().Range(func(k string, v pcommon.Value) bool {
		switch v.Type() {
		case pcommon.ValueTypeStr:
			out[k] = v.Str()
		case pcommon.ValueTypeInt:
			out[k] = strconv.FormatInt(v.Int(), 10)
		case pcommon.ValueTypeDouble:
			out[k] = strconv.FormatFloat(v.Double(), 'f', -1, 64)
		case pcommon.ValueTypeBool:
			out[k] = strconv.FormatBool(v.Bool())
		default:
			panic("missing case")
		}
		return true
	})

	return out
}

func (f *fluentLogsForwarder) convertLogToJSON(lr plog.LogRecord) []byte {
	rec := map[string]string{
		"time": time.Unix(0, int64(lr.Timestamp())).Format("02/01/2006:15:04:05Z"),
	}
	rec["log"] = lr.Body().Str()

	lr.Attributes().Range(func(k string, v pcommon.Value) bool {
		switch v.Type() {
		case pcommon.ValueTypeStr:
			rec[k] = v.Str()
		case pcommon.ValueTypeInt:
			rec[k] = strconv.FormatInt(v.Int(), 10)
		case pcommon.ValueTypeDouble:
			rec[k] = strconv.FormatFloat(v.Double(), 'f', -1, 64)
		case pcommon.ValueTypeBool:
			rec[k] = strconv.FormatBool(v.Bool())
		default:
			panic("missing case")
		}
		return true
	})
	b, err := json.Marshal(rec)
	if err != nil {
		panic("failed to write log: " + err.Error())
	}
	return b
}

func (f *fluentLogsForwarder) Flush() {
	_ = f.dataFile.Sync()
}

func (f *fluentLogsForwarder) GenConfigYAMLStr() string {
	return fmt.Sprintf(`
  fluentforward:
    endpoint: 127.0.0.1:%d`, f.Port)
}

func (f *fluentLogsForwarder) ProtocolName() string {
	return "fluentforward"
}
