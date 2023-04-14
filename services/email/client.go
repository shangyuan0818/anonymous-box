package email

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/mailservice"
)

// NewMailServiceClient creates a new MailServiceClient.
func NewMailServiceClient(ctx context.Context, resolver discovery.Resolver) (mailservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-mail-service-client")
	defer span.End()

	return mailservice.NewClient(
		"email-service",
		client.WithResolver(resolver),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
