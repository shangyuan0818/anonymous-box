package verify

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("verify-service",
		fx.Provide(NewVerifyServiceImpl),
		fx.Provide(NewVerifyServiceClient),
	)
}
