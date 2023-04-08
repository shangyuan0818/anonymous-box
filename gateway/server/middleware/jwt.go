package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
)

func JwtParser(forceLogin bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx, span := tracer.Start(ctx, "jwt-parser")
		defer span.End()

		authorization := string(c.GetHeader("Authorization"))
		if authorization == "" {
			if forceLogin {
				c.AbortWithStatusJSON(401, serializer.ResponseError(serializer.ErrorUnauthorized))
				return
			} else {
				return
			}
		}

		authorization = strings.TrimPrefix(authorization, "Bearer ")
		c.Set("token", authorization)
	}
}
