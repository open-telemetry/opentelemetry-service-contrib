// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"
	"errors"
	"fmt"
	"net"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

var errIPNotFound = errors.New("no IP address found in the resource attributes")

// newGeoIPProcessor creates a new instance of geoIPProcessor with the specified fields.
type geoIPProcessor struct {
	providers []provider.GeoIPProvider
	fields    []string
}

func newGeoIPProcessor(fields []string) *geoIPProcessor {
	return &geoIPProcessor{
		fields: fields,
	}
}

// ipFromResourceAttributes extracts an IP address from the given resource's attributes based on the specified fields.
// It returns the first IP address if found, or an error if no valid IP address is found.
func ipFromResourceAttributes(fields []string, resource pcommon.Resource) (net.IP, error) {
	for _, field := range fields {
		if ipField, found := resource.Attributes().Get(field); found {
			ipAttribute := net.ParseIP(ipField.AsString())
			if ipAttribute == nil {
				return nil, fmt.Errorf("could not parse ip address %s", ipField.AsString())
			} else {
				return ipAttribute, nil
			}
		}
	}

	return nil, errIPNotFound
}

// geoLocation fetches geolocation information for the given IP address using the configured providers.
// It returns a set of attributes containing the geolocation data, or an error if the location could not be determined.
func (g *geoIPProcessor) geoLocation(ctx context.Context, ip net.IP) (attribute.Set, error) {
	allAttributes := attribute.EmptySet()
	for _, provider := range g.providers {
		geoAttributes, err := provider.Location(ctx, ip)
		if err != nil {
			return attribute.Set{}, err
		}
		*allAttributes = attribute.NewSet(append(allAttributes.ToSlice(), geoAttributes.ToSlice()...)...)
	}

	return *allAttributes, nil
}

// processResource processes a single resource by adding geolocation attributes based on the found IP address.
func (g *geoIPProcessor) processResource(ctx context.Context, resource pcommon.Resource) error {
	ipAddr, err := ipFromResourceAttributes(g.fields, resource)
	if err != nil {
		// TODO: log IP error not found
		if errors.Is(err, errIPNotFound) {
			return nil
		}
		return err
	}

	attributes, err := g.geoLocation(ctx, ipAddr)
	if err != nil {
		return err
	}

	for _, geoAttr := range attributes.ToSlice() {
		resource.Attributes().PutStr(string(geoAttr.Key), geoAttr.Value.AsString())
	}

	return nil
}

func (g *geoIPProcessor) processMetrics(ctx context.Context, ms pmetric.Metrics) (pmetric.Metrics, error) {
	rm := ms.ResourceMetrics()
	for i := 0; i < rm.Len(); i++ {
		err := g.processResource(ctx, rm.At(i).Resource())
		if err != nil {
			return ms, err
		}
	}
	return ms, nil
}

func (g *geoIPProcessor) processTraces(ctx context.Context, ts ptrace.Traces) (ptrace.Traces, error) {
	rt := ts.ResourceSpans()
	for i := 0; i < rt.Len(); i++ {
		err := g.processResource(ctx, rt.At(i).Resource())
		if err != nil {
			return ts, err
		}
	}
	return ts, nil
}

func (g *geoIPProcessor) processLogs(ctx context.Context, ls plog.Logs) (plog.Logs, error) {
	rl := ls.ResourceLogs()
	for i := 0; i < rl.Len(); i++ {
		err := g.processResource(ctx, rl.At(i).Resource())
		if err != nil {
			return ls, err
		}
	}
	return ls, nil
}
