package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type AboutCompanyKeynoteServiceInterface interface {
	CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error
	FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error)
	FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error)
	FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error)
}
type aboutCompanyKeynoteService struct {
	aboutCompanyKeynoteRepo repository.AboutCompanyKeynoteInterface
	aboutCompanyRepo        repository.AboutCompanyInterface
}

// CreateAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	_, err := c.aboutCompanyRepo.FetchByIDAboutCompany(ctx, req.AboutCompanyID)
	if err != nil {
		log.Errorf("[SERVICE] CreateAboutCompanyKeynote - 1: %v", err)
		return err
	}
	return c.aboutCompanyKeynoteRepo.CreateAboutCompanyKeynote(ctx, req)
}

// EditByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	_, err := c.aboutCompanyRepo.FetchByIDAboutCompany(ctx, req.AboutCompanyID)
	if err != nil {
		log.Errorf("[SERVICE] EditByIDAboutCompanyKeynote - 1: %v", err)
		return err
	}
	return c.aboutCompanyKeynoteRepo.EditByIDAboutCompanyKeynote(ctx, req)
}

// DeleteByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error {
	return c.aboutCompanyKeynoteRepo.DeleteByIDAboutCompanyKeynote(ctx, id)
}

// FetchAllAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error) {
	return c.aboutCompanyKeynoteRepo.FetchAllAboutCompanyKeynote(ctx)
}

// FetchByIDAboutCompanyKeynote implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error) {
	return c.aboutCompanyKeynoteRepo.FetchByIDAboutCompanyKeynote(ctx, id)
}

// FetchByCompanyID implements AboutCompanyKeynoteServiceInterface.
func (c *aboutCompanyKeynoteService) FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error) {
	return c.aboutCompanyKeynoteRepo.FetchByCompanyID(ctx, companyId)
}

func NewAboutCompanyKeynoteService(aboutCompanyKeynoteRepo repository.AboutCompanyKeynoteInterface, aboutCompanyRepo repository.AboutCompanyInterface) AboutCompanyKeynoteServiceInterface {
	return &aboutCompanyKeynoteService{
		aboutCompanyKeynoteRepo: aboutCompanyKeynoteRepo,
		aboutCompanyRepo:        aboutCompanyRepo,
	}
}
