package handlers

import (
	"lets-go-chat/internal/services"
	"net/http"
	"lets-go-chat/internal/models"
	"encoding/json"
	"log"
)

func GetActiveUsers(w http.ResponseWriter, r *http.Request) {
	count :=services.GetTotalActiveUsers()
	userResponse := models.ActiveUserResponse{
		Count: count,
	}

	err := json.NewEncoder(w).Encode(&userResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}


