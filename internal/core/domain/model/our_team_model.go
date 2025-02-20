package model

import (
	"time"

	"gorm.io/gorm"
)

type OurTeam struct {
	ID        int64 `gorm:"id,primaryKey"`
	Name      string
	Role      string
	PathPhoto string
	Tagline   string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
