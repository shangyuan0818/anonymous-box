package hashids

import (
	"context"
	"errors"
	"github.com/speps/go-hashids/v2"
	"go.opentelemetry.io/otel"
	"strconv"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
)

var tracer = otel.Tracer("internal.hashids")

var (
	ErrInvalidHash = errors.New("invalid hash")
)

type Service interface {
	getClient(ctx context.Context) (*hashids.HashID, error)
	Encode(ctx context.Context, id uint64, t HashType) (string, error)
	Decode(ctx context.Context, hash string, t HashType) (uint64, error)
}

type service struct {
	SettingRepo repo.SettingRepo
}

func NewService(svc service) Service {
	return &svc
}

func (s *service) getClient(ctx context.Context) (*hashids.HashID, error) {
	ctx, span := tracer.Start(ctx, "hashids-get-client")
	defer span.End()

	settings, err := s.SettingRepo.ListByNames(ctx, []string{
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
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Encode implements Service.Encode
func (s *service) Encode(ctx context.Context, id uint64, t HashType) (string, error) {
	ctx, span := tracer.Start(ctx, "hashids-encode")
	defer span.End()

	c, err := s.getClient(ctx)
	if err != nil {
		return "", err
	}

	return c.EncodeInt64([]int64{int64(id), int64(t)})
}

// Decode implements Service.Decode
func (s *service) Decode(ctx context.Context, hash string, t HashType) (uint64, error) {
	ctx, span := tracer.Start(ctx, "hashids-decode")
	defer span.End()

	c, err := s.getClient(ctx)
	if err != nil {
		return 0, err
	}

	ids, err := c.DecodeInt64WithError(hash)
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
