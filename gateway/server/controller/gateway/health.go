package gateway

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
)

// Health implements Service.Health
func (ctr *Controller) Health(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "health")
	defer span.End()

	c.JSON(200, serializer.ResponseSuccess(nil))
}
