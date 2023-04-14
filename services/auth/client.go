package auth

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/authservice"
)

// NewAuthServiceClient creates a new AuthServiceClient.
func NewAuthServiceClient(ctx context.Context, resolver discovery.Resolver) (authservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-auth-service-client")
	defer span.End()

	return authservice.NewClient("auth-service", client.WithResolver(resolver))
}
