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

package statsdreceiver

import (
	"context"
	"errors"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/asserts/opentelemetry-collector-contrib/internal/common/testutil"
	"github.com/asserts/opentelemetry-collector-contrib/receiver/statsdreceiver/transport"
	"github.com/asserts/opentelemetry-collector-contrib/receiver/statsdreceiver/transport/client"
)

func Test_statsdreceiver_New(t *testing.T) {
	defaultConfig := createDefaultConfig().(*Config)
	type args struct {
		config       Config
		nextConsumer consumer.Metrics
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "nil_nextConsumer",
			args: args{
				config: *defaultConfig,
			},
			wantErr: component.ErrNilNextConsumer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(receivertest.NewNopCreateSettings(), tt.args.config, tt.args.nextConsumer)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_statsdreceiver_Start(t *testing.T) {
	type args struct {
		config       Config
		nextConsumer consumer.Metrics
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "unsupported transport",
			args: args{
				config: Config{
					NetAddr: confignet.NetAddr{
						Endpoint:  "localhost:8125",
						Transport: "unknown",
					},
				},
				nextConsumer: consumertest.NewNop(),
			},
			wantErr: errors.New("unsupported transport \"unknown\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver, err := New(receivertest.NewNopCreateSettings(), tt.args.config, tt.args.nextConsumer)
			require.NoError(t, err)
			err = receiver.Start(context.Background(), componenttest.NewNopHost())
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestStatsdReceiver_Flush(t *testing.T) {
	ctx := context.Background()
	cfg := createDefaultConfig().(*Config)
	nextConsumer := consumertest.NewNop()
	rcv, err := New(receivertest.NewNopCreateSettings(), *cfg, nextConsumer)
	assert.NoError(t, err)
	r := rcv.(*statsdReceiver)
	var metrics = pmetric.NewMetrics()
	assert.Nil(t, r.Flush(ctx, metrics, nextConsumer))
	assert.NoError(t, r.Start(ctx, componenttest.NewNopHost()))
	assert.NoError(t, r.Shutdown(ctx))
}

func Test_statsdreceiver_EndToEnd(t *testing.T) {
	addr := testutil.GetAvailableLocalAddress(t)
	host, portStr, err := net.SplitHostPort(addr)
	require.NoError(t, err)
	port, err := strconv.Atoi(portStr)
	require.NoError(t, err)

	tests := []struct {
		name     string
		configFn func() *Config
		clientFn func(t *testing.T) *client.StatsD
	}{
		{
			name: "default_config with 4s interval",
			configFn: func() *Config {
				return &Config{
					NetAddr: confignet.NetAddr{
						Endpoint:  defaultBindEndpoint,
						Transport: defaultTransport,
					},
					AggregationInterval: 4 * time.Second,
				}
			},
			clientFn: func(t *testing.T) *client.StatsD {
				c, err := client.NewStatsD(client.UDP, host, port)
				require.NoError(t, err)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.configFn()
			cfg.NetAddr.Endpoint = addr
			sink := new(consumertest.MetricsSink)
			rcv, err := New(receivertest.NewNopCreateSettings(), *cfg, sink)
			require.NoError(t, err)
			r := rcv.(*statsdReceiver)

			mr := transport.NewMockReporter(1)
			r.reporter = mr

			require.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))
			defer func() {
				assert.NoError(t, r.Shutdown(context.Background()))
			}()

			statsdClient := tt.clientFn(t)

			statsdMetric := client.Metric{
				Name:  "test.metric",
				Value: "42",
				Type:  "c",
			}
			err = statsdClient.SendMetric(statsdMetric)
			require.NoError(t, err)

			time.Sleep(5 * time.Second)
			mdd := sink.AllMetrics()
			require.Len(t, mdd, 1)
			require.Equal(t, 1, mdd[0].ResourceMetrics().Len())
			require.Equal(t, 1, mdd[0].ResourceMetrics().At(0).ScopeMetrics().Len())
			require.Equal(t, 1, mdd[0].ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().Len())
			metric := mdd[0].ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().At(0)
			assert.Equal(t, statsdMetric.Name, metric.Name())
			assert.Equal(t, pmetric.MetricTypeSum, metric.Type())
			require.Equal(t, 1, metric.Sum().DataPoints().Len())
			assert.NotEqual(t, 0, metric.Sum().DataPoints().At(0).Timestamp())
			assert.NotEqual(t, 0, metric.Sum().DataPoints().At(0).StartTimestamp())
			assert.Less(t, metric.Sum().DataPoints().At(0).StartTimestamp(), metric.Sum().DataPoints().At(0).Timestamp())

			// Send the same metric again to ensure that the timestamps of successive data points
			// are aligned.
			statsdMetric.Value = "43"
			err = statsdClient.SendMetric(statsdMetric)
			require.NoError(t, err)

			time.Sleep(5 * time.Second)
			mddAfter := sink.AllMetrics()
			require.Len(t, mddAfter, 2)
			metricAfter := mddAfter[1].ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().At(0)
			require.Equal(t, metric.Sum().DataPoints().At(0).Timestamp(), metricAfter.Sum().DataPoints().At(0).StartTimestamp())
		})
	}
}
