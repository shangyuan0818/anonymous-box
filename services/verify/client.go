package verify

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/verifyservice"
)

// NewVerifyServiceClient returns a new instance of the VerifyClient
func NewVerifyServiceClient(ctx context.Context, resolver discovery.Resolver) (verifyservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-verify-service-client")
	defer span.End()

	return verifyservice.NewClient(
		ServiceName,
		client.WithResolver(resolver),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
