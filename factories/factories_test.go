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

package factories

import (
	"context"
	"testing"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/models"
	"github.com/open-telemetry/opentelemetry-service/processor"
	"github.com/open-telemetry/opentelemetry-service/receiver"
)

type ExampleReceiverFactory struct {
}

// Type gets the type of the Receiver config created by this factory.
func (f *ExampleReceiverFactory) Type() string {
	return "examplereceiver"
}

// CustomUnmarshaler returns nil because we don't need custom unmarshaling for this factory.
func (f *ExampleReceiverFactory) CustomUnmarshaler() CustomUnmarshaler {
	return nil
}

// CreateDefaultConfig creates the default configuration for the Receiver.
func (f *ExampleReceiverFactory) CreateDefaultConfig() models.Receiver {
	return nil
}

// CreateTraceReceiver creates a trace receiver based on this config.
func (f *ExampleReceiverFactory) CreateTraceReceiver(
	ctx context.Context,
	logger *zap.Logger,
	cfg models.Receiver,
	nextConsumer consumer.TraceConsumer,
) (receiver.TraceReceiver, error) {
	// Not used for this test, just return nil
	return nil, nil
}

// CreateMetricsReceiver creates a metrics receiver based on this config.
func (f *ExampleReceiverFactory) CreateMetricsReceiver(
	logger *zap.Logger,
	cfg models.Receiver,
	consumer consumer.MetricsConsumer,
) (receiver.MetricsReceiver, error) {
	// Not used for this test, just return nil
	return nil, nil
}

func TestRegisterReceiverFactory(t *testing.T) {
	f := ExampleReceiverFactory{}
	err := RegisterReceiverFactory(&f)
	if err != nil {
		t.Fatalf("cannot register factory")
	}

	if &f != GetReceiverFactory(f.Type()) {
		t.Fatalf("cannot find factory")
	}

	// Verify that attempt to register a factory with duplicate name panics
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		err = RegisterReceiverFactory(&f)
	}()

	if !panicked {
		t.Fatalf("must panic on double registration")
	}
}

type ExampleExporterFactory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *ExampleExporterFactory) Type() string {
	return "exampleexporter"
}

// CreateDefaultConfig creates the default configuration for the Exporter.
func (f *ExampleExporterFactory) CreateDefaultConfig() models.Exporter {
	return nil
}

// CreateTraceExporter creates a trace exporter based on this config.
func (f *ExampleExporterFactory) CreateTraceExporter(cfg models.Exporter) (consumer.TraceConsumer, StopFunc, error) {
	return nil, nil, nil
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *ExampleExporterFactory) CreateMetricsExporter(cfg models.Exporter) (consumer.MetricsConsumer, StopFunc, error) {
	return nil, nil, nil
}

func TestRegisterExporterFactory(t *testing.T) {
	f := ExampleExporterFactory{}
	err := RegisterExporterFactory(&f)
	if err != nil {
		t.Fatalf("cannot register factory")
	}

	if &f != GetExporterFactory(f.Type()) {
		t.Fatalf("cannot find factory")
	}

	// Verify that attempt to register a factory with duplicate name panics
	paniced := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				paniced = true
			}
		}()

		err = RegisterExporterFactory(&f)
	}()

	if !paniced {
		t.Fatalf("must panic on double registration")
	}
}

type ExampleProcessorFactory struct {
}

// Type gets the type of the Processor config created by this factory.
func (f *ExampleProcessorFactory) Type() string {
	return "exampleoption"
}

// CreateDefaultConfig creates the default configuration for the Processor.
func (f *ExampleProcessorFactory) CreateDefaultConfig() models.Processor {
	return nil
}

// CreateTraceProcessor creates a trace processor based on this config.
func (f *ExampleProcessorFactory) CreateTraceProcessor(
	logger *zap.Logger,
	nextConsumer consumer.TraceConsumer,
	cfg models.Processor,
) (processor.TraceProcessor, error) {
	return nil, ErrDataTypeIsNotSupported
}

// CreateMetricsProcessor creates a metrics processor based on this config.
func (f *ExampleProcessorFactory) CreateMetricsProcessor(
	logger *zap.Logger,
	nextConsumer consumer.MetricsConsumer,
	cfg models.Processor,
) (processor.MetricsProcessor, error) {
	return nil, ErrDataTypeIsNotSupported
}

func TestRegisterProcessorFactory(t *testing.T) {
	f := ExampleProcessorFactory{}
	err := RegisterProcessorFactory(&f)
	if err != nil {
		t.Fatalf("cannot register factory")
	}

	if &f != GetProcessorFactory(f.Type()) {
		t.Fatalf("cannot find factory")
	}

	// Verify that attempt to register a factory with duplicate name panics
	paniced := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				paniced = true
			}
		}()

		err = RegisterProcessorFactory(&f)
	}()

	if !paniced {
		t.Fatalf("must panic on double registration")
	}
}
