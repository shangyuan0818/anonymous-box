package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/authservice"
	"github.com/star-horizon/anonymous-box-saas/services/auth"
)

var (
	tracer = otel.Tracer("main")
	ctx    = context.Background()
	app    *fx.App
)

func init() {
	ctx, span := tracer.Start(ctx, "init")
	defer span.End()

	app = bootstrap.InitApp(
		ctx,
		auth.ServiceName,
		fx.Invoke(bootstrap.RunService(authservice.NewServer)),
	)
}

func main() {
	_, span := tracer.Start(ctx, "main")
	defer span.End()

	app.Run()
}
