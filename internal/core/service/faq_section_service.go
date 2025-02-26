package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type FaqSectionServiceInterface interface {
	CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error)
	FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error)
	EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	DeleteByIDFaqSection(ctx context.Context, id int64) error
}

type faqSectionService struct {
	faqSectionRepo repository.FaqSectionRepositoryInterface
}

// CreateFaqSection implements FaqSectionServiceInterface.
func (c *faqSectionService) CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	return c.faqSectionRepo.CreateFaqSection(ctx, req)
}

// DeleteByIDFaqSection implements FaqSectionServiceInterface.
func (c *faqSectionService) DeleteByIDFaqSection(ctx context.Context, id int64) error {
	return c.faqSectionRepo.DeleteByIDFaqSection(ctx, id)
}

// EditByIDFaqSection implements FaqSectionServiceInterface.
func (c *faqSectionService) EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	return c.faqSectionRepo.EditByIDFaqSection(ctx, req)
}

// FetchAllFaqSection implements FaqSectionServiceInterface.
func (c *faqSectionService) FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error) {
	return c.faqSectionRepo.FetchAllFaqSection(ctx)
}

// FetchByIDFaqSection implements FaqSectionServiceInterface.
func (c *faqSectionService) FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error) {
	return c.faqSectionRepo.FetchByIDFaqSection(ctx, id)
}

func NewFaqSectionService(faqSectionRepo repository.FaqSectionRepositoryInterface) FaqSectionServiceInterface {
	return &faqSectionService{
		faqSectionRepo: faqSectionRepo,
	}
}
