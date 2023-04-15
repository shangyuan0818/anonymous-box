package v1

import (
	"context"

	"github.com/cloudwego/hertz/pkg/route"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("gateway-service.server.controller.dash.box.v1")

type BoxApiV1RouterGroup struct {
	*route.RouterGroup
}

type Services struct {
	fx.In
}

type Params struct {
	fx.In
	RouterGroup *BoxApiV1RouterGroup
}

func BindRoutes(ctx context.Context, svc Services, params Params) {
	ctx, span := tracer.Start(ctx, "bind-routes")
	defer span.End()

	//boxapi := params.Server.Group("/dash/v1/box")
	{

	}
}
