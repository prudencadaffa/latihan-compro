package model

import (
	"time"

	"gorm.io/gorm"
)

type ServiceSection struct {
	ID        int64 `gorm:"id,primaryKey"`
	PathIcon  string
	Name      string
	Tagline   string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
