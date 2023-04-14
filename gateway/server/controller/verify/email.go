package verify

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	verifyapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
)

// ApplyEmailVerify implements Service.ApplyEmailVerify
func (ctr *Controller) ApplyEmailVerify(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "apply-email-verify")
	defer span.End()

	var payload verifyapi.ApplyEmailVerifyRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	resp, err := ctr.VerifySvcClient.ApplyEmailVerify(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
