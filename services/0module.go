package services

import (
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/services/dash-auth"
	"github.com/star-horizon/anonymous-box-saas/services/dash-comment"
	"github.com/star-horizon/anonymous-box-saas/services/dash-verify"
	"github.com/star-horizon/anonymous-box-saas/services/dash-website"
	"github.com/star-horizon/anonymous-box-saas/services/email"
	"github.com/star-horizon/anonymous-box-saas/services/email-consumer"
)

func Module() fx.Option {
	return fx.Module(
		"services",
		dash_auth.Module(),
		dash_comment.Module(),
		email.Module(),
		email_consumer.Module(),
		dash_verify.Module(),
		dash_website.Module(),
	)
}
