package app

import (
	"context"
	"latihan-compro/config"
	"latihan-compro/internal/adapter/handler"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/core/service"
	"latihan-compro/utils/auth"
	"latihan-compro/utils/validator"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	en "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	jwt := auth.NewJwt(cfg)
	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo, cfg, jwt)

	e := echo.New()
	e.Use(middleware.CORS())
	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator

	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	handler.NewUserHandler(e, userService)

	//Start the server
	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}

		err := e.Start(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	// Block until a signal is received.
	<-quit
	log.Println("server shutdown of 5 second.")

	// Shutdown with gracefully, waiting max 5 seccond for current processing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
