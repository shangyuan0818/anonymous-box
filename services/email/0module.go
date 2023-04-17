package email

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewEmailServiceImpl),
		fx.Provide(NewEmailServiceClient),
	)
}
