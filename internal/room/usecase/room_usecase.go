package usecase

import (
	"github.com/babon21/hotel-management/internal/domain"
)

// RoomUsecase represent the room's usecases
type RoomUsecase interface {
	GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error)
	Add(room *domain.Room) error
	Delete(roomId string) error
}

type roomUsecase struct {
	roomRepo RoomRepository
}

// NewRoomUsecase will create new an roomUsecase object representation of domain.RoomUsecase interface
func NewRoomUsecase(roomRepository RoomRepository) RoomUsecase {
	return &roomUsecase{
		roomRepo: roomRepository,
	}
}

func (useCase *roomUsecase) GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error) {
	list, err := useCase.roomRepo.GetList(sortField, sortOrder)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (useCase *roomUsecase) Add(room *domain.Room) error {
	return useCase.roomRepo.Save(room)
}

func (useCase *roomUsecase) Delete(roomId string) error {
	if !useCase.roomRepo.CheckExistence(roomId) {
		return domain.ErrNotFound
	}
	return useCase.roomRepo.Remove(roomId)
}
