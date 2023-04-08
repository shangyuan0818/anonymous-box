package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal"
	"github.com/star-horizon/anonymous-box-saas/internal/database"
	"github.com/star-horizon/anonymous-box-saas/internal/redis"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
	"github.com/star-horizon/anonymous-box-saas/services/email"
	"github.com/star-horizon/anonymous-box-saas/services/verify"
	"github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api/verifyservice"
)

const serviceName = "verify-service"

var (
	ctx    = context.Background()
	tracer = otel.Tracer("main")
	app    *fx.App
)

func init() {
	ctx, span := tracer.Start(ctx, "init")
	defer span.End()

	app = fx.New(
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			serviceName,
		),
		internal.InfraModule(),
		redis.Module(),
		fx.Provide(cache.NewRedisDriver),
		database.Module(),
		email.Module(), // use client to send email
		verify.Module(),
		fx.Invoke(run),
	)
}

func run(ctx context.Context, svc api.VerifyService, lc fx.Lifecycle, r registry.Registry) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	svr := verifyservice.NewServer(svc, server.WithRegistry(r), server.WithRegistryInfo(&registry.Info{
		ServiceName: serviceName,
	}))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "start")
			defer span.End()

			go func() {
				if err := svr.Run(); err != nil {
					logrus.WithContext(ctx).WithError(err).Fatal("run server failed")
					return
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "stop")
			defer span.End()

			return svr.Stop()
		},
	})
}

func main() {
	_, span := tracer.Start(ctx, "main")
	defer span.End()

	app.Run()
}
