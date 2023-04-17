package mq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.mq")

func InitMQ(ctx context.Context, e *config.MqEnv) (*amqp.Channel, error) {
	ctx, span := tracer.Start(ctx, "InitMQ")
	defer span.End()

	// init mq
	conn, err := amqp.DialConfig(fmt.Sprintf("amqp://%s:%s@%s:%d/", e.User, e.Password, e.Host, e.Port), amqp.Config{
		Vhost: e.Vhost,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("init mq failed")
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("init mq channel failed")
		return nil, err
	}

	return ch, nil
}
