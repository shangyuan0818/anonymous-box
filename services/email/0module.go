package email

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("email-service",
		fx.Provide(NewEmailServiceImpl),
		fx.Provide(NewEmailServiceClient),
	)
}
