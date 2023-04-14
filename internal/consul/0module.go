package consul

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("internal.consul")

func Module() fx.Option {
	return fx.Module(
		"consul",
		fx.Provide(NewConfig),
		fx.Provide(NewClient),
		fx.Provide(NewKitexConsulRegistry),
		fx.Provide(NewKitexConsulResolver),
		fx.Provide(NewHertzConsulRegistry),
	)
}
