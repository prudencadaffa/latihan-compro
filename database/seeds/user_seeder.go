package seeds

import (
	"latihan-compro/internal/core/domain/model"
	"latihan-compro/utils/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	admin := model.User{
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: bytes,
	}

	if err = db.FirstOrCreate(&admin, model.User{Email: "admin@mail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	} else {
		log.Info().Msg("Admin user has been seeded")
	}
}
