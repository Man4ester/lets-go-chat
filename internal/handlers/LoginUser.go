package handlers

import (
	"encoding/json"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest models.LoginUserRequest
	_ = json.NewDecoder(r.Body).Decode(&userLoginRequest)
	user, err := rep.GetUserByUserName(userLoginRequest.UserName)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
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
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&userLoginResponse)
}
