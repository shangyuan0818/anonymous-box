package v1

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"gateway.dash.box.v1",
		fx.Invoke(BindRoutes),
	)
}
