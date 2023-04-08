package consul

import (
	"context"

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
