package main

import (
	"net/http"
	"strconv"
	"log"
	"flag"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"lets-go-chat/configs"
	"lets-go-chat/internal/handlers"
	rep "lets-go-chat/internal/repositories"
)

func main() {

	configFile := flag.String("config", "", "Configuration file in JSON-format")
	flag.Parse()

	config, err := configs.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("can't load configuration")
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBConfig.DBHost, config.DBConfig.DBPort, config.DBConfig.DBUser, config.DBConfig.DBPassword, config.DBConfig.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("can't connect to DB")
	}

	defer db.Close()

	rep.NewUsersRepository(db)

	r := mux.NewRouter()
	r.HandleFunc("/v1/user", handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/v1/user/login", handlers.LoginUser).Methods(http.MethodPost)
	err =  http.ListenAndServe(":"+strconv.Itoa(config.ServerPort), r)
	if err != nil {
		log.Fatal("server failed to add listener")
	}

}