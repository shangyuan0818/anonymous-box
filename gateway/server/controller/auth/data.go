package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
)

// GetAuthData implements Service.GetAuthData
func (ctr *Controller) GetAuthData(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "get-auth-data")
	defer span.End()

	resp, err := ctr.AuthSvcClient.GetServerAuthData(ctx, &api.AuthToken{
		Token: c.GetString("token"),
	})
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		span.RecordError(err)
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
