package usecase

import (
	"github.com/babon21/hotel-management/internal/domain"
)

// BookingUsecase represent the booking's usecases
type BookingUsecase interface {
	GetList(roomId string) ([]domain.Booking, error)
	Add(booking *domain.Booking) (string, error)
	Delete(bookingId string) error
}

type bookingUsecase struct {
	bookingRepository BookingRepository
	roomChecker       domain.ExistenceChecker
}

// NewBookingUsecase will create new an bookingUsecase object representation of BookingUsecase interface
func NewBookingUsecase(bookingRepository BookingRepository, roomChecker domain.ExistenceChecker) BookingUsecase {
	return &bookingUsecase{
		bookingRepository: bookingRepository,
		roomChecker:       roomChecker,
	}
}

func (useCase *bookingUsecase) GetList(roomId string) ([]domain.Booking, error) {
	return useCase.bookingRepository.GetList(roomId)
}

func (useCase *bookingUsecase) Add(booking *domain.Booking) (string, error) {
	if !useCase.roomChecker.CheckExistence(booking.RoomId) {
		return "", domain.ErrNotFound
	}
	bookingId, err := useCase.bookingRepository.Save(booking)
	return bookingId, err
}

func (useCase *bookingUsecase) Delete(bookingId string) error {
	if !useCase.bookingRepository.CheckBookingExists(bookingId) {
		return domain.ErrNotFound
	}
	return useCase.bookingRepository.Remove(bookingId)
}
