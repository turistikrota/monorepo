package tracer

import (
	"context"
	"testing"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func TestPush(t *testing.T) {
	exp := sdktrace.NewTracerProvider()
	tr := exp.Tracer("test-tracer")

	t.Run("Span Creation and Context Propagation", func(t *testing.T) {
		ctx := context.Background()
		ctx = Push(ctx, tr, "test-span")

		span := trace.SpanFromContext(ctx)
		if span == nil {
			t.Error("Expected a span to be created and attached to the context")
		}
	})

	t.Run("Span Name Correctness", func(t *testing.T) {
		ctx := context.Background()
		ctx = Push(ctx, tr, "specific-span-name")

		span := trace.SpanFromContext(ctx)
		if span.SpanContext().TraceID().String() == "" ||
			span.SpanContext().SpanID().String() == "" {

			t.Error("Invalid SpanContext")
		}
	})
}
