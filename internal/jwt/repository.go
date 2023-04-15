package jwt

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
)

var tracer = otel.Tracer("internal.jwt")

type Service interface {
	GenerateToken(ctx context.Context, userID uint64) (string, error)
	ParseToken(ctx context.Context, token string) (uint64, error)
}

type service struct {
	fx.In
	SettingRepo repo.SettingRepo
}

func NewService(svc service) Service {
	return &svc
}
