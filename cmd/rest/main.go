package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lets-go-chat/configs"
	"lets-go-chat/internal/handlers"
	"net/http"
	"os"
	"strconv"
)


var configuration configs.Configuration

func main(){
	err :=loadConfiguration()
	if err != nil {
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/v1/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/v1/user/login", handlers.LoginUser).Methods("POST")
	http.ListenAndServe(":" + strconv.Itoa(configuration.ServerPort), r)

}

func loadConfiguration() error{
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		return err
	}
	return nil
}