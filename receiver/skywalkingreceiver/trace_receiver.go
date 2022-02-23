// Copyright  OpenTelemetry Authors
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

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/obsreport"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	v1 "skywalking.apache.org/repo/goapi/satellite/data/v1"
)

// configuration defines the behavior and the ports that
// the Skywalking receiver will use.
type configuration struct {
	CollectorHTTPPort           int
	CollectorHTTPSettings       confighttp.HTTPServerSettings
	CollectorGRPCPort           int
	CollectorGRPCServerSettings configgrpc.GRPCServerSettings
}

// Receiver type is used to receive spans that were originally intended to be sent to Jaeger.
// This receiver is basically a Jaeger collector.
type swReceiver struct {
	nextConsumer consumer.Traces
	id           config.ComponentID

	config *configuration

	grpc            *grpc.Server
	collectorServer *http.Server

	goroutines sync.WaitGroup

	settings component.ReceiverCreateSettings

	grpcObsrecv *obsreport.Receiver
	httpObsrecv *obsreport.Receiver
}

const (
	collectorHTTPTransport = "http"
	grpcTransport          = "grpc"

	protobufFormat = "protobuf"
)

// newSkywalkingReceiver creates a TracesReceiver that receives traffic as a Skywalking collector
func newSkywalkingReceiver(
	id config.ComponentID,
	config *configuration,
	nextConsumer consumer.Traces,
	set component.ReceiverCreateSettings,
) *swReceiver {
	return &swReceiver{
		config:       config,
		nextConsumer: nextConsumer,
		id:           id,
		settings:     set,
		grpcObsrecv: obsreport.NewReceiver(obsreport.ReceiverSettings{
			ReceiverID:             id,
			Transport:              grpcTransport,
			ReceiverCreateSettings: set,
		}),
		httpObsrecv: obsreport.NewReceiver(obsreport.ReceiverSettings{
			ReceiverID:             id,
			Transport:              collectorHTTPTransport,
			ReceiverCreateSettings: set,
		}),
	}
}

func (sr *swReceiver) collectorGRPCAddr() string {
	var port int
	if sr.config != nil {
		port = sr.config.CollectorGRPCPort
	}
	return fmt.Sprintf(":%d", port)
}

func (sr *swReceiver) collectorGRPCEnabled() bool {
	return sr.config != nil && sr.config.CollectorGRPCPort > 0
}

func (sr *swReceiver) collectorHTTPEnabled() bool {
	return sr.config != nil && sr.config.CollectorHTTPPort > 0
}

func (sr *swReceiver) Start(_ context.Context, host component.Host) error {

	return sr.startCollector(host)
}

func (sr *swReceiver) Shutdown(ctx context.Context) error {
	var errs error

	if sr.collectorServer != nil {
		if cerr := sr.collectorServer.Shutdown(ctx); cerr != nil {
			errs = multierr.Append(errs, cerr)
		}
	}
	if sr.grpc != nil {
		sr.grpc.GracefulStop()
	}

	sr.goroutines.Wait()
	return errs
}

func (sr *swReceiver) startCollector(host component.Host) error {
	if !sr.collectorGRPCEnabled() && !sr.collectorHTTPEnabled() {
		return nil
	}

	if sr.collectorHTTPEnabled() {
		cln, cerr := sr.config.CollectorHTTPSettings.ToListener()
		if cerr != nil {
			return fmt.Errorf("failed to bind to Collector address %q: %v",
				sr.config.CollectorHTTPSettings.Endpoint, cerr)
		}

		nr := mux.NewRouter()
		nr.HandleFunc("/v3/segments", sr.httpHandler).Methods(http.MethodPost)
		sr.collectorServer, cerr = sr.config.CollectorHTTPSettings.ToServer(host, sr.settings.TelemetrySettings, nr)
		if cerr != nil {
			return cerr
		}

		sr.goroutines.Add(1)
		go func() {
			defer sr.goroutines.Done()
			if errHTTP := sr.collectorServer.Serve(cln); !errors.Is(errHTTP, http.ErrServerClosed) && errHTTP != nil {
				host.ReportFatalError(errHTTP)
			}
		}()
	}

	if sr.collectorGRPCEnabled() {
		opts, err := sr.config.CollectorGRPCServerSettings.ToServerOption(host, sr.settings.TelemetrySettings)
		if err != nil {
			return fmt.Errorf("failed to build the options for the Jaeger gRPC Collector: %v", err)
		}

		sr.grpc = grpc.NewServer(opts...)
		gaddr := sr.collectorGRPCAddr()
		gln, gerr := net.Listen("tcp", gaddr)
		if gerr != nil {
			return fmt.Errorf("failed to bind to gRPC address %q: %v", gaddr, gerr)
		}

		//api_v2.RegisterCollectorServiceServer(sr.grpc, sr)

		if gerr != nil {
			return fmt.Errorf("failed to create collector strategy store: %v", gerr)
		}

		sr.goroutines.Add(1)
		go func() {
			defer sr.goroutines.Done()
			if errGrpc := sr.grpc.Serve(gln); !errors.Is(errGrpc, grpc.ErrServerStopped) && errGrpc != nil {
				host.ReportFatalError(errGrpc)
			}
		}()
	}

	return nil
}

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func (sr *swReceiver) httpHandler(rsp http.ResponseWriter, r *http.Request) {
	rsp.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := &Response{Status: failing, Msg: err.Error()}
		ResponseWithJSON(rsp, response, http.StatusBadRequest)
		return
	}

	e := &v1.SniffData{
		Name:      httpEventName,
		Timestamp: time.Now().UnixNano() / 1e6,
		Meta:      nil,
		Type:      v1.SniffType_TracingType,
		Remote:    true,
		Data: &v1.SniffData_Segment{
			Segment: b,
		},
	}

	//TODO: convert to otel trace.
	sr.settings.Logger.Debug(e.String())

}

func ResponseWithJSON(rsp http.ResponseWriter, response *Response, code int) {
	rsp.WriteHeader(code)
	_ = json.NewEncoder(rsp).Encode(response)
}
