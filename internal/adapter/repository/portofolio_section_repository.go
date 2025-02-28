package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioSectionRepositoryInterface interface {
	CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error)
	FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error)
	EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error
	DeleteByIDPortofolioSection(ctx context.Context, id int64) error
}
type portofolioSectionRepository struct {
	DB *gorm.DB
}

// CreatePortofolioSection implements PortofolioSectionInterface.
func (h *portofolioSectionRepository) CreatePortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	modelPortofolioSection := model.PortofolioSection{
		Thumbnail: &req.Thumbnail,
		Name:      req.Name,
		Tagline:   req.Tagline,
	}

	if err = h.DB.Create(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioSection - 1: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioSection implements PortofolioSectionInterface.
func (h *portofolioSectionRepository) FetchAllPortofolioSection(ctx context.Context) ([]entity.PortofolioSectionEntity, error) {
	modelPortofolioSection := []model.PortofolioSection{}
	if err = h.DB.Select("id", "thumbnail", "tagline", "name").Find(&modelPortofolioSection).Order("created_at DESC").Error; err != nil {
		log.Errorf("[REPOSITORY] FetchAllPortofolioSection - 1: %v", err)
		return nil, err
	}

	var portofolioSectionRepositoryEntities []entity.PortofolioSectionEntity
	for _, v := range modelPortofolioSection {
		portofolioSectionRepositoryEntities = append(portofolioSectionRepositoryEntities, entity.PortofolioSectionEntity{
			ID:        v.ID,
			Thumbnail: *v.Thumbnail,
			Name:      v.Name,
			Tagline:   v.Tagline,
		})
	}

	return portofolioSectionRepositoryEntities, nil
}

// FetchByIDPortofolioSection implements PortofolioSectionInterface.
func (h *portofolioSectionRepository) FetchByIDPortofolioSection(ctx context.Context, id int64) (*entity.PortofolioSectionEntity, error) {
	modelPortofolioSection := model.PortofolioSection{}
	if err = h.DB.Select("id", "thumbnail", "tagline", "name").Where("id = ?", id).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] FetchByIDPortofolioSection - 1: %v", err)
		return nil, err
	}

	return &entity.PortofolioSectionEntity{
		ID:        modelPortofolioSection.ID,
		Thumbnail: *modelPortofolioSection.Thumbnail,
		Name:      modelPortofolioSection.Name,
		Tagline:   modelPortofolioSection.Tagline,
	}, nil
}

// EditByIDPortofolioSection implements PortofolioSectionInterface.
func (h *portofolioSectionRepository) EditByIDPortofolioSection(ctx context.Context, req entity.PortofolioSectionEntity) error {
	modelPortofolioSection := model.PortofolioSection{}

	if err = h.DB.Where("id =?", req.ID).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioSection - 1: %v", err)
		return err
	}
	modelPortofolioSection.Name = req.Name
	modelPortofolioSection.Tagline = req.Tagline
	modelPortofolioSection.Thumbnail = &req.Thumbnail

	if err = h.DB.Save(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioSection - 2: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioSection implements PortofolioSectionInterface.
func (h *portofolioSectionRepository) DeleteByIDPortofolioSection(ctx context.Context, id int64) error {
	modelPortofolioSection := model.PortofolioSection{}

	if err = h.DB.Where("id = ?", id).First(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioSection - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelPortofolioSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioSection - 2: %v", err)
		return err
	}
	return nil
}
func NewPortofolioSectionRepository(DB *gorm.DB) PortofolioSectionRepositoryInterface {
	return &portofolioSectionRepository{
		DB: DB,
	}
}
