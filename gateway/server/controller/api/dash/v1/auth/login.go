package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	authapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
)

// UsernameLogin implements Service.UsernameLogin
func (ctr *Controller) UsernameLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "email-login")
	defer span.End()

	var payload authapi.UsernameLoginRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		span.RecordError(err)
		return
	}

	resp, err := ctr.AuthSvcClient.UsernameLogin(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		span.RecordError(err)
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

// EmailLogin implements Service.EmailLogin
func (ctr *Controller) EmailLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "email-login")
	defer span.End()

	var payload authapi.EmailLoginRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		span.RecordError(err)
		return
	}

	resp, err := ctr.AuthSvcClient.EmailLogin(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		span.RecordError(err)
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
