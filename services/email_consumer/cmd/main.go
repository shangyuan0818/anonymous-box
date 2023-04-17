package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/emailserviceconsumer"
	"github.com/star-horizon/anonymous-box-saas/services/email_consumer"
)

var tracer = otel.Tracer("main")

func main() {
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	app := bootstrap.InitApp(
		ctx,
		email_consumer.ServiceName,
		fx.Invoke(email_consumer.RunConsumer),
		fx.Invoke(bootstrap.RunService(emailserviceconsumer.NewServer)),
	)

	app.Run()
}
