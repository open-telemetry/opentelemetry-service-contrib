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

package zipkinreceiver

import (
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/census-instrumentation/opencensus-service/translator/trace"

	"go.opencensus.io/trace"

	zipkinmodel "github.com/openzipkin/zipkin-go/model"
	zipkinproto "github.com/openzipkin/zipkin-go/proto/v2"

	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	agenttracepb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/trace/v1"
	tracepb "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"

	"github.com/census-instrumentation/opencensus-service/data"
	"github.com/census-instrumentation/opencensus-service/internal"
	"github.com/census-instrumentation/opencensus-service/receiver"
)

// ZipkinReceiver type is used to handle spans received in the Zipkin format.
type ZipkinReceiver struct {
	// mu protects the fields of this struct
	mu sync.Mutex

	// addr is the address onto which the HTTP server will be bound
	addr string

	spanSink receiver.TraceReceiverSink

	startOnce sync.Once
	stopOnce  sync.Once
	server    *http.Server
}

var _ receiver.TraceReceiver = (*ZipkinReceiver)(nil)
var _ http.Handler = (*ZipkinReceiver)(nil)

// New creates a new zipkinreceiver.ZipkinReceiver reference.
func New(address string) (*ZipkinReceiver, error) {
	zr := &ZipkinReceiver{addr: address}
	return zr, nil
}

const defaultAddress = ":9411"

func (zr *ZipkinReceiver) address() string {
	addr := zr.addr
	if addr == "" {
		addr = defaultAddress
	}
	return addr
}

var (
	errAlreadyStarted = errors.New("already started")
	errAlreadyStopped = errors.New("already stopped")
)

// StartTraceReception spins up the receiver's HTTP server and makes the receiver start its processing.
func (zr *ZipkinReceiver) StartTraceReception(ctx context.Context, spanSink receiver.TraceReceiverSink) error {
	zr.mu.Lock()
	defer zr.mu.Unlock()

	var err = errAlreadyStarted

	zr.startOnce.Do(func() {
		ln, lerr := net.Listen("tcp", zr.address())
		if lerr != nil {
			err = lerr
			return
		}

		server := &http.Server{Handler: zr}
		go func() {
			_ = server.Serve(ln)
		}()

		zr.spanSink = spanSink
		zr.server = server

		err = nil
	})

	return err
}

// v1ToTraceSpans parses Zipkin v1 JSON traces and converts them to OpenCensus Proto spans.
func (zr *ZipkinReceiver) v1ToTraceSpans(blob []byte) (reqs []*agenttracepb.ExportTraceServiceRequest, err error) {
	return tracetranslator.ZipkinV1JSONBatchToOCProto(blob)
}

// v2ToTraceSpans parses Zipkin v2 JSON or Protobuf traces and converts them to OpenCensus Proto spans.
func (zr *ZipkinReceiver) v2ToTraceSpans(blob []byte, hdr http.Header) (reqs []*agenttracepb.ExportTraceServiceRequest, err error) {
	// This flag's reference is from:
	//      https://github.com/openzipkin/zipkin-go/blob/3793c981d4f621c0e3eb1457acffa2c1cc591384/proto/v2/zipkin.proto#L154
	debugWasSet := hdr.Get("X-B3-Flags") == "1"

	var zipkinSpans []*zipkinmodel.SpanModel

	// Zipkin can send protobuf via http
	switch hdr.Get("Content-Type") {
	// TODO: (@odeke-em) record the unique types of Content-Type uploads
	case "application/x-protobuf":
		zipkinSpans, err = zipkinproto.ParseSpans(blob, debugWasSet)

	default: // By default, we'll assume using JSON
		zipkinSpans, err = zr.deserializeFromJSON(blob, debugWasSet)
	}

	if err != nil {
		return nil, err
	}

	// *commonpb.Node instances have unique addresses hence
	// for grouping within a map, we'll use the .String() value
	byNodeGrouping := make(map[string][]*tracepb.Span)
	uniqueNodes := make([]*commonpb.Node, 0, len(zipkinSpans))
	// Now translate them into tracepb.Span
	for _, zspan := range zipkinSpans {
		span, node, err := zipkinSpanToTraceSpan(zspan)
		// TODO:(@odeke-em) record errors
		if err == nil && span != nil {
			key := node.String()
			if _, alreadyAdded := byNodeGrouping[key]; !alreadyAdded {
				uniqueNodes = append(uniqueNodes, node)
			}
			byNodeGrouping[key] = append(byNodeGrouping[key], span)
		}
	}

	for _, node := range uniqueNodes {
		key := node.String()
		spans := byNodeGrouping[key]
		if len(spans) == 0 {
			// Should never happen but nonetheless be cautious
			// not to send blank spans.
			continue
		}
		reqs = append(reqs, &agenttracepb.ExportTraceServiceRequest{
			Node:  node,
			Spans: spans,
		})
		delete(byNodeGrouping, key)
	}

	return reqs, nil
}

