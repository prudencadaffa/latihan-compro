package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OurTeamInterface interface {
	CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error)
	FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error)
	EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	DeleteByIDOurTeam(ctx context.Context, id int64) error
}
type ourTeamRepository struct {
	DB *gorm.DB
}

// CreateOurTeam implements OurTeamInterface.
func (h *ourTeamRepository) CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	modelOurTeam := model.OurTeam{
		Name:      req.Name,
		Role:      req.Role,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}

	if err = h.DB.Create(&modelOurTeam).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateOurTeam - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDOurTeam implements OurTeamInterface.
func (h *ourTeamRepository) DeleteByIDOurTeam(ctx context.Context, id int64) error {
	modelOurTeam := model.OurTeam{}

	err = h.DB.Where("id = ?", id).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDOurTeam - 1: %v", err)
		return err
	}

	err = h.DB.Delete(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDOurTeam - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDOurTeam implements OurTeamInterface.
func (h *ourTeamRepository) EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	modelOurTeam := model.OurTeam{}

	err = h.DB.Where("id =?", req.ID).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDOurTeam - 1: %v", err)
		return err
	}
	modelOurTeam.Name = req.Name
	modelOurTeam.Role = req.Role
	modelOurTeam.PathPhoto = req.PathPhoto
	modelOurTeam.Tagline = req.Tagline
	err = h.DB.Save(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDOurTeam - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllOurTeam implements OurTeamInterface.
func (h *ourTeamRepository) FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error) {
	modelOurTeam := []model.OurTeam{}
	err = h.DB.Select("id", "name", "role", "path_photo", "tagline").Find(&modelOurTeam).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllOurTeam - 1: %v", err)
		return nil, err
	}

	var ourTeamRepositoryEntities []entity.OurTeamEntity
	for _, v := range modelOurTeam {
		ourTeamRepositoryEntities = append(ourTeamRepositoryEntities, entity.OurTeamEntity{
			ID:        v.ID,
			Name:      v.Name,
			PathPhoto: v.PathPhoto,
			Tagline:   v.Tagline,
			Role:      v.Role,
		})
	}

	return ourTeamRepositoryEntities, nil
}

// FetchByIDOurTeam implements OurTeamInterface.
func (h *ourTeamRepository) FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error) {
	modelOurTeam := model.OurTeam{}
	err = h.DB.Select("id", "name", "role", "path_photo", "tagline").Where("id = ?", id).First(&modelOurTeam).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDOurTeam - 1: %v", err)
		return nil, err
	}

	return &entity.OurTeamEntity{
		ID:        modelOurTeam.ID,
		Name:      modelOurTeam.Name,
		PathPhoto: modelOurTeam.PathPhoto,
		Tagline:   modelOurTeam.Tagline,
		Role:      modelOurTeam.Role,
	}, nil
}

func NewOurTeamRepository(DB *gorm.DB) OurTeamInterface {
	return &ourTeamRepository{
		DB: DB,
	}
}
