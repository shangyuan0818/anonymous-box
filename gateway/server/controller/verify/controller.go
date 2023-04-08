package verify

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api/verifyservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.verify")

type Controller struct {
	fx.In
	VerifySvcClient verifyservice.Client
}

type Service interface {
	ApplyEmailVerify(ctx context.Context, c *app.RequestContext)
}

func NewController(impl Controller) Service {
	return &impl
}
