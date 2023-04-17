package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.redis")

func NewRedis(ctx context.Context, e *config.RedisEnv) (*redis.Client, error) {
	ctx, span := tracer.Start(ctx, "init-redis")
	defer span.End()

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
