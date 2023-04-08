package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/dto"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/vo"
	authapi "github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api"
)

// EmailLogin implements Service.EmailLogin
func (ctr *Controller) EmailLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "email-login")
	defer span.End()

	var payload dto.EmailAuth
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	req := &authapi.EmailAuthRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}

	resp, err := ctr.AuthSvcClient.EmailAuth(ctx, req)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.AuthToken{
		Token: resp.Token,
	}))
}

// UsernameLogin implements Service.UsernameLogin
func (ctr *Controller) UsernameLogin(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "username-login")
	defer span.End()

	var payload dto.UsernameAuth
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	req := &authapi.UsernameAuthRequest{
		Username: payload.Username,
		Password: payload.Password,
	}

	resp, err := ctr.AuthSvcClient.UsernameAuth(ctx, req)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.AuthToken{
		Token: resp.Token,
	}))
}
