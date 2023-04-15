package jwt

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("jwt",
		fx.Provide(NewService),
	)
}
