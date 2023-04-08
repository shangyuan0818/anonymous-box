package email

import (
	"context"

	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"

	"github.com/star-horizon/anonymous-box-saas/services/email/kitex_gen/api/mailservice"
)

// NewMailServiceClient creates a new MailServiceClient.
func NewMailServiceClient(ctx context.Context) (mailservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-mail-service-client")
	defer span.End()

	return mailservice.NewClient("email-service", client.WithResolver(dns.NewDNSResolver()))
}
