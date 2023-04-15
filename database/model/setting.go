package model

import "encoding/gob"

type Setting struct {
	Name  string      `gorm:"primaryKey"`
	Type  SettingType `gorm:"not null;index"`
	Value string      `gorm:"not null"`
}

type SettingType string

const (
	SettingTypeSystem SettingType = "system" // 系统设置，不可修改
	SettingTypeBasic  SettingType = "basic"  // 基础设置
	SettingTypeAuth   SettingType = "auth"   // 认证设置
	SettingTypeEmail  SettingType = "email"  // 邮件设置
)

func init() {
	gob.Register(Setting{})
}
