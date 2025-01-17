package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type Hot struct {
	Total int64
	Limit int64
}

var attrLimit = metric.WithAttributes(Error("limit"))

func (h *Hot) Write(ctx context.Context, m Metrics) {
	total := h.Total - h.Limit
	if total > 0 {
		m.Datapoints().Add(ctx, total)
	}
	if h.Limit > 0 {
		m.Datapoints().Add(ctx, h.Limit, attrLimit)
	}
}
