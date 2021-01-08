package domain

// Room ...
type Room struct {
	ID          int64  `json:"id"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	DateAdded   string `json:"date_added"`
}

// RoomUsecase represent the room's usecases
type RoomUsecase interface {
	GetList(sortField string, isAscendingOrder bool) ([]Room, error)
	Add(room *Room) (int64, error)
	Delete(roomId int64) error
}

// RoomRepository represent the room's repository contract
type RoomRepository interface {
	FindById(roomId int64) ([]Room, error)
	Save(room *Room) error
	Remove(roomId int64) error
}
