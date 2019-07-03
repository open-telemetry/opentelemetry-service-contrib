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

package vmmetricsreceiver

import (
	"context"
	"errors"
	"runtime"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/models"
	"github.com/open-telemetry/opentelemetry-service/receiver"
)

// This file implements factory for VMMetrics receiver.

var _ = receiver.RegisterFactory(&factory{})

const (
	// The value of "type" key in configuration.
	typeStr = "vmmetrics"
)

// factory is the factory for receiver.
type factory struct {
}

// Type gets the type of the Receiver config created by this factory.
func (f *factory) Type() string {
	return typeStr
}

// CustomUnmarshaler returns custom unmarshaler for this config.
func (f *factory) CustomUnmarshaler() receiver.CustomUnmarshaler {
	return nil
}

// CreateDefaultConfig creates the default configuration for receiver.
func (f *factory) CreateDefaultConfig() models.Receiver {
	return &ConfigV2{
		ReceiverSettings: models.ReceiverSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
	}
}

// CreateTraceReceiver creates a trace receiver based on provided config.
func (f *factory) CreateTraceReceiver(
	ctx context.Context,
	logger *zap.Logger,
	cfg models.Receiver,
	nextConsumer consumer.TraceConsumer,
) (receiver.TraceReceiver, error) {
	// VMMetrics does not support traces
	return nil, models.ErrDataTypeIsNotSupported
}

// CreateMetricsReceiver creates a metrics receiver based on provided config.
func (f *factory) CreateMetricsReceiver(
	logger *zap.Logger,
	config models.Receiver,
	consumer consumer.MetricsConsumer,
) (receiver.MetricsReceiver, error) {

	if runtime.GOOS != "linux" {
		return nil, errors.New("vmmetrics receiver is only supported on linux")
	}
	cfg := config.(*ConfigV2)

	vmc, err := NewVMMetricsCollector(cfg.ScrapeInterval, cfg.MountPoint, cfg.ProcessMountPoint, cfg.MetricPrefix, consumer)
	if err != nil {
		return nil, err
	}

	vmr := &Receiver{
		vmc: vmc,
	}
	return vmr, nil
}
