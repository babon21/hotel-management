package usecase

import (
	"github.com/babon21/hotel-management/internal/domain"
)

// BookingRepository represent the booking's repository contract
type BookingRepository interface {
	GetList(roomId string) ([]domain.Booking, error)
	Save(booking *domain.Booking) (string, error)
	Remove(bookingId string) error
	CheckBookingExists(bookingId string) bool
}
