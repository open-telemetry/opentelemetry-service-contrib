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

package loggingexporter

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exporterhelper"
)

type traceDataBuffer struct {
	str strings.Builder
}

func (b *traceDataBuffer) logEntry(format string, a ...interface{}) {
	b.str.WriteString(fmt.Sprintf(format, a...))
	b.str.WriteString("\n")
}

func (b *traceDataBuffer) logAttr(label string, value string) {
	b.logEntry("    %-15s: %s", label, value)
}

func (b *traceDataBuffer) logMap(label string, data map[string]string) {
	if len(data) == 0 {
		return
	}

	b.logEntry("%s:", label)
	for label, value := range data {
		b.logEntry("     -> %s: %s", label, value)
	}
}

type loggingExporter struct {
	logger *zap.Logger
	name   zap.Field
	debug  bool
}

func (s *loggingExporter) pushTraceData(
	ctx context.Context,
	td consumerdata.TraceData,
) (int, error) {
	buf := traceDataBuffer{}

	resourceInfo := ""
	nodesInfo := ""

	if td.Resource != nil {
		resourceInfo = fmt.Sprintf(", resource \"%s\" (%d labels)", td.Resource.Type, len(td.Resource.Labels))
	}

	if td.Node != nil {
		nodesInfo = fmt.Sprintf(", node service: %s", td.Node.ServiceInfo.Name)
	}

	buf.logEntry("TraceData with %d spans%s%s", len(td.Spans), nodesInfo, resourceInfo)

	if s.debug {
		if td.Resource != nil {
			buf.logMap("Resource labels", td.Resource.Labels)
		}

		if td.Node != nil {
			id := td.Node.Identifier
			if id != nil {
				buf.logEntry("%20s: %s", "HostName", id.HostName)
				buf.logEntry("%20s: %d", "PID", id.Pid)
			}
			li := td.Node.LibraryInfo
			if li != nil {
				buf.logEntry("%20s: %s", "Library language", li.Language.String())
				buf.logEntry("%20s: %s", "Core library version", li.CoreLibraryVersion)
				buf.logEntry("%20s: %s", "Exporter version", li.ExporterVersion)
			}
			buf.logMap("Node attributes", td.Node.Attributes)
		}
	}

	s.logger.Info(buf.str.String(), s.name)

	if s.debug {
		for i, span := range td.Spans {
			buf = traceDataBuffer{}
			buf.logEntry("Span #%d", i)
			if span == nil {
				buf.logEntry("* Empty span")
				continue
			}

			buf.logAttr("Trace ID", hex.EncodeToString(span.TraceId))
			buf.logAttr("Parent ID", hex.EncodeToString(span.ParentSpanId))
			buf.logAttr("ID", hex.EncodeToString(span.SpanId))
			buf.logAttr("Name", span.Name.Value)
			buf.logAttr("Kind", span.Kind.String())
			buf.logAttr("Start time", span.StartTime.String())
			buf.logAttr("End time", span.EndTime.String())
			if span.Status != nil {
				buf.logAttr("Status code", strconv.Itoa(int(span.Status.Code)))
				buf.logAttr("Status message", span.Status.Message)
			}

			if span.Attributes != nil {
				buf.logAttr("Span attributes", "")
				for attr, value := range span.Attributes.AttributeMap {
					v := ""
					ts := value.GetStringValue()

					if ts != nil {
						v = ts.Value
					} else {
						// For other types, just use the proto compact form rather than digging into series of checks
						v = value.String()
					}

					buf.logEntry("         -> %s: %s", attr, v)
				}
			}

			s.logger.Debug(buf.str.String(), s.name)
		}
	}

	return 0, nil
}

// NewTraceExporter creates an exporter.TraceExporter that just drops the
// received data and logs debugging messages.
func NewTraceExporter(config configmodels.Exporter, level string, logger *zap.Logger) (exporter.TraceExporter, error) {
	s := &loggingExporter{
		debug:  level == "debug",
		name:   zap.String("exporter", config.Name()),
		logger: logger,
	}

	return exporterhelper.NewTraceExporter(
		config,
		s.pushTraceData,
		exporterhelper.WithShutdown(logger.Sync),
	)
}

// NewMetricsExporter creates an exporter.MetricsExporter that just drops the
// received data and logs debugging messages.
func NewMetricsExporter(config configmodels.Exporter, logger *zap.Logger) (exporter.MetricsExporter, error) {
	typeLog := zap.String("type", config.Type())
	nameLog := zap.String("name", config.Name())
	return exporterhelper.NewMetricsExporter(
		config,
		func(ctx context.Context, md consumerdata.MetricsData) (int, error) {
			logger.Info("MetricsExporter", typeLog, nameLog, zap.Int("#metrics", len(md.Metrics)))
			// TODO: Add ability to record the received data
			return 0, nil
		},
		exporterhelper.WithShutdown(logger.Sync),
	)
}
