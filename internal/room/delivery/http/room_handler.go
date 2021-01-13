package http

import (
	"errors"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/room/usecase"
	"github.com/babon21/hotel-management/internal/utils"
	"github.com/babon21/hotel-management/pkg/delivery/http/api"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// RoomHandler  represent the httphandler for room
type RoomHandler struct {
	RoomUsecase usecase.RoomUsecase
}

// NewRoomHandler will initialize the rooms/ resources endpoint
func NewRoomHandler(e *echo.Echo, us usecase.RoomUsecase) {
	handler := &RoomHandler{
		RoomUsecase: us,
	}
	e.GET("/rooms", handler.GetRoomList)
	e.POST("/rooms", handler.Add)
	e.DELETE("/rooms/:id", handler.Delete)
}

// GetRoomList will fetch the room based on given params
func (a *RoomHandler) GetRoomList(c echo.Context) error {
	sortParam := c.QueryParam("sort_by")
	orderParam := c.QueryParam("order_by")

	sortField := formSortField(sortParam)
	if sortField == usecase.ErrField {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: "sort_by field must be price or date_added"}, "  ")
	}

	sortOrder := formSortOrder(orderParam)
	if sortOrder == usecase.ErrOrder {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: "order_by field must be asc or desc"}, "  ")
	}

	rooms, err := a.RoomUsecase.GetList(sortField, sortOrder)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	return c.JSONPretty(http.StatusOK, rooms, "  ")
}

func formSortOrder(orderValue string) usecase.SortOrder {
	switch orderValue {
	case string(usecase.AscOrder):
		return usecase.AscOrder
	case string(usecase.DescOrder):
		return usecase.DescOrder
	default:
		return usecase.ErrOrder
	}
}

func formSortField(sortValue string) usecase.SortField {
	switch sortValue {
	case string(usecase.PriceField):
		return usecase.PriceField
	case string(usecase.DateAddedField):
		return usecase.DateAddedField
	default:
		return usecase.ErrField
	}
}

// Add will store the room by given request body
func (a *RoomHandler) Add(c echo.Context) (err error) {
	var request api.AddRoomRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	if err := isRequestValid(&request); err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	room := domain.Room{
		Price:       request.Price,
		Description: request.Description,
	}

	id, err := a.RoomUsecase.Add(&room)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	response := api.AddRoomResponse{Id: id}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func isRequestValid(request *api.AddRoomRequest) error {
	if !utils.IsNumeric(request.Price) {
		return errors.New("price is not number")
	}
	return nil
}

// Delete will delete room by given param
func (a *RoomHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	err := a.RoomUsecase.Delete(id)
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
