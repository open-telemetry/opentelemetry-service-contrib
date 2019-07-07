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

package prometheusexporter

import (
	"net"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/config/configerror"
	"github.com/open-telemetry/opentelemetry-service/config/configmodels"
	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/exporter"
	"github.com/orijtech/prometheus-go-metrics-exporter"
)

var _ = exporter.RegisterFactory(&factory{})

const (
	// The value of "type" key in configuration.
	typeStr = "prometheus"
)

// factory is the factory for Prometheus exporter.
type factory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for exporter.
func (f *factory) CreateDefaultConfig() configmodels.Exporter {
	return &ConfigV2{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		ConstLabels: map[string]string{},
	}
}

// CreateTraceExporter creates a trace exporter based on this config.
func (f *factory) CreateTraceExporter(logger *zap.Logger, config configmodels.Exporter) (consumer.TraceConsumer, exporter.StopFunc, error) {
	return nil, nil, configerror.ErrDataTypeIsNotSupported
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *factory) CreateMetricsExporter(logger *zap.Logger, cfg configmodels.Exporter) (consumer.MetricsConsumer, exporter.StopFunc, error) {
	pcfg := cfg.(*ConfigV2)

	addr := strings.TrimSpace(pcfg.Endpoint)
	if addr == "" {
		return nil, nil, errBlankPrometheusAddress
	}

	opts := prometheus.Options{
		Namespace:   pcfg.Namespace,
		ConstLabels: pcfg.ConstLabels,
	}
	pe, err := prometheus.New(opts)
	if err != nil {
		return nil, nil, err
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	// The Prometheus metrics exporter has to run on the provided address
	// as a server that'll be scraped by Prometheus.
	mux := http.NewServeMux()
	mux.Handle("/metrics", pe)

	srv := &http.Server{Handler: mux}
	go func() {
		_ = srv.Serve(ln)
	}()

	pexp := &prometheusExporter{exporter: pe}

	return pexp, ln.Close, nil
}
