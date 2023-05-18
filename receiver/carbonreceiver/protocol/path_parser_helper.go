// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package protocol // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
)

// PathParser implements the code needed to handle only the <metric_path> part of
// a Carbon metric line:
//
//	<metric_path> <metric_value> <metric_timestamp>
//
// See https://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol,
// for more information.
//
// The type PathParserHelper implements the common code for parsers that differ
// only by the way that they handle the <metric_path>.
type PathParser interface {
	// ParsePath parses the <metric_path> of a Carbon line (see Parse function
	// for description of the full line). The results of parsing the path are
	// stored on the parsedPath struct. Implementers of the interface can assume
	// that the PathParserHelper will never pass nil when calling this method.
	ParsePath(path string, parsedPath *ParsedPath) error
}

// ParsedPath holds the result of parsing the <metric_path> with the ParsePath
// method on the PathParser interface.
type ParsedPath struct {
	// MetricName extracted/generated by the parser.
	MetricName string
	// LabelKeys extracted/generated by the parser.
	LabelKeys []*metricspb.LabelKey
	// LabelValues extracted/generated by the parser.
	LabelValues []*metricspb.LabelValue
	// MetricType instructs the helper to generate the metric as the specified
	// TargetMetricType.
	MetricType TargetMetricType
}

type TargetMetricType string

// Values for enum TargetMetricType.
const (
	DefaultMetricType    = TargetMetricType("")
	GaugeMetricType      = TargetMetricType("gauge")
	CumulativeMetricType = TargetMetricType("cumulative")
)

// PathParserHelper implements the common code to parse a Carbon line taking a
// PathParser to implement a full parser.
type PathParserHelper struct {
	pathParser PathParser
}

var _ Parser = (*PathParserHelper)(nil)

// NewParser creates a new Parser instance that receives plaintext
// Carbon data.
func NewParser(pathParser PathParser) (Parser, error) {
	if pathParser == nil {
		return nil, errors.New("nil pathParser")
	}
	return &PathParserHelper{
		pathParser: pathParser,
	}, nil
}

// Parse receives the string with plaintext data, aka line, in the Carbon
// format and transforms it to the collector metric format. See
// https://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol.
//
// The expected line is a text line in the following format:
//
//	"<metric_path> <metric_value> <metric_timestamp>"
//
// The <metric_path> is where there are variations that require selection
// of specialized parsers to handle them, but include the metric name and
// labels/dimensions for the metric.
//
// The <metric_value> is the textual representation of the metric value.
//
// The <metric_timestamp> is the Unix time text of when the measurement was
// made.
func (pph *PathParserHelper) Parse(line string) (*metricspb.Metric, error) {
	parts := strings.SplitN(line, " ", 4)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid carbon metric [%s]", line)
	}

	path := parts[0]
	valueStr := parts[1]
	timestampStr := parts[2]

	parsedPath := ParsedPath{}
	err := pph.pathParser.ParsePath(path, &parsedPath)
	if err != nil {
		return nil, fmt.Errorf("invalid carbon metric [%s]: %w", line, err)
	}

	unixTime, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid carbon metric time [%s]: %w", line, err)
	}

	var metricType metricspb.MetricDescriptor_Type
	point := metricspb.Point{
		Timestamp: convertUnixSec(unixTime),
	}
	intVal, err := strconv.ParseInt(valueStr, 10, 64)
	if err == nil {
		if parsedPath.MetricType == CumulativeMetricType {
			metricType = metricspb.MetricDescriptor_CUMULATIVE_INT64
		} else {
			metricType = metricspb.MetricDescriptor_GAUGE_INT64
		}
		point.Value = &metricspb.Point_Int64Value{Int64Value: intVal}
	} else {
		dblVal, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid carbon metric value [%s]: %w", line, err)
		}
		if parsedPath.MetricType == CumulativeMetricType {
			metricType = metricspb.MetricDescriptor_CUMULATIVE_DOUBLE
		} else {
			metricType = metricspb.MetricDescriptor_GAUGE_DOUBLE
		}
		point.Value = &metricspb.Point_DoubleValue{DoubleValue: dblVal}
	}

	metric := buildMetricForSinglePoint(
		parsedPath.MetricName,
		metricType,
		parsedPath.LabelKeys,
		parsedPath.LabelValues,
		&point)
	return metric, nil
}
