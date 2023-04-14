package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint64       `gorm:"primarykey"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	DeletedAt sql.NullTime `gorm:"index"`
}
