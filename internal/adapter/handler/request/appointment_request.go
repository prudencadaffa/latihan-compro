package request

type AppointmentRequest struct {
	ServiceID   int64   `json:"service_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Brief       string  `json:"brief" validate:"required"`
	Budget      float64 `json:"budget" validate:"required"`
	MeetAt      string  `json:"meet_at" validate:"required"`
}
