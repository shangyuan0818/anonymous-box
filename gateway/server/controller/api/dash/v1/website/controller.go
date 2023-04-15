package website

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/websiteservice"
)

var tracer = otel.Tracer("gateway-service.server.controller.dash.dash.v1.website")

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
