package domain

// Room ...
type Room struct {
	ID          string `json:"id"`
	Price       string `json:"price"`
	Description string `json:"desc"`
	DateAdded   string `json:"date_added" db:"date_added"`
}
