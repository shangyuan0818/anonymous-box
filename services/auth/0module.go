package auth

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal/jwt"
)

func Module() fx.Option {
	return fx.Module("auth-service",
		jwt.Module(),
		fx.Provide(NewAuthServiceImpl),
		fx.Provide(NewAuthServiceClient),
	)
}
