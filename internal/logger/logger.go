package logger

import (
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/star-horizon/anonymous-box-saas/config"
)

var tracer = otel.Tracer("internal.logger")

func NewLogger(ctx context.Context, e *config.LoggerEnv) (*logrus.Logger, error) {
	ctx, span := tracer.Start(ctx, "init-logger")
	defer span.End()

	l := logrus.StandardLogger()

	level, err := logrus.ParseLevel(e.Level)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to parse log level")
		return nil, err
	}

	l.SetLevel(level)
	l.SetOutput(os.Stdout)

	hlog.SetLogger(hertzlogrus.NewLogger(hertzlogrus.WithLogger(l)))

	l.SetFormatter(&logrus.TextFormatter{})

	return l, nil
}
