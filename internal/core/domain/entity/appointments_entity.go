package entity

import "time"

type AppointmentsEntity struct {
	ID          int64
	ServiceID   int64
	Name        string
	PhoneNumber string
	Email       string
	Brief       string
	Budget      float64
	MeetAt      time.Time
	ServiceName string
}
