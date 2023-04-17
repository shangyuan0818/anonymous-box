package trace

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/uptrace-go/extra/otellogrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.trace")

func NewTracerProvider(ctx context.Context, serviceName string, exporter tracesdk.SpanExporter) *tracesdk.TracerProvider {
	ctx, span := tracer.Start(ctx, "new-tracer-provider")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNamespace(config.Namespace),
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(config.Version),
			semconv.ServiceInstanceIDKey.String(config.ServiceInstanceID),
		)),
	)

	return tp
}

func InitTracer(ctx context.Context, tracerProvider *tracesdk.TracerProvider) error {
	ctx, span := tracer.Start(ctx, "init-tracer")
	defer span.End()

	otel.SetTracerProvider(tracerProvider)
	logrus.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	)))

	return nil
}
