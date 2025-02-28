package request

type PortofolioDetailRequest struct {
	Category            string `json:"category" validate:"required"`
	ClientName          string `json:"client_name" validate:"required"`
	ProjectDate         string `json:"project_date" validate:"required"`
	ProjectUrl          string `json:"project_url"`
	Title               string `json:"title" validate:"required"`
	Description         string `json:"description" validate:"required"`
	PortofolioSectionID int64  `json:"portofolio_section_id" validate:"required"`
}
