package trace

import (
	"context"
	"github.com/star-horizon/anonymous-box-saas/config"

	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewExporter(ctx context.Context, e *config.TraceEnv) (tracesdk.SpanExporter, error) {
	ctx, span := tracer.Start(ctx, "trace-new-exporter")
	defer span.End()

	switch e.Exporter {
	case "jaeger":
		return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(e.Endpoint)))
	}

	return nil, nil
}
