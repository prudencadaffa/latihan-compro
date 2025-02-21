package repository

import (
	"context"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	err  error
	code string
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type userRepo struct {
	db *gorm.DB
}

// GetUserByEmail implements UserRepositoryInterface.
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var modelUser model.User

	err = u.db.Select("email", "password", "name", "id").Where("email = ?", email).First(&modelUser).Error
	if err != nil {
		code = "[REPOSITORY] GetUserByEmail - 1"
		log.Err(err).Msg(code)
		return nil, err
	}

	return &entity.UserEntity{
		ID:       modelUser.ID,
		Name:     modelUser.Name,
		Email:    email,
		Password: modelUser.Password,
	}, nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepo{db: db}
}
