package verify

import (
	"context"

	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"

	"github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api/verifyservice"
)

// NewVerifyServiceClient returns a new instance of the VerifyClient
func NewVerifyServiceClient(ctx context.Context) (verifyservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-verify-service-client")
	defer span.End()

	return verifyservice.NewClient("verify-service", client.WithResolver(dns.NewDNSResolver()))
}
