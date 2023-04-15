package main

import (
	"context"
	"github.com/star-horizon/anonymous-box-saas/services/website"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database"
	"github.com/star-horizon/anonymous-box-saas/internal/infra"
	"github.com/star-horizon/anonymous-box-saas/internal/mq"
	"github.com/star-horizon/anonymous-box-saas/internal/redis"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/commentservice"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
	"github.com/star-horizon/anonymous-box-saas/services/comment"
)

var (
	ctx    = context.Background()
	tracer = otel.Tracer("main")
	app    *fx.App
)

func init() {
	ctx, span := tracer.Start(ctx, "init")
	defer span.End()

	serviceName := comment.ServiceName

	app = fx.New(
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			serviceName,
		),
		infra.Module(),
		redis.Module(),
		fx.Provide(cache.NewRedisDriver),
		database.Module(),
		mq.Module(),
		website.Module(),
		comment.Module(),
		fx.Invoke(run),
	)
}

func run(ctx context.Context, lc fx.Lifecycle, svc dash.CommentService, r registry.Registry) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	svr := commentservice.NewServer(
		svc,
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: comment.ServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
	)

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

			if err := svr.Stop(); err != nil {
				logrus.WithContext(ctx).WithError(err).Fatal("stop server failed")
				return err
			}

			return nil
		},
	})
}

func main() {
	_, span := tracer.Start(ctx, "main")
	defer span.End()

	app.Run()
}
