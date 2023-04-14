package consul

import (
	"context"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type env struct {
	Addr   string `default:"localhost:8500"`
	Scheme string `default:"http"`
	Token  string `default:""`
}

func getEnv(ctx context.Context) (*env, error) {
	ctx, span := tracer.Start(ctx, "get-env")
	defer span.End()

	var e env
	if err := envconfig.Process("CONSUL", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to process env config")
		return nil, err
	}

	return &e, nil
}

// NewConfig returns a consul config.
func NewConfig(ctx context.Context) (*consulapi.Config, error) {
	ctx, span := tracer.Start(ctx, "new-config")
	defer span.End()

	e, err := getEnv(ctx)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get consul env")
		return nil, err
	}

	config := consulapi.DefaultConfig()
	config.Address = e.Addr
	config.Scheme = e.Scheme
	config.Token = e.Token

	return config, nil
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
