package server

import (
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/auth"
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/verify"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewServer),

		fx.Provide(auth.NewController),
		fx.Provide(verify.NewController),

		fx.Invoke(controller.BindRoutes),
	)
}
