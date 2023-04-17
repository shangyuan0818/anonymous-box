package trace

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"trace",
		fx.Provide(NewExporter),
		fx.Provide(NewTracerProvider),
		fx.Invoke(InitTracer),
		fx.Invoke(KitexOpenTelemetryProvider),
	)
}
