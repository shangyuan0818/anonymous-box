package main

import (
	"github.com/star-horizon/anonymous-box-saas/config"
	"github.com/star-horizon/anonymous-box-saas/database/model"
)

var defaultSettings = []model.Setting{
	{Name: "system_version", Value: config.Version, Type: model.SettingTypeSystem},

	{Name: "app_name", Value: "Anonymous Box", Type: model.SettingTypeBasic},
	{Name: "app_description", Value: "Anonymous Box SaaS Service", Type: model.SettingTypeBasic},
	{Name: "app_logo", Value: "https://example.com/logo.png", Type: model.SettingTypeBasic},
	{Name: "app_url", Value: "https://example.com", Type: model.SettingTypeBasic},
	{Name: "app_hashids_salt", Value: "salt", Type: model.SettingTypeBasic},
	{Name: "app_hashids_min_length", Value: "8", Type: model.SettingTypeBasic},
	{Name: "app_hashids_alphabet", Value: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", Type: model.SettingTypeBasic},

	{Name: "auth_jwt_secret", Value: "secret", Type: model.SettingTypeAuth},
	{Name: "auth_jwt_expires", Value: "3600", Type: model.SettingTypeAuth},
	{Name: "auth_allow_register", Value: "true", Type: model.SettingTypeAuth},
	{Name: "auth_require_email_verify", Value: "true", Type: model.SettingTypeAuth},

	{Name: "email_from_address", Value: "no-reply@localhost", Type: model.SettingTypeEmail},
	{Name: "email_from_name", Value: "Anonymous Box", Type: model.SettingTypeEmail},

	{Name: "email_template_verify_code_content_type", Value: "text/plain", Type: model.SettingTypeEmail},
	{Name: "email_template_verify_code", Value: "", Type: model.SettingTypeEmail},
	{Name: "email_template_register", Value: "", Type: model.SettingTypeEmail},
	{Name: "email_template_reset_password", Value: "", Type: model.SettingTypeEmail},
}
