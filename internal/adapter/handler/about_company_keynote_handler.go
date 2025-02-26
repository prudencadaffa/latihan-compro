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

type AboutCompanyKeynoteHandlerInterface interface {
	CreateAboutCompanyKeynote(c echo.Context) error
	FetchAllAboutCompanyKeynote(c echo.Context) error
	FetchByIDAboutCompanyKeynote(c echo.Context) error
	EditByIDAboutCompanyKeynote(c echo.Context) error
	DeleteByIDAboutCompanyKeynote(c echo.Context) error
	FetchByCompanyID(c echo.Context) error
}

type aboutCompanyKeynoteHandler struct {
	aboutCompanyKeynoteService service.AboutCompanyKeynoteServiceInterface
}

// FetchByCompanyID implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) FetchByCompanyID(c echo.Context) error {
	var (
		resp                    = response.DefaultSuccessResponse{}
		respError               = response.ErrorResponseDefault{}
		ctx                     = c.Request().Context()
		respAboutCompanyKeynote = []response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByCompanyID - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompany := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompany)
	if err != nil {
		log.Errorf("[HANDLER] FetchByCompanyID - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	results, err := cs.aboutCompanyKeynoteService.FetchByCompanyID(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByCompanyID - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respAboutCompanyKeynote = append(respAboutCompanyKeynote, response.AboutCompanyKeynoteResponse{
			ID:                      val.ID,
			AboutCompanyID:          val.AboutCompanyID,
			Keynote:                 val.Keynote,
			PathImage:               val.PathImage,
			AboutCompanyDescription: val.AboutCompanyDescription,
		})
	}

	resp.Meta.Message = "Success fetch by about company id"
	resp.Meta.Status = true
	resp.Data = respAboutCompanyKeynote
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// CreateAboutCompanyKeynote implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) CreateAboutCompanyKeynote(c echo.Context) error {
	var (
		req       = request.AboutCompanyKeynoteRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyKeynoteEntity{
		AboutCompanyID: req.AboutCompanyID,
		Keynote:        req.Keynote,
		PathImage:      req.PathImage,
	}

	err = cs.aboutCompanyKeynoteService.CreateAboutCompanyKeynote(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAboutCompanyKeynote - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDAboutCompanyKeynote implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) DeleteByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDAboutCompanyKeynote - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.aboutCompanyKeynoteService.DeleteByIDAboutCompanyKeynote(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDAboutCompanyKeynote implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) EditByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		req       = request.AboutCompanyKeynoteRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDAboutCompanyKeynote - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompanyKeynote - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AboutCompanyKeynoteEntity{
		ID:             id,
		AboutCompanyID: req.AboutCompanyID,
		Keynote:        req.Keynote,
		PathImage:      req.PathImage,
	}

	err = cs.aboutCompanyKeynoteService.EditByIDAboutCompanyKeynote(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDAboutCompanyKeynote - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit about company keynote"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllAboutCompanyKeynote implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) FetchAllAboutCompanyKeynote(c echo.Context) error {
	var (
		resp                    = response.DefaultSuccessResponse{}
		respError               = response.ErrorResponseDefault{}
		ctx                     = c.Request().Context()
		respAboutCompanyKeynote = []response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllAboutCompanyKeynote - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.aboutCompanyKeynoteService.FetchAllAboutCompanyKeynote(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respAboutCompanyKeynote = append(respAboutCompanyKeynote, response.AboutCompanyKeynoteResponse{
			ID:                      val.ID,
			AboutCompanyID:          val.AboutCompanyID,
			Keynote:                 val.Keynote,
			PathImage:               val.PathImage,
			AboutCompanyDescription: val.AboutCompanyDescription,
		})
	}

	resp.Meta.Message = "Success fetch all about company keynote"
	resp.Meta.Status = true
	resp.Data = respAboutCompanyKeynote
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAboutCompanyKeynote implements AboutCompanyKeynoteHandlerInterface.
func (cs *aboutCompanyKeynoteHandler) FetchByIDAboutCompanyKeynote(c echo.Context) error {
	var (
		resp                    = response.DefaultSuccessResponse{}
		respError               = response.ErrorResponseDefault{}
		ctx                     = c.Request().Context()
		respAboutCompanyKeynote = response.AboutCompanyKeynoteResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDAboutCompanyKeynote - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAboutCompanyKeynote := c.Param("id")
	id, err := conv.StringToInt64(idAboutCompanyKeynote)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAboutCompanyKeynote - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.aboutCompanyKeynoteService.FetchByIDAboutCompanyKeynote(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAboutCompanyKeynote - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAboutCompanyKeynote.ID = result.ID
	respAboutCompanyKeynote.AboutCompanyID = result.AboutCompanyID
	respAboutCompanyKeynote.Keynote = result.Keynote
	respAboutCompanyKeynote.PathImage = result.PathImage
	respAboutCompanyKeynote.AboutCompanyDescription = result.AboutCompanyDescription
	resp.Meta.Message = "Success fetch about company keynote by ID"
	resp.Meta.Status = true
	resp.Data = respAboutCompanyKeynote
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewAboutCompanyKeynoteHandler(e *echo.Echo, aboutCompanyKeynoteService service.AboutCompanyKeynoteServiceInterface, cfg *config.Config) AboutCompanyKeynoteHandlerInterface {
	h := &aboutCompanyKeynoteHandler{
		aboutCompanyKeynoteService: aboutCompanyKeynoteService,
	}

	mid := middleware.NewMiddleware(cfg)

	aboutCompanyKeynoteApp := e.Group("/about-company-keynotes")
	adminApp := aboutCompanyKeynoteApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateAboutCompanyKeynote)
	adminApp.GET("", h.FetchAllAboutCompanyKeynote)
	adminApp.GET("/:id", h.FetchByIDAboutCompanyKeynote)
	adminApp.PUT("/:id", h.EditByIDAboutCompanyKeynote)
	adminApp.DELETE("/:id", h.DeleteByIDAboutCompanyKeynote)

	keynoteApp := adminApp.Group("/keynotes")
	keynoteApp.GET("/:id", h.FetchByCompanyID)

	return h
}
