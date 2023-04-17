package trace

import (
	"context"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/config"
)

func KitexOpenTelemetryProvider(ctx context.Context, tracerProvider *tracesdk.TracerProvider, serviceName string, lc fx.Lifecycle) {
	ctx, span := tracer.Start(ctx, "kitex-opentelemetry-provider")
	defer span.End()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithSdkTracerProvider(tracerProvider),
		provider.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNamespaceKey.String(config.Namespace),
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(config.Version),
			semconv.ServiceInstanceIDKey.String(config.ServiceInstanceID),
		)),
		provider.WithInsecure(),
		provider.WithEnableTracing(true),
	)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "kitex-opentelemetry-provider-on-stop")
			defer span.End()

			return p.Shutdown(ctx)
		},
	})
}
