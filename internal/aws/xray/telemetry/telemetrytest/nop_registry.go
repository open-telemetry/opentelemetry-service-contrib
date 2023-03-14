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

package telemetrytest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry/telemetrytest"

import (
	"go.opentelemetry.io/collector/component"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

func NewNopRegistry() telemetry.Registry {
	return nopRegistryInstance
}

type nopRegistry struct {
	recorder telemetry.Recorder
}

var _ telemetry.Registry = (*nopRegistry)(nil)

var nopRegistryInstance = &nopRegistry{
	recorder: NewNopRecorder(),
}

func (n nopRegistry) Register(component.ID, telemetry.Config, awsxray.XRayClient, ...telemetry.RecorderOption) telemetry.Recorder {
	return n.recorder
}

func (n nopRegistry) Set(component.ID, telemetry.Recorder) {
}

func (n nopRegistry) Get(component.ID) telemetry.Recorder {
	return n.recorder
}
