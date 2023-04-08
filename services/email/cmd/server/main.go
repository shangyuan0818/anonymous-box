package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/anonymous-box-saas/internal/database"
	"github.com/ahdark-services/anonymous-box-saas/internal/logger"
	"github.com/ahdark-services/anonymous-box-saas/internal/mq"
	"github.com/ahdark-services/anonymous-box-saas/internal/redis"
	"github.com/ahdark-services/anonymous-box-saas/internal/trace"
	"github.com/ahdark-services/anonymous-box-saas/pkg/cache"
	"github.com/ahdark-services/anonymous-box-saas/services/email"
	"github.com/ahdark-services/anonymous-box-saas/services/email/kitex_gen/api"
	"github.com/ahdark-services/anonymous-box-saas/services/email/kitex_gen/api/mailservice"
)

const serviceName = "email-service"

var (
	ctx    = context.Background()
	tracer = otel.Tracer("main")
	app    *fx.App
)

func init() {
	ctx, span := tracer.Start(ctx, "init")
	defer span.End()

	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			serviceName,
		),
		logger.Module(),
		trace.Module(),
		redis.Module(),
		fx.Provide(cache.NewRedisDriver),
		database.Module(),
		mq.Module(),
		email.Module(),
		fx.Invoke(run),
	}

	app = fx.New(opts...)
}

func run(ctx context.Context, lc fx.Lifecycle, svc api.MailService) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	svr := mailservice.NewServer(svc)

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
