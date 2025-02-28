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

type PortofolioTestimonialHandlerInterface interface {
	CreatePortofolioTestimonial(c echo.Context) error
	FetchAllPortofolioTestimonial(c echo.Context) error
	FetchByIDPortofolioTestimonial(c echo.Context) error
	EditByIDPortofolioTestimonial(c echo.Context) error
	DeleteByIDPortofolioTestimonial(c echo.Context) error

	FetchAllPortofolioTestimonialHome(c echo.Context) error
}
type portofolioTestimonialHandler struct {
	portofolioTestimonialService service.PortofolioTestimonialServiceInterface
}

// CreatePortofolioTestimonial implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) CreatePortofolioTestimonial(c echo.Context) error {
	var (
		req       = request.PortofolioTestimonialRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioTestimonialEntity{
		Thumbnail:         req.Thumbnail,
		Message:           req.Message,
		ClientName:        req.ClientName,
		Role:              req.Role,
		PortofolioSection: entity.PortofolioSectionEntity{ID: req.PortofolioSectionID},
	}

	err = cs.portofolioTestimonialService.CreatePortofolioTestimonial(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioTestimonial - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllPortofolioTestimonial implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) FetchAllPortofolioTestimonial(c echo.Context) error {
	var (
		resp                      = response.DefaultSuccessResponse{}
		respError                 = response.ErrorResponseDefault{}
		ctx                       = c.Request().Context()
		respPortofolioTestimonial = []response.PortofolioTestimonialResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllPortofolioTestimonial - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.portofolioTestimonialService.FetchAllPortofolioTestimonial(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respPortofolioTestimonial = append(respPortofolioTestimonial, response.PortofolioTestimonialResponse{
			ID:                val.ID,
			Thumbnail:         val.Thumbnail,
			Message:           val.Message,
			ClientName:        val.ClientName,
			Role:              val.Role,
			PortofolioSection: response.PortofolioSectionResponse{Name: val.PortofolioSection.Name},
		})
	}

	resp.Meta.Message = "Success fetch all portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = respPortofolioTestimonial
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioTestimonial implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) FetchByIDPortofolioTestimonial(c echo.Context) error {
	var (
		resp                      = response.DefaultSuccessResponse{}
		respError                 = response.ErrorResponseDefault{}
		ctx                       = c.Request().Context()
		respPortofolioTestimonial = response.PortofolioTestimonialResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDPortofolioTestimonial - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.portofolioTestimonialService.FetchByIDPortofolioTestimonial(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioTestimonial.ID = result.ID
	respPortofolioTestimonial.Thumbnail = result.Thumbnail
	respPortofolioTestimonial.Message = result.Message
	respPortofolioTestimonial.ClientName = result.ClientName
	respPortofolioTestimonial.Role = result.Role
	respPortofolioTestimonial.PortofolioSection.ID = result.PortofolioSection.ID
	respPortofolioTestimonial.PortofolioSection.Name = result.PortofolioSection.Name
	respPortofolioTestimonial.PortofolioSection.Thumbnail = result.PortofolioSection.Thumbnail

	resp.Meta.Message = "Success fetch portofolio testimonial by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioTestimonial
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDPortofolioTestimonial implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) EditByIDPortofolioTestimonial(c echo.Context) error {
	var (
		req       = request.PortofolioTestimonialRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDPortofolioTestimonial - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioTestimonial - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioTestimonialEntity{
		ID:                id,
		Thumbnail:         req.Thumbnail,
		Message:           req.Message,
		ClientName:        req.ClientName,
		Role:              req.Role,
		PortofolioSection: entity.PortofolioSectionEntity{ID: req.PortofolioSectionID},
	}

	err = cs.portofolioTestimonialService.EditByIDPortofolioTestimonial(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioTestimonial - 6: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDPortofolioTestimonial implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) DeleteByIDPortofolioTestimonial(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDPortofolioTestimonial - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioTestimonial := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioTestimonial)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioTestimonial - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.portofolioTestimonialService.DeleteByIDPortofolioTestimonial(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioTestimonial - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete portofolio testimonial"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// FetchAllPortofolioTestimonialHome implements PortofolioTestimonialHandlerInterface.
func (cs *portofolioTestimonialHandler) FetchAllPortofolioTestimonialHome(c echo.Context) error {
	var (
		respTestimonials = []response.PortofolioTestimonialResponse{}
		resp             = response.DefaultSuccessResponse{}
		respError        = response.ErrorResponseDefault{}
		ctx              = c.Request().Context()
	)

	results, err := cs.portofolioTestimonialService.FetchAllPortofolioTestimonial(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPortofolioTestimonialHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	for _, val := range results {
		respTestimonials = append(respTestimonials, response.PortofolioTestimonialResponse{
			ID:         val.ID,
			Thumbnail:  val.Thumbnail,
			Message:    val.Message,
			ClientName: val.ClientName,
			Role:       val.Role,
			PortofolioSection: response.PortofolioSectionResponse{
				Name: val.PortofolioSection.Name,
			},
		})
	}

	resp.Meta.Message = "Success fetch all portofolio testimonial home"
	resp.Meta.Status = true
	resp.Data = respTestimonials
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}
func NewPortofolioTestimonialHandler(e *echo.Echo, portofolioTestimonialService service.PortofolioTestimonialServiceInterface, cfg *config.Config) PortofolioTestimonialHandlerInterface {
	h := &portofolioTestimonialHandler{
		portofolioTestimonialService: portofolioTestimonialService,
	}

	mid := middleware.NewMiddleware(cfg)

	portofolioTestimonialApp := e.Group("/portofolio-testimonials")
	portofolioTestimonialApp.GET("", h.FetchAllPortofolioTestimonialHome)

	adminApp := portofolioTestimonialApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreatePortofolioTestimonial)
	adminApp.GET("", h.FetchAllPortofolioTestimonial)
	adminApp.GET("/:id", h.FetchByIDPortofolioTestimonial)
	adminApp.PUT("/:id", h.EditByIDPortofolioTestimonial)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioTestimonial)

	return h
}
