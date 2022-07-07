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

package translation

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/alias"
)

// noModify is used when there is no definition present for a given change
type noModify struct{}

var _ Modifier = (*noModify)(nil)

func (noModify) UpdateAttrs(attrs pcommon.Map)                 {}
func (noModify) RevertAttrs(attrs pcommon.Map)                 {}
func (noModify) UpdateAttrsIf(match string, attrs pcommon.Map) {}
func (noModify) RevertAttrsIf(match string, attrs pcommon.Map) {}
func (noModify) UpdateSignal(signal alias.Signal)              {}
func (noModify) RevertSignal(signal alias.Signal)              {}
