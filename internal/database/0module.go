package database

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal/database/repo"
)

func Module() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(NewDB),
		fx.Provide(NewQuery),

		fx.Provide(repo.NewSettingRepo),
		fx.Provide(repo.NewUserRepo),
	)
}
