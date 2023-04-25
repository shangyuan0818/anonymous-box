package v1

import (
	"context"

	"github.com/cloudwego/hertz/pkg/route"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1/auth"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1/comment"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1/verify"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1/website"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/middleware"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/authservice"
)

var tracer = otel.Tracer("services.gateway.server.controller")

type DashApiV1RouterGroup struct {
	*route.RouterGroup
}

type Services struct {
	fx.In
	Auth    auth.Service
	Verify  verify.Service
	Website website.Service
	Comment comment.Service
}

type Params struct {
	fx.In
	RouterGroup *DashApiV1RouterGroup
	AuthClient  authservice.Client
}

func BindRoutes(ctx context.Context, svc Services, params Params) {
	ctx, span := tracer.Start(ctx, "bind-routes")
	defer span.End()

	api := params.RouterGroup
	api.Use(middleware.JwtParser())
	api.Use(middleware.AuthDataParser(params.AuthClient))
	{
		authService := api.Group("/auth")
		{
			authService.GET("", middleware.MustAuth(), svc.Auth.GetAuthData)

			authService.POST("/login", svc.Auth.Login)
			authService.POST("/register", middleware.MustNotAuth(), svc.Auth.Register)
			authService.POST("/change-password", middleware.MustAuth(), svc.Auth.ChangePassword)
			authService.POST("/reset-password", middleware.MustNotAuth(), svc.Auth.ResetPassword)
		}

		verifyService := api.Group("/verify")
		{
			verifyService.POST("/email", svc.Verify.ApplyEmailVerify)
		}

		websiteService := api.Group("/website", middleware.MustAuth())
		{
			websiteService.GET("", svc.Website.ListWebsites)
			websiteService.GET("/:id", svc.Website.GetWebsite)
			websiteService.POST("", svc.Website.CreateWebsite)
			websiteService.PUT("/:id", svc.Website.UpdateWebsite)
		}

		commentService := api.Group("/comment", middleware.MustAuth())
		{
			commentService.GET("", svc.Comment.ListComments)
			commentService.GET("/:id", svc.Comment.GetComment)
			commentService.DELETE("/:id", svc.Comment.DeleteComment)
		}
	}
}
