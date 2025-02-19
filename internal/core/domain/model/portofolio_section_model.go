package model

import (
	"time"

	"gorm.io/gorm"
)

type PortofolioSection struct {
	ID        int64 `gorm:"id,primaryKey"`
	Name      string
	Tagline   string
	Thumbnail *string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
