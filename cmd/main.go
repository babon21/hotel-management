package main

import (
	"fmt"
	bookingRepository "github.com/babon21/hotel-management/internal/booking/repository/postgres"
	bookingUseCase "github.com/babon21/hotel-management/internal/booking/usecase"
	roomRepository "github.com/babon21/hotel-management/internal/room/repository/postgres"
	roomUseCase "github.com/babon21/hotel-management/internal/room/usecase"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	bookingHttp "github.com/babon21/hotel-management/internal/booking/delivery/http"
	roomHttp "github.com/babon21/hotel-management/internal/room/delivery/http"
	//"github.com/labstack/echo"
	//"github.com/spf13/viper"
	//"net/url"
	//"time"
)

func main() {
	dbUser := "postgres"
	dbPassword := "postgres"
	//host := "postgres"
	host := "localhost"
	dbPort := "5432"
	dbName := "postgres"

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", dbUser, dbPassword, host, dbPort, dbName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	e := echo.New()
	roomRepo := roomRepository.NewPostgresRoomRepository(db)
	roomUsecase := roomUseCase.NewRoomUsecase(roomRepo)
	roomHttp.NewRoomHandler(e, roomUsecase)

	bookingRepo := bookingRepository.NewPostgresBookingRepository(db)
	bookingUsecase := bookingUseCase.NewBookingUsecase(bookingRepo, roomRepo)
	bookingHttp.NewBookingHandler(e, bookingUsecase)

	log.Fatal().Msg(e.Start(":9090").Error())
}
