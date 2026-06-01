package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebConn(c *gin.Context) {

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDValue.(int)
	if !ok || userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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
