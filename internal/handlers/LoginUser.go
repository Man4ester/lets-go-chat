package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
	"log"
	"time"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"lets-go-chat/pkg/jwt"
	"lets-go-chat/internal/services"
)

type UserLogin struct {
	Repo rep.UserRepository
}

func (uc UserLogin) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	user, err := uc.Repo.GetUserByUserName(userLoginRequest.UserName)
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
		Url: "ws://"+ r.Host + "/v1/chat/ws.rtm.start?token="+token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Rate-Limit", "2")
	w.Header().Add("X-Expires-After", time.Now().UTC().String())
	err = json.NewEncoder(w).Encode(&userLoginResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	services.RegisterToken(token)
	w.WriteHeader(http.StatusFound)
}
