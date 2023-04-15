package verify

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/verifyservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.dash.dash.v1.verify")

type Controller struct {
	fx.In
	VerifySvcClient verifyservice.Client
}

type Service interface {
	ApplyEmailVerify(ctx context.Context, c *app.RequestContext) // ApplyEmailVerify provides a method to apply email verify.
}

func NewController(impl Controller) Service {
	return &impl
}
