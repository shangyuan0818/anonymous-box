package mq

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"mq",
		fx.Provide(NewConnection),
		fx.Provide(NewChannel),
	)
}
