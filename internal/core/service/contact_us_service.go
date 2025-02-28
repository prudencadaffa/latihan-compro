package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type ContactUsServiceInterface interface {
	CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error
	FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error)
	FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error)
	EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error
	DeleteByIDContactUs(ctx context.Context, id int64) error
}
type contactUsService struct {
	contactUsRepo repository.ContactUsInterface
}

// CreateContactUs implements ContactUsServiceInterface.
func (c *contactUsService) CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	return c.contactUsRepo.CreateContactUs(ctx, req)
}

// FetchAllContactUs implements ContactUsServiceInterface.
func (c *contactUsService) FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error) {
	return c.contactUsRepo.FetchAllContactUs(ctx)
}

// FetchByIDContactUs implements ContactUsServiceInterface.
func (c *contactUsService) FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error) {
	return c.contactUsRepo.FetchByIDContactUs(ctx, id)
}

// EditByIDContactUs implements ContactUsServiceInterface.
func (c *contactUsService) EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	return c.contactUsRepo.EditByIDContactUs(ctx, req)
}

// DeleteByIDContactUs implements ContactUsServiceInterface.
func (c *contactUsService) DeleteByIDContactUs(ctx context.Context, id int64) error {
	return c.contactUsRepo.DeleteByIDContactUs(ctx, id)
}
func NewContactUsService(contactUsRepo repository.ContactUsInterface) ContactUsServiceInterface {
	return &contactUsService{
		contactUsRepo: contactUsRepo,
	}
}
