package postgres

import (
	"github.com/babon21/hotel-management/internal/booking/usecase"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/jmoiron/sqlx"
)

type postgresBookingRepository struct {
	Conn *sqlx.DB
}

// NewPostgresBookingRepository will create an object that represent the BookingRepository interface
func NewPostgresBookingRepository(conn *sqlx.DB) usecase.BookingRepository {
	return &postgresBookingRepository{conn}
}

func (repo *postgresBookingRepository) CheckBookingExists(bookingId string) bool {
	var booking domain.Booking
	err := repo.Conn.Get(&booking, "SELECT * FROM booking WHERE id = $1", bookingId)
	if err != nil {
		return false
	}
	return true
}

func (repo *postgresBookingRepository) GetList(roomId string) ([]domain.Booking, error) {
	getListQuery := "SELECT id,room_id,start_date,expiration_date FROM booking WHERE room_id = $1 ORDER BY start_date ASC"
	bookings := make([]domain.Booking, 0, 1)
	err := repo.Conn.Select(&bookings, getListQuery, roomId)
	return bookings, err
}

func (repo *postgresBookingRepository) Save(booking *domain.Booking) error {
	var id string
	err := repo.Conn.QueryRow("INSERT INTO booking(room_id, start_date, expiration_date) VALUES ($1, $2, $3) RETURNING id", booking.RoomId, booking.StartDate, booking.ExpirationDate).Scan(&id)
	booking.ID = id
	return err
}

func (repo *postgresBookingRepository) Remove(bookingId string) error {
	deleteQuery := "DELETE FROM booking WHERE id = $1"
	_, err := repo.Conn.Exec(deleteQuery, bookingId)
	return err
}
