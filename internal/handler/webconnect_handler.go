package handler

import (
	"log"
	"net/http"
	"tipodikayayagoda/internal/middleware"

	"github.com/gin-gonic/gin"
)

func WebConn(c *gin.Context) {
	user, _ := r.Context().Value(middleware.UserKey).(middleware.UserContext)
	if user.ID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка WebSocket:", err)
		return
	}

	GlobalHub.mutex.Lock()
	if oldWs, exists := GlobalHub.clients[user.ID]; exists {
		oldWs.Close()
	}
	GlobalHub.clients[user.ID] = ws
	GlobalHub.mutex.Unlock()

	defer func() {
		GlobalHub.mutex.Lock()
		delete(GlobalHub.clients, user.ID)
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
