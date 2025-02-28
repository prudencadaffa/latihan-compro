package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AppointmentRepositoryInterface interface {
	CreateAppointment(ctx context.Context, req entity.AppointmentEntity) (string, error)
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	DeleteByIDAppointment(ctx context.Context, id int64) error
}
type appointmentRepository struct {
	DB *gorm.DB
}

// FetchByIDAppointment implements AppointmentRepositoryInterface.
func (h *appointmentRepository) FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error) {
	panic("unimplemented")
}

// CreateAppointment implements AppointmentRepositoryInterface.
func (h *appointmentRepository) CreateAppointment(ctx context.Context, req entity.AppointmentEntity) (string, error) {
	modelAppointment := model.Appointment{
		ServiceID:   req.ServiceID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Brief:       req.Brief,
		Budget:      req.Budget,
		MeetAt:      req.MeetAt,
	}

	if err = h.DB.Create(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateAppointment - 1: %v", err)
		return "", err
	}

	return modelAppointment.Email, nil

}

// FetchAllAppointment implements AppointmentInterface.
func (h *appointmentRepository) FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error) {
	rows, err := h.DB.
		Table("appointments as a").
		Select("a.id", "a.name", "a.email", "a.budget", "ss.name").
		Joins("inner join service_sections as ss on ss.id = a.service_id").
		Where("a.deleted_at IS NULL").
		Rows()
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllAppointment - 1: %v", err)
		return nil, err
	}

	var appointmentRepositoryEntities []entity.AppointmentEntity
	for rows.Next() {
		var appointment entity.AppointmentEntity
		err = rows.Scan(&appointment.ID, &appointment.Name, &appointment.Email, &appointment.Budget, &appointment.ServiceName)
		if err != nil {
			log.Errorf("[REPOSITORY] FetchAllAppointment - 2: %v", err)
			return nil, err
		}
		appointmentRepositoryEntities = append(appointmentRepositoryEntities, appointment)
	}

	return appointmentRepositoryEntities, nil
}

// DeleteByIDAppointment implements AppointmentInterface.
func (h *appointmentRepository) DeleteByIDAppointment(ctx context.Context, id int64) error {
	modelAppointment := model.Appointment{}

	if err = h.DB.Where("id = ?", id).First(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment - 1: %v", err)
		return err
	}

	if err = h.DB.Delete(&modelAppointment).Error; err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDAppointment - 2: %v", err)
		return err
	}
	return nil
}
func NewAppointmentRepository(DB *gorm.DB) AppointmentRepositoryInterface {
	return &appointmentRepository{
		DB: DB,
	}
}
