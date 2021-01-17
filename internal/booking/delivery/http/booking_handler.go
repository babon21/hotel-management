package http

import (
	"errors"
	"github.com/babon21/hotel-management/internal/booking/usecase"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/pkg/delivery/http/api"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// BookingHandler  represent the httphandler for booking
type BookingHandler struct {
	BookingUsecase usecase.BookingUsecase
}

// NewBookingHandler will initialize the bookings/ resources endpoint
func NewBookingHandler(e *echo.Echo, us usecase.BookingUsecase) {
	handler := &BookingHandler{
		BookingUsecase: us,
	}
	e.GET("/bookings", handler.GetBookingList)
	e.POST("/bookings", handler.Add)
	e.DELETE("/bookings/:id", handler.Delete)
}

// GetBookingList will fetch the booking based on given params
func (a *BookingHandler) GetBookingList(c echo.Context) error {
	roomId := c.QueryParam("room_id")
	if roomId == "" {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: errors.New("room_id param not found").Error()}, "  ")
	}

	bookings, err := a.BookingUsecase.GetList(roomId)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	return c.JSONPretty(http.StatusOK, bookings, "  ")
}

// Add will store the room by given request body
func (a *BookingHandler) Add(c echo.Context) (err error) {
	var request api.AddBookingRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	if err = isRequestValid(&request); err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	booking := domain.Booking{
		RoomId:         request.RoomId,
		StartDate:      request.StartDate,
		ExpirationDate: request.ExpirationDate,
	}

	if err := a.BookingUsecase.Add(&booking); err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	response := api.AddRoomResponse{Id: booking.ID}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func isRequestValid(request *api.AddBookingRequest) error {
	dateLayout := "2006-01-02"

	_, err := time.Parse(dateLayout, request.StartDate)
	if err != nil {
		return errors.New("start_date field isn't valid, must be in year-month-day format")
	}

	_, err = time.Parse(dateLayout, request.ExpirationDate)
	if err != nil {
		return errors.New("expiration_date field isn't valid, must be in year-month-day format")
	}

	return nil
}

// Delete will delete room by given param
func (a *BookingHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	err := a.BookingUsecase.Delete(id)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error().Msg(err.Error())
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
