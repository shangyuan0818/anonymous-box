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

func NewConnection(ctx context.Context, e *config.MqEnv) (*amqp.Connection, error) {
	ctx, span := tracer.Start(ctx, "new-amqp-connection")
	defer span.End()

	// init mq
	conn, err := amqp.DialConfig(fmt.Sprintf("amqp://%s:%s@%s:%d/", e.User, e.Password, e.Host, e.Port), amqp.Config{
		Vhost: e.Vhost,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("init mq failed")
		return nil, err
	}

	return conn, nil
}

func NewChannel(ctx context.Context, conn *amqp.Connection) (*amqp.Channel, error) {
	ctx, span := tracer.Start(ctx, "new-amqp-channel")
	defer span.End()

	ch, err := conn.Channel()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("init mq channel failed")
		return nil, err
	}

	return ch, nil
}
