package mq

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("internal.mq")

type env struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5672"`
	User     string `default:"guest"`
	Password string `default:"guest"`
	Vhost    string `default:"/"`
}

func InitMQ(ctx context.Context) (*amqp.Channel, error) {
	ctx, span := tracer.Start(ctx, "InitMQ")
	defer span.End()

	// init env
	var e env
	if err := envconfig.Process("MQ", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("init env failed")
	}

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
