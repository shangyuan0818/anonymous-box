package main

import (
	"context"

	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/gateway/server"
	"github.com/star-horizon/anonymous-box-saas/internal"
	"github.com/star-horizon/anonymous-box-saas/services/auth"
	"github.com/star-horizon/anonymous-box-saas/services/verify"
)

const serviceName = "gateway"

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
		auth.Module(),   // use client
		verify.Module(), // use client
		server.Module(), // use controller
		fx.Invoke(run),
	)
}

func run(ctx context.Context, svr *hertzserver.Hertz, lc fx.Lifecycle) {
	ctx, span := tracer.Start(ctx, "run-gateway")
	defer span.End()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "start-gateway")
			defer span.End()

			go func() {
				if err := svr.Run(); err != nil {
					logrus.WithContext(ctx).WithError(err).Fatal("gateway server error")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "stop-gateway")
			defer span.End()

			if err := svr.Shutdown(ctx); err != nil {
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
