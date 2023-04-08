package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api/authservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.auth")

type Controller struct {
	fx.In
	AuthSvcClient authservice.Client
}

type Service interface {
	EmailLogin(ctx context.Context, c *app.RequestContext)
	UsernameLogin(ctx context.Context, c *app.RequestContext)
	Register(ctx context.Context, c *app.RequestContext)
	ChangePassword(ctx context.Context, c *app.RequestContext)
	ResetPassword(ctx context.Context, c *app.RequestContext)
}

func NewController(impl Controller) Service {
	return &impl
}
