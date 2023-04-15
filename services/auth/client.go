package auth

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/authservice"
)

// NewAuthServiceClient creates a new AuthServiceClient.
func NewAuthServiceClient(ctx context.Context, resolver discovery.Resolver) (authservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-auth-service-client")
	defer span.End()

	return authservice.NewClient(
		"auth-service",
		client.WithResolver(resolver),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
