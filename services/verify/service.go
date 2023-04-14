package verify

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/mailservice"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

var tracer = otel.Tracer("verify-service")

// VerifyServiceImpl implements the last service interface defined in the IDL.
type VerifyServiceImpl struct {
	fx.In
	MailSvcClient mailservice.Client
	SettingRepo   repo.SettingRepo
	Cache         cache.Driver
}

// NewVerifyServiceImpl creates a new VerifyServiceImpl.
func NewVerifyServiceImpl(impl VerifyServiceImpl) api.VerifyService {
	return &impl
}

var (
	ErrVerifyCodeNotFound = errors.New("verify code not found")
	ErrVerifyCodeInvalid  = errors.New("verify code invalid")
)

// VerifyEmail implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) VerifyEmail(ctx context.Context, req *api.VerifyEmailRequest) (*emptypb.Empty, error) {
	ctx, span := tracer.Start(ctx, "verify-email")
	defer span.End()

	v, exist := s.Cache.Get(ctx, fmt.Sprint("verify_service::email_verify_code::", req.GetEmail()))
	if !exist {
		return nil, ErrVerifyCodeNotFound
	}

	if v != req.GetCode() {
		return nil, ErrVerifyCodeInvalid
	}

	return &emptypb.Empty{}, nil
}
