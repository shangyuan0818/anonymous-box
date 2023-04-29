package comment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
)

func (ctr *Controller) ListComments(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "list-comments")
	defer span.End()

	span.AddEvent("get-auth-data")
	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*dash.ServerAuthDataResponse)

	span.AddEvent("parse-payload")
	var payload dash.ListCommentsRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()

	span.AddEvent("call-comment-service")
	resp, err := ctr.CommentSvcClient.ListComments(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
