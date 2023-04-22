package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewEnvConfig[T any](prefix string, out T) func(ctx context.Context) (*T, error) {
	if err := envconfig.Process(prefix, &out); err != nil {
		return func(ctx context.Context) (*T, error) {
			ctx, span := tracer.Start(ctx, fmt.Sprintf("new-%s-env-config", strings.ToLower(prefix)))
			defer span.End()

			span.RecordError(err)

			return nil, err
		}
	}

	return func(ctx context.Context) (*T, error) {
		ctx, span := tracer.Start(ctx, fmt.Sprintf("new-%s-env-config", strings.ToLower(prefix)), trace.WithAttributes(
			attribute.String("prefix", prefix),
			attribute.String("out", fmt.Sprint(out)),
		))
		defer span.End()

		return &out, nil
	}
}
