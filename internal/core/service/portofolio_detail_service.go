package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type PortofolioDetailServiceInterface interface {
	CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error)
	FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error)
	EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	DeleteByIDPortofolioDetail(ctx context.Context, id int64) error

	FetchDetailPotofolioByPortoID(ctx context.Context, portoID int64) (*entity.PortofolioDetailEntity, error)
}
type portofolioDetailService struct {
	portofolioDetailRepo  repository.PortofolioDetailRepositoryInterface
	portofolioSectionRepo repository.PortofolioSectionRepositoryInterface
}

// CreatePortofolioDetail implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	_, err = c.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] CreatePortofolioDetail - 1: %v", err)
		return err
	}
	return c.portofolioDetailRepo.CreatePortofolioDetail(ctx, req)
}

// FetchAllPortofolioDetail implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error) {
	return c.portofolioDetailRepo.FetchAllPortofolioDetail(ctx)
}

// FetchByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error) {
	return c.portofolioDetailRepo.FetchByIDPortofolioDetail(ctx, id)
}

// EditByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	_, err = c.portofolioSectionRepo.FetchByIDPortofolioSection(ctx, req.PortofolioSection.ID)
	if err != nil {
		log.Errorf("[SERVICE] EditByIDPortofolioDetail - 1: %v", err)
		return err
	}
	return c.portofolioDetailRepo.EditByIDPortofolioDetail(ctx, req)
}

// DeleteByIDPortofolioDetail implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) DeleteByIDPortofolioDetail(ctx context.Context, id int64) error {
	return c.portofolioDetailRepo.DeleteByIDPortofolioDetail(ctx, id)
}

// FetchDetailPotofolioByPortoID implements PortofolioDetailServiceInterface.
func (c *portofolioDetailService) FetchDetailPotofolioByPortoID(ctx context.Context, portoID int64) (*entity.PortofolioDetailEntity, error) {
	return c.portofolioDetailRepo.FetchDetailPotofolioByPortoID(ctx, portoID)
}
func NewPortofolioDetailService(portofolioDetailRepo repository.PortofolioDetailRepositoryInterface, portofolioSectionRepo repository.PortofolioSectionRepositoryInterface) PortofolioDetailServiceInterface {
	return &portofolioDetailService{
		portofolioDetailRepo:  portofolioDetailRepo,
		portofolioSectionRepo: portofolioSectionRepo,
	}
}
