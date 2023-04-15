package comment

import (
	"context"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/commentservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.api.v1.comment")

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

func (ctr *Controller) GetComment(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "get-comment")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	payload := api.GetCommentRequest{
		UserId: authData.GetId(),
		Id:     id,
	}

	resp, err := ctr.CommentSvcClient.GetComment(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

func (ctr *Controller) ListComments(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "list-comments")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	var payload api.ListCommentsRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()

	resp, err := ctr.CommentSvcClient.ListComments(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

func (ctr *Controller) DeleteComment(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "delete-comment")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	payload := api.DeleteCommentRequest{
		UserId: authData.GetId(),
		Id:     id,
	}

	resp, err := ctr.CommentSvcClient.DeleteComment(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
