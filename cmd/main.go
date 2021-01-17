package main

import (
	"fmt"
	"github.com/babon21/hotel-management/internal/http/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"

	bookingRepository "github.com/babon21/hotel-management/internal/booking/repository/postgres"
	bookingUseCase "github.com/babon21/hotel-management/internal/booking/usecase"
	"github.com/babon21/hotel-management/internal/config"
	roomRepository "github.com/babon21/hotel-management/internal/room/repository/postgres"
	roomUseCase "github.com/babon21/hotel-management/internal/room/usecase"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"

	bookingHttp "github.com/babon21/hotel-management/internal/booking/delivery/http"
	roomHttp "github.com/babon21/hotel-management/internal/room/delivery/http"
)

func main() {
	conf := config.Init()

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", conf.Database.Username,
		conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.AccessLogMiddleware)
	roomRepo := roomRepository.NewPostgresRoomRepository(db)
	roomUsecase := roomUseCase.NewRoomUsecase(roomRepo)
	roomHttp.NewRoomHandler(e, roomUsecase)

	bookingRepo := bookingRepository.NewPostgresBookingRepository(db)
	bookingUsecase := bookingUseCase.NewBookingUsecase(bookingRepo, roomRepo)
	bookingHttp.NewBookingHandler(e, bookingUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
