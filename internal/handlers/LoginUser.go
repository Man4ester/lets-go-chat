package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"time"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := rep.GetUserByUserName(userLoginRequest.UserName)
	if errors.Is(err, rep.UserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userAuth := hasher.CheckPasswordHash(userLoginRequest.Password, user.Password)
	if !userAuth {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userLoginResponse := models.LoginUserResponse{
		Url: "redirect to user",
	}

	err = json.NewEncoder(w).Encode(&userLoginResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("X-Rate-Limit", "2")
	w.Header().Add("X-Expires-After", time.Now().UTC().String())
	w.WriteHeader(http.StatusFound)
}
