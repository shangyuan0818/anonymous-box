package model

import (
	"database/sql"
	"encoding/gob"
	"net"
)

type Comment struct {
	Model
	WebsiteRefer uint64 `gorm:"not null"`
	SenderIP     net.IP `gorm:"not null"`

	Name    sql.NullString `gorm:"default:null"`
	Email   sql.NullString `gorm:"default:null"`
	Url     sql.NullString `gorm:"default:null"`
	Content string         `gorm:"not null;default:''"`
}

func init() {
	gob.Register(Comment{})
}
