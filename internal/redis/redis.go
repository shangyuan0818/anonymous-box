package redis

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

type env struct {
	Network  string `default:"tcp"`
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Username string `default:""`
	Password string `default:""`
	DB       int    `default:"0"`
}

var tracer = otel.Tracer("internal.redis")

func NewRedis(ctx context.Context) (*redis.Client, error) {
	ctx, span := tracer.Start(ctx, "init-redis")
	defer span.End()

	var e env
	if err := envconfig.Process("REDIS", &e); err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Network:  e.Network,
		Addr:     fmt.Sprintf("%s:%d", e.Host, e.Port),
		Username: e.Username,
		Password: e.Password,
		DB:       e.DB,
	})

	if err := redisotel.InstrumentTracing(client); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to instrument tracing")
		return nil, err
	}

	return client, nil
}
