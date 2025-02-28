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

type ServiceSectionHandlerInterface interface {
	CreateServiceSection(c echo.Context) error
	FetchAllServiceSection(c echo.Context) error
	FetchByIDServiceSection(c echo.Context) error
	EditByIDServiceSection(c echo.Context) error
	DeleteByIDServiceSection(c echo.Context) error

	FetchAllServiceHome(c echo.Context) error
}

type serviceSectionHandler struct {
	serviceSectionService service.ServiceSectionServiceInterface
}

// CreateServiceSection implements ServiceSectionHandlerInterface.

func (cs *serviceSectionHandler) CreateServiceSection(c echo.Context) error {
	var (
		req       = request.ServiceSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateServiceSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceSectionEntity{
		PathIcon: req.PathIcon,
		Name:     req.Name,
		Tagline:  req.Tagline,
	}

	err = cs.serviceSectionService.CreateServiceSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateServiceSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllServiceSection implements ServiceSectionHandlerInterface.
func (cs *serviceSectionHandler) FetchAllServiceSection(c echo.Context) error {
	var (
		resp               = response.DefaultSuccessResponse{}
		respError          = response.ErrorResponseDefault{}
		ctx                = c.Request().Context()
		respServiceSection = []response.ServiceSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllServiceSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.serviceSectionService.FetchAllServiceSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respServiceSection = append(respServiceSection, response.ServiceSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			Tagline:  val.Tagline,
			PathIcon: val.PathIcon,
		})
	}

	resp.Meta.Message = "Success fetch all service section"
	resp.Meta.Status = true
	resp.Data = respServiceSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDServiceSection implements ServiceSectionHandlerInterface.
func (cs *serviceSectionHandler) FetchByIDServiceSection(c echo.Context) error {
	var (
		resp               = response.DefaultSuccessResponse{}
		respError          = response.ErrorResponseDefault{}
		ctx                = c.Request().Context()
		respServiceSection = response.ServiceSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDServiceSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceSection := c.Param("id")
	id, err := conv.StringToInt64(idServiceSection)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.serviceSectionService.FetchByIDServiceSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respServiceSection.ID = result.ID
	respServiceSection.Name = result.Name
	respServiceSection.Tagline = result.Tagline
	respServiceSection.PathIcon = result.PathIcon
	resp.Meta.Message = "Success fetch service section by ID"
	resp.Meta.Status = true
	resp.Data = respServiceSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDServiceSection implements ServiceSectionHandlerInterface.
func (cs *serviceSectionHandler) EditByIDServiceSection(c echo.Context) error {
	var (
		req       = request.ServiceSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDServiceSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceSection := c.Param("id")
	id, err := conv.StringToInt64(idServiceSection)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDServiceSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceSectionEntity{
		ID:       id,
		PathIcon: req.PathIcon,
		Name:     req.Name,
		Tagline:  req.Tagline,
	}

	err = cs.serviceSectionService.EditByIDServiceSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDServiceSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDServiceSection implements ServiceSectionHandlerInterface.
func (cs *serviceSectionHandler) DeleteByIDServiceSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDServiceSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceSection := c.Param("id")
	id, err := conv.StringToInt64(idServiceSection)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDServiceSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.serviceSectionService.DeleteByIDServiceSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDServiceSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete service section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// FetchAllServiceHome implements ServiceSectionHandlerInterface.
func (cs *serviceSectionHandler) FetchAllServiceHome(c echo.Context) error {
	var (
		respServices = []response.ServiceSectionResponse{}
		resp         = response.DefaultSuccessResponse{}
		respError    = response.ErrorResponseDefault{}
		ctx          = c.Request().Context()
	)

	results, err := cs.serviceSectionService.FetchAllServiceSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllServiceHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respServices = append(respServices, response.ServiceSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			Tagline:  val.Tagline,
			PathIcon: val.PathIcon,
		})
	}
	resp.Data = respServices
	resp.Meta.Message = "Success fetch all service home"
	resp.Meta.Status = true
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewServiceSectionHandler(e *echo.Echo, serviceSectionService service.ServiceSectionServiceInterface, cfg *config.Config) ServiceSectionHandlerInterface {
	h := &serviceSectionHandler{
		serviceSectionService: serviceSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	serviceSectionApp := e.Group("/service-sections")
	serviceSectionApp.GET("", h.FetchAllServiceHome)

	adminApp := serviceSectionApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateServiceSection)
	adminApp.GET("", h.FetchAllServiceSection)
	adminApp.GET("/:id", h.FetchByIDServiceSection)
	adminApp.PUT("/:id", h.EditByIDServiceSection)
	adminApp.DELETE("/:id", h.DeleteByIDServiceSection)

	return h
}
