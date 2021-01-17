package usecase

import (
	"github.com/babon21/hotel-management/internal/domain"
)

// RoomRepository represent the room's repository contract
type RoomRepository interface {
	GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error)
	Save(room *domain.Room) error
	Remove(roomId string) error
	domain.ExistenceChecker
	//CheckRoomExists(roomId string) bool
}
