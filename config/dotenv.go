package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadDotEnv(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "load-dotenv")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}
}
