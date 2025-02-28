package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AboutCompanyInterface interface {
	CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error)
	FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error)
	EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error
	DeleteByIDAboutCompany(ctx context.Context, id int64) error
	FetchAllCompanyAndKeynote(ctx context.Context) (*entity.AboutCompanyEntity, error)
}

type aboutCompanyRepository struct {
	DB *gorm.DB
}

// FetchAllCompanyAndKeynote implements AboutCompanyInterface.
func (h *aboutCompanyRepository) FetchAllCompanyAndKeynote(ctx context.Context) (*entity.AboutCompanyEntity, error) {
	modelAboutCompany := model.AboutCompany{}
	err := h.DB.Select("id", "description").Order("created_at DESC").Limit(1).Find(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllCompanyAndKeynote - 1: %v", err)
		return nil, err
	}

	var aboutCompanyRepositoryEntities entity.AboutCompanyEntity
	var aboutCompanyKeynoteModel []model.AboutCompanyKeynote
	err = h.DB.Select("id", "keypoint", "path_image", "about_company_id").Where("about_company_id = ?", modelAboutCompany.ID).Find(&aboutCompanyKeynoteModel).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllCompanyAndKeynote - 2: %v", err)
		return nil, err
	}

	var aboutCompanyKeynoteEntity []entity.AboutCompanyKeynoteEntity
	for _, val := range aboutCompanyKeynoteModel {
		aboutCompanyKeynoteEntity = append(aboutCompanyKeynoteEntity, entity.AboutCompanyKeynoteEntity{
			ID:             val.ID,
			AboutCompanyID: modelAboutCompany.ID,
			Keynote:        val.Keypoint,
			PathImage:      *val.PathImage,
		})
	}

	aboutCompanyRepositoryEntities.ID = modelAboutCompany.ID
	aboutCompanyRepositoryEntities.Description = modelAboutCompany.Description
	aboutCompanyRepositoryEntities.Keynote = aboutCompanyKeynoteEntity

	return &aboutCompanyRepositoryEntities, nil
}

// CreateAboutCompany implements AboutCompanyInterface.
func (h *aboutCompanyRepository) CreateAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	modelAboutCompany := model.AboutCompany{
		Description: req.Description,
	}

	if err := h.DB.Create(&modelAboutCompany).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAboutCompany - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompanyRepository) DeleteByIDAboutCompany(ctx context.Context, id int64) error {
	modelAboutCompany := model.AboutCompany{}
	err := h.DB.Where("id = ?", id).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompany - 1: %v", err)
		return err
	}

	err = h.DB.Delete(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompany - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompanyRepository) EditByIDAboutCompany(ctx context.Context, req entity.AboutCompanyEntity) error {
	modelAboutCompany := model.AboutCompany{}
	err := h.DB.Where("id =?", req.ID).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDAboutCompany - 1: %v", err)
		return err
	}
	modelAboutCompany.Description = req.Description

	err = h.DB.Save(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDAboutCompany - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAboutCompany implements AboutCompanyInterface.
func (h *aboutCompanyRepository) FetchAllAboutCompany(ctx context.Context) ([]entity.AboutCompanyEntity, error) {
	modelAboutCompany := []model.AboutCompany{}
	err := h.DB.Select("id", "description").Order("created_at DESC").Find(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAboutCompany - 1: %v", err)
		return nil, err
	}

	var aboutCompanyRepositoryEntities []entity.AboutCompanyEntity
	for _, v := range modelAboutCompany {
		aboutCompanyRepositoryEntities = append(aboutCompanyRepositoryEntities, entity.AboutCompanyEntity{
			ID:          v.ID,
			Description: v.Description,
		})
	}

	return aboutCompanyRepositoryEntities, nil
}

// FetchByIDAboutCompany implements AboutCompanyInterface.
func (h *aboutCompanyRepository) FetchByIDAboutCompany(ctx context.Context, id int64) (*entity.AboutCompanyEntity, error) {
	modelAboutCompany := model.AboutCompany{}
	err := h.DB.Select("id", "description").Where("id = ?", id).First(&modelAboutCompany).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDAboutCompany - 1: %v", err)
		return nil, err
	}

	return &entity.AboutCompanyEntity{
		ID:          modelAboutCompany.ID,
		Description: modelAboutCompany.Description,
	}, nil
}

func NewAboutCompanyRepository(DB *gorm.DB) AboutCompanyInterface {
	return &aboutCompanyRepository{
		DB: DB,
	}
}
