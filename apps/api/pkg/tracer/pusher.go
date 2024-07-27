package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func Push(ctx context.Context, t trace.Tracer, name string) context.Context {
	ctx, span := t.Start(ctx, name)
	defer span.End()
	return ctx
}
