package model

import (
	"time"

	"gorm.io/gorm"
)

type ContactUs struct {
	ID           int64 `gorm:"id,primaryKey"`
	CompanyName  string
	LocationName string
	Address      string
	PhoneNumber  string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Tabler interface {
	TableName() string
}

func (ContactUs) TableName() string {
	return "contact_us"
}
