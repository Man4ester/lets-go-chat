package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"lets-go-chat/pkg/jwt"
	"lets-go-chat/internal/services"
)

var upgrader = websocket.Upgrader{}

func WsRTMStart(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	keys, _ := r.URL.Query()["token"]
	token :=keys[0]
	userName, _ := jwt.DecodeJWT(token)
	services.AddUserToCache(userName)

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	log.Println("Disconnected user:" + userName)
	services.RemoveUserFromCache(userName)

}
