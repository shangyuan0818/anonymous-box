package v1

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/v1/auth"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/v1/comment"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/v1/verify"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/v1/website"
)

func Module() fx.Option {
	return fx.Module(
		"gateway.api.v1",
		fx.Provide(auth.NewController),
		fx.Provide(verify.NewController),
		fx.Provide(website.NewController),
		fx.Provide(comment.NewController),

		fx.Invoke(BindRoutes),
	)
}
