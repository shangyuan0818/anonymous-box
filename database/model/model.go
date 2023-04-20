package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint64       `gorm:"<-:create;primarykey"`
	CreatedAt time.Time    `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	DeletedAt sql.NullTime `gorm:"index"`
}