func (zr *ZipkinReceiver) deserializeFromJSON(jsonBlob []byte, debugWasSet bool) (zs []*zipkinmodel.SpanModel, err error) {
	if err = json.Unmarshal(jsonBlob, &zs); err != nil {
		return nil, err
	}
	return zs, nil
}

// StopTraceReception tells the receiver that should stop reception,
// giving it a chance to perform any necessary clean-up and shutting down
// its HTTP server.
func (zr *ZipkinReceiver) StopTraceReception(ctx context.Context) error {
	var err = errAlreadyStopped
	zr.stopOnce.Do(func() {
		err = zr.server.Close()
	})
	return err
}

// processBodyIfNecessary checks the "Content-Encoding" HTTP header and if
// a compression such as "gzip", "deflate", "zlib", is found, the body will
// be uncompressed accordingly or return the body untouched if otherwise.
// Clients such as Zipkin-Java do this behavior e.g.
//    send "Content-Encoding":"gzip" of the JSON content.
func processBodyIfNecessary(req *http.Request) io.Reader {
	switch req.Header.Get("Content-Encoding") {
	default:
		return req.Body

	case "gzip":
		return gunzippedBodyIfPossible(req.Body)

	case "deflate", "zlib":
		return zlibUncompressedbody(req.Body)
	}
}

func gunzippedBodyIfPossible(r io.Reader) io.Reader {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		// Just return the old body as was
		return r
	}
	return gzr
}

func zlibUncompressedbody(r io.Reader) io.Reader {
	zr, err := zlib.NewReader(r)
	if err != nil {
		// Just return the old body as was
		return r
	}
	return zr
}

// The ZipkinReceiver receives spans from endpoint /api/v2 as JSON,
// unmarshals them and sends them along to the spansink.Sink.
func (zr *ZipkinReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Trace this method
	ctx, span := trace.StartSpan(context.Background(), "ZipkinReceiver.Export")
	defer span.End()

	// If the starting RPC has a parent span, then add it as a parent link.
	parentCtx := r.Context()
	internal.SetParentLink(parentCtx, span)

	pr := processBodyIfNecessary(r)
	slurp, err := ioutil.ReadAll(pr)
	if c, ok := pr.(io.Closer); ok {
		_ = c.Close()
	}
	_ = r.Body.Close()

	// Now deserialize and process the spans.
	asZipkinv1 := r.URL != nil && strings.Contains(r.URL.Path, "api/v1/spans")

	var ereqs []*agenttracepb.ExportTraceServiceRequest

	var receiverNameTag string
	if asZipkinv1 {
		ereqs, err = zr.v1ToTraceSpans(slurp)
		receiverNameTag = "zipkinV1"
	} else {
		ereqs, err = zr.v2ToTraceSpans(slurp, r.Header)
		receiverNameTag = "zipkinV2"
	}

	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: err.Error(),
		})
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	spansMetricsFn := internal.NewReceivedSpansRecorderStreaming(ctx, receiverNameTag)
	// Now translate them into TraceData
	for _, ereq := range ereqs {
		zr.spanSink.ReceiveTraceData(ctx, data.TraceData{Node: ereq.Node, Spans: ereq.Spans})
		// We MUST unconditionally record metrics from this reception.
		spansMetricsFn(ereq.Node, ereq.Spans)
	}

	// Finally send back the response "Accepted" as
	// required at https://zipkin.io/zipkin-api/#/default/post_spans
	w.WriteHeader(http.StatusAccepted)
}

