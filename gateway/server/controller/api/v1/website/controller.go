package website

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/gateway/serializer"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/websiteservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.api.v1.website")

type Controller struct {
	fx.In
	WebsiteSvcClient websiteservice.Client
}

type Service interface {
	CreateWebsite(ctx context.Context, c *app.RequestContext) // CreateWebsite provides a method to create website.
	GetWebsite(ctx context.Context, c *app.RequestContext)    // GetWebsite provides a method to get website.
	UpdateWebsite(ctx context.Context, c *app.RequestContext) // UpdateWebsite provides a method to update website.
	ListWebsites(ctx context.Context, c *app.RequestContext)  // ListWebsites provides a method to list websites.
}

func NewController(impl Controller) Service {
	return &impl
}

func (ctr *Controller) CreateWebsite(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "create-website")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	var payload api.CreateWebsiteRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()

	resp, err := ctr.WebsiteSvcClient.CreateWebsite(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

func (ctr *Controller) GetWebsite(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "get-website")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	payload := api.GetWebsiteRequest{
		UserId: authData.GetId(),
		Id:     id,
	}

	resp, err := ctr.WebsiteSvcClient.GetWebsite(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

func (ctr *Controller) UpdateWebsite(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "update-website")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}

	var payload api.UpdateWebsiteRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()
	payload.Id = id

	resp, err := ctr.WebsiteSvcClient.UpdateWebsite(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}

func (ctr *Controller) ListWebsites(ctx context.Context, c *app.RequestContext) {
	ctx, span := tracer.Start(ctx, "list-websites")
	defer span.End()

	a, exist := c.Get("auth_data")
	if !exist {
		c.JSON(401, serializer.ErrorUnauthorized)
		return
	}
	authData := a.(*api.ServerAuthDataResponse)

	var payload api.ListWebsitesRequest
	if err := c.Bind(&payload); err != nil {
		c.JSON(400, serializer.ResponseError(err))
		return
	}
	payload.UserId = authData.GetId()

	resp, err := ctr.WebsiteSvcClient.ListWebsites(ctx, &payload)
	if err != nil {
		c.JSON(500, serializer.ResponseError(err))
		return
	}

	c.JSON(200, serializer.ResponseSuccess(resp))
}
