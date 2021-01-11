package usecase

import (
	"github.com/babon21/hotel-management/domain"
)

// RoomUsecase represent the room's usecases
type RoomUsecase interface {
	GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error)
	Add(room *domain.Room) (uint64, error)
	Delete(roomId int64) error
}

type roomUsecase struct {
	roomRepo RoomRepository
	//contextTimeout time.Duration
}

// NewRoomUsecase will create new an roomUsecase object representation of domain.RoomUsecase interface
func NewRoomUsecase(roomRepository RoomRepository) RoomUsecase {
	return &roomUsecase{
		roomRepo: roomRepository,
		//contextTimeout: timeout,
	}
}

func (useCase *roomUsecase) GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error) {
	return useCase.roomRepo.GetList(sortField, sortOrder)
}

func (useCase *roomUsecase) Add(room *domain.Room) (uint64, error) {
	roomId, err := useCase.roomRepo.Save(room)
	return roomId, err
}

func (useCase *roomUsecase) Delete(roomId int64) error {
	// TODO check exists room with specified roomId
	return useCase.roomRepo.Remove(roomId)
}
