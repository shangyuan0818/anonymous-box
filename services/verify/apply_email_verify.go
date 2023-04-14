package verify

import (
	"context"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	emailapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/pkg/util"
)

type VerifyEmailTemplate struct {
	EmailAddress string
	VerifyCode   string

	AppName        string
	AppDescription string
	AppUrl         string
}

// generateVerifyCode generates a random 6-digit number
func (s *VerifyServiceImpl) generateVerifyCode(ctx context.Context, email string) (string, error) {
	ctx, span := tracer.Start(ctx, "generate-verify-code")
	defer span.End()

	code := util.RandString(6)
	code = strings.ToUpper(code)

	if err := s.Cache.Set(ctx, fmt.Sprint("verify_service::email_verify_code::", email), code, 5*time.Minute); err != nil {
		return "", err
	}

	return code, nil
}

// ApplyEmailVerify implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) ApplyEmailVerify(ctx context.Context, req *api.ApplyEmailVerifyRequest) (*api.ApplyEmailVerifyResponse, error) {
	ctx, span := tracer.Start(ctx, "apply-email-verify", trace.WithAttributes(
		attribute.String("params.email", req.GetEmail()),
	))
	defer span.End()

	logger := logrus.WithContext(ctx).WithField("params.email", req.GetEmail())

	// get app name from setting table
	appName, err := s.SettingRepo.GetByName(ctx, "app_name")
	if err != nil {
		logger.WithError(err).Error("query setting failed")
		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, err
	}

	settings, err := s.SettingRepo.ListByNames(ctx, []string{
		"app_name",
		"app_description",
		"app_url",
		"email_template_verify_code",
		"email_template_verify_code_content_type",
	})

	code, err := s.generateVerifyCode(ctx, req.GetEmail())
	if err != nil {
		logger.WithError(err).Error("generate verify code failed")
		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, err
	}

	content, err := util.RenderTemplate(settings["email_template_verify_code"], &VerifyEmailTemplate{
		EmailAddress:   req.GetEmail(),
		VerifyCode:     code,
		AppName:        settings["app_name"],
		AppDescription: settings["app_description"],
		AppUrl:         settings["app_url"],
	})
	if err != nil {
		logger.WithError(err).Error("render email template failed")
		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, err
	}

	// send email via email service api
	if r, err := s.MailSvcClient.SendMail(ctx, &emailapi.SendMailRequest{
		Type: lo.Switch[string, emailapi.MailType](settings["email_template_verify_code_content_type"]).
			Case("text/plain", emailapi.MailType_MAIL_TYPE_TEXT).
			Case("text/html", emailapi.MailType_MAIL_TYPE_HTML).
			Default(emailapi.MailType_MAIL_TYPE_UNKNOWN),
		To:      req.GetEmail(),
		Subject: fmt.Sprintf("[%s] 验证你的电子邮件地址", appName),
		Body:    content,
	}); err != nil {
		logger.WithError(err).Error("send email failed")

		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, err
	} else if !r.GetSuccess() {
		logger.Warn("send email failed, unknown reason")

		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, errors.New("send email failed, unknown reason")
	}

	return &api.ApplyEmailVerifyResponse{
		Email: req.Email,
		Ok:    true,
	}, nil
}
