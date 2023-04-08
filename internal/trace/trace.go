package trace

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/uptrace-go/extra/otellogrus"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"

	"github.com/ahdark-services/anonymous-box-saas/internal/config"
)

var tracer = otel.Tracer("internal.trace")

type env struct {
	Endpoint string `default:"http://localhost:14268/api/traces"`
}

func InitTracer(ctx context.Context, serviceName string) error {
	ctx, span := tracer.Start(ctx, "init-tracer")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}

	var e env
	if err := envconfig.Process("TRACE", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to process envconfig")
		return err
	}

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(e.Endpoint)))
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create jaeger exporter")
		return err
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

	otel.SetTracerProvider(tp)
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
