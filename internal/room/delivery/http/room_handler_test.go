package http_test

import (
	"encoding/json"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/domain/mocks"
	roomHttp "github.com/babon21/hotel-management/internal/room/delivery/http"
	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRoomList(t *testing.T) {
	var mockRoom domain.Room
	err := faker.FakeData(&mockRoom)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoomUsecase)
	mockListRoom := make([]domain.Room, 0)
	mockListRoom = append(mockListRoom, mockRoom)

	mockUCase.On("GetList", mock.Anything, mock.Anything).Return(mockListRoom, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/rooms?sort_by=price&order_by=asc", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := roomHttp.RoomHandler{
		RoomUsecase: mockUCase,
	}
	err = handler.GetRoomList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetRoomListError(t *testing.T) {
	mockUCase := new(mocks.RoomUsecase)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/rooms?sort_by=price&order_by=abracadabra", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := roomHttp.RoomHandler{
		RoomUsecase: mockUCase,
	}
	err = handler.GetRoomList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestAdd(t *testing.T) {
	mockRoom := domain.Room{
		Price:       "100",
		Description: "desc",
	}

	tempMockRoom := mockRoom
	mockUCase := new(mocks.RoomUsecase)

	j, err := json.Marshal(tempMockRoom)
	assert.NoError(t, err)

	mockUCase.On("Add", mock.AnythingOfType("*domain.Room")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/rooms", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := roomHttp.RoomHandler{
		RoomUsecase: mockUCase,
	}
	err = handler.Add(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockRoom domain.Room
	err := faker.FakeData(&mockRoom)
	assert.NoError(t, err)

	mockUCase := new(mocks.RoomUsecase)

	num := mockRoom.ID

	mockUCase.On("Delete", num).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/rooms/"+num, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("rooms/:id")
	c.SetParamNames("id")
	c.SetParamValues(num)
	handler := roomHttp.RoomHandler{
		RoomUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}
