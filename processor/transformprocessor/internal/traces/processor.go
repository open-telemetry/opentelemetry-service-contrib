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

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/traces"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/tql"
)

type Processor struct {
	queries []tql.Query
	logger  *zap.Logger
}

func NewProcessor(statements []string, functions map[string]interface{}, settings component.ProcessorCreateSettings) (*Processor, error) {
	queries, err := tql.ParseQueries(statements, functions, ParsePath, ParseEnum)
	if err != nil {
		return nil, err
	}
	return &Processor{
		queries: queries,
		logger:  settings.Logger,
	}, nil
}

func (p *Processor) ProcessTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	ctx := spanTransformContext{}
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		rspans := td.ResourceSpans().At(i)
		ctx.resource = rspans.Resource()
		for j := 0; j < rspans.ScopeSpans().Len(); j++ {
			sspan := rspans.ScopeSpans().At(j)
			ctx.il = sspan.Scope()
			spans := sspan.Spans()
			for k := 0; k < spans.Len(); k++ {
				ctx.span = spans.At(k)

				for _, statement := range p.queries {
					if statement.Condition(ctx) {
						statement.Function(ctx)
					}
				}
			}
		}
	}
	return td, nil
}
