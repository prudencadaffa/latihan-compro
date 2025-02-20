package model

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	ID          int64 `gorm:"id,primaryKey"`
	ServiceID   int64
	Name        string
	PhoneNumber string
	Email       string
	Brief       string
	Budget      float64
	MeetAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
