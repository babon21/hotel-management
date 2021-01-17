package usecase_test

import (
	"errors"
	"github.com/babon21/hotel-management/internal/booking/usecase"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetList(t *testing.T) {
	mockBookingRepo := new(mocks.BookingRepository)
	mockChecker := new(mocks.ExistenceChecker)
	mockBooking := domain.Booking{
		ID: "1",
	}

	mockListBooking := make([]domain.Booking, 0)
	mockListBooking = append(mockListBooking, mockBooking)

	t.Run("success", func(t *testing.T) {
		mockChecker.On("CheckExistence", mock.AnythingOfType("string")).Return(true).Once()
		mockBookingRepo.On("GetList", mock.AnythingOfType("string")).Return(mockListBooking, nil).Once()

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		list, err := u.GetList("1")

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListBooking))
		mockBookingRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBookingRepo.On("GetList", mock.AnythingOfType("string")).
			Return(nil, errors.New("Unexpected Error")).Once()

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		list, err := u.GetList("1")

		assert.Error(t, err)
		assert.Nil(t, list)
		mockBookingRepo.AssertExpectations(t)
	})
}

func TestAdd(t *testing.T) {
	mockBookingRepo := new(mocks.BookingRepository)
	mockChecker := new(mocks.ExistenceChecker)

	mockBooking := domain.Booking{
		RoomId: "1",
	}

	t.Run("success", func(t *testing.T) {
		tempMockBooking := mockBooking

		mockChecker.On("CheckExistence", mock.AnythingOfType("string")).Return(true).Once()
		mockBookingRepo.On("Save", mock.Anything).Return(nil)

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		err := u.Add(&tempMockBooking)

		assert.NoError(t, err)
		assert.Equal(t, mockBooking.RoomId, tempMockBooking.RoomId)
		mockBookingRepo.AssertExpectations(t)
	})

	t.Run("room-is-not-exist", func(t *testing.T) {
		tempMockBooking := mockBooking

		mockChecker.On("CheckExistence", mock.AnythingOfType("string")).Return(false).Once()
		mockBookingRepo.On("Save", mock.Anything).Return(nil)

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		err := u.Add(&tempMockBooking)

		assert.Error(t, err)
		assert.Equal(t, mockBooking.RoomId, tempMockBooking.RoomId)
		mockBookingRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockBookingRepo := new(mocks.BookingRepository)
	mockChecker := new(mocks.ExistenceChecker)

	t.Run("success", func(t *testing.T) {
		mockBookingRepo.On("CheckBookingExists", mock.AnythingOfType("string")).Return(true).Once()
		mockBookingRepo.On("Remove", mock.AnythingOfType("string")).Return(nil).Once()

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		err := u.Delete("1")

		assert.NoError(t, err)
		mockBookingRepo.AssertExpectations(t)
	})
	t.Run("booking-is-not-exist", func(t *testing.T) {
		mockBookingRepo.On("CheckBookingExists", mock.AnythingOfType("string")).Return(false).Once()

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		err := u.Delete("1")

		assert.Error(t, err)
		mockBookingRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockBookingRepo.On("CheckBookingExists", mock.AnythingOfType("string")).Return(true).Once()
		mockBookingRepo.On("Remove", mock.AnythingOfType("string")).Return(errors.New("Unexpected Error"))

		u := usecase.NewBookingUsecase(mockBookingRepo, mockChecker)
		err := u.Delete("1")

		assert.Error(t, err)
		mockBookingRepo.AssertExpectations(t)
	})
}
