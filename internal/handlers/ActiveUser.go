package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"lets-go-chat/internal/services"
	"lets-go-chat/internal/models"
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


