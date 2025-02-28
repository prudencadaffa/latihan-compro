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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PortofolioDetailHandlerInterface interface {
	CreatePortofolioDetail(c echo.Context) error
	FetchAllPortofolioDetail(c echo.Context) error
	FetchByIDPortofolioDetail(c echo.Context) error
	EditByIDPortofolioDetail(c echo.Context) error
	DeleteByIDPortofolioDetail(c echo.Context) error

	FetchDetailPotofolioByPortoID(c echo.Context) error
}
type portofolioDetailHandler struct {
	portofolioDetailService service.PortofolioDetailServiceInterface
}

// CreatePortofolioDetail implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) CreatePortofolioDetail(c echo.Context) error {
	var (
		req       = request.PortofolioDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringProjectDate, err := time.Parse("2006-01-02", req.ProjectDate)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}
	reqEntity := entity.PortofolioDetailEntity{
		Category:    req.Category,
		ClientName:  req.ClientName,
		ProjectDate: stringProjectDate,
		ProjectUrl:  req.ProjectUrl,
		Title:       req.Title,
		Description: req.Description,
		PortofolioSection: entity.PortofolioSectionEntity{
			ID: req.PortofolioSectionID,
		},
	}

	err = cs.portofolioDetailService.CreatePortofolioDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreatePortofolioDetail - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create portofolio detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllPortofolioDetail implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) FetchAllPortofolioDetail(c echo.Context) error {
	var (
		resp                 = response.DefaultSuccessResponse{}
		respError            = response.ErrorResponseDefault{}
		ctx                  = c.Request().Context()
		respPortofolioDetail = []response.PortofolioDetailResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllPortofolioDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.portofolioDetailService.FetchAllPortofolioDetail(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respPortofolioDetail = append(respPortofolioDetail, response.PortofolioDetailResponse{
			ID:          val.ID,
			Category:    val.Category,
			ClientName:  val.ClientName,
			ProjectDate: val.ProjectDate.Format("02 January 2006"),
			ProjectUrl:  val.ProjectUrl,
			Title:       val.Title,
			Description: val.Description,
			PortofolioSection: response.PortofolioSectionResponse{
				ID:        val.PortofolioSection.ID,
				Name:      val.PortofolioSection.Name,
				Thumbnail: val.PortofolioSection.Thumbnail,
			},
		})
	}

	resp.Meta.Message = "Success fetch all portofolio detail"
	resp.Meta.Status = true
	resp.Data = respPortofolioDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDPortofolioDetail implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) FetchByIDPortofolioDetail(c echo.Context) error {
	var (
		resp                 = response.DefaultSuccessResponse{}
		respError            = response.ErrorResponseDefault{}
		ctx                  = c.Request().Context()
		respPortofolioDetail = response.PortofolioDetailResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDPortofolioDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.portofolioDetailService.FetchByIDPortofolioDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDPortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respPortofolioDetail.ID = result.ID
	respPortofolioDetail.Category = result.Category
	respPortofolioDetail.ClientName = result.ClientName
	respPortofolioDetail.ProjectDate = result.ProjectDate.Format("02 January 2006")
	respPortofolioDetail.ProjectUrl = result.ProjectUrl
	respPortofolioDetail.Title = result.Title
	respPortofolioDetail.Description = result.Description
	respPortofolioDetail.PortofolioSection.ID = result.PortofolioSection.ID
	respPortofolioDetail.PortofolioSection.Name = result.PortofolioSection.Name
	respPortofolioDetail.PortofolioSection.Thumbnail = result.PortofolioSection.Thumbnail

	resp.Meta.Message = "Success fetch portofolio detail by ID"
	resp.Meta.Status = true
	resp.Data = respPortofolioDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDPortofolioDetail implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) EditByIDPortofolioDetail(c echo.Context) error {
	var (
		req       = request.PortofolioDetailRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringProjectDate, err := time.Parse("2006-01-02", req.ProjectDate)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.PortofolioDetailEntity{
		ID:          id,
		Category:    req.Category,
		ClientName:  req.ClientName,
		ProjectDate: stringProjectDate,
		ProjectUrl:  req.ProjectUrl,
		Title:       req.Title,
		Description: req.Description,
		PortofolioSection: entity.PortofolioSectionEntity{
			ID: req.PortofolioSectionID,
		},
	}

	err = cs.portofolioDetailService.EditByIDPortofolioDetail(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDPortofolioDetail - 6: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit portofolio detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDPortofolioDetail implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) DeleteByIDPortofolioDetail(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDPortofolioDetail - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idPortofolioDetail := c.Param("id")
	id, err := conv.StringToInt64(idPortofolioDetail)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioDetail - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.portofolioDetailService.DeleteByIDPortofolioDetail(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDPortofolioDetail - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete portofolio detail"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// FetchDetailPotofolioByPortoID implements PortofolioDetailHandlerInterface.
func (cs *portofolioDetailHandler) FetchDetailPotofolioByPortoID(c echo.Context) error {
	var (
		respDetail = response.PortofolioDetailResponse{}
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
	)
	idPorto := c.Param("id")
	id, err := conv.StringToInt64(idPorto)
	if err != nil {
		log.Errorf("[HANDLER] FetchDetailPotofolioByPortoID - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.portofolioDetailService.FetchDetailPotofolioByPortoID(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchDetailPotofolioByPortoID - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	respDetail.ID = result.ID
	respDetail.Category = result.Category
	respDetail.ClientName = result.ClientName
	respDetail.ProjectDate = result.ProjectDate.Format("02 January 2006")
	respDetail.ProjectUrl = result.ProjectUrl
	respDetail.Title = result.Title
	respDetail.Description = result.Description
	respDetail.PortofolioSection.ID = result.PortofolioSection.ID
	respDetail.PortofolioSection.Name = result.PortofolioSection.Name
	respDetail.PortofolioSection.Thumbnail = result.PortofolioSection.Thumbnail
	resp.Meta.Message = "Success"
	resp.Meta.Status = true
	resp.Data = respDetail
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}
func NewPortofolioDetailHandler(e *echo.Echo, portofolioDetailService service.PortofolioDetailServiceInterface, cfg *config.Config) PortofolioDetailHandlerInterface {
	h := &portofolioDetailHandler{
		portofolioDetailService: portofolioDetailService,
	}

	mid := middleware.NewMiddleware(cfg)

	portofolioDetailApp := e.Group("/portofolio-details")

	portofolioDetailApp.GET("/:id", h.FetchDetailPotofolioByPortoID)

	adminApp := portofolioDetailApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreatePortofolioDetail)
	adminApp.GET("", h.FetchAllPortofolioDetail)
	adminApp.GET("/:id", h.FetchByIDPortofolioDetail)
	adminApp.PUT("/:id", h.EditByIDPortofolioDetail)
	adminApp.DELETE("/:id", h.DeleteByIDPortofolioDetail)

	return h
}
