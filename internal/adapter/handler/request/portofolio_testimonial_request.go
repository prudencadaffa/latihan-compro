package request

type PortofolioTestimonialRequest struct {
	Thumbnail           string `json:"thumbnail" validate:"required"`
	Message             string `json:"message" validate:"required"`
	ClientName          string `json:"client_name" validate:"required"`
	Role                string `json:"role" validate:"required"`
	PortofolioSectionID int64  `json:"portofolio_section_id" validate:"required"`
}
