package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type FaqSectionRepositoryInterface interface {
	CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error)
	FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error)
	EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error
	DeleteByIDFaqSection(ctx context.Context, id int64) error
}

type faqSectionRepository struct {
	DB *gorm.DB
}

// CreateFaqSection implements FaqSectionInterface.
func (h *faqSectionRepository) CreateFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	modelFaqSection := model.FaqSection{
		Description: req.Description,
		Title:       req.Title,
	}

	if err = h.DB.Create(&modelFaqSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateFaqSection - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDFaqSection implements FaqSectionInterface.
func (h *faqSectionRepository) DeleteByIDFaqSection(ctx context.Context, id int64) error {
	modelFaqSection := model.FaqSection{}

	err = h.DB.Where("id = ?", id).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDFaqSection - 1: %v", err)
		return err
	}

	err = h.DB.Delete(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDFaqSection - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDFaqSection implements FaqSectionInterface.
func (h *faqSectionRepository) EditByIDFaqSection(ctx context.Context, req entity.FaqSectionEntity) error {
	modelFaqSection := model.FaqSection{}

	err = h.DB.Where("id =?", req.ID).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDFaqSection - 1: %v", err)
		return err
	}
	modelFaqSection.Description = req.Description
	modelFaqSection.Title = req.Title

	err = h.DB.Save(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDFaqSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllFaqSection implements FaqSectionInterface.
func (h *faqSectionRepository) FetchAllFaqSection(ctx context.Context) ([]entity.FaqSectionEntity, error) {
	modelFaqSection := []model.FaqSection{}
	err = h.DB.Select("id", "title", "description").Find(&modelFaqSection).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllFaqSection - 1: %v", err)
		return nil, err
	}

	var faqSectionRepositoryEntities []entity.FaqSectionEntity
	for _, v := range modelFaqSection {
		faqSectionRepositoryEntities = append(faqSectionRepositoryEntities, entity.FaqSectionEntity{
			ID:          v.ID,
			Description: v.Description,
			Title:       v.Title,
		})
	}

	return faqSectionRepositoryEntities, nil
}

// FetchByIDFaqSection implements FaqSectionInterface.
func (h *faqSectionRepository) FetchByIDFaqSection(ctx context.Context, id int64) (*entity.FaqSectionEntity, error) {
	modelFaqSection := model.FaqSection{}
	err = h.DB.Select("id", "title", "description").Where("id = ?", id).First(&modelFaqSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDFaqSection - 1: %v", err)
		return nil, err
	}

	return &entity.FaqSectionEntity{
		ID:          modelFaqSection.ID,
		Description: modelFaqSection.Description,
		Title:       modelFaqSection.Title,
	}, nil
}

func NewFaqSectionRepository(DB *gorm.DB) FaqSectionRepositoryInterface {
	return &faqSectionRepository{
		DB: DB,
	}
}
