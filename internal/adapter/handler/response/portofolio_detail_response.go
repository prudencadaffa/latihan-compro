package response

type PortofolioDetailResponse struct {
	ID                int64                     `json:"id"`
	Category          string                    `json:"category"`
	ClientName        string                    `json:"client_name"`
	ProjectDate       string                    `json:"project_date"`
	ProjectUrl        string                    `json:"project_url"`
	Title             string                    `json:"title"`
	Description       string                    `json:"description"`
	PortofolioSection PortofolioSectionResponse `json:"portofolio_section"`
}
