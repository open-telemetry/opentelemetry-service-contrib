// Copyright 2020, OpenTelemetry Authors
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

package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExcludeMetricsRule(t *testing.T) {
	rule := GetExcludeMetricsRule([]string{"m1", "m2"})
	assert.Equal(t, rule.Action, ActionDropMetrics)
	assert.Equal(t, 2, len(rule.MetricNames))
	assert.False(t, rule.MetricNames["m0"])
	assert.True(t, rule.MetricNames["m1"])
}
