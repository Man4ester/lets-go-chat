package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
	"time"
	"fmt"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	userRep := rep.GetUSerRepository()
	user, err := userRep.GetUserByUserName(userLoginRequest.UserName)
	if errors.Is(err, rep.UserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
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
