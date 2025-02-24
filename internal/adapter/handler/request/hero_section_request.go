package request

type HeroSectionRequest struct {
	Heading    string `json:"heading" validate:"required"`
	SubHeading string `json:"subheading" validate:"required"`
	PathVideo  string `json:"path_video"`
	Banner     string `json:"banner" validate:"required"`
}
