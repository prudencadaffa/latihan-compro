package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type HeroSectionInterface interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error)
	EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	DeleteByIDHeroSection(ctx context.Context, id int64) error
}

type heroSection struct {
	DB *gorm.DB
}

// CreateHeroSection implements HeroSectionInterface.
func (h *heroSection) CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	modelHeroSection := model.HeroSection{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  &req.PathVideo,
		PathBanner: req.Banner,
	}

	if err = h.DB.Create(&modelHeroSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateHeroSection - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) DeleteByIDHeroSection(ctx context.Context, id int64) error {
	modelHeroSection := model.HeroSection{}

	err = h.DB.Where("id = ?", id).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDHeroSection - 1: %v", err)
		return err
	}

	err = h.DB.Delete(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDHeroSection - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	modelHeroSection := model.HeroSection{}

	err = h.DB.Where("id =?", req.ID).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDHeroSection - 1: %v", err)
		return err
	}
	modelHeroSection.Heading = req.Heading
	modelHeroSection.SubHeading = req.SubHeading
	modelHeroSection.PathVideo = &req.PathVideo
	modelHeroSection.PathBanner = req.Banner
	err = h.DB.Save(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDHeroSection - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllHeroSection implements HeroSectionInterface.
func (h *heroSection) FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error) {
	modelHeroSection := []model.HeroSection{}
	err = h.DB.Select("id", "heading", "sub_heading", "path_video", "path_banner").Find(&modelHeroSection).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllHeroSection - 1: %v", err)
		return nil, err
	}

	var heroSectionEntities []entity.HeroSectionEntity
	for _, v := range modelHeroSection {
		heroSectionEntities = append(heroSectionEntities, entity.HeroSectionEntity{
			ID:         v.ID,
			Heading:    v.Heading,
			SubHeading: v.SubHeading,
			PathVideo:  *v.PathVideo,
			Banner:     v.PathBanner,
		})
	}

	return heroSectionEntities, nil
}

// FetchByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error) {
	modelHeroSection := model.HeroSection{}
	err = h.DB.Where("id = ?", id).First(&modelHeroSection).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDHeroSection - 1: %v", err)
		return nil, err
	}

	return &entity.HeroSectionEntity{
		ID:         modelHeroSection.ID,
		Heading:    modelHeroSection.Heading,
		SubHeading: modelHeroSection.SubHeading,
		PathVideo:  *modelHeroSection.PathVideo,
		Banner:     modelHeroSection.PathBanner,
	}, nil
}

func NewHeroSectionRepository(DB *gorm.DB) HeroSectionInterface {
	return &heroSection{
		DB: DB,
	}
}
