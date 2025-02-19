package entity

type AboutCompanyEntity struct {
	ID          int64
	Description string
	Keynote     []AboutCompanyKeynoteEntity
}
