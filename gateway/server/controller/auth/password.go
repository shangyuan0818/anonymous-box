package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/dto"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/vo"
	authapi "github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api"
)

// ChangePassword implements Service.ChangePassword
func (ctr *Controller) ChangePassword(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "change-password")
	defer span.End()

	var payload dto.ChangePassword
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	token := c.GetString("token")
	if token == "" {
		c.JSON(401, serializer.ResponseErrorMsg("token is empty"))
		return
	}

	req := &authapi.ChangePasswordRequest{
		Token:       token,
		OldPassword: payload.Password,
		NewPassword: payload.NewPassword,
	}

	resp, err := ctr.AuthSvcClient.ChangePassword(ctx, req)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.AuthToken{
		Token: resp.Token,
	}))
}

// ResetPassword implements Service.ResetPassword
func (ctr *Controller) ResetPassword(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "reset-password")
	defer span.End()

	var payload dto.ResetPassword
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	req := &authapi.ResetPasswordRequest{
		Email:            payload.Email,
		VerificationCode: payload.Code,
		NewPassword:      payload.NewPassword,
	}

	resp, err := ctr.AuthSvcClient.ResetPassword(ctx, req)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.AuthToken{
		Token: resp.Token,
	}))
}
