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

package ecsobserver

import (
	"context"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
)

var _ component.Extension = (*ecsObserver)(nil)

// ecsObserver implements component.ServiceExtension interface.
type ecsObserver struct {
	logger *zap.Logger
	sd     *ServiceDiscovery

	// for Shutdown
	cancel func()
}

// Start runs the service discovery in backeground
func (e *ecsObserver) Start(_ context.Context, host component.Host) error {
	e.logger.Info("Starting ECSDiscovery")
	// Ignore the ctx parameter as it is not for long running operation
	ctx, cancel := context.WithCancel(context.Background())
	e.cancel = cancel
	go func() {
		if err := e.sd.RunAndWriteFile(ctx); err != nil {
			e.logger.Error("ECSDiscovery stopped by error", zap.Error(err))
		}
	}()
	return nil
}

func (e *ecsObserver) Shutdown(ctx context.Context) error {
	e.logger.Info("Stopping ECSDiscovery")
	e.cancel()
	return nil
}
