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

package transport

import (
	"bytes"
	"context"
	"io"
	"net"
	"strings"
	"sync"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumerdata"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"
)

type udpServer struct {
	wg         sync.WaitGroup
	packetConn net.PacketConn
}

var _ (Server) = (*udpServer)(nil)

// NewUDPServer creates a transport.Server using UDP as its transport.
func NewUDPServer(addr string) (Server, error) {
	packetConn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, err
	}

	u := udpServer{
		packetConn: packetConn,
	}
	return &u, nil
}

func (u *udpServer) ListenAndServe(
	parser protocol.Parser,
	nextConsumer consumer.MetricsConsumerOld,
) error {
	if parser == nil || nextConsumer == nil {
		return errNilListenAndServeParameters
	}

	buf := make([]byte, 65527) // max size for udp packet body (assuming ipv6)
	for {
		n, _, err := u.packetConn.ReadFrom(buf)
		if n > 0 {
			u.wg.Add(1)
			bufCopy := make([]byte, n)
			copy(bufCopy, buf)
			go func() {
				u.handlePacket(parser, nextConsumer, bufCopy)
				u.wg.Done()
			}()
		}
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if netErr.Temporary() {
					continue
				}
			}
			return err
		}
	}
}

func (u *udpServer) Close() error {
	err := u.packetConn.Close()
	u.wg.Wait()
	return err
}

func (u *udpServer) handlePacket(
	p protocol.Parser,
	nextConsumer consumer.MetricsConsumerOld,
	data []byte,
) {
	ctx := context.Background()
	var numReceivedTimeseries, numInvalidTimeseries int
	var metrics []*metricspb.Metric
	buf := bytes.NewBuffer(data)
	for {
		bytes, err := buf.ReadBytes((byte)('\n'))
		if err == io.EOF {
			if len(bytes) == 0 {
				// Completed without errors.
				break
			}
		}
		line := strings.TrimSpace(string(bytes))
		if line != "" {
			numReceivedTimeseries++
			metric, err := p.Parse(line)
			if err != nil {
				numInvalidTimeseries++
				continue
			}

			metrics = append(metrics, metric)
		}
	}

	md := consumerdata.MetricsData{
		Metrics: metrics,
	}
	// TODO: handle error?
	nextConsumer.ConsumeMetricsData(ctx, md)
}
