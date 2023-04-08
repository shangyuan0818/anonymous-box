package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gopkg.in/mail.v2"

	"github.com/star-horizon/anonymous-box-saas/internal/database/dal"
	"github.com/star-horizon/anonymous-box-saas/internal/database/model"
)

const (
	mqExchangeName = "email-exchange"
	mqQueueName    = "email-queue"
)

func run(ctx context.Context, ch *amqp.Channel, q *dal.Query, lc fx.Lifecycle) error {
	ctx, span := tracer.Start(ctx, "init-consumer")
	defer span.End()

	// declare exchange
	if err := ch.ExchangeDeclare(mqExchangeName, "direct", true, false, false, false, nil); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("declare exchange failed")
		return err
	}

	// declare queue
	queue, err := ch.QueueDeclare(mqQueueName, true, false, false, false, nil)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("declare queue failed")
		return err
	}

	if err := ch.QueueBind(queue.Name, "email.send", mqExchangeName, false, nil); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("bind queue failed")
		return err
	}

	settingSlice, err := q.WithContext(ctx).
		Setting.Where(q.Setting.Type.Eq(string(model.SettingTypeEmail))).
		Find()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("query setting failed")
		return err
	}

	settingMap := lo.SliceToMap(settingSlice, func(setting *model.Setting) (string, string) {
		return setting.Name, setting.Value
	})

	// init mail dialer
	dialer := mail.NewDialer(settingMap["email_host"], 465, settingMap["email_username"], settingMap["email_password"])

	// ssl
	if v, err := strconv.ParseBool(settingMap["email_ssl"]); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("parse ssl failed")
		return err
	} else {
		dialer.SSL = v
	}

	// tls
	if v, err := strconv.ParseBool(settingMap["email_tls"]); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("parse tls failed")
		return err
	} else {
		if v {
			dialer.StartTLSPolicy = mail.MandatoryStartTLS
		} else {
			dialer.StartTLSPolicy = mail.NoStartTLS
		}
	}

	consumerName := fmt.Sprintf("email-consumer-%s", uuid.New().String())

	// consume
	msgs, err := ch.Consume(queue.Name, consumerName, true, false, false, false, nil)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("consume queue failed")
		return err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "start-consumer")
			defer span.End()

			logrus.WithContext(ctx).Info("start consume email queue")
			go func() {
				for msg := range msgs {
					func(msg amqp.Delivery) {
						ctx, span := tracer.Start(context.Background(), "email-consume")
						defer span.End()

						// parse gob
						buff := bytes.NewBuffer(msg.Body)
						dec := gob.NewDecoder(buff)
						var email mail.Message
						if err := dec.Decode(&email); err != nil {
							logrus.WithContext(ctx).WithError(err).Error("decode email failed")
							msg.Nack(false, false)
							return
						}

						// send email
						if err := dialer.DialAndSend(&email); err != nil {
							logrus.WithContext(ctx).WithError(err).Error("send email failed")
							msg.Nack(false, true)
							return
						}

						msg.Ack(false)
					}(msg)
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
