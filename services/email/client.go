package email

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/emailservice"
)

// NewEmailServiceClient creates a new MailServiceClient.
func NewEmailServiceClient(ctx context.Context, resolver discovery.Resolver) (emailservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-email-service-client")
	defer span.End()

	return emailservice.NewClient(
		ServiceName,
		client.WithResolver(resolver),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
