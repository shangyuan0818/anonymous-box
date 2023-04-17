package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/emailservice"
	"github.com/star-horizon/anonymous-box-saas/services/email"
)

var tracer = otel.Tracer("main")

func main() {
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	app := bootstrap.InitApp(
		ctx,
		email.ServiceName,
		fx.Invoke(bootstrap.RunService(emailservice.NewServer)),
	)

	app.Run()
}
