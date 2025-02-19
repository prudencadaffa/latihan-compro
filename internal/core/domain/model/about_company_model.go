package model

import (
	"time"

	"gorm.io/gorm"
)

type AboutCompany struct {
	ID          int64 `gorm:"id,primaryKey"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
