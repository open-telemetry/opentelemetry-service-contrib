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

package httpforwarder

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/extension/extensionhelper"
)

const (
	// The value of extension "type" in configuration.
	typeStr configmodels.Type = "http_forwarder"

	// Default endpoints to bind to.
	defaultEndpoint = ":6060"
)

// NewFactory creates a factory for HostObserver extension.
func NewFactory() component.ExtensionFactory {
	return extensionhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		createExtension)
}

func createDefaultConfig() configmodels.Extension {
	return &Config{
		ExtensionSettings: configmodels.ExtensionSettings{
			TypeVal: typeStr,
			NameVal: string(typeStr),
		},
		Ingress: confighttp.HTTPServerSettings{
			Endpoint: defaultEndpoint,
		},
		Egress: confighttp.HTTPClientSettings{
			Timeout: 10 * time.Second,
		},
	}
}

func createExtension(
	_ context.Context,
	_ component.ExtensionCreateParams,
	_ configmodels.Extension,
) (component.ServiceExtension, error) {
	// TODO: Return httpForwarder
	return nil, nil
}
