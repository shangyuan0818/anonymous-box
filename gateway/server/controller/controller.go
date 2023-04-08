package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/auth"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/gateway"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/verify"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/middleware"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("services.gateway.server.controller")

type Params struct {
	fx.In
	Auth    auth.Service
	Verify  verify.Service
	Gateway gateway.Service
}

func BindRoutes(ctx context.Context, r *server.Hertz, svc Params) {
	ctx, span := tracer.Start(ctx, "bind-routes")
	defer span.End()

	authService := r.Group("/auth")
	{
		authService.POST("/login/email", svc.Auth.EmailLogin)
		authService.POST("/login/username", svc.Auth.UsernameLogin)
		authService.POST("/register", svc.Auth.Register)
		authService.POST("/change-password", middleware.JwtParser(true), svc.Auth.ChangePassword)
		authService.POST("/reset-password", svc.Auth.ResetPassword)
	}

	verifyService := r.Group("/verify")
	{
		verifyService.POST("/email", svc.Verify.ApplyEmailVerify)
	}

	gatewayService := r.Group("/gateway")
	{
		gatewayService.GET("/health", svc.Gateway.Health)
	}
}
