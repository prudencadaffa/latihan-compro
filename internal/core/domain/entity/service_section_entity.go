package entity

type ServiceSectionEntity struct {
	ID            int64
	PathIcon      string
	Name          string
	Tagline       string
	ServiceDetail ServiceDetailEntity
}
