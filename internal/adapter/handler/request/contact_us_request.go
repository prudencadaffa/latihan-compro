package request

type ContactUsRequest struct {
	CompanyName  string `json:"company_name" validate:"required"`
	LocationName string `json:"location_name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
}
