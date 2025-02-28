package response

type PortofolioSectionResponse struct {
	ID        int64  `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Name      string `json:"name"`
	Tagline   string `json:"tagline"`
}
