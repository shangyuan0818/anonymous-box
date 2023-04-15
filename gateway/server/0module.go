package server

import (
	"go.uber.org/fx"

	boxv1 "github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/box/v1"
	dashv1 "github.com/star-horizon/anonymous-box-saas/gateway/server/controller/api/dash/v1"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewServer),
		fx.Provide(
			NewDashApiV1RouterGroup,
			NewBoxApiV1RouterGroup,
		),

		dashv1.Module(), // for /dash/v1
		boxv1.Module(),  // for /dash/box/v1
	)
}
