package entity

import "time"

type PortofolioDetailEntity struct {
	ID                int64
	Category          string
	ClientName        string
	ProjectDate       time.Time
	ProjectUrl        string
	Title             string
	Description       string
	PortofolioSection PortofolioSectionEntity
}
