package verify

import (
	"context"
	"errors"
	"fmt"
	"github.com/star-horizon/anonymous-box-saas/database/repo"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

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
)

// VerifyEmail implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) VerifyEmail(ctx context.Context, req *api.VerifyEmailRequest) (resp *api.VerifyEmailResponse, err error) {
	ctx, span := tracer.Start(ctx, "verify-email")
	defer span.End()

	v, exist := s.Cache.Get(ctx, fmt.Sprint("verify_service::email_verify_code::", req.GetEmail()))
	if !exist {
		return &api.VerifyEmailResponse{
			Ok: false,
		}, ErrVerifyCodeNotFound
	}

	if v != req.GetCode() {
		return &api.VerifyEmailResponse{
			Ok: false,
		}, ErrVerifyCodeNotFound
	}

	return &api.VerifyEmailResponse{
		Ok: true,
	}, nil
}
