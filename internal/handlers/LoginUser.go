package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
	"time"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"lets-go-chat/pkg/jwt"
	"log"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	userRep := rep.GetUserRepository()
	user, err := userRep.GetUserByUserName(userLoginRequest.UserName)
	if errors.Is(err, rep.UserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	userAuth := hasher.CheckPasswordHash(userLoginRequest.Password, user.Password)
	if !userAuth {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	token, err := jwt.GenerateJWT(userLoginRequest.UserName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	userLoginResponse := models.LoginUserResponse{
		Url: "ws://fancy-chat.io/ws&token="+token, //ws://fancy-chat.io/ws&token=one-time-token
	}

	err = json.NewEncoder(w).Encode(&userLoginResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("X-Rate-Limit", "2")
	w.Header().Add("X-Expires-After", time.Now().UTC().String())
	w.WriteHeader(http.StatusFound)
}
