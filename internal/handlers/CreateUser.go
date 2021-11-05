package handlers

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(userRequest.UserName) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHashed, err := hasher.HashPassword(userRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		userResponse := models.CreateUserResponse{
			UserName: userRequest.UserName,
			Id:       userId.String(),
		}

		user := models.User{
			Id:       userResponse.Id,
			UserName: userResponse.UserName,
			Password: passwordHashed,
		}
		err = rep.SaveUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&userResponse)
	}

}
