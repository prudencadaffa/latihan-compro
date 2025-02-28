package request

type ServiceSectionRequest struct {
	Name     string `json:"name" validate:"required"`
	Tagline  string `json:"tagline" validate:"required"`
	PathIcon string `json:"path_icon"`
}
