package server

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	hertzregistry "github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("gateway-service.server")

func NewServer(ctx context.Context, reg hertzregistry.Registry, serviceName string, e *config.ServerEnv) (*server.Hertz, error) {
	ctx, span := tracer.Start(ctx, "new-server")
	defer span.End()

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
	s.Use(gzip.Gzip(gzip.DefaultCompression))

	{
		c := cors.Config{
			AllowOrigins:     e.AllowOrigins,
			AllowMethods:     e.AllowMethods,
			AllowHeaders:     e.AllowHeaders,
			ExposeHeaders:    e.ExposeHeaders,
			AllowCredentials: e.AllowCredentials,
			MaxAge:           time.Duration(e.MaxAge) * time.Second,
		}
		if e.AllowOrigins == nil || lo.SomeBy(e.AllowOrigins, func(s string) bool { return s == "*" }) {
			c.AllowOrigins = nil
			c.AllowAllOrigins = true
		}

		s.Use(cors.New(c))
	}

	return s, nil
}
