package api

type AddRoomRequest struct {
	Price       string `json:"price"`
	Description string `json:"desc"`
}

type AddRoomResponse struct {
	Id string `json:"id"`
}

//type AddBookingRequest struct {
//
//}
//
//type AddBookingResponse struct {
//
//}
