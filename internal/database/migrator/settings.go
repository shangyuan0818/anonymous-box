package main

import (
	"github.com/ahdark-services/anonymous-box-saas/internal/config"
	"github.com/ahdark-services/anonymous-box-saas/internal/database/model"
)

var defaultSettings = []model.Setting{
	{Name: "system_version", Value: config.Version, Type: model.SettingTypeSystem},

	{Name: "app_name", Value: "Anonymous Box", Type: model.SettingTypeBasic},
	{Name: "app_description", Value: "Anonymous Box SaaS Service", Type: model.SettingTypeBasic},
	{Name: "app_logo", Value: "https://example.com/logo.png", Type: model.SettingTypeBasic},
	{Name: "app_url", Value: "https://example.com", Type: model.SettingTypeBasic},

	{Name: "auth_jwt_secret", Value: "secret", Type: model.SettingTypeAuth},
	{Name: "auth_jwt_expires", Value: "3600", Type: model.SettingTypeAuth},

	{Name: "email_host", Value: "smtp.example.com", Type: model.SettingTypeEmail},
	{Name: "email_port", Value: "465", Type: model.SettingTypeEmail},
	{Name: "email_username", Value: "username", Type: model.SettingTypeEmail},
	{Name: "email_password", Value: "password", Type: model.SettingTypeEmail},
	{Name: "email_from", Value: "no-reply@localhost", Type: model.SettingTypeEmail},
	{Name: "email_from_name", Value: "Anonymous Box", Type: model.SettingTypeEmail},
	{Name: "email_tls", Value: "false", Type: model.SettingTypeEmail},
	{Name: "email_ssl", Value: "false", Type: model.SettingTypeEmail},

	{Name: "email_template_register", Value: "", Type: model.SettingTypeEmail},
	{Name: "email_template_reset_password", Value: "", Type: model.SettingTypeEmail},
}
