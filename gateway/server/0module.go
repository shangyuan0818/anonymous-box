package server

import (
	"github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/v1"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewServer),
		v1.Module(),
	)
}
