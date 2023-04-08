package email

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"gopkg.in/mail.v2"

	"github.com/star-horizon/anonymous-box-saas/internal/database/dal"
	"github.com/star-horizon/anonymous-box-saas/internal/database/model"
	"github.com/star-horizon/anonymous-box-saas/services/email/kitex_gen/api"
)

var tracer = otel.Tracer("email-service")

// MailServiceImpl implements the last service interface defined in the IDL.
type MailServiceImpl struct {
	fx.In

	Q  *dal.Query
	MQ *amqp.Channel
}

// NewMailServiceImpl creates a new MailServiceImpl.
func NewMailServiceImpl(impl MailServiceImpl) api.MailService {
	return &impl
}

// SendMail implements the MailServiceImpl interface.
func (s *MailServiceImpl) SendMail(ctx context.Context, req *api.SendMailRequest) (*api.SendMailResponse, error) {
	ctx, span := tracer.Start(ctx, "send-mail")
	defer span.End()

	settingSlice, err := s.Q.Setting.WithContext(ctx).
		Where(s.Q.Setting.Type.Eq(string(model.SettingTypeEmail))).
		Find()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("query setting failed")
		return nil, err
	}

	settingMap := lo.SliceToMap(settingSlice, func(setting *model.Setting) (string, string) {
		return setting.Name, setting.Value
	})

	m := mail.NewMessage(
		mail.SetEncoding(mail.Base64),
		mail.SetCharset("UTF-8"),
	)

	m.SetAddressHeader("From", settingMap["email_from"], settingMap["email_from_name"])
	m.SetHeader("To", req.GetTo())
	m.SetHeader("Subject", req.GetSubject())
	switch req.Type {
	case api.MailType_MAIL_TYPE_UNKNOWN:
		m.SetBody("text/plain", req.GetBody())
	case api.MailType_MAIL_TYPE_HTML:
		m.SetBody("text/html", req.GetBody())
	case api.MailType_MAIL_TYPE_TEXT:
		m.SetBody("text/plain", req.GetBody())
	}

	buff := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buff).Encode(m); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("encode mail message failed")
		return &api.SendMailResponse{
			Success: false,
		}, err
	}

	if err := s.MQ.PublishWithContext(
		ctx,
		"email",
		"email.send",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/x-gob",
			Body:        buff.Bytes(),
			Timestamp:   time.Now(),
		},
	); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("publish message failed")
		return nil, err
	}

	return &api.SendMailResponse{
		Success: true,
	}, nil
}
