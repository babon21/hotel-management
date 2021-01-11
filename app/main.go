package main

import (
	//"database/sql"
	"fmt"
	"github.com/babon21/hotel-management/room/repository/postgres"
	"github.com/babon21/hotel-management/room/usecase"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	roomHttp "github.com/babon21/hotel-management/room/delivery/http"
	//"github.com/labstack/echo"
	//"github.com/spf13/viper"
	//"log"
	//"net/url"
	//"time"
	//_articleHttpDelivery "github.com/babon21/hotel-management/room/delivery/http"
	//_articleHttpDeliveryMiddleware "github.com/babon21/hotel-management/room/delivery/http/middleware"
	//_articleRepo "github.com/babon21/hotel-management/room/repository/mysql"
	//_articleUcase "github.com/babon21/hotel-management/room/usecase"
	//_authorRepo "github.com/babon21/hotel-management/author/repository/mysql"
)

func main() {
	dbUser := "postgres"
	dbPassword := "postgres"
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
	roomRepo := postgres.NewPostgresRoomRepository(db)
	roomUseCase := usecase.NewRoomUsecase(roomRepo)
	roomHttp.NewRoomHandler(e, roomUseCase)
	log.Fatal().Msg(e.Start(":9090").Error())
}

//func init() {
//	viper.SetConfigFile(`config.json`)
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//
//	if viper.GetBool(`debug`) {
//		log.Println("Service RUN on DEBUG mode")
//	}
//}

//func main() {
//	dbHost := viper.GetString(`database.host`)
//	dbPort := viper.GetString(`database.port`)
//	dbUser := viper.GetString(`database.user`)
//	dbPass := viper.GetString(`database.pass`)
//	dbName := viper.GetString(`database.name`)
//	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
//	val := url.Values{}
//	val.Add("parseTime", "1")
//	val.Add("loc", "Asia/Jakarta")
//	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
//	dbConn, err := sql.Open(`mysql`, dsn)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = dbConn.Ping()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func() {
//		err := dbConn.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	e := echo.New()
//	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
//	e.Use(middL.CORS)
//	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
//	ar := _articleRepo.NewMysqlArticleRepository(dbConn)
//
//	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
//	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
//	_articleHttpDelivery.NewArticleHandler(e, au)
//
//	log.Fatal(e.Start(viper.GetString("server.address")))
//}
