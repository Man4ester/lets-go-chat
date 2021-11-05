package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"lets-go-chat/configs"
	"lets-go-chat/internal/handlers"
	"net/http"
	"strconv"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/v1/user", handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/v1/user/login", handlers.LoginUser).Methods(http.MethodPost)
	http.ListenAndServe(":"+strconv.Itoa(configs.Config.ServerPort), r)

}
