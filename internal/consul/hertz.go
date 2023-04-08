package consul

import (
	"context"

	hertzregistry "github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/hashicorp/consul/api"
	hertzconsul "github.com/hertz-contrib/registry/consul"
	"github.com/sirupsen/logrus"
)

func NewHertzConsulRegistry(ctx context.Context) (hertzregistry.Registry, error) {
	ctx, span := tracer.Start(ctx, "new-hertz-consul-registry")
	defer span.End()

	e, err := getEnv(ctx)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get env")
		return nil, err
	}

	client, err := api.NewClient(&api.Config{
		Address: e.Addr,
		Scheme:  e.Scheme,
		Token:   e.Token,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create consul client")
		return nil, err
	}

	r := hertzconsul.NewConsulRegister(client)

	return r, nil
}
