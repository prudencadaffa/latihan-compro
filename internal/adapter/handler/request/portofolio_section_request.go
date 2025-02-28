package request

type PortofolioSectionRequest struct {
	Thumbnail string `json:"thumbnail" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Tagline   string `json:"tagline" validate:"required"`
}
