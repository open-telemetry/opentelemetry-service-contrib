// Copyright 2018, OpenCensus Authors
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

package observabilitytest

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"

	tracepb "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	"github.com/census-instrumentation/opencensus-service/data"
	"github.com/census-instrumentation/opencensus-service/exporter"
	"github.com/census-instrumentation/opencensus-service/internal"
	"github.com/census-instrumentation/opencensus-service/observability"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	fakeReceiverName = "fake_receiver_trace"
	fakeExporterName = "fake_exporter_trace"
)

type nopMetricsExporter int

var _ view.Exporter = (*nopMetricsExporter)(nil)

func (cme *nopMetricsExporter) ExportView(vd *view.Data) {}

// CheckRecordedMetricsForTraceExporter checks that the given TraceExporter records the correct set of metrics with the correct
// set of tags by sending some TraceData to the exporter. The exporter should be able to handle the requests.
func CheckRecordedMetricsForTraceExporter(t *testing.T, te exporter.TraceExporter, exporterTagName string) {
	// Register a nop metrics exporter for the OC library.
	nmp := new(nopMetricsExporter)
	view.RegisterExporter(nmp)
	defer view.UnregisterExporter(nmp)

	// Now for the stats exporter
	if err := view.Register(observability.ViewExporterReceivedSpans, observability.ViewExporterDroppedSpans); err != nil {
		t.Fatalf("Failed to register all views: %v", err)
	}
	defer view.Unregister(observability.ViewExporterReceivedSpans, observability.ViewExporterDroppedSpans)

	now := time.Now().UTC()
	spans := []*tracepb.Span{
		{
			TraceId:      []byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E},
			SpanId:       []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7},
			ParentSpanId: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Name:         &tracepb.TruncatableString{Value: "ServerSpan"},
			Kind:         tracepb.Span_SERVER,
			StartTime:    internal.TimeToTimestamp(now.Add(-15 * time.Millisecond)),
			EndTime:      internal.TimeToTimestamp(now),
			Status:       &tracepb.Status{Code: int32(0), Message: "OK"},
			Tracestate:   &tracepb.Span_Tracestate{},
			Links: &tracepb.Span_Links{
				Link: []*tracepb.Span_Link{
					{
						TraceId: []byte{0x4F, 0x4E, 0x4D, 0x4C, 0x4B, 0x4A, 0x49, 0x48, 0x47, 0x46, 0x45, 0x44, 0x43, 0x42, 0x41, 0x40},
						SpanId:  []byte{0x7F, 0x7E, 0x7D, 0x7C, 0x7B, 0x7A, 0x79, 0x78},
						Type:    tracepb.Span_Link_PARENT_LINKED_SPAN,
					},
				},
			},
		},
		{
			TraceId:      []byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E},
			SpanId:       []byte{0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F},
			ParentSpanId: []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7},
			Name:         &tracepb.TruncatableString{Value: "LocalSpan"},
			Kind:         tracepb.Span_SPAN_KIND_UNSPECIFIED,
			StartTime:    internal.TimeToTimestamp(now.Add(-15 * time.Millisecond)),
			EndTime:      internal.TimeToTimestamp(now),
			Status:       &tracepb.Status{Code: int32(0), Message: "OK"},
			Tracestate:   &tracepb.Span_Tracestate{},
		},
	}
	td := data.TraceData{Spans: spans}
	ctx := observability.ContextWithReceiverName(context.Background(), fakeReceiverName)
	const numBatches = 7
	for i := 0; i < numBatches; i++ {
		if err := te.ProcessTraceData(ctx, td); err != nil {
			t.Fatalf("Want nil got %v", err)
		}
	}

	checkValueForExporterView(t, observability.ViewExporterReceivedSpans, exporterTagName, int64(numBatches*len(spans)))
	checkValueForExporterView(t, observability.ViewExporterDroppedSpans, exporterTagName, 0)
}

func checkValueForExporterView(t *testing.T, v *view.View, exporterTagName string, value int64) {
	rows, err := view.RetrieveData(v.Name)
	if err != nil {
		t.Fatalf("Error retrieving view data.")
	}

	wantTags := []tag.Tag{
		{Key: observability.TagKeyReceiver, Value: fakeReceiverName},
		{Key: observability.TagKeyExporter, Value: exporterTagName},
	}
	// Make sure the vector is sorted by tag keys.
	sort.SliceStable(wantTags, func(i, j int) bool {
		return wantTags[i].Key.Name() < wantTags[j].Key.Name()
	})

	for _, row := range rows {
		// Make sure the vecotr is sorted by tag keys.
		sort.SliceStable(row.Tags, func(i, j int) bool {
			return row.Tags[i].Key.Name() < row.Tags[j].Key.Name()
		})
		if reflect.DeepEqual(wantTags, row.Tags) {
			sum := row.Data.(*view.SumData)
			if float64(value) != sum.Value {
				t.Fatalf("Want %v got %v", float64(value), sum.Value)
			}
			// We found the result
			return
		}
	}
	t.Fatalf("Could not find wantTags: %s in rows %v", wantTags, rows)
}
