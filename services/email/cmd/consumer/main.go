package main

import (
	"context"
	"encoding/gob"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"gopkg.in/mail.v2"

	"github.com/ahdark-services/anonymous-box-saas/internal/database"
	"github.com/ahdark-services/anonymous-box-saas/internal/logger"
	"github.com/ahdark-services/anonymous-box-saas/internal/mq"
	"github.com/ahdark-services/anonymous-box-saas/internal/redis"
	"github.com/ahdark-services/anonymous-box-saas/internal/trace"
	"github.com/ahdark-services/anonymous-box-saas/pkg/cache"
)

const (
	serviceName = "email-consumer"
)

var (
	ctx    = context.Background()
	tracer = otel.Tracer("main")
	app    *fx.App
)

func init() {
	ctx, span := tracer.Start(ctx, "init")
	defer span.End()

	gob.Register(mail.Message{})

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
		mq.Module(),
		fx.Invoke(run),
	)
}

func main() {
	_, span := tracer.Start(ctx, "main")
	defer span.End()

	app.Run()
}
