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

package attributesprocessor // import "github.com/asserts/opentelemetry-collector-contrib/processor/attributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/asserts/opentelemetry-collector-contrib/internal/coreinternal/attraction"
	"github.com/asserts/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/asserts/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

type logAttributesProcessor struct {
	logger   *zap.Logger
	attrProc *attraction.AttrProc
	skipExpr expr.BoolExpr[ottllog.TransformContext]
}

// newLogAttributesProcessor returns a processor that modifies attributes of a
// log record. To construct the attributes processors, the use of the factory
// methods are required in order to validate the inputs.
func newLogAttributesProcessor(logger *zap.Logger, attrProc *attraction.AttrProc, skipExpr expr.BoolExpr[ottllog.TransformContext]) *logAttributesProcessor {
	return &logAttributesProcessor{
		logger:   logger,
		attrProc: attrProc,
		skipExpr: skipExpr,
	}
}

func (a *logAttributesProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rs := rls.At(i)
		ilss := rs.ScopeLogs()
		resource := rs.Resource()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			logs := ils.LogRecords()
			library := ils.Scope()
			for k := 0; k < logs.Len(); k++ {
				lr := logs.At(k)
				if a.skipExpr != nil {
					skip, err := a.skipExpr.Eval(ctx, ottllog.NewTransformContext(lr, library, resource))
					if err != nil {
						return ld, err
					}
					if skip {
						continue
					}
				}

				a.attrProc.Process(ctx, a.logger, lr.Attributes())
			}
		}
	}
	return ld, nil
}
