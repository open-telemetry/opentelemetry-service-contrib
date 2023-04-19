// Copyright The OpenTelemetry Authors
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

package webhookeventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"

import (
	"bufio"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/receiver"

	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

var (
	errNilLogsConsumer      = errors.New("Missing a logs consumer")
	errMissingEndpoint      = errors.New("Missing a receiver endpoint")
	errInvalidRequestMethod = errors.New("Invalid method. Valid method is POST")
	errInvalidEncodingType  = errors.New("Invalid encoding type")
	errEmptyResponseBody    = errors.New("Request body content length is zero")
	errCosnumerError        = errors.New("Error with log consumer")
)

type eventReceiver struct {
	settings    receiver.CreateSettings
	cfg         *Config
	logConsumer consumer.Logs
	server      *http.Server
	shutdownWG  sync.WaitGroup
	obsrecv     *obsreport.Receiver
	gzipPool    *sync.Pool
}

func newLogsReceiver(params receiver.CreateSettings, cfg Config, consumer consumer.Logs) (receiver.Logs, error) {
	if consumer == nil {
		return nil, errNilLogsConsumer
	}

	if cfg.Endpoint == "" {
		return nil, errMissingEndpoint
	}

	transport := "http"
	if cfg.TLSSetting != nil {
		transport = "https"
	}

	obsrecv, err := obsreport.NewReceiver(obsreport.ReceiverSettings{
		ReceiverID:             params.ID,
		Transport:              transport,
		ReceiverCreateSettings: params,
	})

	if err != nil {
		return nil, err
	}

	// create eventReceiver instance
	er := &eventReceiver{
		settings:    params,
		cfg:         &cfg,
		logConsumer: consumer,
		obsrecv:     obsrecv,
		gzipPool:    &sync.Pool{New: func() interface{} { return new(gzip.Reader) }},
	}

	return er, nil
}

// start function manages receiver startup tasks. part of the receiver.Logs interface.
func (er *eventReceiver) Start(ctx context.Context, host component.Host) error {
	// noop if not nil. if start has not been called before these values should be nil.
	if er.server != nil && er.server.Handler != nil {
		return nil
	}

	// create listener from config
	ln, err := er.cfg.HTTPServerSettings.ToListener()
	if err != nil {
		return err
	}

	// set up router.
	router := httprouter.New()
    
	router.POST(er.cfg.Path, er.handleReq)

    // webhook server standup and configuration
	er.server, err = er.cfg.HTTPServerSettings.ToServer(host, er.settings.TelemetrySettings, router)
	if err != nil {
		return err
	}

	readTimeout, err := time.ParseDuration(er.cfg.ReadTimeout + "ms")
	if err != nil {
		return err
	}

	writeTimeout, err := time.ParseDuration(er.cfg.WriteTimeout + "ms")
	if err != nil {
		return err
	}

	// set timeouts
	er.server.ReadHeaderTimeout = readTimeout
	er.server.WriteTimeout = writeTimeout

	// shutdown
	er.shutdownWG.Add(1)
	go func() {
		defer er.shutdownWG.Done()
		if errHTTP := er.server.Serve(ln); !errors.Is(errHTTP, http.ErrServerClosed) && errHTTP != nil {
			host.ReportFatalError(errHTTP)
		}
	}()

	return nil
}

// stop function manages receiver shutdown tasks. part of the receiver.Logs interface.
func (er *eventReceiver) Shutdown(ctx context.Context) error {
	err := er.server.Close()
	er.shutdownWG.Wait()
	return err
}

// handle incoming request from webhook. On success returns a 200 response code to the webhook
func (er *eventReceiver) handleReq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	ctx = er.obsrecv.StartLogsOp(ctx)

	if r.Method != http.MethodPost {
		er.failBadReq(ctx, w, http.StatusBadRequest, errInvalidRequestMethod)
		return
	}

	encoding := r.Header.Get("Content-Encoding")
	// only support gzip if encoding header is set.
	if encoding != "" && encoding != "gzip" {
		er.failBadReq(ctx, w, http.StatusUnsupportedMediaType, errInvalidEncodingType)
		return
	}

	if r.ContentLength == 0 {
		er.obsrecv.EndLogsOp(ctx, typeStr, 0, nil)
		er.failBadReq(ctx, w, http.StatusBadRequest, errEmptyResponseBody)
	}

	bodyReader := r.Body
    // gzip encoded case
	if encoding == "gzip" || encoding == "x-gzip" {
		reader := er.gzipPool.Get().(*gzip.Reader)
		err := reader.Reset(bodyReader)

		if err != nil {
			er.failBadReq(ctx, w, http.StatusBadRequest, err)
			_, _ = io.ReadAll(r.Body)
			_ = r.Body.Close()
			return
		}
		bodyReader = reader
		defer er.gzipPool.Put(reader)
	}

    // finish reading the body into a log 
	sc := bufio.NewScanner(bodyReader)
	ld, numLogs := reqToLog(sc, r.URL.Query(), er.cfg)
	consumerErr := er.logConsumer.ConsumeLogs(ctx, ld)

	_ = bodyReader.Close()

	if consumerErr != nil {
		er.failBadReq(ctx, w, http.StatusInternalServerError, consumerErr)
	} else {
		w.WriteHeader(http.StatusOK)
		er.obsrecv.EndLogsOp(ctx, typeStr, numLogs, nil)
	}
}

// write response on a failed/bad request. Generates a small json body based on the thrown by
// the handle func and the appropriate http status code. many webhooks will either log these responses or
// notify webhook users should a none 2xx code be detected.
func (er *eventReceiver) failBadReq(ctx context.Context,
	w http.ResponseWriter,
	httpStatusCode int,
	err error) error {
	jsonResp, err := jsoniter.Marshal(err.Error())
	if err != nil {
		return err
	}

	// write response to webhook
	w.WriteHeader(httpStatusCode)
	if len(jsonResp) > 0 {
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write(jsonResp)
		if err != nil {
			er.settings.Logger.Warn("Failed to write json response", zap.Error(err))
		}
	}

	// log bad webhook request if debug is enabled
	if er.settings.Logger.Core().Enabled(zap.DebugLevel) {
		msg := string(jsonResp)
		er.settings.Logger.Debug(msg, zap.Int("http_status_code", httpStatusCode), zap.Error(err))
	}

	return nil
}

