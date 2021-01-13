package domain

// Booking ...
type Booking struct {
	ID             string `json:"id"`
	RoomId         string `json:"-" db:"room_id"`
	StartDate      string `json:"start_date" db:"start_date"`
	ExpirationDate string `json:"expiration_date" db:"expiration_date"`
}
