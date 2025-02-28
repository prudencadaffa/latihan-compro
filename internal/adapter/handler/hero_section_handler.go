package handler

import (
	"latihan-compro/config"
	"latihan-compro/internal/adapter/handler/request"
	"latihan-compro/internal/adapter/handler/response"
	"latihan-compro/internal/core/domain/entity"
	"latihan-compro/internal/core/service"
	"latihan-compro/utils/conv"
	"latihan-compro/utils/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type HeroSectionHandlerInterface interface {
	CreateHeroSection(c echo.Context) error
	FetchAllHeroSection(c echo.Context) error
	FetchByIDHeroSection(c echo.Context) error
	EditByIDHeroSection(c echo.Context) error
	DeleteByIDHeroSection(c echo.Context) error

	FetchHeroDataHome(c echo.Context) error
}
type heroSectionHandler struct {
	heroSectionService service.HeroSectionServiceInterface
}

// CreateHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) CreateHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
	}

	err = h.heroSectionService.CreateHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchAllHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respHero  = []response.HeroSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}
	results, err := h.heroSectionService.FetchAllHeroSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respHero = append(respHero, response.HeroSectionResponse{
			ID:         val.ID,
			Heading:    val.Heading,
			SubHeading: val.SubHeading,
			PathVideo:  val.PathVideo,
			Banner:     val.Banner,
		})
	}

	resp.Meta.Message = "Success fetch all hero section"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchByIDHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
		respHero  = response.HeroSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := h.heroSectionService.FetchByIDHeroSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respHero.ID = result.ID
	respHero.Heading = result.Heading
	respHero.SubHeading = result.SubHeading
	respHero.PathVideo = result.PathVideo
	respHero.Banner = result.Banner
	resp.Meta.Message = "Success fetch hero section by ID"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) EditByIDHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		ID:         id,
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
	}

	err = h.heroSectionService.EditByIDHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDHeroSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) DeleteByIDHeroSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idHero := c.Param("id")
	id, err := conv.StringToInt64(idHero)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = h.heroSectionService.DeleteByIDHeroSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// FetchHeroDataHome implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchHeroDataHome(c echo.Context) error {
	var (
		respHero  = response.HeroSectionResponse{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	results, err := h.heroSectionService.FetchAllHeroSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchHeroDataHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respHero.Banner = results[0].Banner
	respHero.Heading = results[0].Heading
	respHero.SubHeading = results[0].SubHeading
	respHero.PathVideo = results[0].PathVideo
	respHero.ID = results[0].ID

	resp.Meta.Message = "Success fetch hero data home"
	resp.Meta.Status = true
	resp.Data = respHero
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}
func NewHeroSectionHandler(c *echo.Echo, cfg *config.Config, heroSectionService service.HeroSectionServiceInterface) HeroSectionHandlerInterface {
	heroHandler := &heroSectionHandler{
		heroSectionService: heroSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	heroApp := c.Group("/hero-sections")

	heroApp.GET("", heroHandler.FetchHeroDataHome)

	adminApp := heroApp.Group("/admin", mid.CheckToken())
	adminApp.GET("", heroHandler.FetchAllHeroSection)
	adminApp.POST("", heroHandler.CreateHeroSection)
	adminApp.GET("/:id", heroHandler.FetchByIDHeroSection)
	adminApp.PUT("/:id", heroHandler.EditByIDHeroSection)
	adminApp.DELETE("/:id", heroHandler.DeleteByIDHeroSection)

	return heroHandler
}
