package comment

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/commentservice"
)

func NewCommentServiceClient(ctx context.Context, resolver discovery.Resolver) (commentservice.Client, error) {
	ctx, span := tracer.Start(ctx, "new-email-service-client")
	defer span.End()

	return commentservice.NewClient(
		ServiceName,
		client.WithResolver(resolver),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
