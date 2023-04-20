package dash_auth

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewAuthServiceImpl),
		fx.Provide(NewAuthServiceClient),
	)
}
