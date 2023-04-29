package website

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
)

func (ctr *Controller) UpdateWebsite(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "update-website")
	defer span.End()

	span.AddEvent("parse-payload")
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

	span.AddEvent("parse-payload")
	var payload dash.UpdateWebsiteRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()
	payload.Id = id

	span.AddEvent("call-website-service")
	resp, err := ctr.WebsiteSvcClient.UpdateWebsite(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
