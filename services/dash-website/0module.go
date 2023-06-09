package dash_website

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewWebsiteService),
		fx.Provide(NewWebsiteServiceClient),
	)
}
