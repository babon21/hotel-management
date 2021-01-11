package usecase

import (
	"github.com/babon21/hotel-management/domain"
)

// RoomRepository represent the room's repository contract
type RoomRepository interface {
	GetList(sortField SortField, sortOrder SortOrder) ([]domain.Room, error)
	Save(room *domain.Room) (uint64, error)
	Remove(roomId int64) error
}
