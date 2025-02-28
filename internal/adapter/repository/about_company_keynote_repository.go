package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type AboutCompanyKeynoteInterface interface {
	CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error)
	FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error)
	EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error
	DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error
	FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error)
}

type aboutCompanyKeynoteRepository struct {
	DB *gorm.DB
}

// FetchByCompanyID implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) FetchByCompanyID(ctx context.Context, companyId int64) ([]entity.AboutCompanyKeynoteEntity, error) {
	rows, err := h.DB.Table("about_company_keynotes as ack").
		Select("ack.id", "ack.keypoint", "ack.about_company_id", "ack.path_image", "ac.description").
		Joins("inner join about_company as ac on ac.id = ack.about_company_id").
		Where("ack.about_company_id = ? AND ack.deleted_at IS NULL", companyId).
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByCompanyID - 1: %v", err)
		return nil, err
	}

	var aboutCompanyKeynoteRepositoryEntities []entity.AboutCompanyKeynoteEntity
	for rows.Next() {
		aboutCompanyKeynote := entity.AboutCompanyKeynoteEntity{}
		err = rows.Scan(&aboutCompanyKeynote.ID, &aboutCompanyKeynote.Keynote, &aboutCompanyKeynote.AboutCompanyID, &aboutCompanyKeynote.PathImage, &aboutCompanyKeynote.AboutCompanyDescription)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchByCompanyID - 2: %v", err)
			return nil, err
		}
		aboutCompanyKeynoteRepositoryEntities = append(aboutCompanyKeynoteRepositoryEntities, aboutCompanyKeynote)
	}

	return aboutCompanyKeynoteRepositoryEntities, nil
}

// CreateAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) CreateAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{
		AboutCompanyID: req.AboutCompanyID,
		Keypoint:       req.Keynote,
		PathImage:      &req.PathImage,
	}

	if err = h.DB.Create(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAboutCompanyKeynote - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) DeleteByIDAboutCompanyKeynote(ctx context.Context, id int64) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{}

	if err = h.DB.Where("id = ?", id).First(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompanyKeynote - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAboutCompanyKeynote - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) EditByIDAboutCompanyKeynote(ctx context.Context, req entity.AboutCompanyKeynoteEntity) error {
	modelAboutCompanyKeynote := model.AboutCompanyKeynote{}

	if err = h.DB.Where("id =?", req.ID).First(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDAboutCompanyKeynote - 1: %v", err)
		return err
	}
	modelAboutCompanyKeynote.AboutCompanyID = req.AboutCompanyID
	modelAboutCompanyKeynote.Keypoint = req.Keynote
	modelAboutCompanyKeynote.PathImage = &req.PathImage

	if err = h.DB.Save(&modelAboutCompanyKeynote).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDAboutCompanyKeynote - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) FetchAllAboutCompanyKeynote(ctx context.Context) ([]entity.AboutCompanyKeynoteEntity, error) {
	rows, err := h.DB.Table("about_company_keynotes as ack").
		Select("ack.id", "ack.keypoint", "ack.about_company_id", "ack.path_image", "ac.description").
		Joins("inner join about_company as ac on ac.id = ack.about_company_id").
		Where("ack.deleted_at IS NULL").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAboutCompanyKeynote - 1: %v", err)
		return nil, err
	}

	var aboutCompanyKeynoteRepositoryEntities []entity.AboutCompanyKeynoteEntity
	for rows.Next() {
		aboutCompanyKeynote := entity.AboutCompanyKeynoteEntity{}
		err = rows.Scan(&aboutCompanyKeynote.ID, &aboutCompanyKeynote.Keynote, &aboutCompanyKeynote.AboutCompanyID, &aboutCompanyKeynote.PathImage, &aboutCompanyKeynote.AboutCompanyDescription)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAboutCompanyKeynote - 2: %v", err)
			return nil, err
		}
		aboutCompanyKeynoteRepositoryEntities = append(aboutCompanyKeynoteRepositoryEntities, aboutCompanyKeynote)
	}

	return aboutCompanyKeynoteRepositoryEntities, nil
}

// FetchByIDAboutCompanyKeynote implements AboutCompanyKeynoteInterface.
func (h *aboutCompanyKeynoteRepository) FetchByIDAboutCompanyKeynote(ctx context.Context, id int64) (*entity.AboutCompanyKeynoteEntity, error) {
	rows, err := h.DB.Table("about_company_keynotes as ack").
		Select("ack.id", "ack.keypoint", "ack.about_company_id", "ack.path_image", "ac.description").
		Joins("inner join about_company as ac on ac.id = ack.about_company_id").
		Where("ack.id = ? AND ack.deleted_at IS NULL", id).
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDAboutCompanyKeynote - 1: %v", err)
		return nil, err
	}

	respEntity := entity.AboutCompanyKeynoteEntity{}
	for rows.Next() {
		err = rows.Scan(&respEntity.ID, &respEntity.Keynote, &respEntity.AboutCompanyID, &respEntity.PathImage, &respEntity.AboutCompanyDescription)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchByIDAboutCompanyKeynote - 2: %v", err)
			return nil, err
		}
	}
	return &respEntity, nil
}

func NewAboutCompanyKeynoteRepository(DB *gorm.DB) AboutCompanyKeynoteInterface {
	return &aboutCompanyKeynoteRepository{
		DB: DB,
	}
}
