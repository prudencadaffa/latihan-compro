package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ServiceSectionRepositoryInterface interface {
	CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error)
	FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error)
	EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error
	DeleteByIDServiceSection(ctx context.Context, id int64) error
}
type serviceSectionRepository struct {
	DB *gorm.DB
}

// CreateServiceSection implements ServiceSectionInterface.
func (h *serviceSectionRepository) CreateServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	modelServiceSection := model.ServiceSection{
		PathIcon: req.PathIcon,
		Name:     req.Name,
		Tagline:  req.Tagline,
	}
	if err := h.DB.Create(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateServiceSection - 1: %v", err)
		return err

	}
	return nil
}

// FetchAllServiceSection implements ServiceSectionInterface.
func (h *serviceSectionRepository) FetchAllServiceSection(ctx context.Context) ([]entity.ServiceSectionEntity, error) {
	modelServiceSection := []model.ServiceSection{}
	if err = h.DB.Select("id", "path_icon", "tagline", "name").Find(&modelServiceSection).Order("created_at DESC").Error; err != nil {
		log.Errorf("[REPOSITORY] FetchAllServiceSection - 1: %v", err)
		return nil, err
	}

	var serviceSectionRepositoryEntities []entity.ServiceSectionEntity
	for _, v := range modelServiceSection {
		serviceSectionRepositoryEntities = append(serviceSectionRepositoryEntities, entity.ServiceSectionEntity{
			ID:       v.ID,
			PathIcon: v.PathIcon,
			Name:     v.Name,
			Tagline:  v.Tagline,
		})
	}

	return serviceSectionRepositoryEntities, nil
}

// FetchByIDServiceSection implements ServiceSectionInterface.
func (h *serviceSectionRepository) FetchByIDServiceSection(ctx context.Context, id int64) (*entity.ServiceSectionEntity, error) {
	modelServiceSection := model.ServiceSection{}
	if err = h.DB.Select("id", "path_icon", "tagline", "name").Where("id = ?", id).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] FetchByIDServiceSection - 1: %v", err)
		return nil, err
	}

	return &entity.ServiceSectionEntity{
		ID:       modelServiceSection.ID,
		PathIcon: modelServiceSection.PathIcon,
		Name:     modelServiceSection.Name,
		Tagline:  modelServiceSection.Tagline,
	}, nil
}

// EditByIDServiceSection implements ServiceSectionInterface.
func (h *serviceSectionRepository) EditByIDServiceSection(ctx context.Context, req entity.ServiceSectionEntity) error {
	modelServiceSection := model.ServiceSection{}
	if err := h.DB.Where("id =?", req.ID).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDServiceSection - 1: %v", err)
		return err
	}

	modelServiceSection.PathIcon = req.PathIcon
	modelServiceSection.Name = req.Name
	modelServiceSection.Tagline = req.Tagline

	if err := h.DB.Save(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDServiceSection - 2: %v", err)
		return err
	}

	return nil
}

// DeleteByIDServiceSection implements ServiceSectionInterface.

func (h *serviceSectionRepository) DeleteByIDServiceSection(ctx context.Context, id int64) error {
	modelServiceSection := model.ServiceSection{}
	if err := h.DB.Where("id =?", id).First(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceSection - 1: %v", err)
		return err
	}

	if err := h.DB.Delete(&modelServiceSection).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDServiceSection - 2: %v", err)
		return err
	}

	return nil
}

func NewServiceSectionRepository(DB *gorm.DB) ServiceSectionRepositoryInterface {
	return &serviceSectionRepository{
		DB: DB,
	}
}
