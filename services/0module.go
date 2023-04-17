package services

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/services/auth"
	"github.com/star-horizon/anonymous-box-saas/services/comment"
	"github.com/star-horizon/anonymous-box-saas/services/email"
	"github.com/star-horizon/anonymous-box-saas/services/email_consumer"
	"github.com/star-horizon/anonymous-box-saas/services/verify"
	"github.com/star-horizon/anonymous-box-saas/services/website"
)

func Module() fx.Option {
	return fx.Module(
		"services",
		auth.Module(),
		comment.Module(),
		email.Module(),
		email_consumer.Module(),
		verify.Module(),
		website.Module(),
	)
}
