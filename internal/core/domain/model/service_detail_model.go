package model

import (
	"time"

	"gorm.io/gorm"
)

type ServiceDetail struct {
	ID          int64 `gorm:"id,primaryKey"`
	ServiceID   int64
	PathImage   string
	Title       string
	Description string
	PathPdf     *string
	PathDocx    *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
