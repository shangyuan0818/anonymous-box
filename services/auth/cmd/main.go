package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal/database"
	"github.com/star-horizon/anonymous-box-saas/internal/logger"
	"github.com/star-horizon/anonymous-box-saas/internal/redis"
	"github.com/star-horizon/anonymous-box-saas/internal/trace"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
	"github.com/star-horizon/anonymous-box-saas/services/auth"
	"github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api/authservice"
)

const serviceName = "auth-service"

var (
	tracer = otel.Tracer("main")
	ctx    = context.Background()
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
		logger.Module(),
		trace.Module(),
		redis.Module(),
		fx.Provide(cache.NewRedisDriver),
		database.Module(),
		auth.Module(),
		fx.Invoke(run),
	)
}

func run(ctx context.Context, svc api.AuthService, lc fx.Lifecycle) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	svr := authservice.NewServer(svc)

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
