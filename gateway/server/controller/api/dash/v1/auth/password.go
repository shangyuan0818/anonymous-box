package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	authapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
)

// ChangePassword implements Service.ChangePassword
func (ctr *Controller) ChangePassword(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "change-password")
	defer span.End()

	var payload authapi.ChangePasswordRequest
	if err := c.Bind(&payload); err != nil {
		span.RecordError(err)
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.Token = c.GetString("token")

	resp, err := ctr.AuthSvcClient.ChangePassword(ctx, &payload)
	if err != nil {
		span.RecordError(err)
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

// ResetPassword implements Service.ResetPassword
func (ctr *Controller) ResetPassword(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "reset-password")
	defer span.End()

	var payload authapi.ResetPasswordRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	resp, err := ctr.AuthSvcClient.ResetPassword(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
