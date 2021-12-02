// Copyright  The OpenTelemetry Authors
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

package metadataparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadataparser"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type Label struct {
	Name       string             `yaml:"name"`
	ColumnName string             `yaml:"column_name"`
	ValueType  metadata.ValueType `yaml:"value_type"`
}

func (label Label) toLabelValueMetadata() (metadata.LabelValueMetadata, error) {
	return metadata.NewLabelValueMetadata(label.Name, label.ColumnName, label.ValueType)
}
