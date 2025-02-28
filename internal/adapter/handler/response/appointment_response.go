package response

type AppointmentResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Brief       string  `json:"brief"`
	Budget      float64 `json:"budget"`
	MeetAt      string  `json:"meet_at"`
	ServiceName string  `json:"service_name"`
	ServiceID   int64   `json:"service_id"`
}
