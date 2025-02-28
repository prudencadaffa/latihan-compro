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

type AppointmentHandlerInterface interface {
	CreateAppointment(c echo.Context) error
	FetchAllAppointment(c echo.Context) error
	FetchByIDAppointment(c echo.Context) error
	DeleteByIDAppointment(c echo.Context) error
}
type appointmentHandler struct {
	appointmentService service.AppointmentServiceInterface
}

// CreateAppointment implements AppointmentHandlerInterface.
func (cs *appointmentHandler) CreateAppointment(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		req       = request.AppointmentRequest{}
		ctx       = c.Request().Context()
	)

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 1: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	stringProjectDate, err := time.Parse("2006-01-02", req.MeetAt)
	if err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.AppointmentEntity{
		ServiceID:   req.ServiceID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Brief:       req.Brief,
		Budget:      req.Budget,
		MeetAt:      stringProjectDate,
	}

	err = cs.appointmentService.CreateAppointment(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateAppointment - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create appointment"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// FetchAllAppointment implements AppointmentHandlerInterface.
func (cs *appointmentHandler) FetchAllAppointment(c echo.Context) error {
	var (
		resp            = response.DefaultSuccessResponse{}
		respError       = response.ErrorResponseDefault{}
		ctx             = c.Request().Context()
		respAppointment = []response.AppointmentResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchAllAppointment - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	results, err := cs.appointmentService.FetchAllAppointment(ctx)
	if err != nil {
		log.Errorf("[HANDLER] FetchAllAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	for _, val := range results {
		respAppointment = append(respAppointment, response.AppointmentResponse{
			ID:          val.ID,
			Name:        val.Name,
			PhoneNumber: val.PhoneNumber,
			Email:       val.Email,
			Brief:       val.Brief,
			Budget:      val.Budget,
			MeetAt:      val.MeetAt.Format("02 Jan 2006 15:04:05"),
			ServiceName: val.ServiceName,
			ServiceID:   val.ServiceID,
		})
	}

	resp.Meta.Message = "Success fetch all appointment"
	resp.Meta.Status = true
	resp.Data = respAppointment
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// FetchByIDAppointment implements AppointmentHandlerInterface.
func (cs *appointmentHandler) FetchByIDAppointment(c echo.Context) error {
	var (
		resp            = response.DefaultSuccessResponse{}
		respError       = response.ErrorResponseDefault{}
		ctx             = c.Request().Context()
		respAppointment = response.AppointmentResponse{}
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] FetchByIDAppointment - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAppointment := c.Param("id")
	id, err := conv.StringToInt64(idAppointment)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	result, err := cs.appointmentService.FetchByIDAppointment(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] FetchByIDAppointment - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	respAppointment.ID = result.ID
	respAppointment.Name = result.Name
	respAppointment.PhoneNumber = result.PhoneNumber
	respAppointment.Email = result.Email
	respAppointment.Brief = result.Brief
	respAppointment.Budget = result.Budget
	respAppointment.MeetAt = result.MeetAt.Format("02 Jan 2006 15:04:05")
	respAppointment.ServiceName = result.ServiceName
	respAppointment.ServiceID = result.ServiceID
	resp.Meta.Message = "Success fetch appointment by ID"
	resp.Meta.Status = true
	resp.Data = respAppointment
	resp.Pagination = nil
	return c.JSON(http.StatusOK, resp)
}

// DeleteByIDAppointment implements AppointmentHandlerInterface.
func (cs *appointmentHandler) DeleteByIDAppointment(c echo.Context) error {
	var (
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] DeleteByIDAppointment - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	idAppointment := c.Param("id")
	id, err := conv.StringToInt64(idAppointment)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAppointment - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	err = cs.appointmentService.DeleteByIDAppointment(ctx, id)
	if err != nil {
		log.Errorf("[HANDLER] DeleteByIDAppointment - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}
	resp.Meta.Message = "Success delete appointment"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil

	return c.JSON(http.StatusOK, resp)
}
func NewAppointmentHandler(e *echo.Echo, appointmentService service.AppointmentServiceInterface, cfg *config.Config) AppointmentHandlerInterface {
	h := &appointmentHandler{
		appointmentService: appointmentService,
	}

	mid := middleware.NewMiddleware(cfg)

	appointmentApp := e.Group("/appointments")
	appointmentApp.POST("", h.CreateAppointment)

	adminApp := appointmentApp.Group("/admin", mid.CheckToken())

	adminApp.GET("", h.FetchAllAppointment)
	adminApp.GET("/:id", h.FetchByIDAppointment)
	adminApp.DELETE("/:id", h.DeleteByIDAppointment)

	return h
}
