package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/bootstrap"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/commentservice"
	"github.com/star-horizon/anonymous-box-saas/services/dash-comment"
)

var tracer = otel.Tracer("main")

func main() {
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	app := bootstrap.InitApp(
		ctx,
		dash_comment.ServiceName,
		fx.Invoke(bootstrap.RunService(commentservice.NewServer)),
	)

	app.Run()
}
