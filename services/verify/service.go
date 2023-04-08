package verify

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/anonymous-box-saas/internal/database/repo"
	"github.com/ahdark-services/anonymous-box-saas/pkg/cache"
	"github.com/ahdark-services/anonymous-box-saas/services/email/kitex_gen/api/mailservice"
	"github.com/ahdark-services/anonymous-box-saas/services/verify/kitex_gen/api"
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

// VerifyEmail implements the VerifyServiceImpl interface.
func (s *VerifyServiceImpl) VerifyEmail(ctx context.Context, req *api.VerifyEmailRequest) (resp *api.VerifyEmailResponse, err error) {
	// TODO: Your code here...
	return
}
