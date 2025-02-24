package response

type HeroSectionResponse struct {
	ID         int64  `json:"id"`
	Heading    string `json:"heading"`
	SubHeading string `json:"subheading"`
	PathVideo  string `json:"path_video"`
	Banner     string `json:"banner"`
}
