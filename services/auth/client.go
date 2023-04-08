package auth

import (
	"context"

	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"

	"github.com/ahdark-services/anonymous-box-saas/services/auth/kitex_gen/api/authservice"
)

// NewAuthServiceClient creates a new AuthServiceClient.
func NewAuthServiceClient(ctx context.Context) (authservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-auth-service-client")
	defer span.End()

	return authservice.NewClient("auth-service", client.WithResolver(dns.NewDNSResolver()))
}
