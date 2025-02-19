package model

import (
	"time"

	"gorm.io/gorm"
)

type FaqSection struct {
	ID          int64 `gorm:"id,primaryKey"`
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
