package model

import "encoding/gob"

type User struct {
	Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Location  string `gorm:"not null"`
	Bio       string `gorm:"not null"`
}

func init() {
	gob.Register(User{})
}
