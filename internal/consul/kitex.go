package consul

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
	kitexregistry "github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	kitexconsul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
)

func NewKitexConsulRegistry(ctx context.Context) (kitexregistry.Registry, error) {
	ctx, span := tracer.Start(ctx, "new-kitex-consul-registry")
	defer span.End()

	e, err := getEnv(ctx)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get env")
		return nil, err
	}

	r, err := kitexconsul.NewConsulRegisterWithConfig(&api.Config{
		Address: e.Addr,
		Scheme:  e.Scheme,
		Token:   e.Token,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create kitex consul register")
		return nil, err
	}

	return r, nil
}

func NewKitexConsulResolver(ctx context.Context) (discovery.Resolver, error) {
	ctx, span := tracer.Start(ctx, "new-kitex-consul-resolver")
	defer span.End()

	e, err := getEnv(ctx)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get env")
		return nil, err
	}

	r, err := kitexconsul.NewConsulResolverWithConfig(&api.Config{
		Address: e.Addr,
		Scheme:  e.Scheme,
		Token:   e.Token,
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create kitex consul resolver")
		return nil, err
	}

	return r, nil
}