var (
	errNilZipkinSpan = errors.New("non-nil Zipkin span expected")
	errZeroTraceID   = errors.New("trace id is zero")
	errZeroID        = errors.New("id is zero")
)

func zTraceIDToOCProtoTraceID(zTraceID zipkinmodel.TraceID) ([]byte, error) {
	if zTraceID.High == 0 && zTraceID.Low == 0 {
		return nil, errZeroTraceID
	}
	traceID := make([]byte, 16)
	binary.BigEndian.PutUint64(traceID[:8], uint64(zTraceID.High))
	binary.BigEndian.PutUint64(traceID[8:], uint64(zTraceID.Low))
	return traceID, nil
}

func zSpanIDToOCProtoSpanID(id zipkinmodel.ID) ([]byte, error) {
	if id == 0 {
		return nil, errZeroID
	}
	spanID := make([]byte, 8)
	binary.BigEndian.PutUint64(spanID, uint64(id))
	return spanID, nil
}

func zipkinSpanToTraceSpan(zs *zipkinmodel.SpanModel) (*tracepb.Span, *commonpb.Node, error) {
	if zs == nil {
		return nil, nil, errNilZipkinSpan
	}

	node := nodeFromZipkinEndpoints(zs)
	traceID, err := zTraceIDToOCProtoTraceID(zs.TraceID)
	if err != nil {
		return nil, node, fmt.Errorf("TraceID: %v", err)
	}
	spanID, err := zSpanIDToOCProtoSpanID(zs.ID)
	if err != nil {
		return nil, node, fmt.Errorf("SpanID: %v", err)
	}
	var parentSpanID []byte
	if zs.ParentID != nil {
		parentSpanID, err = zSpanIDToOCProtoSpanID(*zs.ParentID)
		if err != nil {
			return nil, node, fmt.Errorf("ParentSpanID: %v", err)
		}
	}

	pbs := &tracepb.Span{
		TraceId:      traceID,
		SpanId:       spanID,
		ParentSpanId: parentSpanID,
		Name:         &tracepb.TruncatableString{Value: zs.Name},
		StartTime:    internal.TimeToTimestamp(zs.Timestamp),
		EndTime:      internal.TimeToTimestamp(zs.Timestamp.Add(zs.Duration)),
		Kind:         zipkinSpanKindToProtoSpanKind(zs.Kind),
		Status:       extractProtoStatus(zs),
		Attributes:   zipkinTagsToTraceAttributes(zs.Tags),
		TimeEvents:   zipkinAnnotationsToProtoTimeEvents(zs.Annotations),
	}

	return pbs, node, nil
}

func nodeFromZipkinEndpoints(zs *zipkinmodel.SpanModel) *commonpb.Node {
	if zs.LocalEndpoint == nil && zs.RemoteEndpoint == nil {
		return nil
	}

	node := new(commonpb.Node)

	// Retrieve and make use of the local endpoint
	if lep := zs.LocalEndpoint; lep != nil {
		node.ServiceInfo = &commonpb.ServiceInfo{
			Name: lep.ServiceName,
		}
		node.Attributes = zipkinEndpointIntoAttributes(lep, node.Attributes, func(s string) string { return s })
	}

	// Retrieve and make use of the remote endpoint
	if rep := zs.RemoteEndpoint; rep != nil {
		// For remoteEndpoint, our goal is to prefix its fields with "zipkin.remoteEndpoint."
		// For example becoming:
		// {
		//      "zipkin.remoteEndpoint.ipv4": "192.168.99.101",
		//      "zipkin.remoteEndpoint.port": "9000"
		//      "zipkin.remoteEndpoint.serviceName": "backend",
		// }
		node.Attributes = zipkinEndpointIntoAttributes(rep, node.Attributes, func(s string) string {
			return "zipkin.remoteEndpoint." + s
		})
	}
	return node
}

var blankIP net.IP

func zipkinEndpointIntoAttributes(ep *zipkinmodel.Endpoint, into map[string]string, prefixFunc func(string) string) map[string]string {
	if into == nil {
		into = make(map[string]string)
	}
	if ep.IPv4 != nil && !ep.IPv4.Equal(blankIP) {
		into[prefixFunc("ipv4")] = ep.IPv4.String()
	}
	if ep.IPv6 != nil && !ep.IPv6.Equal(blankIP) {
		into[prefixFunc("ipv6")] = ep.IPv6.String()
	}
	if ep.Port > 0 {
		into[prefixFunc("port")] = fmt.Sprintf("%d", ep.Port)
	}
	if serviceName := ep.ServiceName; serviceName != "" {
		into[prefixFunc("serviceName")] = serviceName
	}
	return into
}

