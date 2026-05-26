package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/repository"

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

func getUserIDFromContext(r *http.Request) int {
	val := r.Context().Value(middelware.UserKey)
	if val == nil {
		return 0
	}

	userCtx, ok := val.(middelware.UserContext)
	if !ok {
		return 0
	}

	return userCtx.ID
}

func WebConn(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка WebSocket:", err)
		return
	}

	GlobalHub.mutex.Lock()
	if oldWs, exists := GlobalHub.clients[userID]; exists {
		oldWs.Close()
	}
	GlobalHub.clients[userID] = ws
	GlobalHub.mutex.Unlock()

	defer func() {
		GlobalHub.mutex.Lock()
		delete(GlobalHub.clients, userID)
		GlobalHub.mutex.Unlock()
		ws.Close()
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

// ХЕНДЛЕР: Вывод списка уведомлений
func GetNotificationsList(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notificationsFromDB, err := repository.GetNotificationPoId(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": notificationsFromDB,
	})
}

// ХЕНДЛЕР: Пометка ОДНОГО уведомления как прочитанного
func MarkSingleNotificationRead(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	notifIDStr := pathParts[len(pathParts)-1]
	notifID, err := strconv.Atoi(notifIDStr)
	if err != nil {
		http.Error(w, "Неверный ID уведомления", http.StatusBadRequest)
		return
	}

	err = repository.MarkNotificationAsRead(notifID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
