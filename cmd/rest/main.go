package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"lets-go-chat/internal/handlers"
	"net/http"
)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	http.Handle("/", r)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}