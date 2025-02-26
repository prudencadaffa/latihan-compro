package response

type AboutCompanyKeynoteResponse struct {
	ID                      int64  `json:"id"`
	AboutCompanyID          int64  `json:"about_company_id"`
	Keynote                 string `json:"keynote"`
	PathImage               string `json:"path_image"`
	AboutCompanyDescription string `json:"about_company_description"`
}
