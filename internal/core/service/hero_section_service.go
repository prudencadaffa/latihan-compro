package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type HeroSectionServiceInterface interface {
	CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error)
	EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error
	DeleteByIDHeroSection(ctx context.Context, id int64) error
}

type heroSectionService struct {
	heroSectionRepo repository.HeroSectionInterface
}

// CreateHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) CreateHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	return h.heroSectionRepo.CreateHeroSection(ctx, req)
}

// FetchAllHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) FetchAllHeroSection(ctx context.Context) ([]entity.HeroSectionEntity, error) {
	return h.heroSectionRepo.FetchAllHeroSection(ctx)
}

// FetchByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) FetchByIDHeroSection(ctx context.Context, id int64) (*entity.HeroSectionEntity, error) {
	return h.heroSectionRepo.FetchByIDHeroSection(ctx, id)
}

// EditByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) EditByIDHeroSection(ctx context.Context, req entity.HeroSectionEntity) error {
	return h.heroSectionRepo.EditByIDHeroSection(ctx, req)
}

// DeleteByIDHeroSection implements HeroSectionServiceInterface.
func (h *heroSectionService) DeleteByIDHeroSection(ctx context.Context, id int64) error {
	return h.heroSectionRepo.DeleteByIDHeroSection(ctx, id)
}

func NewHeroSectionService(heroSectionRepo repository.HeroSectionInterface) HeroSectionServiceInterface {
	return &heroSectionService{
		heroSectionRepo: heroSectionRepo,
	}
}
