package database

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/star-horizon/anonymous-box-saas/internal/database/dal"
)

type env struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	User     string `default:"postgres"`
	Password string `default:"postgres"`
	Database string `default:"postgres"`
	SSLMode  string `default:"disable"`
	TimeZone string `default:"Asia/Shanghai" envconfig:"TZ"`
}

var tracer = otel.Tracer("internal.database")

func NewDB(ctx context.Context) (*gorm.DB, error) {
	ctx, span := tracer.Start(ctx, "init-db")
	defer span.End()

	if err := godotenv.Load(); err != nil {
		logrus.WithContext(ctx).WithError(err).Warn("failed to load .env")
	}

	var e env
	if err := envconfig.Process("DB", &e); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to process env")
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		e.Host,
		e.User,
		e.Password,
		e.Database,
		e.Port,
		e.SSLMode,
		e.TimeZone,
	)), &gorm.Config{})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to open database")
		return nil, err
	}

	db = db.WithContext(ctx)

	// add otel plugin
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithAttributes(
		semconv.DBSystemPostgreSQL,
		semconv.DBNameKey.String(e.Database),
		semconv.DBUserKey.String(e.User),
	))); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to add otel plugin")
		return nil, err
	}

	return db, nil
}

func NewQuery(ctx context.Context, db *gorm.DB) (*dal.Query, error) {
	ctx, span := tracer.Start(ctx, "init-query")
	defer span.End()

	return dal.Use(db), nil
}
