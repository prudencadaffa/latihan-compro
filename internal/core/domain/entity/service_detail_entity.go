package entity

type ServiceDetailEntity struct {
	ID          int64
	ServiceID   int64
	PathImage   string
	Title       string
	Description string
	PathPdf     *string
	PathDocx    *string
	ServiceName string
}
