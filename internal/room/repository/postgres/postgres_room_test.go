package postgres_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/room/repository/postgres"
	"github.com/babon21/hotel-management/internal/room/usecase"
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

	mockArticles := []domain.Room{
		{ID: "1", Price: "10", Description: "desc1"},
		{ID: "2", Price: "20", Description: "desc2"},
	}

	rows := sqlmock.NewRows([]string{"id", "price", "description", "date_added"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Price, mockArticles[0].Description, mockArticles[0].DateAdded).
		AddRow(mockArticles[1].ID, mockArticles[1].Price, mockArticles[1].Description, mockArticles[1].DateAdded)

	sortField := usecase.PriceField
	sortOrder := usecase.AscOrder
	query := "SELECT id,price,description,date_added FROM room ORDER BY price ASC"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresRoomRepository(db)

	list, err := a.GetList(sortField, sortOrder)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
