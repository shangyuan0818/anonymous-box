package infra

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal/consul"
	"github.com/star-horizon/anonymous-box-saas/internal/logger"
	"github.com/star-horizon/anonymous-box-saas/internal/trace"
)

func Module() fx.Option {
	return fx.Options(
		logger.Module(),
		consul.Module(),
		trace.Module(),
	)
}
