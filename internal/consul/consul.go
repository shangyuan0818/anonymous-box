package consul

import (
	"context"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"

	"github.com/star-horizon/anonymous-box-saas/config"
)

// NewConfig returns a consul config.
func NewConfig(ctx context.Context, e *config.ConsulEnv) (*consulapi.Config, error) {
	ctx, span := tracer.Start(ctx, "new-config")
	defer span.End()

	c := consulapi.DefaultConfig()
	c.Address = e.Addr
	c.Scheme = e.Scheme
	c.Token = e.Token

	return c, nil
}

// NewClient returns a consul client.
func NewClient(ctx context.Context, config *consulapi.Config) (*consulapi.Client, error) {
	ctx, span := tracer.Start(ctx, "new-client")
	defer span.End()

	client, err := consulapi.NewClient(config)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create consul client")
		return nil, err
	}

	return client, nil
}
