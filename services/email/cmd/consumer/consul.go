package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"

	consulapi "github.com/hashicorp/consul/api"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/pkg/util"
)

func register(ctx context.Context, client *consulapi.Client, lc fx.Lifecycle) error {
	ctx, span := tracer.Start(ctx, "register-consul")
	defer span.End()

	ip, err := util.GetLocalIP()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("get local ip error")
		return err
	}

	registration := &consulapi.AgentServiceRegistration{
		Kind: consulapi.ServiceKindTypical,
		ID:   fmt.Sprintf("%s:%s", serviceName, ip),
		Name: serviceName,
	}

	agent := client.Agent()
	if agent == nil {
		return errors.New("consul agent is nil")
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "consul-register")
			defer span.End()

			return agent.ServiceRegister(registration)
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "consul-deregister")
			defer span.End()

			return agent.ServiceDeregister(registration.ID)
		},
	})

	return nil
}
