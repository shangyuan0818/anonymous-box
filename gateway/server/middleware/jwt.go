package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/authservice"
)

func JwtParser() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx, span := tracer.Start(ctx, "jwt-parser")
		defer span.End()

		authorization := strings.TrimSpace(string(c.GetHeader("Authorization")))
		if lo.IsEmpty(authorization) {
			span.SetStatus(codes.Unset, "authorization header not found")
			return
		}

		authorization = strings.TrimPrefix(authorization, "Bearer ")

		c.Set("token", authorization)
		span.SetStatus(codes.Ok, "token parsed")
	}
}

func AuthDataParser(authClient authservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx, span := tracer.Start(ctx, "auth-data-parser")
		defer span.End()

		token := strings.TrimSpace(c.GetString("token"))
		if lo.IsEmpty(token) {
			span.SetStatus(codes.Unset, "token not found")
			return
		}

		span.SetAttributes(
			attribute.String("token", token),
		)

		res, err := authClient.GetServerAuthData(ctx, &api.AuthToken{
			Token: token,
		})
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "auth service error")
			return
		}

		c.Set("auth_data", res)
		span.SetStatus(codes.Ok, "auth data parsed")
		return
	}
}

func MustAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx, span := tracer.Start(ctx, "must-auth")
		defer span.End()

		_, exist := c.Get("auth_data")
		if !exist {
			c.AbortWithStatusJSON(401, serializer.ResponseError(serializer.ErrorUnauthorized))
			return
		}
	}
}

func MustNotAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx, span := tracer.Start(ctx, "must-not-auth")
		defer span.End()

		_, exist := c.Get("auth_data")
		if exist {
			c.AbortWithStatusJSON(403, serializer.ResponseError(serializer.ErrorPermissionDenied))
			return
		}
	}
}
