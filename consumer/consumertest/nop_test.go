// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consumertest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/model/pdata"
)

func TestNop(t *testing.T) {
	nc := NewNop()
	require.NotNil(t, nc)
	assert.NotPanics(t, nc.unexported)
	assert.NoError(t, nc.ConsumeLogs(context.Background(), pdata.NewLogs()))
	assert.NoError(t, nc.ConsumeMetrics(context.Background(), pdata.NewMetrics()))
	assert.NoError(t, nc.ConsumeTraces(context.Background(), pdata.NewTraces()))
}
