package comment

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
)

func (ctr *Controller) DeleteComment(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "delete-comment")
	defer span.End()

	span.AddEvent("get-auth-data")
	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*dash.ServerAuthDataResponse)

	span.AddEvent("parse-id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	span.AddEvent("generate-payload")
	payload := dash.DeleteCommentRequest{
		UserId: authData.GetId(),
		Id:     id,
	}

	span.AddEvent("call-comment-service")
	resp, err := ctr.CommentSvcClient.DeleteComment(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
