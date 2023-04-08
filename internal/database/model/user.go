package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
}
