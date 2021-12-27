package services

import (
	"encoding/json"
	"log"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"lets-go-chat/internal/models"
)
var (
	rdb *redis.Client
)

var broadcaster = make(chan models.ChatMessage)
var clients = make(map[*websocket.Conn]bool)

func StartRedis(redisURL string){
	rdb = redis.NewClient(&redis.Options{
		Addr: redisURL,
		Password: "",
		DB: 0,
	})
}

func HandleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		storeInRedis(msg)
		messageClients(msg)
	}
}

func RegisterNewClient (c *websocket.Conn){
	clients[c] = true

	if rdb.Exists("chat_messages").Val() != 0 {
		sendPreviousMessages(c)
	}
}

func AddMessage(msg models.ChatMessage) {
	broadcaster <- msg
}

func messageClients(msg models.ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

func messageClient(client *websocket.Conn, msg models.ChatMessage) {
	err := client.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

func sendPreviousMessages(ws *websocket.Conn) {
	chatMessages, err := rdb.LRange("chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// send previous messages
	for _, chatMessage := range chatMessages {
		var msg models.ChatMessage
		json.Unmarshal([]byte(chatMessage), &msg)
		messageClient(ws, msg)
	}
}

func storeInRedis(msg models.ChatMessage) {
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	if err := rdb.RPush("chat_messages", json).Err(); err != nil {
		panic(err)
	}
}
