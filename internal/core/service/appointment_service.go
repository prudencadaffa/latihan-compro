package service

import (
	"context"
	"fmt"
	"latihan-compro/internal/adapter/messaging"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
)

type AppointmentServiceInterface interface {
	FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error)
	FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error)
	DeleteByIDAppointment(ctx context.Context, id int64) error
	CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error
}
type appointmentService struct {
	appointmentRepo repository.AppointmentRepositoryInterface
	sendEmail       messaging.EmailMessagingInterface
}

// CreateAppointment implements AppointmentServiceInterface.
func (c *appointmentService) CreateAppointment(ctx context.Context, req entity.AppointmentEntity) error {
	email, err := c.appointmentRepo.CreateAppointment(ctx, req)
	if err != nil {
		log.Errorf("[SERVICE] CreateAppointment - 1: %v", err)
		return err
	}

	body := fmt.Sprintf("You have received a new appointment request from %s", email)
	err = c.sendEmail.SendEmailAppointment(nil, email, "New Appointment", body)
	if err != nil {
		log.Errorf("[SERVICE] CreateAppointment - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllAppointment implements AppointmentServiceInterface.
func (c *appointmentService) FetchAllAppointment(ctx context.Context) ([]entity.AppointmentEntity, error) {
	return c.appointmentRepo.FetchAllAppointment(ctx)
}

// FetchByIDAppointment implements AppointmentServiceInterface.
func (c *appointmentService) FetchByIDAppointment(ctx context.Context, id int64) (*entity.AppointmentEntity, error) {
	return c.appointmentRepo.FetchByIDAppointment(ctx, id)
}

// DeleteByIDAppointment implements AppointmentServiceInterface.
func (c *appointmentService) DeleteByIDAppointment(ctx context.Context, id int64) error {
	return c.appointmentRepo.DeleteByIDAppointment(ctx, id)
}
func NewAppointmentService(appointmentRepo repository.AppointmentRepositoryInterface, sendEmail messaging.EmailMessagingInterface) AppointmentServiceInterface {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
		sendEmail:       sendEmail,
	}
}
