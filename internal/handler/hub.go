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

type wsClient struct {
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	clients map[int]*wsClient
	mutex   sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int]*wsClient),
	}
}

func (h *Hub) Register(userID int, conn *websocket.Conn) {
	h.mutex.Lock()

	if oldClient, exists := h.clients[userID]; exists {
		close(oldClient.send)
		oldClient.conn.Close()
	}

	client := &wsClient{
		conn: conn,
		send: make(chan []byte, 256),
	}
	h.clients[userID] = client
	h.mutex.Unlock()

	go h.writePump(userID, client)
}

func (h *Hub) Unregister(userID int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if client, exists := h.clients[userID]; exists {
		close(client.send)
		client.conn.Close()
		delete(h.clients, userID)
	}
}

func (h *Hub) SendToUser(userID int, message []byte) bool {
	h.mutex.RLock()
	client, exists := h.clients[userID]
	h.mutex.RUnlock()

	if !exists {
		return false
	}

	select {
	case client.send <- message:
		return true
	default:
		h.Unregister(userID)
		return false
	}
}

func (h *Hub) writePump(userID int, client *wsClient) {
	defer func() {
		h.Unregister(userID)
	}()

	for message := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
