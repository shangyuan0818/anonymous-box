package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/star-horizon/anonymous-box-saas/internal/database/dal"
	"github.com/star-horizon/anonymous-box-saas/internal/database/model"
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

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.WithError(err).Warn("failed to load .env")
	}

	var e env
	if err := envconfig.Process("DB", &e); err != nil {
		logrus.WithError(err).Fatal("failed to process env")
		return
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
		logrus.WithError(err).Fatal("failed to open database")
		return
	}

	if err := db.AutoMigrate(&model.User{}, &model.Setting{}); err != nil {
		logrus.WithError(err).Fatal("failed to migrate database")
		return
	}

	q := dal.Use(db)

	for _, setting := range defaultSettings {
		if setting.Type == model.SettingTypeSystem {
			v, err := q.WithContext(context.TODO()).Setting.
				Where(q.Setting.Name.Eq(setting.Name)).
				First()

			if err == nil && v.Value == setting.Value {
				logrus.Infof("skip default setting: %s", setting.Name)
			} else if err == nil && v.Value != setting.Value {
				if _, err := q.WithContext(context.TODO()).Setting.
					Where(q.Setting.Name.Eq(setting.Name)).
					Update(q.Setting.Value, setting.Value); err != nil {
					logrus.WithError(err).Errorf("failed to update default setting: %s", setting.Name)
				}
			} else {
				if err := q.WithContext(context.TODO()).Setting.
					Where(q.Setting.Name.Eq(setting.Name)).
					Create(&setting); err != nil {
					logrus.WithError(err).Errorf("failed to create default setting: %s", setting.Name)
				}
			}
		} else {
			if _, err := q.WithContext(context.TODO()).Setting.
				Where(q.Setting.Name.Eq(setting.Name)).
				First(); err == nil {
				logrus.Infof("skip default setting: %s", setting.Name)
				continue
			}

			if err := q.WithContext(context.TODO()).Setting.
				Where(q.Setting.Name.Eq(setting.Name)).
				Create(&setting); err != nil {
				logrus.WithError(err).Errorf("failed to create default setting: %s", setting.Name)
			}
		}
	}

	logrus.Info("database initialized successfully")
}
