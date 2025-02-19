package model

import (
	"time"

	"gorm.io/gorm"
)

type ClientSection struct {
	ID        int64 `gorm:"id,primaryKey"`
	Name      string
	PathIcon  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
