package model

import (
	"time"

	"gorm.io/gorm"
)

type PortofolioTestimonial struct {
	ID                  int64 `gorm:"id,primaryKey"`
	PortofolioSectionID int64
	Thumbnail           string
	Message             string
	ClientName          string
	Role                string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