const statusCodeUnknown = 2

func extractProtoStatus(zs *zipkinmodel.SpanModel) *tracepb.Status {
	// The status is stored with the "error" key
	// See https://github.com/census-instrumentation/opencensus-go/blob/1eb9a13c7dd02141e065a665f6bf5c99a090a16a/exporter/zipkin/zipkin.go#L160-L165
	if zs == nil || len(zs.Tags) == 0 {
		return nil
	}
	canonicalCodeStr := zs.Tags["error"]
	message := zs.Tags["opencensus.status_description"]
	if message == "" && canonicalCodeStr == "" {
		return nil
	}
	code, set := canonicalCodesMap[canonicalCodeStr]
	if !set {
		// If not status code was set, then we should use UNKNOWN
		code = statusCodeUnknown
	}
	return &tracepb.Status{
		Message: message,
		Code:    code,
	}
}

var canonicalCodesMap = map[string]int32{
	// https://github.com/googleapis/googleapis/blob/bee79fbe03254a35db125dc6d2f1e9b752b390fe/google/rpc/code.proto#L33-L186
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
	"UNAUTHENTICATED":     16,
}

func zipkinSpanKindToProtoSpanKind(skind zipkinmodel.Kind) tracepb.Span_SpanKind {
	switch strings.ToUpper(string(skind)) {
	case "CLIENT":
		return tracepb.Span_CLIENT
	case "SERVER":
		return tracepb.Span_SERVER
	default:
		return tracepb.Span_SPAN_KIND_UNSPECIFIED
	}
}

func zipkinAnnotationsToProtoTimeEvents(zas []zipkinmodel.Annotation) *tracepb.Span_TimeEvents {
	if len(zas) == 0 {
		return nil
	}
	tevs := make([]*tracepb.Span_TimeEvent, 0, len(zas))
	for _, za := range zas {
		if tev := zipkinAnnotationToProtoAnnotation(za); tev != nil {
			tevs = append(tevs, tev)
		}
	}
	if len(tevs) == 0 {
		return nil
	}
	return &tracepb.Span_TimeEvents{
		TimeEvent: tevs,
	}
}

var blankAnnotation zipkinmodel.Annotation

func zipkinAnnotationToProtoAnnotation(zas zipkinmodel.Annotation) *tracepb.Span_TimeEvent {
	if zas == blankAnnotation {
		return nil
	}
	return &tracepb.Span_TimeEvent{
		Time: internal.TimeToTimestamp(zas.Timestamp),
		Value: &tracepb.Span_TimeEvent_Annotation_{
			Annotation: &tracepb.Span_TimeEvent_Annotation{
				Description: &tracepb.TruncatableString{Value: zas.Value},
			},
		},
	}
}

func zipkinTagsToTraceAttributes(tags map[string]string) *tracepb.Span_Attributes {
	if len(tags) == 0 {
		return nil
	}

	amap := make(map[string]*tracepb.AttributeValue, len(tags))
	for key, value := range tags {
		// We did a translation from "boolean" to "string"
		// in OpenCensus-Go's Zipkin exporter as per
		// https://github.com/census-instrumentation/opencensus-go/blob/1eb9a13c7dd02141e065a665f6bf5c99a090a16a/exporter/zipkin/zipkin.go#L138-L155
		switch value {
		case "true", "false":
			amap[key] = &tracepb.AttributeValue{
				Value: &tracepb.AttributeValue_BoolValue{BoolValue: value == "true"},
			}
		default:
			amap[key] = &tracepb.AttributeValue{
				Value: &tracepb.AttributeValue_StringValue{
					StringValue: &tracepb.TruncatableString{Value: value},
				},
			}
		}
	}
	return &tracepb.Span_Attributes{AttributeMap: amap}
}

func setIfNonEmpty(key, value string, dest map[string]string) {
	if value != "" {
		dest[key] = value
	}
}
