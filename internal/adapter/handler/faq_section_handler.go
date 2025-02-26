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

type FaqSectionHandlerInterface interface {
	CreateFaqSection(c echo.Context) error
	FetchAllFaqSection(c echo.Context) error
	FetchByIDFaqSection(c echo.Context) error
	EditByIDFaqSection(c echo.Context) error
	DeleteByIDFaqSection(c echo.Context) error

	FetchAllFaqSectionHome(c echo.Context) error
}

type faqSectionHandler struct {
	faqSectionService service.FaqSectionServiceInterface
}

// FetchAllFaqSectionHome implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) FetchAllFaqSectionHome(c echo.Context) error {
	var (
		respFaqs  = []response.FaqSectionResponse{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	results, err := cs.faqSectionService.FetchAllFaqSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllFaqSectionHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respFaqs = append(respFaqs, response.FaqSectionResponse{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
		})
	}
	resp.Data = respFaqs
	resp.Meta.Message = "Success fetch all faq section home"
	resp.Meta.Status = true
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// CreateFaqSection implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) CreateFaqSection(c echo.Context) error {
	var (
		req       = request.FaqSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateFaqSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.FaqSectionEntity{
		Title:       req.Title,
		Description: req.Description,
	}

	err = cs.faqSectionService.CreateFaqSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateFaqSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create faq section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDFaqSection implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) DeleteByIDFaqSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDFaqSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idFaqSection := c.Param("id")
	id, err := conv.StringToInt64(idFaqSection)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.faqSectionService.DeleteByIDFaqSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete faq section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDFaqSection implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) EditByIDFaqSection(c echo.Context) error {
	var (
		req       = request.FaqSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDFaqSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idFaqSection := c.Param("id")
	id, err := conv.StringToInt64(idFaqSection)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDFaqSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.FaqSectionEntity{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
	}

	err = cs.faqSectionService.EditByIDFaqSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDFaqSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit faq section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllFaqSection implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) FetchAllFaqSection(c echo.Context) error {
	var (
		resp           = response.DefaultSuccessResponse{}
		respError      = response.ErrorResponseDefault{}
		ctx            = c.Request().Context()
		respFaqSection = []response.FaqSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllFaqSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.faqSectionService.FetchAllFaqSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respFaqSection = append(respFaqSection, response.FaqSectionResponse{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
		})
	}

	resp.Meta.Message = "Success fetch all faq section"
	resp.Meta.Status = true
	resp.Data = respFaqSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDFaqSection implements FaqSectionHandlerInterface.
func (cs *faqSectionHandler) FetchByIDFaqSection(c echo.Context) error {
	var (
		resp           = response.DefaultSuccessResponse{}
		respError      = response.ErrorResponseDefault{}
		ctx            = c.Request().Context()
		respFaqSection = response.FaqSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDFaqSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idFaqSection := c.Param("id")
	id, err := conv.StringToInt64(idFaqSection)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDFaqSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.faqSectionService.FetchByIDFaqSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDFaqSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respFaqSection.ID = result.ID
	respFaqSection.Title = result.Title
	respFaqSection.Description = result.Description
	resp.Meta.Message = "Success fetch hero section by ID"
	resp.Meta.Status = true
	resp.Data = respFaqSection
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewFaqSectionHandler(e *echo.Echo, faqSectionService service.FaqSectionServiceInterface, cfg *config.Config) FaqSectionHandlerInterface {
	h := &faqSectionHandler{
		faqSectionService: faqSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	faqApp := e.Group("/faq-sections")
	faqApp.GET("", h.FetchAllFaqSectionHome)

	adminApp := faqApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateFaqSection)
	adminApp.GET("", h.FetchAllFaqSection)
	adminApp.GET("/:id", h.FetchByIDFaqSection)
	adminApp.PUT("/:id", h.EditByIDFaqSection)
	adminApp.DELETE("/:id", h.DeleteByIDFaqSection)

	return h
}
