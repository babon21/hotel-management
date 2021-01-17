package http_test

import (
	"encoding/json"
	bookingHttp "github.com/babon21/hotel-management/internal/booking/delivery/http"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/domain/mocks"
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

func TestGetBookingList(t *testing.T) {
	var mockBooking domain.Booking
	err := faker.FakeData(&mockBooking)
	assert.NoError(t, err)
	mockUCase := new(mocks.BookingUsecase)
	mockListBooking := make([]domain.Booking, 0)
	mockListBooking = append(mockListBooking, mockBooking)

	mockUCase.On("GetList", mock.Anything, mock.Anything).Return(mockListBooking, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/bookings?room_id=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bookingHttp.BookingHandler{
		BookingUsecase: mockUCase,
	}
	err = handler.GetBookingList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetBookingListError(t *testing.T) {
	mockUCase := new(mocks.BookingUsecase)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/booking?room=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bookingHttp.BookingHandler{
		BookingUsecase: mockUCase,
	}
	err = handler.GetBookingList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestAdd(t *testing.T) {
	mockBooking := domain.Booking{
		RoomId:         "1",
		StartDate:      "2006-01-02",
		ExpirationDate: "2008-01-02",
	}

	tempMockArticle := mockBooking
	mockUCase := new(mocks.BookingUsecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Add", mock.AnythingOfType("*domain.Booking")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/bookings", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := bookingHttp.BookingHandler{
		BookingUsecase: mockUCase,
	}
	err = handler.Add(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockBooking domain.Booking
	err := faker.FakeData(&mockBooking)
	assert.NoError(t, err)

	mockUCase := new(mocks.BookingUsecase)

	num := mockBooking.ID

	mockUCase.On("Delete", num).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/bookings/"+num, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("bookings/:id")
	c.SetParamNames("id")
	c.SetParamValues(num)
	handler := bookingHttp.BookingHandler{
		BookingUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}
