package request

type AboutCompanyRequest struct {
	Description string `json:"description" validate:"required"`
}
