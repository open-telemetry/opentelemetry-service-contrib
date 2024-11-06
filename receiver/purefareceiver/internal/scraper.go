// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal"

import (
	"context"
	"fmt"
	"net/url"
	"time"

	configutil "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

type Scraper interface {
	ToPrometheusReceiverConfig(host component.Host, fact receiver.Factory) ([]*config.ScrapeConfig, error)
}

type ScraperType string

const (
	ScraperTypeArray       ScraperType = "array"
	ScraperTypeHosts       ScraperType = "hosts"
	ScraperTypeDirectories ScraperType = "directories"
	ScraperTypePods        ScraperType = "pods"
	ScraperTypeVolumes     ScraperType = "volumes"
)

type scraper struct {
	scraperType        ScraperType
	endpoint           string
	insecureskipverify bool
	configs            []ScraperConfig
	scrapeInterval     time.Duration
	labels             model.LabelSet
}

func NewScraper(_ context.Context,
	scraperType ScraperType,
	endpoint string,
	insecureskipverify bool,
	configs []ScraperConfig,
	scrapeInterval time.Duration,
	labels model.LabelSet,
) Scraper {
	return &scraper{
		scraperType:        scraperType,
		endpoint:           endpoint,
		insecureskipverify: insecureskipverify,
		configs:            configs,
		scrapeInterval:     scrapeInterval,
		labels:             labels,
	}
}

func (h *scraper) ToPrometheusReceiverConfig(host component.Host, _ receiver.Factory) ([]*config.ScrapeConfig, error) {
	scrapeCfgs := []*config.ScrapeConfig{}

	for _, arr := range h.configs {
		u, err := url.Parse(h.endpoint)
		if err != nil {
			return nil, err
		}

		bearerToken, err := RetrieveBearerToken(arr.Auth, host.GetExtensions())
		if err != nil {
			return nil, err
		}

		httpConfig := configutil.HTTPClientConfig{}
		httpConfig.BearerToken = configutil.Secret(bearerToken)
		httpConfig.TLSConfig.InsecureSkipVerify = h.insecureskipverify

		scrapeConfig := &config.ScrapeConfig{
			HTTPClientConfig: httpConfig,
			ScrapeProtocols:  config.DefaultGlobalConfig.ScrapeProtocols,
			ScrapeInterval:   model.Duration(h.scrapeInterval),
			ScrapeTimeout:    model.Duration(h.scrapeInterval),
			JobName:          fmt.Sprintf("%s/%s/%s", "purefa", h.scraperType, arr.Address),
			HonorTimestamps:  true,
			Scheme:           u.Scheme,
			MetricsPath:      fmt.Sprintf("/metrics/%s", h.scraperType),
			Params: url.Values{
				"endpoint":  {arr.Address},
				"namespace": {"purefa"},
			},

			ServiceDiscoveryConfigs: discovery.Configs{
				&discovery.StaticConfig{
					{
						Targets: []model.LabelSet{
							{model.AddressLabel: model.LabelValue(u.Host)},
						},
						Labels: h.labels,
					},
				},
			},
		}

		scrapeCfgs = append(scrapeCfgs, scrapeConfig)
	}

	return scrapeCfgs, nil
}
