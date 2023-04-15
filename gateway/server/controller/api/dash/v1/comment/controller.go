package comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/commentservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.dash.dash.v1.comment")

type Controller struct {
	fx.In
	CommentSvcClient commentservice.Client
}

type Service interface {
	GetComment(ctx context.Context, c *app.RequestContext)    // GetComment provides a method to get comment.
	ListComments(ctx context.Context, c *app.RequestContext)  // ListComments provides a method to list comments.
	DeleteComment(ctx context.Context, c *app.RequestContext) // DeleteComment provides a method to delete comment.
}

func NewController(impl Controller) Service {
	return &impl
}
