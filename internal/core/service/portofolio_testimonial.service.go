package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type PortofolioTestimonialServiceInterface interface {
	CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error)
	FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error)
	EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error
}
type portofolioTestimonialService struct {
	portofolioTestimonialRepo repository.PortofolioTestimonialRepositoryInterface
	portofolioSectionRepo     repository.PortofolioSectionRepositoryInterface
}

// CreatePortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (c *portofolioTestimonialService) CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	_, err = c.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] CreatePortofolioTestimonial - 1: %v", err)
		return err
	}
	return c.portofolioTestimonialRepo.CreatePortofolioTestimonial(ctx, req)
}

// FetchAllPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (c *portofolioTestimonialService) FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error) {
	return c.portofolioTestimonialRepo.FetchAllPortofolioTestimonial(ctx)
}

// FetchByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (c *portofolioTestimonialService) FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error) {
	return c.portofolioTestimonialRepo.FetchByIDPortofolioTestimonial(ctx, id)
}

// EditByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (c *portofolioTestimonialService) EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	_, err = c.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] EditByIDPortofolioTestimonial - 1: %v", err)
		return err
	}
	return c.portofolioTestimonialRepo.EditByIDPortofolioTestimonial(ctx, req)
}

// DeleteByIDPortofolioTestimonial implements PortofolioTestimonialServiceInterface.
func (c *portofolioTestimonialService) DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error {
	return c.portofolioTestimonialRepo.DeleteByIDPortofolioTestimonial(ctx, id)
}
func NewPortofolioTestimonialService(portofolioTestimonialRepo repository.PortofolioTestimonialRepositoryInterface, portofolioSectionRepo repository.PortofolioSectionRepositoryInterface) PortofolioTestimonialServiceInterface {
	return &portofolioTestimonialService{
		portofolioTestimonialRepo: portofolioTestimonialRepo,
		portofolioSectionRepo:     portofolioSectionRepo,
	}
}
