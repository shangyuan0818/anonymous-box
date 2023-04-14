package consul

import (
	"context"

	hertzregistry "github.com/cloudwego/hertz/pkg/app/server/registry"
	consulapi "github.com/hashicorp/consul/api"
	hertzconsul "github.com/hertz-contrib/registry/consul"
)

func NewHertzConsulRegistry(ctx context.Context, client *consulapi.Client) hertzregistry.Registry {
	ctx, span := tracer.Start(ctx, "new-hertz-consul-registry")
	defer span.End()

	return hertzconsul.NewConsulRegister(client)
}
