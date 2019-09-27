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

package zpagesextension

import (
	"net"
	"net/http"

	"go.opencensus.io/zpages"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector/extension"
)

type zpagesExtension struct {
	config Config
	logger *zap.Logger
	server http.Server
}

var _ (extension.ServiceExtension) = (*zpagesExtension)(nil)

func (zpe *zpagesExtension) Start(host extension.Host) error {
	zPagesMux := http.NewServeMux()
	zpages.Handle(zPagesMux, "/debug")

	// Start the listener here so we can have earlier failure if port is
	// already in use.
	ln, err := net.Listen("tcp", zpe.config.Endpoint)
	if err != nil {
		return err
	}

	zpe.logger.Info("Starting zPages extension", zap.Any("config", zpe.config))
	zpe.server = http.Server{Handler: zPagesMux}
	go func() {
		if err := zpe.server.Serve(ln); err != nil && err != http.ErrServerClosed {
			host.ReportFatalError(err)
		}
	}()

	return nil
}

func (zpe *zpagesExtension) Shutdown() error {
	return zpe.server.Close()
}

func newServer(config Config, logger *zap.Logger) (*zpagesExtension, error) {
	zpe := &zpagesExtension{
		config: config,
		logger: logger,
	}

	return zpe, nil
}
