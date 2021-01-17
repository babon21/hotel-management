package usecase_test

import (
	"errors"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/domain/mocks"
	"github.com/babon21/hotel-management/internal/room/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetList(t *testing.T) {
	mockRoomRepo := new(mocks.RoomRepository)
	mockRoom := domain.Room{
		ID:          "1",
		Price:       "10",
		Description: "desc",
	}

	mockListRoom := make([]domain.Room, 0)
	mockListRoom = append(mockListRoom, mockRoom)

	t.Run("success", func(t *testing.T) {
		mockRoomRepo.On("GetList", mock.Anything, mock.Anything).Return(mockListRoom, nil).Once()
		u := usecase.NewRoomUsecase(mockRoomRepo)
		list, err := u.GetList(usecase.PriceField, usecase.AscOrder)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListRoom))

		mockRoomRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRoomRepo.On("GetList", mock.Anything, mock.Anything).
			Return(nil, errors.New("Unexpected Error")).Once()
		u := usecase.NewRoomUsecase(mockRoomRepo)
		list, err := u.GetList(usecase.PriceField, usecase.AscOrder)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockRoomRepo.AssertExpectations(t)
	})
}

func TestAdd(t *testing.T) {
	mockRoomRepo := new(mocks.RoomRepository)
	mockRoom := domain.Room{
		Price:       "10",
		Description: "desc",
	}

	t.Run("success", func(t *testing.T) {
		tempMockRoom := mockRoom
		tempMockRoom.ID = "0"

		mockRoomRepo.On("Save", mock.Anything).Return(nil)
		u := usecase.NewRoomUsecase(mockRoomRepo)
		err := u.Add(&tempMockRoom)

		assert.NoError(t, err)
		assert.Equal(t, mockRoom.Description, tempMockRoom.Description)
		mockRoomRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRoomRepo := new(mocks.RoomRepository)

	t.Run("success", func(t *testing.T) {
		mockRoomRepo.On("CheckExistence", mock.AnythingOfType("string")).Return(true).Once()
		mockRoomRepo.On("Remove", mock.AnythingOfType("string")).Return(nil).Once()

		u := usecase.NewRoomUsecase(mockRoomRepo)
		err := u.Delete("1")

		assert.NoError(t, err)
		mockRoomRepo.AssertExpectations(t)
	})
	t.Run("room-is-not-exist", func(t *testing.T) {
		mockRoomRepo.On("CheckExistence", mock.AnythingOfType("string")).Return(false).Once()

		u := usecase.NewRoomUsecase(mockRoomRepo)
		err := u.Delete("1")

		assert.Error(t, err)
		mockRoomRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockRoomRepo.On("CheckExistence", mock.AnythingOfType("string")).Return(true).Once()
		mockRoomRepo.On("Remove", mock.AnythingOfType("string")).Return(errors.New("Unexpected Error"))

		u := usecase.NewRoomUsecase(mockRoomRepo)
		err := u.Delete("1")

		assert.Error(t, err)
		mockRoomRepo.AssertExpectations(t)
	})
}
