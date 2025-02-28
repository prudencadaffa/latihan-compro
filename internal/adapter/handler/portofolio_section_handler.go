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

type PortofolioSectionHandlerInterface interface {
	CreatePortofolioSection(c echo.Context) error
	FetchAllPortofolioSection(c echo.Context) error
	FetchByIDPortofolioSection(c echo.Context) error
	EditByIDPortofolioSection(c echo.Context) error
	DeleteByIDPortofolioSection(c echo.Context) error

	FetchAllPortofolioHome(c echo.Context) error
}
type portofolioSectionHandler struct {
	portofolioSectionService service.PortofolioSectionServiceInterface
}

// CreatePortofolioSection implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) CreatePortofolioSection(c echo.Context) error {
	var (
		req       = request.PortofolioSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioSectionEntity{
		Thumbnail: req.Thumbnail,
		Name:      req.Name,
		Tagline:   req.Tagline,
	}

	err = cs.portofolioSectionService.CreatePortofolioSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllPortofolioSection implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) FetchAllPortofolioSection(c echo.Context) error {
	var (
		resp                  = response.DefaultSuccessResponse{}
		respError             = response.ErrorResponseDefault{}
		ctx                   = c.Request().Context()
		respPortofolioSection = []response.PortofolioSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllPortofolioSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.portofolioSectionService.FetchAllPortofolioSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respPortofolioSection = append(respPortofolioSection, response.PortofolioSectionResponse{
			ID:        val.ID,
			Name:      val.Name,
			Tagline:   val.Tagline,
			Thumbnail: val.Thumbnail,
		})
	}

	resp.Meta.Message = "Success fetch all portofolio section"
	resp.Meta.Status = true
	resp.Data = respPortofolioSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioSection implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) FetchByIDPortofolioSection(c echo.Context) error {
	var (
		resp                  = response.DefaultSuccessResponse{}
		respError             = response.ErrorResponseDefault{}
		ctx                   = c.Request().Context()
		respPortofolioSection = response.PortofolioSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDPortofolioSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioSection := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioSection)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.portofolioSectionService.FetchByIDPortofolioSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioSection.ID = result.ID
	respPortofolioSection.Name = result.Name
	respPortofolioSection.Tagline = result.Tagline
	respPortofolioSection.Thumbnail = result.Thumbnail
	resp.Meta.Message = "Success fetch portofolio section by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDPortofolioSection implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) EditByIDPortofolioSection(c echo.Context) error {
	var (
		req       = request.PortofolioSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDPortofolioSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioSection := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioSection)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioSectionEntity{
		ID:        id,
		Thumbnail: req.Thumbnail,
		Name:      req.Name,
		Tagline:   req.Tagline,
	}

	err = cs.portofolioSectionService.EditByIDPortofolioSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDPortofolioSection implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) DeleteByIDPortofolioSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDPortofolioSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioSection := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioSection)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.portofolioSectionService.DeleteByIDPortofolioSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete portofolio section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// FetchAllPortofolioHome implements PortofolioSectionHandlerInterface.
func (cs *portofolioSectionHandler) FetchAllPortofolioHome(c echo.Context) error {
	var (
		respPortofolios = []response.PortofolioSectionResponse{}
		resp            = response.DefaultSuccessResponse{}
		respError       = response.ErrorResponseDefault{}
		ctx             = c.Request().Context()
	)

	results, err := cs.portofolioSectionService.FetchAllPortofolioSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPortofolioHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respPortofolios = append(respPortofolios, response.PortofolioSectionResponse{
			ID:        val.ID,
			Name:      val.Name,
			Tagline:   val.Tagline,
			Thumbnail: val.Thumbnail,
		})
	}
	resp.Meta.Message = "Success fetch all portofolio home"
	resp.Meta.Status = true
	resp.Data = respPortofolios
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}
func NewPortofolioSectionHandler(e *echo.Echo, portofolioSectionService service.PortofolioSectionServiceInterface, cfg *config.Config) PortofolioSectionHandlerInterface {
	h := &portofolioSectionHandler{
		portofolioSectionService: portofolioSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	portofolioSectionApp := e.Group("/portofolio-sections")
	portofolioSectionApp.GET("", h.FetchAllPortofolioHome)

	adminApp := portofolioSectionApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreatePortofolioSection)
	adminApp.GET("", h.FetchAllPortofolioSection)
	adminApp.GET("/:id", h.FetchByIDPortofolioSection)
	adminApp.PUT("/:id", h.EditByIDPortofolioSection)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioSection)

	return h
}
