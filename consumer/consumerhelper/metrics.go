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

package consumerhelper

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
)

// ConsumeMetricsFunc is a helper function that is similar to ConsumeMetrics.
type ConsumeMetricsFunc func(ctx context.Context, ld pdata.Metrics) error

// ConsumeMetrics calls f(ctx, ld).
func (f ConsumeMetricsFunc) ConsumeMetrics(ctx context.Context, ld pdata.Metrics) error {
	return f(ctx, ld)
}

type baseMetrics struct {
	*baseConsumer
	ConsumeMetricsFunc
}

// NewMetrics returns a consumer.Metrics configured with the provided options.
func NewMetrics(consume ConsumeMetricsFunc, options ...Option) (consumer.Metrics, error) {
	if consume == nil {
		return nil, errNilFunc
	}
	return &baseMetrics{
		baseConsumer:       newBaseConsumer(options...),
		ConsumeMetricsFunc: consume,
	}, nil
}
