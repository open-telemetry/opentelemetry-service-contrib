// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

var _ exporter.Metrics = (*metricExporterImp)(nil)

type metricExporterImp struct {
	loadBalancer *loadBalancer
	routingKey   routingKey

	stopped    bool
	shutdownWg sync.WaitGroup
}

func newMetricsExporter(params exporter.Settings, cfg component.Config) (*metricExporterImp, error) {
	exporterFactory := otlpexporter.NewFactory()

	lb, err := newLoadBalancer(params, cfg, func(ctx context.Context, endpoint string) (component.Component, error) {
		oCfg := buildExporterConfig(cfg.(*Config), endpoint)
		return exporterFactory.CreateMetricsExporter(ctx, params, &oCfg)
	})
	if err != nil {
		return nil, err
	}

	metricExporter := metricExporterImp{loadBalancer: lb, routingKey: svcRouting}

	switch cfg.(*Config).RoutingKey {
	case svcRoutingStr, "":
		// default case for empty routing key
		metricExporter.routingKey = svcRouting
	case resourceRoutingStr:
		metricExporter.routingKey = resourceRouting
	case metricNameRoutingStr:
		metricExporter.routingKey = metricNameRouting
	default:
		return nil, fmt.Errorf("unsupported routing_key: %q", cfg.(*Config).RoutingKey)
	}
	return &metricExporter, nil

}

func (e *metricExporterImp) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e *metricExporterImp) Start(ctx context.Context, host component.Host) error {
	return e.loadBalancer.Start(ctx, host)
}

func (e *metricExporterImp) Shutdown(ctx context.Context) error {
	err := e.loadBalancer.Shutdown(ctx)
	e.stopped = true
	e.shutdownWg.Wait()
	return err
}

func (e *metricExporterImp) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	var batches map[string]pmetric.Metrics

	switch e.routingKey {
	case svcRouting:
		var err error
		batches, err = splitMetricsByResourceServiceName(md)
		if err != nil {
			return err
		}
	case resourceRouting:
		batches = splitMetricsByResourceID(md)
	case metricNameRouting:
		batches = splitMetricsByMetricName(md)
	}

	// Now assign each batch to an exporter, and merge as we go
	metricsByExporter := map[*wrappedExporter]pmetric.Metrics{}
	exporterEndpoints := map[*wrappedExporter]string{}

	for routingID, mds := range batches {
		exp, endpoint, err := e.loadBalancer.exporterAndEndpoint([]byte(routingID))
		if err != nil {
			return err
		}

		expMetrics, ok := metricsByExporter[exp]
		if !ok {
			exp.consumeWG.Add(1)
			expMetrics = pmetric.NewMetrics()
			metricsByExporter[exp] = expMetrics
			exporterEndpoints[exp] = endpoint
		}

		metrics.Merge(expMetrics, mds)
	}

	var errs error
	for exp, mds := range metricsByExporter {
		start := time.Now()
		err := exp.ConsumeMetrics(ctx, mds)
		duration := time.Since(start)

		exp.consumeWG.Done()
		errs = multierr.Append(errs, err)

		if err == nil {
			_ = stats.RecordWithTags(
				ctx,
				[]tag.Mutator{tag.Upsert(endpointTagKey, exporterEndpoints[exp]), successTrueMutator},
				mBackendLatency.M(duration.Milliseconds()))
		} else {
			_ = stats.RecordWithTags(
				ctx,
				[]tag.Mutator{tag.Upsert(endpointTagKey, exporterEndpoints[exp]), successFalseMutator},
				mBackendLatency.M(duration.Milliseconds()))
		}
	}

	return errs
}

func splitMetricsByResourceServiceName(md pmetric.Metrics) (map[string]pmetric.Metrics, error) {
	results := map[string]pmetric.Metrics{}

	for i := 0; i < md.ResourceMetrics().Len(); i++ {
		rm := md.ResourceMetrics().At(i)

		svc, ok := rm.Resource().Attributes().Get(conventions.AttributeServiceName)
		if !ok {
			return nil, errors.New("unable to get service name")
		}

		newMD := pmetric.NewMetrics()
		rmClone := newMD.ResourceMetrics().AppendEmpty()
		rm.CopyTo(rmClone)

		key := svc.Str()
		existing, ok := results[key]
		if ok {
			metrics.Merge(existing, newMD)
		} else {
			results[key] = newMD
		}
	}

	return results, nil
}

func splitMetricsByResourceID(md pmetric.Metrics) map[string]pmetric.Metrics {
	results := map[string]pmetric.Metrics{}

	for i := 0; i < md.ResourceMetrics().Len(); i++ {
		rm := md.ResourceMetrics().At(i)

		newMD := pmetric.NewMetrics()
		rmClone := newMD.ResourceMetrics().AppendEmpty()
		rm.CopyTo(rmClone)

		key := identity.OfResource(rm.Resource()).String()
		existing, ok := results[key]
		if ok {
			metrics.Merge(existing, newMD)
		} else {
			results[key] = newMD
		}
	}

	return results
}

func splitMetricsByMetricName(md pmetric.Metrics) map[string]pmetric.Metrics {
	results := map[string]pmetric.Metrics{}

	for i := 0; i < md.ResourceMetrics().Len(); i++ {
		rm := md.ResourceMetrics().At(i)

		for j := 0; j < rm.ScopeMetrics().Len(); j++ {
			sm := rm.ScopeMetrics().At(j)

			for k := 0; k < sm.Metrics().Len(); k++ {
				m := sm.Metrics().At(k)

				newMD, mClone := cloneMetricWithoutType(rm, sm, m)
				m.CopyTo(mClone)

				key := m.Name()
				existing, ok := results[key]
				if ok {
					metrics.Merge(existing, newMD)
				} else {
					results[key] = newMD
				}
			}
		}
	}

	return results
}

func cloneMetricWithoutType(rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric) (md pmetric.Metrics, mClone pmetric.Metric) {
	md = pmetric.NewMetrics()

	rmClone := md.ResourceMetrics().AppendEmpty()
	rm.Resource().CopyTo(rmClone.Resource())
	rmClone.SetSchemaUrl(rm.SchemaUrl())

	smClone := rmClone.ScopeMetrics().AppendEmpty()
	sm.Scope().CopyTo(smClone.Scope())
	smClone.SetSchemaUrl(sm.SchemaUrl())

	mClone = smClone.Metrics().AppendEmpty()
	mClone.SetName(m.Name())
	mClone.SetDescription(m.Description())
	mClone.SetUnit(m.Unit())

	return md, mClone
}
