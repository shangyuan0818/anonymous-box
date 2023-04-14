package server

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/server"
	hertzregistry "github.com/cloudwego/hertz/pkg/app/server/registry"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("gateway-service.server")

type env struct {
	Debug            bool     `envconfig:"DEBUG" default:"false"`
	AllowOrigins     []string `envconfig:"CORS_ALLOW_ORIGINS" split_words:"true" default:"*"`
	AllowMethods     []string `envconfig:"CORS_ALLOW_METHODS" split_words:"true" default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowHeaders     []string `envconfig:"CORS_ALLOW_HEADERS" split_words:"true" default:"Authorization,Content-Type"`
	ExposeHeaders    []string `envconfig:"CORS_EXPOSE_HEADERS" split_words:"true" default:""`
	AllowCredentials bool     `envconfig:"CORS_ALLOW_CREDENTIALS" split_words:"true" default:"true"`
	MaxAge           int      `envconfig:"CORS_MAX_AGE" split_words:"true" default:"3600"`
}

func NewServer(ctx context.Context, reg hertzregistry.Registry, serviceName string) (*server.Hertz, error) {
	ctx, span := tracer.Start(ctx, "new-server")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).Warn("failed to load .env file")
	}

	var e env
	if err := envconfig.Process("SERVER", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to process envconfig")
		return nil, err
	}

	serverTracer, cfg := hertztracing.NewServerTracer()
	s := server.Default(
		serverTracer,
		server.WithHostPorts(":8080"),
		server.WithRegistry(reg, &hertzregistry.Info{
			ServiceName: serviceName,
			Weight:      10,
			Tags:        nil,
		}),
		server.WithKeepAlive(true),
	)
	s.Use(hertztracing.ServerMiddleware(cfg))

	return s, nil
}
