package model

import "encoding/gob"

type Website struct {
	Model
	UserRefer uint64 `gorm:"not null"`

	Key      string `gorm:"default:null"`
	IsPublic bool   `gorm:"not null;default:false"`

	Name                  string `gorm:"not null;default:''"`
	Description           string `gorm:"not null;default:''"`
	AvatarIcon            string `gorm:"not null;default:''"`
	Background            string `gorm:"not null;default:''"`
	Language              string `gorm:"not null;default:''"`
	AllowAnonymousComment bool   `gorm:"not null;default:true"`
}

func init() {
	gob.Register(Website{})
}
