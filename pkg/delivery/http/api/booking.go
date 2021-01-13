package api

type AddBookingRequest struct {
	RoomId         string `json:"room_id"`
	StartDate      string `json:"start_date"`
	ExpirationDate string `json:"expiration_date"`
}

type AddBookingResponse struct {
	BookingId string `json:"booking_id"`
}
