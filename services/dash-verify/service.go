package dash_verify

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/emailservice"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

var tracer = otel.Tracer(ServiceName)

const ServiceName = "dash-verify-service"

// VerifyServiceImpl implements the last service interface defined in the IDL.
type VerifyServiceImpl struct {
	fx.In
	EmailSvcClient emailservice.Client
	SettingRepo    repo.SettingRepo
	Cache          cache.Driver
}

// NewVerifyServiceImpl creates a new VerifyServiceImpl.
func NewVerifyServiceImpl(impl VerifyServiceImpl) dash.VerifyService {
	return &impl
}

var (
	ErrVerifyCodeNotFound = errors.New("verify code not found")
	ErrVerifyCodeInvalid  = errors.New("verify code invalid")
)

// VerifyEmail implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) VerifyEmail(ctx context.Context, req *dash.VerifyEmailRequest) (*base.Empty, error) {
	ctx, span := tracer.Start(ctx, "verify-email")
	defer span.End()

	v, exist := s.Cache.Get(ctx, fmt.Sprint(ServiceName, ":email_verify_code:", req.GetEmail()))
	if !exist {
		return nil, ErrVerifyCodeNotFound
	}

	if v != req.GetCode() {
		return nil, ErrVerifyCodeInvalid
	}

	return &base.Empty{}, nil
}
