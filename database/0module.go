package database

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/internal/hashids"
)

func Module() fx.Option {
	return fx.Module(
		"database",
		hashids.Module(),

		fx.Provide(NewDB),
		fx.Provide(NewQuery),

		fx.Provide(repo.NewSettingRepo),
		fx.Provide(repo.NewUserRepo),
		fx.Provide(repo.NewWebsiteRepo),
		fx.Provide(repo.NewCommentRepo),
	)
}
