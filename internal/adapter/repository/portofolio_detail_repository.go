package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PortofolioDetailRepositoryInterface interface {
	CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error)
	FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error)
	EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error
	DeleteByIDPortofolioDetail(ctx context.Context, id int64) error

	FetchDetailPotofolioByPortoID(ctx context.Context, portoID int64) (*entity.PortofolioDetailEntity, error)
}
type portofolioDetailRepository struct {
	DB *gorm.DB
}

// CreatePortofolioDetail implements PortofolioDetailInterface.
func (h *portofolioDetailRepository) CreatePortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	modelPortofolioDetail := model.PortofolioDetail{
		PortofolioSectionID: req.PortofolioSection.ID,
		Category:            req.Category,
		ClientName:          req.ClientName,
		ProjectDate:         req.ProjectDate,
		ProjectUrl:          &req.ProjectUrl,
		Title:               req.Title,
		Description:         req.Description,
	}

	if err = h.DB.Create(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePortofolioDetail - 1: %v", err)
		return err
	}
	return nil
}

// FetchAllPortofolioDetail implements PortofolioDetailInterface.
func (h *portofolioDetailRepository) FetchAllPortofolioDetail(ctx context.Context) ([]entity.PortofolioDetailEntity, error) {
	rows, err := h.DB.
		Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name", "pd.project_date", "ps.name").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd.deleted_at IS NULL").
		Order("created_at DESC").
		Rows()

	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllPortofolioDetail - 1: %v", err)
		return nil, err
	}

	var portofolioDetailRepositoryEntities []entity.PortofolioDetailEntity
	for rows.Next() {
		portofolioDetail := entity.PortofolioDetailEntity{}
		err = rows.Scan(&portofolioDetail.ID,
			&portofolioDetail.Title,
			&portofolioDetail.Category,
			&portofolioDetail.ClientName,
			&portofolioDetail.ProjectDate,
			&portofolioDetail.PortofolioSection.Name)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllPortofolioDetail - 2: %v", err)
			return nil, err
		}

		portofolioDetailRepositoryEntities = append(portofolioDetailRepositoryEntities, portofolioDetail)
	}

	return portofolioDetailRepositoryEntities, nil
}

// FetchByIDPortofolioDetail implements PortofolioDetailInterface.
func (h *portofolioDetailRepository) FetchByIDPortofolioDetail(ctx context.Context, id int64) (*entity.PortofolioDetailEntity, error) {
	rows, err := h.DB.
		Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name", "pd.project_date", "pd.description", "pd.project_url", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("pd.id =? AND pd.deleted_at IS NULL", id).
		Order("pd.created_at DESC").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDPortofolioDetail - 1: %v", err)
		return nil, err
	}

	var portofolioDetailEntity entity.PortofolioDetailEntity
	for rows.Next() {
		err = rows.Scan(&portofolioDetailEntity.ID,
			&portofolioDetailEntity.Title,
			&portofolioDetailEntity.Category,
			&portofolioDetailEntity.ClientName,
			&portofolioDetailEntity.ProjectDate,
			&portofolioDetailEntity.Description,
			&portofolioDetailEntity.ProjectUrl,
			&portofolioDetailEntity.PortofolioSection.ID,
			&portofolioDetailEntity.PortofolioSection.Name,
			&portofolioDetailEntity.PortofolioSection.Thumbnail)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchByIDPortofolioDetail - 2: %v", err)
			return nil, err
		}
	}

	return &portofolioDetailEntity, nil
}

// EditByIDPortofolioDetail implements PortofolioDetailInterface.
func (h *portofolioDetailRepository) EditByIDPortofolioDetail(ctx context.Context, req entity.PortofolioDetailEntity) error {
	modelPortofolioDetail := model.PortofolioDetail{}

	if err = h.DB.Where("id =?", req.ID).First(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioDetail - 1: %v", err)
		return err
	}
	modelPortofolioDetail.Title = req.Title
	modelPortofolioDetail.Description = req.Description
	modelPortofolioDetail.Category = req.Category
	modelPortofolioDetail.ClientName = req.ClientName
	modelPortofolioDetail.ProjectDate = req.ProjectDate
	modelPortofolioDetail.ProjectUrl = &req.ProjectUrl
	modelPortofolioDetail.PortofolioSectionID = req.PortofolioSection.ID

	if err = h.DB.Save(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] EditByIDPortofolioDetail - 2: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPortofolioDetail implements PortofolioDetailInterface.
func (h *portofolioDetailRepository) DeleteByIDPortofolioDetail(ctx context.Context, id int64) error {
	modelPortofolioDetail := model.PortofolioDetail{}

	if err = h.DB.Where("id = ?", id).First(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioDetail - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelPortofolioDetail).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPortofolioDetail - 2: %v", err)
		return err
	}
	return nil
}

// FetchDetailPotofolioByPortoID implements PortofolioDetailRepositoryInterface.
func (h *portofolioDetailRepository) FetchDetailPotofolioByPortoID(ctx context.Context, portoID int64) (*entity.PortofolioDetailEntity, error) {
	rows, err := h.DB.
		Table("portofolio_details as pd").
		Select("pd.id", "pd.title", "pd.category", "pd.client_name",
			"pd.project_date", "pd.description", "pd.project_url", "ps.id", "ps.name", "ps.thumbnail").
		Joins("inner join portofolio_sections as ps on ps.id = pd.portofolio_section_id").
		Where("ps.id =? AND pd.deleted_at IS NULL", portoID).
		Order("ps.created_at DESC").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchDetailPotofolioByPortoID - 1: %v", err)
		return nil, err
	}

	var portofolioDetailEntity entity.PortofolioDetailEntity
	for rows.Next() {
		err = rows.Scan(&portofolioDetailEntity.ID,
			&portofolioDetailEntity.Title,
			&portofolioDetailEntity.Category,
			&portofolioDetailEntity.ClientName,
			&portofolioDetailEntity.ProjectDate,
			&portofolioDetailEntity.Description,
			&portofolioDetailEntity.ProjectUrl,
			&portofolioDetailEntity.PortofolioSection.ID,
			&portofolioDetailEntity.PortofolioSection.Name,
			&portofolioDetailEntity.PortofolioSection.Thumbnail)

		if err != nil {
			log.Errorf("[REPOSITORY] FetchDetailPotofolioByPortoID - 2: %v", err)
			return nil, err
		}
	}

	return &portofolioDetailEntity, nil
}
func NewPortofolioDetailRepository(DB *gorm.DB) PortofolioDetailRepositoryInterface {
	return &portofolioDetailRepository{
		DB: DB,
	}
}
