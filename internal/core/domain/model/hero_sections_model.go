package model

import (
	"time"

	"gorm.io/gorm"
)

type HeroSection struct {
	ID         int64 `gorm:"id,primaryKey"`
	Heading    string
	SubHeading string
	PathVideo  *string
	PathBanner string
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
