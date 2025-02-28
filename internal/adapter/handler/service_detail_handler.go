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

type ServiceDetailHandlerInterface interface {
	CreateServiceDetail(c echo.Context) error
	FetchAllServiceDetail(c echo.Context) error
	FetchByIDServiceDetail(c echo.Context) error
	EditByIDServiceDetail(c echo.Context) error
	DeleteByIDServiceDetail(c echo.Context) error

	FetchServiceDetailByServiceID(c echo.Context) error
}

type serviceDetailHandler struct {
	serviceDetailService service.ServiceDetailServiceInterface
}

// CreateServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) CreateServiceDetail(c echo.Context) error {
	var (
		req       = request.ServiceDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateServiceDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceDetailEntity{
		ServiceID:   req.ServiceID,
		PathImage:   req.PathImage,
		Title:       req.Title,
		Description: req.Description,
		PathPdf:     req.PathPdf,
		PathDocx:    req.PathDocx,
	}

	err = cs.serviceDetailService.CreateServiceDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateServiceDetail - 4: %v", err)
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

// FetchAllServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) FetchAllServiceDetail(c echo.Context) error {
	var (
		resp              = response.DefaultSuccessResponse{}
		respError         = response.ErrorResponseDefault{}
		ctx               = c.Request().Context()
		respServiceDetail = []response.ServiceDetailResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllServiceDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.serviceDetailService.FetchAllServiceDetail(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respServiceDetail = append(respServiceDetail, response.ServiceDetailResponse{
			ID:          val.ID,
			ServiceID:   val.ServiceID,
			PathImage:   val.PathImage,
			Title:       val.Title,
			Description: val.Description,
			PathPdf:     val.PathPdf,
			PathDocx:    val.PathDocx,
			ServiceName: val.ServiceName,
		})
	}

	resp.Meta.Message = "Success fetch all service detail"
	resp.Meta.Status = true
	resp.Data = respServiceDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) FetchByIDServiceDetail(c echo.Context) error {
	var (
		resp              = response.DefaultSuccessResponse{}
		respError         = response.ErrorResponseDefault{}
		ctx               = c.Request().Context()
		respServiceDetail = response.ServiceDetailResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDServiceDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceDetail := c.Param("id")
	id, err := conv.StringToInt64(idServiceDetail)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.serviceDetailService.FetchByIDServiceDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respServiceDetail.ID = result.ID
	respServiceDetail.ServiceID = result.ServiceID
	respServiceDetail.PathImage = result.PathImage
	respServiceDetail.Title = result.Title
	respServiceDetail.Description = result.Description
	respServiceDetail.PathPdf = result.PathPdf
	respServiceDetail.PathDocx = result.PathDocx
	respServiceDetail.ServiceName = result.ServiceName
	resp.Meta.Message = "Success fetch service section by ID"
	resp.Meta.Status = true
	resp.Data = respServiceDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) EditByIDServiceDetail(c echo.Context) error {
	var (
		req       = request.ServiceDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDServiceDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceDetail := c.Param("id")
	id, err := conv.StringToInt64(idServiceDetail)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDServiceDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDServiceDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ServiceDetailEntity{
		ID:          id,
		ServiceID:   req.ServiceID,
		PathImage:   req.PathImage,
		Title:       req.Title,
		Description: req.Description,
		PathPdf:     req.PathPdf,
		PathDocx:    req.PathDocx,
	}

	err = cs.serviceDetailService.EditByIDServiceDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDServiceDetail - 5: %v", err)
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

// DeleteByIDServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) DeleteByIDServiceDetail(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDServiceDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceDetail := c.Param("id")
	id, err := conv.StringToInt64(idServiceDetail)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDServiceDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.serviceDetailService.DeleteByIDServiceDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDServiceDetail - 3: %v", err)
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

// FetchServiceDetail implements ServiceDetailHandlerInterface.
func (cs *serviceDetailHandler) FetchServiceDetailByServiceID(c echo.Context) error {
	var (
		resp              = response.DefaultSuccessResponse{}
		respError         = response.ErrorResponseDefault{}
		ctx               = c.Request().Context()
		respServiceDetail = response.ServiceDetailResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchServiceDetailByServiceID - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idServiceID := c.Param("id")
	id, err := conv.StringToInt64(idServiceID)
	if err != nil {
		log.Errorf("[HANDLER] FetchServiceDetailByServiceID - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.serviceDetailService.GetByServiceIDDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchServiceDetailByServiceID - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respServiceDetail.ID = result.ID
	respServiceDetail.ServiceID = result.ServiceID
	respServiceDetail.PathImage = result.PathImage
	respServiceDetail.Title = result.Title
	respServiceDetail.Description = result.Description
	respServiceDetail.PathPdf = result.PathPdf
	respServiceDetail.PathDocx = result.PathDocx
	respServiceDetail.ServiceName = result.ServiceName
	resp.Meta.Message = "Success fetch service section by ID"
	resp.Meta.Status = true
	resp.Data = respServiceDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}
func NewServiceDetailHandler(e *echo.Echo, serviceDetailService service.ServiceDetailServiceInterface, cfg *config.Config) ServiceDetailHandlerInterface {
	h := &serviceDetailHandler{
		serviceDetailService: serviceDetailService,
	}

	mid := middleware.NewMiddleware(cfg)

	serviceDetailApp := e.Group("/service-details")
	serviceDetailApp.GET("", h.FetchServiceDetailByServiceID)

	adminApp := serviceDetailApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateServiceDetail)
	adminApp.GET("", h.FetchAllServiceDetail)
	adminApp.GET("/:id", h.FetchByIDServiceDetail)
	adminApp.PUT("/:id", h.EditByIDServiceDetail)
	adminApp.DELETE("/:id", h.DeleteByIDServiceDetail)

	return h
}
