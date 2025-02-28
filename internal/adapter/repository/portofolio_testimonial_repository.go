package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioTestimonialRepositoryInterface interface {
	CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error)
	FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error)
	EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error
	DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error
}
type portofolioTestimonialRepository struct {
	DB *gorm.DB
}

// CreatePortofolioTestimonial implements PortofolioTestimonialInterface.
func (h *portofolioTestimonialRepository) CreatePortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{
		PortofolioSectionID: req.PortofolioSection.ID,
		Thumbnail:           req.Thumbnail,
		Message:             req.Message,
		ClientName:          req.ClientName,
		Role:                req.Role,
	}

	if err = h.DB.Create(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioTestimonial - 1: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioTestimonial implements PortofolioTestimonialInterface.
func (h *portofolioTestimonialRepository) FetchAllPortofolioTestimonial(ctx context.Context) ([]entity.PortofolioTestimonialEntity, error) {
	rows, err := h.DB.
		Table("portofolio_testimonials as pd").
		Select("pd.id", "pd.thumbnail", "pd.message", "pd.client_name", "pd.role", "ps.name").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd.deleted_at IS NULL").
		Order("pd.created_at DESC").
		Rows()

	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllPortofolioTestimonial - 1: %v", err)
		return nil, err
	}

	var portofolioTestimonialRepositoryEntities []entity.PortofolioTestimonialEntity
	for rows.Next() {
		portofolioTestimonial := entity.PortofolioTestimonialEntity{}
		err = rows.Scan(&portofolioTestimonial.ID,
			&portofolioTestimonial.Thumbnail,
			&portofolioTestimonial.Message,
			&portofolioTestimonial.ClientName,
			&portofolioTestimonial.Role,
			&portofolioTestimonial.PortofolioSection.Name)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllPortofolioTestimonial - 2: %v", err)
			return nil, err
		}

		portofolioTestimonialRepositoryEntities = append(portofolioTestimonialRepositoryEntities, portofolioTestimonial)
	}

	return portofolioTestimonialRepositoryEntities, nil
}

// FetchByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (h *portofolioTestimonialRepository) FetchByIDPortofolioTestimonial(ctx context.Context, id int64) (*entity.PortofolioTestimonialEntity, error) {
	rows, err := h.DB.
		Table("portofolio_testimonials as pd").
		Select("pd.id", "pd.thumbnail", "pd.message", "pd.client_name", "pd.role", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd.id =? AND pd.deleted_at IS NULL", id).
		Order("created_at DESC").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDPortofolioTestimonial - 1: %v", err)
		return nil, err
	}

	var portofolioTestimonialEntity entity.PortofolioTestimonialEntity
	for rows.Next() {
		err = rows.Scan(&portofolioTestimonialEntity.ID,
			&portofolioTestimonialEntity.Thumbnail,
			&portofolioTestimonialEntity.Message,
			&portofolioTestimonialEntity.ClientName,
			&portofolioTestimonialEntity.Role,
			&portofolioTestimonialEntity.PortofolioSection.ID,
			&portofolioTestimonialEntity.PortofolioSection.Name,
			&portofolioTestimonialEntity.PortofolioSection.Thumbnail)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchByIDPortofolioTestimonial - 2: %v", err)
			return nil, err
		}
	}

	return &portofolioTestimonialEntity, nil
}

// EditByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (h *portofolioTestimonialRepository) EditByIDPortofolioTestimonial(ctx context.Context, req entity.PortofolioTestimonialEntity) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{}

	if err = h.DB.Where("id =?", req.ID).First(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioTestimonial - 1: %v", err)
		return err
	}
	modelPortofolioTestimonial.Thumbnail = req.Thumbnail
	modelPortofolioTestimonial.Message = req.Message
	modelPortofolioTestimonial.ClientName = req.ClientName
	modelPortofolioTestimonial.Role = req.Role
	modelPortofolioTestimonial.PortofolioSectionID = req.PortofolioSection.ID

	if err = h.DB.Save(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioTestimonial - 2: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioTestimonial implements PortofolioTestimonialInterface.
func (h *portofolioTestimonialRepository) DeleteByIDPortofolioTestimonial(ctx context.Context, id int64) error {
	modelPortofolioTestimonial := model.PortofolioTestimonial{}

	if err = h.DB.Where("id = ?", id).First(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioTestimonial - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelPortofolioTestimonial).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioTestimonial - 2: %v", err)
		return err
	}
	return nil
}
func NewPortofolioTestimonialRepository(DB *gorm.DB) PortofolioTestimonialRepositoryInterface {
	return &portofolioTestimonialRepository{
		DB: DB,
	}
}
