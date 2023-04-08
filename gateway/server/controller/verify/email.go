package verify

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/dto"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer/vo"
	verifyapi "github.com/star-horizon/anonymous-box-saas/services/verify/kitex_gen/api"
)

// ApplyEmailVerify implements Service.ApplyEmailVerify
func (ctr *Controller) ApplyEmailVerify(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "apply-email-verify")
	defer span.End()

	var payload dto.ApplyEmailVerify
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	req := &verifyapi.ApplyEmailVerifyRequest{
		Email: payload.Email,
	}

	resp, err := ctr.VerifySvcClient.ApplyEmailVerify(ctx, req)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(vo.ApplyEmailVerify{
		Email: resp.Email,
		Ok:    resp.Ok,
	}))
}
