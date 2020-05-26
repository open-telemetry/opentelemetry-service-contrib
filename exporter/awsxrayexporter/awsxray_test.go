// Copyright 2019, OpenTelemetry Authors
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

package awsxrayexporter

import (
	"context"
	"encoding/binary"
	"fmt"
	"go.opentelemetry.io/collector/consumer/pdata"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	semconventions "go.opentelemetry.io/collector/translator/conventions"
	"go.uber.org/zap"
)

func TestTraceExport(t *testing.T) {
	traceExporter := initializeTraceExporter()
	ctx := context.Background()
	td := constructSpanData()
	err := traceExporter.ConsumeTraces(ctx, td)
	assert.Nil(t, err)
}

func initializeTraceExporter() component.TraceExporter {
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIASSWVJUY4PZXXXXXX")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "XYrudg2H87u+ADAAq19Wqx3D41a09RsTXXXXXXXX")
	os.Setenv("AWS_DEFAULT_REGION", "us-east-1")
	os.Setenv("AWS_REGION", "us-east-1")
	logger := zap.NewNop()
	factory := Factory{}
	config := factory.CreateDefaultConfig()
	config.(*Config).Region = "us-east-1"
	config.(*Config).LocalMode = true
	mconn := new(mockConn)
	mconn.sn, _ = getDefaultSession(logger)
	traceExporter, err := NewTraceExporter(config, logger, mconn)
	if err != nil {
		panic(err)
	}
	return traceExporter
}

func constructSpanData() pdata.Traces {
	resource := constructResource()

	traces := pdata.NewTraces()
	traces.ResourceSpans().Resize(1)
	rspans := traces.ResourceSpans().At(0)
	resource.CopyTo(rspans.Resource())
	rspans.InstrumentationLibrarySpans().Resize(1)
	ispans := rspans.InstrumentationLibrarySpans().At(0)
	ispans.Spans().Resize(2)
	constructHTTPClientSpan().CopyTo(ispans.Spans().At(0))
	constructHTTPServerSpan().CopyTo(ispans.Spans().At(1))
	return traces
}

func constructResource() pdata.Resource {
	resource := pdata.NewResource()
	resource.Attributes().InsertString(semconventions.AttributeServiceName, "signup_aggregator")
	resource.Attributes().InsertString(semconventions.AttributeContainerName, "signup_aggregator")
	resource.Attributes().InsertString(semconventions.AttributeContainerImage, "otel/signupaggregator")
	resource.Attributes().InsertString(semconventions.AttributeContainerTag, "v1")
	resource.Attributes().InsertString(semconventions.AttributeCloudProvider, "aws")
	resource.Attributes().InsertString(semconventions.AttributeCloudAccount, "999999998")
	resource.Attributes().InsertString(semconventions.AttributeCloudRegion, "us-west-2")
	resource.Attributes().InsertString(semconventions.AttributeCloudZone, "us-west-1b")
	return resource
}

func constructHTTPClientSpan() pdata.Span {
	attributes := make(map[string]interface{})
	attributes[semconventions.AttributeComponent] = semconventions.ComponentTypeHTTP
	attributes[semconventions.AttributeHTTPMethod] = "GET"
	attributes[semconventions.AttributeHTTPURL] = "https://api.example.com/users/junit"
	attributes[semconventions.AttributeHTTPStatusCode] = 200
	endTime := time.Now().Round(time.Second)
	startTime := endTime.Add(-90 * time.Second)
	spanAttributes := constructSpanAttributes(attributes)

	span := pdata.NewSpan()
	span.SetTraceID(newTraceID())
	span.SetSpanID(newSegmentID())
	span.SetParentSpanID(newSegmentID())
	span.SetName("/users/junit")
	span.SetKind(pdata.SpanKindCLIENT)
	span.SetStartTime(pdata.TimestampUnixNano(startTime.UnixNano()))
	span.SetEndTime(pdata.TimestampUnixNano(endTime.UnixNano()))

	status := pdata.NewSpanStatus()
	status.SetCode(0)
	status.SetMessage("OK")
	status.CopyTo(span.Status())

	spanAttributes.CopyTo(span.Attributes())
	return span
}

func constructHTTPServerSpan() pdata.Span {
	attributes := make(map[string]interface{})
	attributes[semconventions.AttributeComponent] = semconventions.ComponentTypeHTTP
	attributes[semconventions.AttributeHTTPMethod] = "GET"
	attributes[semconventions.AttributeHTTPURL] = "https://api.example.com/users/junit"
	attributes[semconventions.AttributeHTTPClientIP] = "192.168.15.32"
	attributes[semconventions.AttributeHTTPStatusCode] = 200
	endTime := time.Now().Round(time.Second)
	startTime := endTime.Add(-90 * time.Second)
	spanAttributes := constructSpanAttributes(attributes)

	span := pdata.NewSpan()
	span.SetTraceID(newTraceID())
	span.SetSpanID(newSegmentID())
	span.SetParentSpanID(newSegmentID())
	span.SetName("/users/junit")
	span.SetKind(pdata.SpanKindSERVER)
	span.SetStartTime(pdata.TimestampUnixNano(startTime.UnixNano()))
	span.SetEndTime(pdata.TimestampUnixNano(endTime.UnixNano()))

	status := pdata.NewSpanStatus()
	status.SetCode(0)
	status.SetMessage("OK")
	status.CopyTo(span.Status())

	spanAttributes.CopyTo(span.Attributes())
	return span
}

func convertTimeToTimestamp(t time.Time) *timestamp.Timestamp {
	if t.IsZero() {
		return nil
	}
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}

func constructSpanAttributes(attributes map[string]interface{}) pdata.AttributeMap {
	attrs := pdata.NewAttributeMap()
	for key, value := range attributes {
		if cast, ok := value.(int); ok {
			attrs.InsertInt(key, int64(cast))
		} else if cast, ok := value.(int64); ok {
			attrs.InsertInt(key, cast)
		} else {
			attrs.InsertString(key, fmt.Sprintf("%v", value))
		}
	}
	return attrs
}

func newTraceID() []byte {
	var r [16]byte
	epoch := time.Now().Unix()
	binary.BigEndian.PutUint32(r[0:4], uint32(epoch))
	_, err := rand.Read(r[4:])
	if err != nil {
		panic(err)
	}
	return r[:]
}

func newSegmentID() []byte {
	var r [8]byte
	_, err := rand.Read(r[:])
	if err != nil {
		panic(err)
	}
	return r[:]
}
