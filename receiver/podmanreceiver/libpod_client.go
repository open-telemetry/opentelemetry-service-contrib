// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows
// +build !windows

package podmanreceiver // import "github.com/asserts/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

var (
	errNoStatsFound = fmt.Errorf("No stats found")
)

type libpodClient struct {
	conn     *http.Client
	endpoint string
}

func newLibpodClient(logger *zap.Logger, cfg *Config) (PodmanClient, error) {
	connection, err := newPodmanConnection(logger, cfg.Endpoint, cfg.SSHKey, cfg.SSHPassphrase)
	if err != nil {
		return nil, err
	}
	c := &libpodClient{
		conn:     connection,
		endpoint: fmt.Sprintf("http://d/v%s/libpod", cfg.APIVersion),
	}
	return c, nil
}

func (c *libpodClient) request(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	if len(params) > 0 {
		req.URL.RawQuery = params.Encode()
	}

	return c.conn.Do(req)
}

func (c *libpodClient) stats(ctx context.Context, options url.Values) ([]containerStats, error) {
	resp, err := c.request(ctx, "/containers/stats", options)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	report := &containerStatsReport{}
	err = json.Unmarshal(bytes, report)
	if err != nil {
		return nil, err
	}
	if report.Error.Message != "" {
		return nil, errors.New(report.Error.Message)
	} else if report.Stats == nil {
		return nil, errNoStatsFound
	}

	return report.Stats, nil
}

func (c *libpodClient) list(ctx context.Context, options url.Values) ([]container, error) {
	resp, err := c.request(ctx, "/containers/json", options)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var report []container
	err = json.Unmarshal(bytes, &report)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (c *libpodClient) ping(ctx context.Context) error {
	resp, err := c.request(ctx, "/_ping", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ping response was %d", resp.StatusCode)
	}
	return nil
}

// events returns a stream of events. It's up to the caller to close the stream by canceling the context.
func (c *libpodClient) events(ctx context.Context, options url.Values) (<-chan event, <-chan error) {
	events := make(chan event)
	errs := make(chan error, 1)

	started := make(chan struct{})
	go func() {
		defer close(errs)

		resp, err := c.request(ctx, "/events", options)
		if err != nil {
			close(started)
			errs <- err
			return
		}
		defer resp.Body.Close()

		dec := json.NewDecoder(resp.Body)
		close(started)
		for {
			var e event
			select {
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			default:
				err := dec.Decode(&e)
				if err != nil {
					errs <- err
					return
				}

				select {
				case events <- e:
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				}
			}
		}
	}()

	<-started

	return events, errs
}
