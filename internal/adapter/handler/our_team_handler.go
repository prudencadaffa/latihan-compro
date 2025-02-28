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

type OurTeamHandlerInterface interface {
	CreateOurTeam(c echo.Context) error
	FetchAllOurTeam(c echo.Context) error
	FetchByIDOurTeam(c echo.Context) error
	EditByIDOurTeam(c echo.Context) error
	DeleteByIDOurTeam(c echo.Context) error

	FetchAllOurTeamHome(c echo.Context) error
}

type ourTeamHandler struct {
	ourTeamService service.OurTeamServiceInterface
}

// FetchAllOurTeamHome implements OurTeamHandlerInterface.
func (h *ourTeamHandler) FetchAllOurTeamHome(c echo.Context) error {
	var (
		respOurTeams = []response.OurTeamResponse{}
		resp         = response.DefaultSuccessResponse{}
		respError    = response.ErrorResponseDefault{}
		ctx          = c.Request().Context()
	)

	results, err := h.ourTeamService.FetchAllOurTeam(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllOurTeamHome - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respOurTeams = append(respOurTeams, response.OurTeamResponse{
			ID:        val.ID,
			Name:      val.Name,
			Role:      val.Role,
			PathPhoto: val.PathPhoto,
			Tagline:   val.Tagline,
		})
	}

	resp.Meta.Message = "Success fetch all our team home"
	resp.Meta.Status = true
	resp.Data = respOurTeams
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// CreateOurTeam implements OurTeamHandlerInterface.
func (h *ourTeamHandler) CreateOurTeam(c echo.Context) error {
	var (
		req       = request.OurTeamRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateOurTeam - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.OurTeamEntity{
		Name:      req.Name,
		Role:      req.Role,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}

	err = h.ourTeamService.CreateOurTeam(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateOurTeam - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create our team"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDOurTeam implements OurTeamHandlerInterface.
func (h *ourTeamHandler) DeleteByIDOurTeam(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDOurTeam - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idOurTeam := c.Param("id")
	id, err := conv.StringToInt64(idOurTeam)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = h.ourTeamService.DeleteByIDOurTeam(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete our team"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}

// EditByIDOurTeam implements OurTeamHandlerInterface.
func (h *ourTeamHandler) EditByIDOurTeam(c echo.Context) error {
	var (
		req       = request.OurTeamRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] EditByIDOurTeam - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idOurTeam := c.Param("id")
	id, err := conv.StringToInt64(idOurTeam)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] EditByIDOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] EditByIDOurTeam - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.OurTeamEntity{
		ID:        id,
		Name:      req.Name,
		Role:      req.Role,
		PathPhoto: req.PathPhoto,
		Tagline:   req.Tagline,
	}

	err = h.ourTeamService.EditByIDOurTeam(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] EditByIDOurTeam - 5: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success edit our team"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchAllOurTeam implements OurTeamHandlerInterface.
func (h *ourTeamHandler) FetchAllOurTeam(c echo.Context) error {
	var (
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
		respOurTeam = []response.OurTeamResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllOurTeam - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := h.ourTeamService.FetchAllOurTeam(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respOurTeam = append(respOurTeam, response.OurTeamResponse{
			ID:        val.ID,
			Name:      val.Name,
			Role:      val.Role,
			PathPhoto: val.PathPhoto,
			Tagline:   val.Tagline,
		})
	}

	resp.Meta.Message = "Success fetch all our team"
	resp.Meta.Status = true
	resp.Data = respOurTeam
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDOurTeam implements OurTeamHandlerInterface.
func (h *ourTeamHandler) FetchByIDOurTeam(c echo.Context) error {
	var (
		resp        = response.DefaultSuccessResponse{}
		respError   = response.ErrorResponseDefault{}
		ctx         = c.Request().Context()
		respOurTeam = response.OurTeamResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDOurTeam - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idOurTeam := c.Param("id")
	id, err := conv.StringToInt64(idOurTeam)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDOurTeam - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := h.ourTeamService.FetchByIDOurTeam(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDOurTeam - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respOurTeam.ID = result.ID
	respOurTeam.Name = result.Name
	respOurTeam.Role = result.Role
	respOurTeam.PathPhoto = result.PathPhoto
	respOurTeam.Tagline = result.Tagline
	resp.Meta.Message = "Success fetch our team by ID"
	resp.Meta.Status = true
	resp.Data = respOurTeam
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

func NewOurTeamHandler(c *echo.Echo, cfg *config.Config, ourTeamService service.OurTeamServiceInterface) OurTeamHandlerInterface {
	heroHandler := &ourTeamHandler{
		ourTeamService: ourTeamService,
	}

	mid := middleware.NewMiddleware(cfg)

	ourTeamApp := c.Group("/our-teams")
	ourTeamApp.GET("", heroHandler.FetchAllOurTeamHome)

	adminApp := ourTeamApp.Group("/admin", mid.CheckToken())
	adminApp.GET("", heroHandler.FetchAllOurTeam)
	adminApp.POST("", heroHandler.CreateOurTeam)
	adminApp.GET("/:id", heroHandler.FetchByIDOurTeam)
	adminApp.PUT("/:id", heroHandler.EditByIDOurTeam)
	adminApp.DELETE("/:id", heroHandler.DeleteByIDOurTeam)

	return heroHandler
}
