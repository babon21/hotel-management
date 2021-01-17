package postgres_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babon21/hotel-management/internal/booking/repository/postgres"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetList(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockBookings := []domain.Booking{
		{ID: "1", RoomId: "10", StartDate: "2006-01-02"},
		{ID: "2", RoomId: "20", ExpirationDate: "2008-01-02"},
	}

	rows := sqlmock.NewRows([]string{"id", "room_id", "start_date", "expiration_date"}).
		AddRow(mockBookings[0].ID, mockBookings[0].RoomId, mockBookings[0].StartDate, mockBookings[0].ExpirationDate).
		AddRow(mockBookings[1].ID, mockBookings[1].RoomId, mockBookings[1].StartDate, mockBookings[1].ExpirationDate)

	query := "SELECT id,room_id,start_date,expiration_date FROM booking WHERE room_id = \\$1 ORDER BY start_date ASC"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresBookingRepository(db)

	list, err := a.GetList("1")
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
