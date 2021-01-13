package postgres

import (
	"fmt"
	"github.com/babon21/hotel-management/internal/domain"
	"github.com/babon21/hotel-management/internal/room/usecase"
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

//func (repo *postgresRoomRepository) CheckRoomExists(roomId string) bool {
//	var booking domain.Room
//	err := repo.Conn.Get(&booking, "SELECT * FROM room WHERE id = $1", roomId)
//	if err != nil {
//		return false
//	}
//	return true
//}

func (repo *postgresRoomRepository) CheckExistence(id string) bool {
	var room domain.Room
	err := repo.Conn.Get(&room, "SELECT * FROM room WHERE id = $1", id)
	if err != nil {
		return false
	}
	return true
}

func (repo *postgresRoomRepository) GetList(sortField usecase.SortField, sortOrder usecase.SortOrder) ([]domain.Room, error) {
	getListQuery := formGetListQuery(sortField, sortOrder)
	rooms := make([]domain.Room, 0, 1)
	err := repo.Conn.Select(&rooms, getListQuery)
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

func (repo *postgresRoomRepository) Save(room *domain.Room) (string, error) {
	var id string
	err := repo.Conn.QueryRow("INSERT INTO room(price, description, date_added) VALUES ($1, $2, $3) RETURNING id", room.Price, room.Description, time.Now()).Scan(&id)
	return id, err
}

func (repo *postgresRoomRepository) Remove(roomId string) error {
	deleteQuery := "DELETE FROM room WHERE id = $1"
	_, err := repo.Conn.Exec(deleteQuery, roomId)
	return err
}
