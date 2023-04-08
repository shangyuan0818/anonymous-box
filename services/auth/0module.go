package auth

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("auth-service",
		fx.Provide(NewAuthServiceImpl),
		fx.Provide(NewAuthServiceClient),
	)
}
