// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"fmt"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

var receiverRegex = regexp.MustCompile(`/receiver/(\w*receiver)`)

func routeWithDefaults(defaultDSType string) func(
	pcommon.Map,
	pcommon.Map,
	pcommon.Map,
	string,
	bool,
	string,
) string {
	return func(
		recordAttr pcommon.Map,
		scopeAttr pcommon.Map,
		resourceAttr pcommon.Map,
		fIndex string,
		otel bool,
		scopeName string,
	) string {
		// Order:
		// 1. read data_stream.* from attributes
		// 2. read elasticsearch.index.* from attributes
		// 3. receiver-based routing
		// 4. use default hardcoded data_stream.*
		dataset, datasetExists := getFromAttributes(dataStreamDataset, defaultDataStreamDataset, recordAttr, scopeAttr, resourceAttr)
		namespace, namespaceExists := getFromAttributes(dataStreamNamespace, defaultDataStreamNamespace, recordAttr, scopeAttr, resourceAttr)
		dataStreamMode := datasetExists || namespaceExists
		if !dataStreamMode {
			prefix, prefixExists := getFromAttributes(indexPrefix, "", resourceAttr, scopeAttr, recordAttr)
			suffix, suffixExists := getFromAttributes(indexSuffix, "", resourceAttr, scopeAttr, recordAttr)
			if prefixExists || suffixExists {
				return fmt.Sprintf("%s%s%s", prefix, fIndex, suffix)
			}
		}

		dataset = sanitizeDataStreamDataset(dataset)
		namespace = sanitizeDataStreamNamespace(namespace)

    // Receiver-based routing
		// For example, hostmetricsreceiver (or hostmetricsreceiver.otel in the OTel output mode)
		// for the scope name
		// github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper
		if submatch := receiverRegex.FindStringSubmatch(scopeName); len(submatch) > 0 {
			receiverName := submatch[1]
			dataset = receiverName
		}

		// The naming convention for datastream is expected to be "logs-[dataset].otel-[namespace]".
		// This is in order to match the built-in logs-*.otel-* index template.
		if otel {
			dataset += ".otel"
		}

		recordAttr.PutStr(dataStreamDataset, dataset)
		recordAttr.PutStr(dataStreamNamespace, namespace)
		recordAttr.PutStr(dataStreamType, defaultDSType)
		return fmt.Sprintf("%s-%s-%s", defaultDSType, dataset, namespace)
	}
}

var (
	// routeLogRecord returns the name of the index to send the log record to according to data stream routing related attributes.
	// This function may mutate record attributes.
	routeLogRecord = routeWithDefaults(defaultDataStreamTypeLogs)

	// routeDataPoint returns the name of the index to send the data point to according to data stream routing related attributes.
	// This function may mutate record attributes.
	routeDataPoint = routeWithDefaults(defaultDataStreamTypeMetrics)

	// routeSpan returns the name of the index to send the span to according to data stream routing related attributes.
	// This function may mutate record attributes.
	routeSpan = routeWithDefaults(defaultDataStreamTypeTraces)

	// routeSpanEvent returns the name of the index to send the span event to according to data stream routing related attributes.
	// This function may mutate record attributes.
	routeSpanEvent = routeWithDefaults(defaultDataStreamTypeLogs)
)
