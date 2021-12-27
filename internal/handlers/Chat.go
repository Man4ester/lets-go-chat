package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"lets-go-chat/pkg/jwt"
	"lets-go-chat/internal/services"
	"lets-go-chat/internal/models"
)


type WsRTM struct {
	Upgrader websocket.Upgrader
}

func (ws WsRTM)WsRTMStart(w http.ResponseWriter, r *http.Request) {
	c, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	keys, _ := r.URL.Query()["token"]
	token :=keys[0]
	userName, _ := jwt.DecodeJWT(token)
	services.AddUserToCache(userName)
	services.RegisterNewClient(c)

	go services.HandleMessages()

	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		var msg  = models.ChatMessage{
			Text: string(message),
			Username: userName,
		}

		services.AddMessage(msg)

		if err != nil {
			log.Println("read:", err)
			break
		}
	}

	log.Println("Disconnected user:" + userName)
	services.RemoveUserFromCache(userName)
}
