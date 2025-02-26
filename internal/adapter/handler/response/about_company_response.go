package response

type AboutCompanyResponse struct {
	ID              int64                         `json:"id"`
	Description     string                        `json:"description"`
	CompanyKeynotes []AboutCompanyKeynoteResponse `json:"company_keynotes"`
}
