package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/authservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.dash.dash.v1.auth")

type Controller struct {
	fx.In
	AuthSvcClient authservice.Client
}

type Service interface {
	UsernameLogin(ctx context.Context, c *app.RequestContext)  // UsernameLogin provides a method to log in with username.
	EmailLogin(ctx context.Context, c *app.RequestContext)     // EmailLogin provides a method to log in with email.
	Register(ctx context.Context, c *app.RequestContext)       // Register provides a method to register a new user.
	ChangePassword(ctx context.Context, c *app.RequestContext) // ChangePassword provides a method to change password.
	ResetPassword(ctx context.Context, c *app.RequestContext)  // ResetPassword provides a method to reset password.
	GetAuthData(ctx context.Context, c *app.RequestContext)    // GetAuthData provides a method to get user's self data.
}

func NewController(impl Controller) Service {
	return &impl
}
