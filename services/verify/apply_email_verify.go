package verify

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/star-horizon/anonymous-box-saas/pkg/util"
	emailapi "github.com/star-horizon/anonymous-box-saas/services/email/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api"
)

type VerifyEmailTemplate struct {
	EmailAddress string
	VerifyCode   string

	AppName        string
	AppDescription string
	AppUrl         string
}

func (s *VerifyServiceImpl) renderEmailTemplate(ctx context.Context, raw string, tpl VerifyEmailTemplate) (string, error) {
	ctx, span := tracer.Start(ctx, "render-email-template")
	defer span.End()

	buff := bytes.NewBuffer(nil)
	t, err := template.New("email").Parse(raw)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("parse email template failed")
		return "", err
	}

	if err := t.Execute(buff, &tpl); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("execute email template failed")
		return "", err
	}

	return buff.String(), nil
}

// generateVerifyCode generates a random 6-digit number
func (s *VerifyServiceImpl) generateVerifyCode(ctx context.Context, email string) (string, error) {
	ctx, span := tracer.Start(ctx, "generate-verify-code")
	defer span.End()

	code := util.RandString(6)
	code = strings.ToUpper(code)

	if err := s.Cache.Set(ctx, fmt.Sprintf("verify_service::email_verify_code:%s", email), code, 60*5); err != nil {
		return "", err
	}

	return code, nil
}

// ApplyEmailVerify implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) ApplyEmailVerify(ctx context.Context, req *api.ApplyEmailVerifyRequest) (*api.ApplyEmailVerifyResponse, error) {
	ctx, span := tracer.Start(ctx, "apply-email-verify")
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
		"email_verify_template",
	})

	code, err := s.generateVerifyCode(ctx, req.GetEmail())
	if err != nil {
		logger.WithError(err).Error("generate verify code failed")
		return &api.ApplyEmailVerifyResponse{
			Email: req.Email,
			Ok:    false,
		}, err
	}

	content, err := s.renderEmailTemplate(ctx, settings["email_verify_template"], VerifyEmailTemplate{
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
		Type:    emailapi.MailType_MAIL_TYPE_HTML,
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
