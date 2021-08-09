// Copyright The OpenTelemetry Authors
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

package docker

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/translator/conventions/v1.5.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

const (
	// TypeStr is type of detector.
	TypeStr = "docker"
)

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
	provider dockerMetadata
	logger   *zap.Logger
}

// NewDetector creates a new system metadata detector
func NewDetector(p component.ProcessorCreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	dockerProvider, err := newDockerMetadata()
	if err != nil {
		return nil, fmt.Errorf("failed creating detector: %w", err)
	}

	return &Detector{provider: dockerProvider, logger: p.Logger}, nil
}

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (pdata.Resource, error) {
	res := pdata.NewResource()
	attrs := res.Attributes()

	osType, err := d.provider.OSType(ctx)
	if err != nil {
		return res, fmt.Errorf("failed getting OS type: %w", err)
	}

	hostname, err := d.provider.Hostname(ctx)
	if err != nil {
		return res, fmt.Errorf("failed getting OS hostname: %w", err)
	}

	attrs.InsertString(conventions.AttributeHostName, hostname)
	attrs.InsertString(conventions.AttributeOSType, osType)

	return res, nil
}
