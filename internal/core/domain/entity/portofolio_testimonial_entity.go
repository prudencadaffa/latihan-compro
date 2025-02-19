package entity

type PortofolioTestimonialEntity struct {
	ID                int64
	Thumbnail         string
	Message           string
	ClientName        string
	Role              string
	PortofolioSection PortofolioSectionEntity
}
