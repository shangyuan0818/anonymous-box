package hashids

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids/v2"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.hashids")

var (
	ErrInvalidHash = errors.New("invalid hash")
)

type Service interface {
	Encode(ctx context.Context, id uint64, t HashType) (string, error)
	Decode(ctx context.Context, hash string, t HashType) (uint64, error)
}

type service struct {
	client *hashids.HashID
}

func NewService(ctx context.Context, e *config.HashidsEnv) (Service, error) {
	ctx, span := tracer.Start(ctx, "hashids-new-service")
	defer span.End()

	c, err := hashids.NewWithData(&hashids.HashIDData{
		MinLength: e.MinLength,
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		Salt:      e.Salt,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create hashids client")
		return nil, err
	}

	return &service{c}, nil
}

// Encode implements Service.Encode
func (s *service) Encode(ctx context.Context, id uint64, t HashType) (string, error) {
	ctx, span := tracer.Start(ctx, "hashids-encode")
	defer span.End()

	return s.client.EncodeInt64([]int64{int64(id), int64(t)})
}

// Decode implements Service.Decode
func (s *service) Decode(ctx context.Context, hash string, t HashType) (uint64, error) {
	ctx, span := tracer.Start(ctx, "hashids-decode")
	defer span.End()

	ids, err := s.client.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}

	if len(ids) != 2 {
		return 0, ErrInvalidHash
	}

	if HashType(ids[1]) != t {
		return 0, ErrInvalidHash
	}

	return uint64(ids[0]), nil
}
