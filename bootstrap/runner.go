package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/config"
)

type NewServiceFunc[T any] func(handler T, opts ...server.Option) server.Server
type ServiceRunnerFunc[T any] func(context.Context, T, fx.Lifecycle, string, registry.Registry, *config.ServiceEnv)

func RunService[T any](serviceFunc NewServiceFunc[T], options ...server.Option) ServiceRunnerFunc[T] {
	return func(ctx context.Context, handler T, lc fx.Lifecycle, serviceName string, reg registry.Registry, e *config.ServiceEnv) {
		ctx, span := tracer.Start(ctx, fmt.Sprintf("run-service:%s", serviceName))
		defer span.End()

		options = append(
			options,
			server.WithRegistry(reg),
			server.WithRegistryInfo(&registry.Info{
				ServiceName: serviceName,
				Weight:      10,
				WarmUp:      10 * time.Second,
			}),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithServiceAddr(utils.NewNetAddr(e.Network, e.Address)),
		)

		svr := serviceFunc(handler, options...)

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				ctx, span := tracer.Start(ctx, fmt.Sprintf("start-service:%s", serviceName))
				defer span.End()

				go func() {
					if err := svr.Run(); err != nil {
						logrus.WithContext(ctx).WithError(err).Panicf("run %s server failed", serviceName)
						return
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				ctx, span := tracer.Start(ctx, fmt.Sprintf("stop-service:%s", serviceName))
				defer span.End()

				if err := svr.Stop(); err != nil {
					logrus.WithContext(ctx).WithError(err).Errorf("stop %s server failed", serviceName)
					return err
				}

				return nil
			},
		})
	}
}
