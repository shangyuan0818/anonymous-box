package comment

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		ServiceName,
		fx.Provide(NewCommentService),
		fx.Provide(NewCommentServiceClient),
	)
}
