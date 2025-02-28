package response

type PortofolioTestimonialResponse struct {
	ID                int64                     `json:"id"`
	Thumbnail         string                    `json:"thumbnail"`
	Message           string                    `json:"message"`
	ClientName        string                    `json:"client_name"`
	Role              string                    `json:"role"`
	PortofolioSection PortofolioSectionResponse `json:"portofolio_section"`
}
