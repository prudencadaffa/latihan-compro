package app

import (
	"context"
	"latihan-compro/config"
	"latihan-compro/internal/adapter/handler"
	"latihan-compro/internal/adapter/messaging"
	"latihan-compro/internal/adapter/repository"
	"latihan-compro/internal/adapter/storage"
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
		log.Fatal("Error connecting to database: %v", err)
		return
	}

	jwt := auth.NewJwt(cfg)
	emailMessage := messaging.NewEmailMessaging(cfg)

	userRepo := repository.NewUserRepository(db.DB)
	heroSectionRepo := repository.NewHeroSectionRepository(db.DB)
	clientSectionRepo := repository.NewClientSectionRepository(db.DB)
	aboutCompanyRepo := repository.NewAboutCompanyRepository(db.DB)
	faqRepo := repository.NewFaqSectionRepository(db.DB)
	ourTeamRepo := repository.NewOurTeamRepository(db.DB)
	aboutCompanyKeynoteRepo := repository.NewAboutCompanyKeynoteRepository(db.DB)
	serviceSectionRepo := repository.NewServiceSectionRepository(db.DB)
	appointmentRepo := repository.NewAppointmentRepository(db.DB)
	portofolioRepo := repository.NewPortofolioSectionRepository(db.DB)
	portofolioDetailRepo := repository.NewPortofolioDetailRepository(db.DB)
	portofolioTestimonialRepo := repository.NewPortofolioTestimonialRepository(db.DB)
	contactUsRepo := repository.NewContactUsRepository(db.DB)
	serviceDetailRepo := repository.NewServiceDetailRepository(db.DB)

	userService := service.NewUserService(userRepo, cfg, jwt)
	heroSectionService := service.NewHeroSectionService(heroSectionRepo)
	clientSectionService := service.NewClientSectionService(clientSectionRepo)
	aboutCompanyService := service.NewAboutCompanyService(aboutCompanyRepo)
	faqService := service.NewFaqSectionService(faqRepo)
	ourTeamService := service.NewOurTeamService(ourTeamRepo)
	aboutCompanyKeynoteService := service.NewAboutCompanyKeynoteService(aboutCompanyKeynoteRepo, aboutCompanyRepo)
	serviceSectionService := service.NewServiceSectionService(serviceSectionRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo, emailMessage)
	portofolioService := service.NewPortofolioSectionService(portofolioRepo)
	portofolioDetailService := service.NewPortofolioDetailService(portofolioDetailRepo, portofolioRepo)
	portofolioTestimonialService := service.NewPortofolioTestimonialService(portofolioTestimonialRepo, portofolioRepo)
	contactUsService := service.NewContactUsService(contactUsRepo)
	serviceDetailService := service.NewServiceDetailService(serviceDetailRepo)

	storageAdapter := storage.NewSupabase(cfg)

	e := echo.New()
	e.Use(middleware.CORS())

	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator

	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	handler.NewUserHandler(e, userService)
	handler.NewUploadImage(e, storageAdapter, cfg)
	handler.NewHeroSectionHandler(e, cfg, heroSectionService)
	handler.NewClientSectionHandler(e, clientSectionService, cfg)
	handler.NewAboutCompanyHandler(e, aboutCompanyService, cfg)
	handler.NewFaqSectionHandler(e, faqService, cfg)
	handler.NewOurTeamHandler(e, cfg, ourTeamService)
	handler.NewAboutCompanyKeynoteHandler(e, aboutCompanyKeynoteService, cfg)
	handler.NewServiceSectionHandler(e, serviceSectionService, cfg)
	handler.NewAppointmentHandler(e, appointmentService, cfg)
	handler.NewPortofolioSectionHandler(e, portofolioService, cfg)
	handler.NewPortofolioDetailHandler(e, portofolioDetailService, cfg)
	handler.NewPortofolioTestimonialHandler(e, portofolioTestimonialService, cfg)
	handler.NewContactUsHandler(e, contactUsService, cfg)
	handler.NewServiceDetailHandler(e, serviceDetailService, cfg)

	// Starting server
	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}

		err := e.Start(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatal("error starting server: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	// Block until a signal is received.
	<-quit

	log.Println("server shutdown of 5 second.")

	// gracefully shutdown the server, waiting max 5 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
