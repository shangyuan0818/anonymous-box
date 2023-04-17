package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/config"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/commentservice"
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

	app = bootstrap.InitApp(
		ctx,
		comment.ServiceName,
		fx.Invoke(run),
	)
}

func run(ctx context.Context, lc fx.Lifecycle, svc dash.CommentService, r registry.Registry, e *config.ServiceEnv) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	svr := commentservice.NewServer(
		svc,
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: comment.ServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServiceAddr(utils.NewNetAddr(e.Network, e.Address)),
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
