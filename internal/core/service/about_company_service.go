package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type AboutCompanyServiceInterface interface {
	CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error)
	FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error)
	EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	DeleteByIDAboutCompany(ctx context.Context, id int64) error
	FetchAllCompanyAndKeynote(ctx context.Context) (*entity.AboutCompanyEntity, error)
}

type aboutCompanyService struct {
	aboutCompanyRepo repository.AboutCompanyInterface
}

// FetchAllCompanyAndKeynote implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) FetchAllCompanyAndKeynote(ctx context.Context) (*entity.AboutCompanyEntity, error) {
	return c.aboutCompanyRepo.FetchAllCompanyAndKeynote(ctx)
}

// CreateAboutCompany implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	return c.aboutCompanyRepo.CreateAboutCompany(ctx, req)
}

// DeleteByIDAboutCompany implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) DeleteByIDAboutCompany(ctx context.Context, id int64) error {
	return c.aboutCompanyRepo.DeleteByIDAboutCompany(ctx, id)
}

// EditByIDAboutCompany implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	return c.aboutCompanyRepo.EditByIDAboutCompany(ctx, req)
}

// FetchAllAboutCompany implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error) {
	return c.aboutCompanyRepo.FetchAllAboutCompany(ctx)
}

// FetchByIDAboutCompany implements AboutCompanyServiceInterface.
func (c *aboutCompanyService) FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error) {
	return c.aboutCompanyRepo.FetchByIDAboutCompany(ctx, id)
}

func NewAboutCompanyService(aboutCompanyRepo repository.AboutCompanyInterface) AboutCompanyServiceInterface {
	return &aboutCompanyService{
		aboutCompanyRepo: aboutCompanyRepo,
	}
}
