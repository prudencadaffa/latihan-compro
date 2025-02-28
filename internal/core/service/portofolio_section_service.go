package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type PortofolioSectionServiceInterface interface {
	CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error)
	FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error)
	EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	DeleteByIDPortofolioSection(ctx context.Context, id int64) error
}

type portofolioSectionService struct {
	portofolioSectionRepo repository.PortofolioSectionRepositoryInterface
}

// CreatePortofolioSection implements PortofolioSectionServiceInterface.
func (c *portofolioSectionService) CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	return c.portofolioSectionRepo.CreatePortofolioSection(ctx, req)
}

// FetchAllPortofolioSection implements PortofolioSectionServiceInterface.
func (c *portofolioSectionService) FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error) {
	return c.portofolioSectionRepo.FetchAllPortofolioSection(ctx)
}

// FetchByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (c *portofolioSectionService) FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error) {
	return c.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, id)
}

// EditByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (c *portofolioSectionService) EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	return c.portofolioSectionRepo.EditByIDPortofolioSection(ctx, req)
}

// DeleteByIDPortofolioSection implements PortofolioSectionServiceInterface.
func (c *portofolioSectionService) DeleteByIDPortofolioSection(ctx context.Context, id int64) error {
	return c.portofolioSectionRepo.DeleteByIDPortofolioSection(ctx, id)
}
func NewPortofolioSectionService(portofolioSectionRepo repository.PortofolioSectionRepositoryInterface) PortofolioSectionServiceInterface {
	return &portofolioSectionService{
		portofolioSectionRepo: portofolioSectionRepo,
	}
}
