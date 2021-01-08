package domain

// Booking ...
type Booking struct {
	ID             int64  `json:"id"`
	RoomId         int64  `json:"room_id"` // TODO exclude from json
	StartDate      string `json:"start_date"`
	ExpirationDate string `json:"expiration_date"`
}

// BookingUsecase represent the booking's usecases
type BookingUsecase interface {
	GetListByRoomId(roomId int64) ([]Booking, error)
	Add(booking *Booking) (int64, error)
	Delete(bookingId int64) error
}

// BookingRepository represent the booking's repository contract
type BookingRepository interface {
	FindByRoomId(roomId int64) ([]Booking, error)
	Save(room *Booking) error
	Remove(bookingId int64) error
}
