package gateway

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("gateway-service.server.controller.gateway")

type Controller struct {
	fx.In
}

type Service interface {
	Health(ctx context.Context, c *app.RequestContext)
}

func NewController(impl Controller) Service {
	return &impl
}
