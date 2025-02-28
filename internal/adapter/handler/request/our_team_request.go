package request

type OurTeamRequest struct {
	Name      string `json:"name" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Tagline   string `json:"tagline" validate:"required"`
	PathPhoto string `json:"path_photo" validate:"required"`
}
