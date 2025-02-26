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

type ClientSectionHandlerInterface interface {
	CreateClientSection(c echo.Context) error
	FetchAllClientSection(c echo.Context) error
	FetchByIDClientSection(c echo.Context) error
	EditByIDClientSection(c echo.Context) error
	DeleteByIDClientSection(c echo.Context) error

	FetchAllClientSectionHome(c echo.Context) error
}
type clientSectionHandler struct {
	clientSectionService service.ClientSectionServiceInterface
}

// CreateClientSection implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) CreateClientSection(c echo.Context) error {
	var (
		req       = request.ClientSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateClientSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateClientSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateClientSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ClientSectionEntity{
		Name:     req.Name,
		PathIcon: req.PathIcon,
	}

	err = cs.clientSectionService.CreateClientSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateClientSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create client section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllClientSection implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) FetchAllClientSection(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = []response.ClientSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllClientSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.clientSectionService.FetchAllClientSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllClientSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClient = append(respClient, response.ClientSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			PathIcon: val.PathIcon,
		})
	}

	resp.Meta.Message = "Success fetch all client section"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDClientSection implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) FetchByIDClientSection(c echo.Context) error {
	var (
		resp       = response.DefaultSuccessResponse{}
		respError  = response.ErrorResponseDefault{}
		ctx        = c.Request().Context()
		respClient = response.ClientSectionResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDClientSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idClient := c.Param("id")
	id, err := conv.StringToInt64(idClient)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDClientSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.clientSectionService.FetchByIDClientSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDClientSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respClient.ID = result.ID
	respClient.Name = result.Name
	respClient.PathIcon = result.PathIcon
	resp.Meta.Message = "Success fetch hero section by ID"
	resp.Meta.Status = true
	resp.Data = respClient
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// EditByIDClientSection implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) EditByIDClientSection(c echo.Context) error {
	var (
		req       = request.ClientSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDClientSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idClient := c.Param("id")
	id, err := conv.StringToInt64(idClient)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDClientSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDClientSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDClientSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.ClientSectionEntity{
		ID:       id,
		Name:     req.Name,
		PathIcon: req.PathIcon,
	}

	err = cs.clientSectionService.EditByIDClientSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDClientSection - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit client section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDClientSection implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) DeleteByIDClientSection(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDClientSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idClient := c.Param("id")
	id, err := conv.StringToInt64(idClient)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDClientSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.clientSectionService.DeleteByIDClientSection(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDClientSection - 3: %v", err)
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

// FetchAllClientSectionHome implements ClientSectionHandlerInterface.
func (cs *clientSectionHandler) FetchAllClientSectionHome(c echo.Context) error {
	var (
		respClients = []response.ClientSectionResponse{}
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
	)

	results, err := cs.clientSectionService.FetchAllClientSection(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllClientSectionHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respClients = append(respClients, response.ClientSectionResponse{
			ID:       val.ID,
			Name:     val.Name,
			PathIcon: val.PathIcon,
		})
	}

	resp.Meta.Message = "Success fetch all client section home"
	resp.Meta.Status = true
	resp.Data = respClients
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewClientSectionHandler(e *echo.Echo, clientSectionService service.ClientSectionServiceInterface, cfg *config.Config) ClientSectionHandlerInterface {
	h := &clientSectionHandler{
		clientSectionService: clientSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	clientApp := e.Group("/client-sections")
	clientApp.GET("", h.FetchAllClientSectionHome)

	adminApp := clientApp.Group("/admin", mid.CheckToken())

	adminApp.POST("", h.CreateClientSection)
	adminApp.GET("", h.FetchAllClientSection)
	adminApp.GET("/:id", h.FetchByIDClientSection)
	adminApp.PUT("/:id", h.EditByIDClientSection)
	adminApp.DELETE("/:id", h.DeleteByIDClientSection)

	return h
}
