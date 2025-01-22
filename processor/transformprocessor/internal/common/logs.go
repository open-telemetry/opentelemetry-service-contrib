// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

var _ consumer.Logs = &logStatements{}

type logStatements struct {
	ottl.StatementSequence[ottllog.TransformContext]
	LogRecordGlobalExpr expr.BoolExpr[ottllog.TransformContext]
	baseGlobalExpressions
}

func (l logStatements) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (l logStatements) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	for i := 0; i < ld.ResourceLogs().Len(); i++ {
		rlogs := ld.ResourceLogs().At(i)

		resourceCtx := ottlresource.NewTransformContext(rlogs.Resource(), plog.NewResourceLogs())
		rCondition, err := l.ResourceGlobalExpr.Eval(ctx, resourceCtx)
		if err != nil {
			return err
		}
		if !rCondition {
			continue
		}

		for j := 0; j < rlogs.ScopeLogs().Len(); j++ {
			slogs := rlogs.ScopeLogs().At(j)

			scopeCtx := ottlscope.NewTransformContext(slogs.Scope(), rlogs.Resource(), plog.NewScopeLogs())
			sCondition, err := l.ScopeGlobalExpr.Eval(ctx, scopeCtx)
			if err != nil {
				return err
			}
			if !sCondition {
				continue
			}

			logs := slogs.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				tCtx := ottllog.NewTransformContext(logs.At(k), slogs.Scope(), rlogs.Resource(), slogs, rlogs)
				condition, err := l.LogRecordGlobalExpr.Eval(ctx, tCtx)
				if err != nil {
					return err
				}
				if condition {
					err := l.Execute(ctx, tCtx)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

type LogParserCollection struct {
	parserCollection
	logParser ottl.Parser[ottllog.TransformContext]
}

type LogParserCollectionOption func(*LogParserCollection) error

func WithLogParser(functions map[string]ottl.Factory[ottllog.TransformContext]) LogParserCollectionOption {
	return func(lp *LogParserCollection) error {
		logParser, err := ottllog.NewParser(functions, lp.settings)
		if err != nil {
			return err
		}
		lp.logParser = logParser
		return nil
	}
}

func WithLogErrorMode(errorMode ottl.ErrorMode) LogParserCollectionOption {
	return func(lp *LogParserCollection) error {
		lp.errorMode = errorMode
		return nil
	}
}

func NewLogParserCollection(settings component.TelemetrySettings, options ...LogParserCollectionOption) (*LogParserCollection, error) {
	rp, err := ottlresource.NewParser(ResourceFunctions(), settings)
	if err != nil {
		return nil, err
	}
	sp, err := ottlscope.NewParser(ScopeFunctions(), settings)
	if err != nil {
		return nil, err
	}
	lpc := &LogParserCollection{
		parserCollection: parserCollection{
			settings:       settings,
			resourceParser: rp,
			scopeParser:    sp,
		},
	}

	for _, op := range options {
		err := op(lpc)
		if err != nil {
			return nil, err
		}
	}

	return lpc, nil
}

func (pc LogParserCollection) ParseContextStatements(contextStatements ContextStatements) (consumer.Logs, error) {
	switch contextStatements.Context {
	case Log:
		parsedStatements, err := pc.logParser.ParseStatements(contextStatements.Statements)
		if err != nil {
			return nil, err
		}
		globalExpr, errGlobalBoolExpr := parseGlobalExpr(
			filterottl.NewBoolExprForLog,
			contextStatements.Conditions,
			pc.parserCollection,
			filterottl.StandardLogFuncs(),
		)
		if errGlobalBoolExpr != nil {
			return nil, errGlobalBoolExpr
		}

		bge, err := pc.parseCommonGlobalExpressions(contextStatements)
		if err != nil {
			return nil, err
		}

		lStatements := ottllog.NewStatementSequence(parsedStatements, pc.settings, ottllog.WithStatementSequenceErrorMode(pc.errorMode))
		return logStatements{
			StatementSequence:     lStatements,
			LogRecordGlobalExpr:   globalExpr,
			baseGlobalExpressions: *bge,
		}, nil
	default:
		statements, err := pc.parseCommonContextStatements(contextStatements)
		if err != nil {
			return nil, err
		}
		return statements, nil
	}
}
