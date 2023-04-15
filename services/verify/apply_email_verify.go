package verify

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	emailapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
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

var (
	RegexEmailAddress = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

	ErrInvalidEmailAddress = errors.New("invalid email address")
)

// ApplyEmailVerify implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) ApplyEmailVerify(ctx context.Context, req *api.ApplyEmailVerifyRequest) (*base.Empty, error) {
	ctx, span := tracer.Start(ctx, "apply-email-verify", trace.WithAttributes(
		attribute.String("params.email", req.GetEmail()),
	))
	defer span.End()

	logger := logrus.WithContext(ctx).WithField("params.email", req.GetEmail())

	// validate email address
	if !RegexEmailAddress.MatchString(req.GetEmail()) {
		logger.WithError(ErrInvalidEmailAddress).Errorf("invalid email address: %s", req.GetEmail())
		return nil, ErrInvalidEmailAddress
	}

	// get app name from setting table
	appName, err := s.SettingRepo.GetByName(ctx, "app_name")
	if err != nil {
		logger.WithError(err).Error("query setting failed")
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	// send email via email service api
	if _, err := s.EmailSvcClient.SendMail(ctx, &emailapi.SendMailRequest{
		Type: lo.Switch[string, emailapi.MailType](settings["email_template_verify_code_content_type"]).
			Case("text/plain", emailapi.MailType_MAIL_TYPE_TEXT).
			Case("text/html", emailapi.MailType_MAIL_TYPE_HTML).
			Default(emailapi.MailType_MAIL_TYPE_UNKNOWN),
		To:      req.GetEmail(),
		Subject: fmt.Sprintf("[%s] 验证你的电子邮件地址", appName),
		Body:    content,
	}); err != nil {
		logger.WithError(err).Error("send email failed")

		return nil, err
	}

	return &base.Empty{}, nil
}
