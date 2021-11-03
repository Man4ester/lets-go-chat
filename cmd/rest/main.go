package main

import (
	"github.com/gorilla/mux"
	handlers "lets-go-chat/internal/handlers"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/v1/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/v1/user/login", handlers.LoginUser).Methods("POST")
	http.ListenAndServe(":8080", r)


}