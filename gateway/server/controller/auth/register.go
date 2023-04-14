package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	authapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
)

// Register implements Service.Register
func (ctr *Controller) Register(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "register")
	defer span.End()

	var payload authapi.RegisterRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	resp, err := ctr.AuthSvcClient.Register(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
