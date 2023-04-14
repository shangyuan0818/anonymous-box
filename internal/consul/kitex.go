package consul

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
	kitexregistry "github.com/cloudwego/kitex/pkg/registry"
	consulapi "github.com/hashicorp/consul/api"
	kitexconsul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
)

func NewKitexConsulRegistry(ctx context.Context, config *consulapi.Config) (kitexregistry.Registry, error) {
	ctx, span := tracer.Start(ctx, "new-kitex-consul-registry")
	defer span.End()

	r, err := kitexconsul.NewConsulRegisterWithConfig(config)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create kitex consul register")
		return nil, err
	}

	return r, nil
}

func NewKitexConsulResolver(ctx context.Context, config *consulapi.Config) (discovery.Resolver, error) {
	ctx, span := tracer.Start(ctx, "new-kitex-consul-resolver")
	defer span.End()

	r, err := kitexconsul.NewConsulResolverWithConfig(config)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create kitex consul resolver")
		return nil, err
	}

	return r, nil
}
