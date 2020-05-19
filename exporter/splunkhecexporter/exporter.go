// Copyright 2020, OpenTelemetry Authors
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

package splunkhecexporter

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.opentelemetry.io/collector/consumer/pdatautil"
	"go.opentelemetry.io/collector/obsreport"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	idleConnTimeout     = 30 * time.Second
	tlsHandshakeTimeout = 10 * time.Second
	dialerTimeout       = 30 * time.Second
	dialerKeepAlive     = 30 * time.Second
)

type splunkExporter struct {
	pushMetricsData func(ctx context.Context, md consumerdata.MetricsData) (droppedTimeSeries int, err error)
	stop            func(ctx context.Context) (err error)
}

type exporterOptions struct {
	url   *url.URL
	token string
}

// New returns a new Splunk exporter.
func New(
	config *Config,
	logger *zap.Logger,
) (component.MetricsExporterOld, error) {

	if config == nil {
		return nil, errors.New("nil config")
	}

	options, err := config.getOptionsFromConfig()
	if err != nil {
		return nil,
			fmt.Errorf("failed to process %q config: %v", config.Name(), err)
	}

	logger.Info("Splunk Config", zap.String("url", options.url.String()))

	if config.Name() == "" {
		config.SetType(typeStr)
		config.SetName(typeStr)
	}

	client := &client{
		url: options.url,
		client: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   dialerTimeout,
					KeepAlive: dialerKeepAlive,
				}).DialContext,
				MaxIdleConns:        int(config.MaxConnections),
				MaxIdleConnsPerHost: int(config.MaxConnections),
				IdleConnTimeout:     idleConnTimeout,
				TLSHandshakeTimeout: tlsHandshakeTimeout,
			},
		},
		logger: logger,
		zippers: sync.Pool{New: func() interface{} {
			return gzip.NewWriter(nil)
		}},
	}

	return splunkExporter{
		pushMetricsData: client.pushMetricsData,
		stop:            client.stop,
	}, nil
}

func (se splunkExporter) Start(context.Context, component.Host) error {
	return nil
}

func (se splunkExporter) Shutdown(ctxt context.Context) error {
	return se.stop(ctxt)
}

func (se splunkExporter) ConsumeMetricsData(ctx context.Context, md consumerdata.MetricsData) error {
	ctx = obsreport.StartMetricsExportOp(ctx, typeStr)
	numDroppedTimeSeries, err := se.pushMetricsData(ctx, md)

	numReceivedTimeSeries, numPoints := pdatautil.TimeseriesAndPointCount(md)

	obsreport.EndMetricsExportOp(ctx, numPoints, numReceivedTimeSeries, numDroppedTimeSeries, err)
	return err
}
