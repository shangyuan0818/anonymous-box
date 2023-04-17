package jwt

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.jwt")

type Service interface {
	GenerateToken(ctx context.Context, userID uint64) (string, error)
	ParseToken(ctx context.Context, token string) (uint64, error)
}

type service struct {
	secret string
	expire time.Duration
}

func NewService(ctx context.Context, e *config.JwtEnv) Service {
	ctx, span := tracer.Start(ctx, "jwt-new-service")
	defer span.End()

	return &service{
		secret: e.Secret,
		expire: time.Duration(e.Expire) * time.Second,
	}
}
