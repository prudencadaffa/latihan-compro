package service

import (
	"context"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"
)

type ServiceDetailServiceInterface interface {
	CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error)
	FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error)
	EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error
	DeleteByIDServiceDetail(ctx context.Context, id int64) error

	GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error)
}

type serviceDetailService struct {
	serviceDetailRepo repository.ServiceDetailRepositoryInterface
}

// CreateServiceDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) CreateServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	return c.serviceDetailRepo.CreateServiceDetail(ctx, req)
}

// FetchAllServiceDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) FetchAllServiceDetail(ctx context.Context) ([]entity.ServiceDetailEntity, error) {
	return c.serviceDetailRepo.FetchAllServiceDetail(ctx)
}

// FetchByIDServiceDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) FetchByIDServiceDetail(ctx context.Context, id int64) (*entity.ServiceDetailEntity, error) {
	return c.serviceDetailRepo.FetchByIDServiceDetail(ctx, id)
}

// EditByIDServiceDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) EditByIDServiceDetail(ctx context.Context, req entity.ServiceDetailEntity) error {
	return c.serviceDetailRepo.EditByIDServiceDetail(ctx, req)
}

// DeleteByIDServiceDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) DeleteByIDServiceDetail(ctx context.Context, id int64) error {
	return c.serviceDetailRepo.DeleteByIDServiceDetail(ctx, id)
}

// GetByServiceIDDetail implements ServiceDetailServiceInterface.
func (c *serviceDetailService) GetByServiceIDDetail(ctx context.Context, serviceId int64) (*entity.ServiceDetailEntity, error) {
	return c.serviceDetailRepo.GetByServiceIDDetail(ctx, serviceId)
}
func NewServiceDetailService(serviceDetailRepo repository.ServiceDetailRepositoryInterface) ServiceDetailServiceInterface {
	return &serviceDetailService{
		serviceDetailRepo: serviceDetailRepo,
	}
}
