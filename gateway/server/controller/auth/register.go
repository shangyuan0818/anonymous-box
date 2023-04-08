package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/dto"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/vo"
	authapi "github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api"
)

// Register implements Service.Register
func (ctr *Controller) Register(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "register")
	defer span.End()

	var payload dto.Register
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	res := &authapi.RegisterRequest{
		Username:         payload.Username,
		Password:         payload.Password,
		Email:            payload.Email,
		VerificationCode: payload.Code,
	}

	resp, err := ctr.AuthSvcClient.Register(ctx, res)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.AuthToken{
		Token: resp.Token,
	}))
}
