package email

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/services/email_consumer"
)

var tracer = otel.Tracer(ServiceName)

const ServiceName = "email-service"

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct {
	fx.In

	Q  *dal.Query
	MQ *amqp.Channel
}

// NewEmailServiceImpl creates a new EmailServiceImpl.
func NewEmailServiceImpl(impl EmailServiceImpl) dash.EmailService {
	return &impl
}

// SendMail implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) SendMail(ctx context.Context, req *dash.SendMailRequest) (*base.Empty, error) {
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

	m := dash.EmailMessage{
		From:    settingMap["email_from_name"],
		To:      req.GetTo(),
		Subject: req.GetSubject(),
		Body:    req.GetBody(),
		ContentType: lo.Switch[dash.MailType, string](req.GetType()).
			Case(dash.MailType_MAIL_TYPE_TEXT, "text/plain").
			Case(dash.MailType_MAIL_TYPE_HTML, "text/html").
			Default("application/octet-stream"),
	}
	span.SetAttributes(
		attribute.String("email.from", m.From),
		attribute.String("email.to", m.To),
		attribute.String("email.subject", m.Subject),
		attribute.String("email.body", m.Body),
		attribute.String("email.content_type", m.ContentType),
	)

	data, err := json.Marshal(&m)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("marshal email message failed")
		return nil, err
	}

	if err := s.MQ.PublishWithContext(
		ctx,
		email_consumer.MQExchangeName,
		email_consumer.MQKeyEmailSend,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
			Headers: amqp.Table{
				"trace-id": span.SpanContext().TraceID().String(),
			},
		},
	); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("publish message failed")
		return nil, err
	}

	return &base.Empty{}, nil
}
