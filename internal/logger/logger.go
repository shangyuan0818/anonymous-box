package logger

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("internal.logger")

type env struct {
	Level string `default:"info"`
}

func InitLogger(ctx context.Context) error {
	ctx, span := tracer.Start(ctx, "init-logger")
	defer span.End()

	if err := godotenv.Load(".env"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}

	var e env
	if err := envconfig.Process("LOG", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to process env")
		return err
	}

	l := logrus.StandardLogger()

	level, err := logrus.ParseLevel(e.Level)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to parse log level")
		return err
	}

	l.SetLevel(level)
	l.SetOutput(os.Stdout)

	hlog.SetLogger(hertzlogrus.NewLogger(hertzlogrus.WithLogger(l)))

	l.SetFormatter(&logrus.TextFormatter{})

	return nil
}
