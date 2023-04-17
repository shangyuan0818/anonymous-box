package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"go.uber.org/fx/fxevent"
)

func FxLogger(ctx context.Context, logger *logrus.Logger) fxevent.Logger {
	ctx, span := tracer.Start(ctx, "fx-logger")
	defer span.End()

	return &fxlogrus.LogrusLogger{
		Logger: logger,
	}
}
