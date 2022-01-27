// Copyright OpenTelemetry Authors
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

package azuremonitorexporter

/*
Contains tests for logexporter.go and log_to_envelope.go
*/

import (
	"context"
	"testing"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/model/otlp"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

const (
	defaultEnvelopeName = "Microsoft.ApplicationInsights.Message"
	defaultdBaseType    = "MessageData"
)

var (
	testLogs = []byte(`{"resourceLogs":[{"resource":{"attributes":[{"key":"service.name","value":{"stringValue":"dotnet"}}]},"instrumentationLibraryLogs":[{"instrumentationLibrary":{},"logs":[{"timeUnixNano":"1643240673066096200","severityText":"Information","name":"FilterModule.Program","body":{"stringValue":"Message Body"},"flags":1,"traceId":"7b20d1349ef9b6d6f9d4d1d4a3ac2e82","spanId":"0c2ad924e1771630"}]}]}]}`)
)

// Tests proper wrapping of a log record to an envelope
func TestLogRecordToEnvelope(t *testing.T) {
	logRecord := getTestLogRecord(t)
	logPacker := getLogPacker()
	envelope := logPacker.LogRecordToEnvelope(logRecord)

	assert.NotNil(t, envelope)
	assert.Equal(t, defaultEnvelopeName, envelope.Name)
	assert.Equal(t, toTime(logRecord.Timestamp()).Format(time.RFC3339Nano), envelope.Time)
	assert.NotNil(t, envelope.Data)
	envelopeData := envelope.Data.(*contracts.Data)
	assert.Equal(t, defaultdBaseType, envelopeData.BaseType)

	assert.NotNil(t, envelopeData.BaseData)

	messageData := envelopeData.BaseData.(*contracts.MessageData)
	assert.Equal(t, messageData.Message, logRecord.Body().StringVal())
	assert.Equal(t, messageData.SeverityLevel, contracts.Information)

	hexTraceId := logRecord.TraceID().HexString()
	assert.Equal(t, messageData.Properties[TRACE_ID_TAG], hexTraceId)
	assert.Equal(t, envelope.Tags[contracts.OperationId], hexTraceId)

	assert.Equal(t, messageData.Properties[SPAN_ID_TAG], logRecord.SpanID().HexString())
	assert.Equal(t, messageData.Properties[CATEGORY_NAME_TAG], logRecord.Name())

}

// Test conversion from logRecord.SeverityText() to contracts.SeverityLevel()
func TestToAiSeverityLevel(t *testing.T) {
	logPacker := getLogPacker()
	for severityText, expectedSeverityLevel := range severityLevelMap {
		severityLevel := logPacker.toAiSeverityLevel(severityText)
		assert.Equal(t, severityLevel, expectedSeverityLevel)
	}
}

// Test onLogData callback for the test logs data
func TestExporterLogDataCallback(t *testing.T) {
	mockTransportChannel := getMockTransportChannel()
	exporter := getLogsExporter(defaultConfig, mockTransportChannel)

	logs := getTestLogs(t)

	assert.NoError(t, exporter.onLogData(context.Background(), logs))

	mockTransportChannel.AssertNumberOfCalls(t, "Send", 1)
}

func getLogsExporter(config *Config, transportChannel transportChannel) *logExporter {
	return &logExporter{
		config,
		transportChannel,
		zap.NewNop(),
	}
}

func getLogPacker() *logPacker {
	return newLogPacker(zap.NewNop())
}

func getTestLogs(t *testing.T) pdata.Logs {
	logsMarshaler := otlp.NewJSONLogsUnmarshaler()
	logs, err := logsMarshaler.UnmarshalLogs(testLogs)
	assert.NoError(t, err, "Can't unmarshal testing logs data -> %s", err)
	return logs
}

func getTestLogRecord(t *testing.T) pdata.LogRecord {
	var logRecord pdata.LogRecord
	logs := getTestLogs(t)
	resourceLogs := logs.ResourceLogs()
	instrumentationLibraryLogs := resourceLogs.At(0).InstrumentationLibraryLogs()
	logRecords := instrumentationLibraryLogs.At(0).Logs()
	logRecord = logRecords.At(0)

	return logRecord
}
