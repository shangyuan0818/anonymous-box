package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	authapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
)

// EmailLogin implements Service.EmailLogin
func (ctr *Controller) EmailLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "email-login")
	defer span.End()

	var payload authapi.EmailAuthRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	resp, err := ctr.AuthSvcClient.EmailAuth(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

// UsernameLogin implements Service.UsernameLogin
func (ctr *Controller) UsernameLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "username-login")
	defer span.End()

	var payload authapi.UsernameAuthRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	resp, err := ctr.AuthSvcClient.UsernameAuth(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
