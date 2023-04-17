package verify

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewVerifyServiceImpl),
		fx.Provide(NewVerifyServiceClient),
	)
}
