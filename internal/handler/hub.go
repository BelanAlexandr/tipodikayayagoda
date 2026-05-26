package handler

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var GlobalHub *Hub

type Hub struct {
	clients map[int]*websocket.Conn
	mutex   sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int]*websocket.Conn),
	}
}

func (h *Hub) SendToUser(userID int, message []byte) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	client, exists := h.clients[userID]
	if !exists {
		return false
	}

	err := client.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		client.Close()
		delete(h.clients, userID)
		return false
	}
	return true
}
