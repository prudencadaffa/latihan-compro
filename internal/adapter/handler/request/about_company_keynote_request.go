package request

type AboutCompanyKeynoteRequest struct {
	AboutCompanyID int64  `json:"about_company_id" validate:"required"`
	Keynote        string `json:"keynote" validate:"required"`
	PathImage      string `json:"path_image"`
}
