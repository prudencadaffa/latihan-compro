package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type OurTeamServiceInterface interface {
	CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error)
	FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error)
	EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error
	DeleteByIDOurTeam(ctx context.Context, id int64) error
}

type ourTeamService struct {
	ourTeamRepo repository.OurTeamInterface
}

// CreateOurTeam implements OurTeamServiceInterface.
func (h *ourTeamService) CreateOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	return h.ourTeamRepo.CreateOurTeam(ctx, req)
}

// DeleteByIDOurTeam implements OurTeamServiceInterface.
func (h *ourTeamService) DeleteByIDOurTeam(ctx context.Context, id int64) error {
	return h.ourTeamRepo.DeleteByIDOurTeam(ctx, id)
}

// EditByIDOurTeam implements OurTeamServiceInterface.
func (h *ourTeamService) EditByIDOurTeam(ctx context.Context, req entity.OurTeamEntity) error {
	return h.ourTeamRepo.EditByIDOurTeam(ctx, req)
}

// FetchAllOurTeam implements OurTeamServiceInterface.
func (h *ourTeamService) FetchAllOurTeam(ctx context.Context) ([]entity.OurTeamEntity, error) {
	return h.ourTeamRepo.FetchAllOurTeam(ctx)
}

// FetchByIDOurTeam implements OurTeamServiceInterface.
func (h *ourTeamService) FetchByIDOurTeam(ctx context.Context, id int64) (*entity.OurTeamEntity, error) {
	return h.ourTeamRepo.FetchByIDOurTeam(ctx, id)
}
func NewOurTeamService(ourTeamRepo repository.OurTeamInterface) OurTeamServiceInterface {
	return &ourTeamService{
		ourTeamRepo: ourTeamRepo,
	}
}
