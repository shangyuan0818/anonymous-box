package hashids

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"hashids",
		fx.Provide(NewService),
	)
}
