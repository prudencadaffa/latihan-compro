package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ContactUsInterface interface {
	CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error
	FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error)
	FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error)
	EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error
	DeleteByIDContactUs(ctx context.Context, id int64) error
}
type contactUsRepository struct {
	DB *gorm.DB
}

// CreateContactUs implements ContactUsInterface.
func (h *contactUsRepository) CreateContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	modelContactUs := model.ContactUs{
		CompanyName:  req.CompanyName,
		LocationName: req.LocationName,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
	}

	if err = h.DB.Create(&modelContactUs).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateContactUs - 1: %v", err)
		return err
	}
	return nil
}

// FetchAllContactUs implements ContactUsInterface.
func (h *contactUsRepository) FetchAllContactUs(ctx context.Context) ([]entity.ContactUsEntity, error) {
	modelContactUs := []model.ContactUs{}
	err = h.DB.Select("id", "location_name", "address", "phone_number", "company_name").Find(&modelContactUs).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllContactUs - 1: %v", err)
		return nil, err
	}

	var contactUsRepositoryEntities []entity.ContactUsEntity
	for _, v := range modelContactUs {
		contactUsRepositoryEntities = append(contactUsRepositoryEntities, entity.ContactUsEntity{
			ID:           v.ID,
			CompanyName:  v.CompanyName,
			LocationName: v.LocationName,
			Address:      v.Address,
			PhoneNumber:  v.PhoneNumber,
		})
	}

	return contactUsRepositoryEntities, nil
}

// FetchByIDContactUs implements ContactUsInterface.
func (h *contactUsRepository) FetchByIDContactUs(ctx context.Context, id int64) (*entity.ContactUsEntity, error) {
	modelContactUs := model.ContactUs{}
	err = h.DB.Select("id", "location_name", "address", "phone_number", "company_name").Where("id = ?", id).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDContactUs - 1: %v", err)
		return nil, err
	}

	return &entity.ContactUsEntity{
		ID:           modelContactUs.ID,
		CompanyName:  modelContactUs.CompanyName,
		LocationName: modelContactUs.LocationName,
		Address:      modelContactUs.Address,
		PhoneNumber:  modelContactUs.PhoneNumber,
	}, nil
}

// EditByIDContactUs implements ContactUsInterface.
func (h *contactUsRepository) EditByIDContactUs(ctx context.Context, req entity.ContactUsEntity) error {
	modelContactUs := model.ContactUs{}

	err = h.DB.Where("id =?", req.ID).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDContactUs - 1: %v", err)
		return err
	}
	modelContactUs.Address = req.Address
	modelContactUs.CompanyName = req.CompanyName
	modelContactUs.PhoneNumber = req.PhoneNumber
	modelContactUs.LocationName = req.LocationName
	err = h.DB.Save(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDContactUs - 2: %v", err)
		return err
	}
	return nil
}

// DeleteByIDContactUs implements ContactUsInterface.
func (h *contactUsRepository) DeleteByIDContactUs(ctx context.Context, id int64) error {
	modelContactUs := model.ContactUs{}

	err = h.DB.Where("id = ?", id).First(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDContactUs - 1: %v", err)
		return err
	}

	err = h.DB.Delete(&modelContactUs).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDContactUs - 2: %v", err)
		return err
	}
	return nil
}
func NewContactUsRepository(DB *gorm.DB) ContactUsInterface {
	return &contactUsRepository{
		DB: DB,
	}
}
