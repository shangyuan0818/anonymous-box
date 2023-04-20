package email_consumer

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewEmailServiceConsumer),
	)
}
