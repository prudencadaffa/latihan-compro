package response

type ServiceSectionResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Tagline  string `json:"tagline"`
	PathIcon string `json:"path_icon"`
}
