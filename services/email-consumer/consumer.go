package email_consumer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"gopkg.in/mail.v2"

	"github.com/star-horizon/anonymous-box-saas/config"
	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/pkg/util"
)

func RunConsumer(ctx context.Context, conn *amqp.Connection, lc fx.Lifecycle, settingRepo repo.SettingRepo, env *config.EmailEnv) error {
	ctx, span := tracer.Start(ctx, "init-consumer")
	defer span.End()

	ch, err := conn.Channel()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("open channel failed")
		return err
	}

	// declare exchange
	if err := ch.ExchangeDeclare(MQExchangeName, "direct", true, false, false, false, nil); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("declare exchange failed")
		return err
	}

	// declare queue
	queue, err := ch.QueueDeclare(MQQueueName, true, false, false, false, nil)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("declare queue failed")
		return err
	}

	if err := ch.QueueBind(queue.Name, MQKeyEmailSend, MQExchangeName, false, nil); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("bind queue failed")
		return err
	}

	// init mail dialer
	dialer := mail.NewDialer(
		env.Host,
		env.Port,
		env.Username,
		env.Password,
	)

	// ssl
	dialer.SSL = env.SSL

	// tls
	if env.TLS {
		dialer.SSL = false
		dialer.TLSConfig = &tls.Config{
			ServerName:         env.Host,
			InsecureSkipVerify: true,
		}
	}

	consumerName := fmt.Sprintf("email-consumer-%s", config.ServiceInstanceID)

	// consume
	msgs, err := ch.Consume(queue.Name, consumerName, false, false, false, false, nil)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("consume queue failed")
		return err
	}

	consume := func(delivery amqp.Delivery) {
		ctx, span := tracer.Start(
			propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}).
				Extract(context.Background(), propagation.MapCarrier(util.InterfaceMapToTypedMap[string](delivery.Headers["trace-context"].(amqp.Table)))),
			"email-consume",
			trace.WithSpanKind(trace.SpanKindConsumer),
			trace.WithAttributes(
				attribute.String("delivery.routing-key", delivery.RoutingKey),
				attribute.String("delivery.exchange", delivery.Exchange),
				attribute.String("delivery.consumer-tag", delivery.ConsumerTag),
				attribute.Int64("delivery.delivery-tag", int64(delivery.DeliveryTag)),
				attribute.Bool("delivery.redelivered", delivery.Redelivered),
			),
		)
		defer span.End()

		// parse gob
		span.AddEvent("Parse Email Message")
		var email dash.EmailMessage
		switch delivery.ContentType {
		case "application/json", "application/x-json":
			if err := json.Unmarshal(delivery.Body, &email); err != nil {
				logrus.WithContext(ctx).WithError(err).Error("decode email failed")
				span.RecordError(err)
				delivery.Nack(false, false)
				return
			}
		default:
			logrus.WithContext(ctx).Errorf("invalid content type: %s", delivery.ContentType)
			delivery.Nack(false, false)
			span.SetStatus(codes.Error, fmt.Sprintf("invalid content type: %s", delivery.ContentType))
			return
		}

		span.AddEvent("Get Sender Address")
		fromAddress, err := settingRepo.GetByName(ctx, "email_from_address")
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("get from address failed")
			span.RecordError(err)
			delivery.Nack(false, false)
			return
		}

		span.AddEvent("Make Email")
		m := mail.NewMessage(
			mail.SetCharset("UTF-8"),
			mail.SetEncoding(mail.Base64),
		)
		m.SetAddressHeader("From", fromAddress, email.From)
		m.SetAddressHeader("To", email.To, email.To)
		m.SetHeader("Subject", email.Subject)
		m.SetBody(email.ContentType, email.Body)

		// send email
		if err := dialer.DialAndSend(m); err != nil {
			logrus.WithContext(ctx).WithError(err).Error("send email failed")
			span.RecordError(err)
			delivery.Nack(false, !delivery.Redelivered)
			return
		}
		span.AddEvent("Send Email", trace.WithAttributes(
			attribute.String("email.from", email.From),
			attribute.String("email.to", email.To),
			attribute.String("email.subject", email.Subject),
		))

		delivery.Ack(false)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "start-consumer")
			defer span.End()

			logrus.WithContext(ctx).Info("start consume email queue")
			go func() {
				for msg := range msgs {
					consume(msg)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "stop-consumer")
			defer span.End()

			if err := ch.Cancel(consumerName, false); err != nil {
				logrus.WithContext(ctx).WithError(err).Error("cancel consumer failed")
			}

			return nil
		},
	})

	return nil
}
