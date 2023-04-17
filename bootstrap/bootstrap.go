package bootstrap

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/config"
	"github.com/star-horizon/anonymous-box-saas/database"
	"github.com/star-horizon/anonymous-box-saas/internal/consul"
	"github.com/star-horizon/anonymous-box-saas/internal/hashids"
	"github.com/star-horizon/anonymous-box-saas/internal/jwt"
	"github.com/star-horizon/anonymous-box-saas/internal/logger"
	"github.com/star-horizon/anonymous-box-saas/internal/mq"
	"github.com/star-horizon/anonymous-box-saas/internal/redis"
	"github.com/star-horizon/anonymous-box-saas/internal/trace"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
	"github.com/star-horizon/anonymous-box-saas/services"
)

var tracer = otel.Tracer("bootstrap")

func InitApp(ctx context.Context, serviceName string, options ...fx.Option) *fx.App {
	ctx, span := tracer.Start(ctx, fmt.Sprintf("init-app-%s", serviceName))
	defer span.End()

	app := fx.New(
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			serviceName,
		),

		config.Module(),
		logger.Module(),
		consul.Module(),
		trace.Module(),

		redis.Module(),
		mq.Module(),
		fx.Provide(cache.NewRedisDriver),
		jwt.Module(),
		hashids.Module(),

		database.Module(),
		services.Module(),

		fx.Options(options...),
	)

	return app
}
