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

type AboutCompanyHandlerInterface interface {
	CreateAboutCompany(c echo.Context) error
	FetchAllAboutCompany(c echo.Context) error
	FetchByIDAboutCompany(c echo.Context) error
	EditByIDAboutCompany(c echo.Context) error
	DeleteByIDAboutCompany(c echo.Context) error
	FetchAllCompanyHome(c echo.Context) error
}

type aboutCompanyHandler struct {
	aboutCompanyService service.AboutCompanyServiceInterface
}

// FetchAllCompanyHome implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) FetchAllCompanyHome(c echo.Context) error {
	var (
		respCompany = response.AboutCompanyResponse{}
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
	)

	result, err := cs.aboutCompanyService.FetchAllCompanyAndKeynote(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllCompanyHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respCompany.ID = result.ID
	respCompany.Description = result.Description
	for _, val := range result.Keynote {
		respCompany.CompanyKeynotes = append(respCompany.CompanyKeynotes, response.AboutCompanyKeynoteResponse{
			ID:             val.ID,
			AboutCompanyID: val.AboutCompanyID,
			Keynote:        val.Keynote,
			PathImage:      val.PathImage,
		})
	}
	resp.Meta.Message = "Success fetch all company home"
	resp.Meta.Status = true
	resp.Data = respCompany
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// CreateAboutCompany implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) CreateAboutCompany(c echo.Context) error {
	var (
		req       = request.AboutCompanyRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateAboutCompany - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyEntity{
		Description: req.Description,
	}

	err = cs.aboutCompanyService.CreateAboutCompany(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAboutCompany - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create about company"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDAboutCompany implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) DeleteByIDAboutCompany(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDAboutCompany - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompany := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompany)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.aboutCompanyService.DeleteByIDAboutCompany(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete client section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDAboutCompany implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) EditByIDAboutCompany(c echo.Context) error {
	var (
		req       = request.AboutCompanyRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDAboutCompany - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompany := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompany)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompany - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyEntity{
		ID:          id,
		Description: req.Description,
	}

	err = cs.aboutCompanyService.EditByIDAboutCompany(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompany - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit about company"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllAboutCompany implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) FetchAllAboutCompany(c echo.Context) error {
	var (
		resp             = response.DefaultSuccessResponse{}
		respError        = response.ErrorResponseDefault{}
		ctx              = c.Request().Context()
		respAboutCompany = []response.AboutCompanyResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllAboutCompany - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.aboutCompanyService.FetchAllAboutCompany(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respAboutCompany = append(respAboutCompany, response.AboutCompanyResponse{
			ID:          val.ID,
			Description: val.Description,
		})
	}

	resp.Meta.Message = "Success fetch all client section"
	resp.Meta.Status = true
	resp.Data = respAboutCompany
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAboutCompany implements AboutCompanyHandlerInterface.
func (cs *aboutCompanyHandler) FetchByIDAboutCompany(c echo.Context) error {
	var (
		resp             = response.DefaultSuccessResponse{}
		respError        = response.ErrorResponseDefault{}
		ctx              = c.Request().Context()
		respAboutCompany = response.AboutCompanyResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDAboutCompany - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompany := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompany)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAboutCompany - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.aboutCompanyService.FetchByIDAboutCompany(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAboutCompany - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAboutCompany.ID = result.ID
	respAboutCompany.Description = result.Description
	resp.Meta.Message = "Success fetch about company by ID"
	resp.Meta.Status = true
	resp.Data = respAboutCompany
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewAboutCompanyHandler(e *echo.Echo, aboutCompanyService service.AboutCompanyServiceInterface, cfg *config.Config) AboutCompanyHandlerInterface {
	h := &aboutCompanyHandler{
		aboutCompanyService: aboutCompanyService,
	}

	mid := middleware.NewMiddleware(cfg)

	aboutCompanyApp := e.Group("/about-company")
	aboutCompanyApp.GET("", h.FetchAllCompanyHome)

	adminApp := aboutCompanyApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateAboutCompany)
	adminApp.GET("", h.FetchAllAboutCompany)
	adminApp.GET("/:id", h.FetchByIDAboutCompany)
	adminApp.PUT("/:id", h.EditByIDAboutCompany)
	adminApp.DELETE("/:id", h.DeleteByIDAboutCompany)

	return h
}
