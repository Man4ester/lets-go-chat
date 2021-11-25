package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"lets-go-chat/internal/models"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/pkg/hasher"
	"log"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	userId, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if len(userRequest.UserName) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHashed, err := hasher.HashPassword(userRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
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

		userRep := rep.GetUserRepository()

		err = userRep.SaveUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err =  json.NewEncoder(w).Encode(&userResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}

}
