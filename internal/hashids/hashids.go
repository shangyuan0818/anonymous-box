package hashids

import (
	"context"
	"errors"
	"strconv"

	"github.com/speps/go-hashids/v2"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
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

func NewService(ctx context.Context, settingRepo repo.SettingRepo) (Service, error) {
	ctx, span := tracer.Start(context.Background(), "new-hashids-service")
	defer span.End()

	settings, err := settingRepo.ListByNames(ctx, []string{
		"app_hashids_salt",
		"app_hashids_min_length",
		"app_hashids_alphabet",
	})
	if err != nil {
		return nil, err
	}

	minLength, err := strconv.Atoi(settings["app_hashids_min_length"])
	if err != nil {
		return nil, err
	}

	c, err := hashids.NewWithData(&hashids.HashIDData{
		MinLength: minLength,
		Alphabet:  settings["app_hashids_alphabet"],
		Salt:      settings["app_hashids_salt"],
	})

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
