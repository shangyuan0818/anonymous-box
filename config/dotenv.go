package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	ctx, span := tracer.Start(context.Background(), "init-dotenv")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}
}
