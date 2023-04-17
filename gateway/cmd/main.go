package main

import (
	"context"
	"fmt"
	"time"

	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/gateway/server"
	"github.com/star-horizon/anonymous-box-saas/pkg/util"
)

const serviceName = "gateway"

var tracer = otel.Tracer("main")

func run(ctx context.Context, svr *hertzserver.Hertz, lc fx.Lifecycle) error {
	ctx, span := tracer.Start(ctx, "run-gateway")
	defer span.End()

	opts := svr.GetOptions()

	ip, err := util.GetLocalIP()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("get local ip error")
		return err
	}

	opts.RegistryInfo.Addr = utils.NewNetAddr("tcp", fmt.Sprintf("%s:%d", ip, 8080))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "start-gateway-http")
			defer span.End()

			svr.OnRun = append(svr.OnRun, func(ctx context.Context) error {
				ctx, span := tracer.Start(ctx, "gateway-http-registry-hook")
				defer span.End()

				go func() {
					// delay register 5s
					time.Sleep(5 * time.Second)
					logrus.WithContext(ctx).Info("gateway server register")
					if err := opts.Registry.Register(opts.RegistryInfo); err != nil {
						logrus.WithContext(ctx).WithError(err).Fatal("gateway server register error")
						return
					}
				}()

				return nil
			})

			go func() {
				if err := svr.Run(); err != nil {
					logrus.WithContext(ctx).WithError(err).Fatal("gateway server error")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "stop-gateway-http")
			defer span.End()

			if err := svr.Shutdown(ctx); err != nil {
				logrus.WithContext(ctx).WithError(err).Error("gateway server shutdown error")
				return err
			}

			return nil
		},
	})

	return nil
}

func main() {
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	app := bootstrap.InitApp(
		ctx,
		serviceName,
		server.Module(),
		fx.Invoke(run),
	)

	app.Run()
}
