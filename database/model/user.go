package model

import "encoding/gob"

type User struct {
	Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
}

func init() {
	gob.Register(User{})
}
