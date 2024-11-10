// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

const (
	// just a guess to avoid too many memory re-allocation
	initAttrCapacity = 15
)

// TODO: It might be nice to consolidate the `fromPdata` in x-ray exporter and
// `toPdata` in this receiver to a common package later

// ToTraces converts X-Ray segment (and its subsegments) to an OT ResourceSpans.
func ToTraces(rawSeg []byte, recorder telemetry.Recorder) (ptrace.Traces, int, error) {
	var seg awsxray.Segment
	err := json.Unmarshal(rawSeg, &seg)
	if err != nil {
		// return 1 as total segment (&subsegments) count
		// because we can't parse the body the UDP packet.
		return ptrace.Traces{}, 1, err
	}
	count := totalSegmentsCount(seg)
	recorder.RecordSegmentsReceived(count)

	err = seg.Validate()
	if err != nil {
		recorder.RecordSegmentsRejected(count)
		return ptrace.Traces{}, count, err
	}

	traceData := ptrace.NewTraces()
	rspanSlice := traceData.ResourceSpans()
	// ## allocate a new ptrace.ResourceSpans for the segment document
	// (potentially with embedded subsegments)
	rspan := rspanSlice.AppendEmpty()

	// ## initialize the fields in a ResourceSpans
	resource := rspan.Resource()
	// each segment (with its subsegments) is generated by one instrument
	// library so only allocate one `ScopeSpans` in the
	// `ScopeSpansSlice`.
	ils := rspan.ScopeSpans().AppendEmpty()
	ils.Spans().EnsureCapacity(count)
	spans := ils.Spans()

	// populating global attributes shared among segment and embedded subsegment(s)
	populateResource(&seg, resource)

	// recursively traverse segment and embedded subsegments
	// to populate the spans. We also need to pass in the
	// TraceID of the root segment in because embedded subsegments
	// do not have that information, but it's needed after we flatten
	// the embedded subsegment to generate independent child spans.
	_, err = segToSpans(seg, seg.TraceID, nil, spans)
	if err != nil {
		recorder.RecordSegmentsRejected(count)
		return ptrace.Traces{}, count, err
	}

	return traceData, count, nil
}

func segToSpans(seg awsxray.Segment, traceID, parentID *string, spans ptrace.SpanSlice) (ptrace.Span, error) {
	span := spans.AppendEmpty()

	err := populateSpan(&seg, traceID, parentID, span)
	if err != nil {
		return ptrace.Span{}, err
	}

	var populatedChildSpan ptrace.Span
	for _, s := range seg.Subsegments {
		populatedChildSpan, err = segToSpans(s,
			traceID, seg.ID,
			spans)
		if err != nil {
			return ptrace.Span{}, err
		}

		if seg.Cause != nil &&
			populatedChildSpan.Status().Code() != ptrace.StatusCodeUnset {
			// if seg.Cause is not nil, then one of the subsegments must contain a
			// HTTP error code. Also, span.Status().Code() is already
			// set to `StatusCodeUnknownError` by `addCause()` in
			// `populateSpan()` above, so here we are just trying to figure out
			// whether we can get an even more specific error code.

			if span.Status().Code() == ptrace.StatusCodeError {
				// update the error code to a possibly more specific code
				span.Status().SetCode(populatedChildSpan.Status().Code())
			}
		}
	}

	return span, nil
}

func populateSpan(seg *awsxray.Segment, traceID, parentID *string, span ptrace.Span) error {
	attrs := span.Attributes()
	attrs.Clear()
	attrs.EnsureCapacity(initAttrCapacity)

	err := addNameAndNamespace(seg, span)
	if err != nil {
		return err
	}

	// decode trace id
	var traceIDBytes [16]byte
	if seg.TraceID == nil {
		// if seg.TraceID is nil, then `seg` must be an embedded subsegment.
		traceIDBytes, err = decodeXRayTraceID(traceID)
		if err != nil {
			return err
		}
	} else {
		traceIDBytes, err = decodeXRayTraceID(seg.TraceID)
		if err != nil {
			return err
		}
	}

	// decode parent id
	var parentIDBytes [8]byte
	if parentID != nil {
		parentIDBytes, err = decodeXRaySpanID(parentID)
		if err != nil {
			return err
		}
	} else if seg.ParentID != nil {
		parentIDBytes, err = decodeXRaySpanID(seg.ParentID)
		if err != nil {
			return err
		}
	}

	// decode span id
	spanIDBytes, err := decodeXRaySpanID(seg.ID)
	if err != nil {
		return err
	}

	span.SetTraceID(traceIDBytes)
	span.SetSpanID(spanIDBytes)

	if parentIDBytes != [8]byte{} {
		span.SetParentSpanID(parentIDBytes)
	} else {
		span.SetKind(ptrace.SpanKindServer)
	}

	addStartTime(seg.StartTime, span)
	addEndTime(seg.EndTime, span)
	addBool(seg.InProgress, awsxray.AWSXRayInProgressAttribute, attrs)
	addString(seg.User, conventions.AttributeEnduserID, attrs)

	addHTTP(seg, span)
	addCause(seg, span)
	addAWSToSpan(seg.AWS, attrs)
	err = addSQLToSpan(seg.SQL, attrs)
	if err != nil {
		return err
	}

	addBool(seg.Traced, awsxray.AWSXRayTracedAttribute, attrs)

	addAnnotations(seg.Annotations, attrs)
	return addMetadata(seg.Metadata, attrs)
}

func populateResource(seg *awsxray.Segment, rs pcommon.Resource) {
	// allocate a new attribute map within the Resource in the ptrace.ResourceSpans allocated above
	attrs := rs.Attributes()
	attrs.Clear()
	attrs.EnsureCapacity(initAttrCapacity)

	addString(seg.Name, conventions.AttributeServiceName, attrs)

	addAWSToResource(seg.AWS, attrs)
	addSdkToResource(seg, attrs)
	if seg.Service != nil {
		addString(seg.Service.Version, conventions.AttributeServiceVersion, attrs)
	}

	addString(seg.ResourceARN, awsxray.AWSXRayResourceARNAttribute, attrs)
}

func totalSegmentsCount(seg awsxray.Segment) int {
	subsegmentCount := 0
	for _, s := range seg.Subsegments {
		subsegmentCount += totalSegmentsCount(s)
	}

	return 1 + subsegmentCount
}

/*
decodeXRayTraceID decodes the traceid from xraysdk
one example of xray format: "1-5f84c7a1-e7d1852db8c4fd35d88bf49a"
decodeXRayTraceID transfers it to "5f84c7a1e7d1852db8c4fd35d88bf49a" and decode it from hex
*/
func decodeXRayTraceID(traceID *string) ([16]byte, error) {
	tid := [16]byte{}

	if traceID == nil {
		return tid, errors.New("traceID is null")
	}
	if len(*traceID) < 35 {
		return tid, errors.New("traceID length is wrong")
	}
	traceIDtoBeDecoded := (*traceID)[2:10] + (*traceID)[11:]

	_, err := hex.Decode(tid[:], []byte(traceIDtoBeDecoded))
	return tid, err
}

// decodeXRaySpanID decodes the spanid from xraysdk
func decodeXRaySpanID(spanID *string) ([8]byte, error) {
	sid := [8]byte{}
	if spanID == nil {
		return sid, errors.New("spanid is null")
	}
	if len(*spanID) != 16 {
		return sid, errors.New("spanID length is wrong")
	}
	_, err := hex.Decode(sid[:], []byte(*spanID))
	return sid, err
}
