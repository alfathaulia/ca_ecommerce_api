package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
	// _ "github.com/lib/pq"
	_userDelivery "github.com/alfathaulia/ca_ecommerce_api/user/delivery/http"
	_userRepo "github.com/alfathaulia/ca_ecommerce_api/user/repository/mysql"
	_userUcase "github.com/alfathaulia/ca_ecommerce_api/user/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	//postgree
	// connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	// dbConn, err := sql.Open(`postgres`, connection)
	//mysql
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepo := _userRepo.NewMysqlUserRepo(dbConn)
	userUcase := _userUcase.NewUserUsecase(userRepo, timeoutContext)

	_userDelivery.NewUserHandler(e, userUcase)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
