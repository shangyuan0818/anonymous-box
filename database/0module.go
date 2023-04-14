package database

import (
	repo2 "github.com/star-horizon/anonymous-box-saas/database/repo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(NewDB),
		fx.Provide(NewQuery),

		fx.Provide(repo2.NewSettingRepo),
		fx.Provide(repo2.NewUserRepo),
	)
}
