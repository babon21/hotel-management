package postgres

import (
	"fmt"
	"github.com/babon21/hotel-management/domain"
	"github.com/babon21/hotel-management/room/usecase"
	"github.com/jmoiron/sqlx"
	"time"
)

type postgresRoomRepository struct {
	Conn *sqlx.DB
}

// NewPostgresRoomRepository will create an object that represent the RoomRepository interface
func NewPostgresRoomRepository(conn *sqlx.DB) usecase.RoomRepository {
	return &postgresRoomRepository{conn}
}

func (repo *postgresRoomRepository) GetList(sortField usecase.SortField, sortOrder usecase.SortOrder) ([]domain.Room, error) {
	getListQuery := formGetListQuery(sortField, sortOrder)
	var rooms []domain.Room
	err := repo.Conn.Select(&rooms, getListQuery)
	for i := range rooms {
		t, _ := time.Parse(time.RFC3339, rooms[i].DateAdded)
		rooms[i].DateAdded = fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
		continue
	}
	return rooms, err
}

func formGetListQuery(sortField usecase.SortField, sortOrder usecase.SortOrder) string {
	var order string
	switch sortOrder {
	case usecase.AscOrder:
		order = "ASC"
	case usecase.DescOrder:
		order = "DESC"
	}

	return fmt.Sprintf("SELECT * FROM room ORDER BY %s %s", sortField, order)
}

func (repo *postgresRoomRepository) Save(room *domain.Room) (uint64, error) {
	var id uint64
	err := repo.Conn.QueryRow("INSERT INTO room(price, description, date_added) VALUES ($1, $2, $3) RETURNING id", room.Price, room.Description, time.Now()).Scan(&id)
	return id, err
}

func (repo *postgresRoomRepository) Remove(roomId int64) error {
	deleteQuery := "DELETE FROM room WHERE id = $1"
	_, err := repo.Conn.Exec(deleteQuery, roomId)
	return err
}
