package handlers

import (
	"net/http"
	"lets-go-chat/pkg/jwt"
	"lets-go-chat/internal/services"
)

func WsRTMStart(w http.ResponseWriter, r *http.Request) {
	keys, _ := r.URL.Query()["token"]
	token :=keys[0]
	userName, _ := jwt.DecodeJWT(token)
	services.AddUserToCache(userName)
}
