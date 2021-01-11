package http

import (
	"github.com/babon21/hotel-management/domain"
	"github.com/babon21/hotel-management/pkg/delivery/http/api"
	"github.com/babon21/hotel-management/room/usecase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// RoomHandler  represent the httphandler for room
type RoomHandler struct {
	RoomUsecase usecase.RoomUsecase
}

// NewRoomHandler will initialize the articles/ resources endpoint
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
		// TODO bad request
	}

	sortOrder := formSortOrder(orderParam)
	if sortOrder == usecase.ErrOrder {
		// TODO bad request
	}

	rooms, err := a.RoomUsecase.GetList(sortField, sortOrder)
	if err != nil {
		// TODO
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

//func isRequestValid(m *domain.Article) (bool, error) {
//	validate := validator.New()
//	err := validate.Struct(m)
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}

// Add will store the room by given request body
func (a *RoomHandler) Add(c echo.Context) (err error) {
	var request api.AddRoomRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	//var ok bool
	//if ok, err = isRequestValid(&request); !ok {
	//	return c.JSON(http.StatusBadRequest, err.Error())
	//}

	room := domain.Room{
		Price:       request.Price,
		Description: request.Description,
	}

	id, err := a.RoomUsecase.Add(&room)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	response := api.AddRoomResponse{Id: strconv.FormatUint(id, 10)}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

// Delete will delete room by given param
func (a *RoomHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)

	err = a.RoomUsecase.Delete(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
