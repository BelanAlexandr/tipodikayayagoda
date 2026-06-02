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
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка WebSocket апгрейда:", err)
		return
	}

	GlobalHub.Register(userID, ws)

	defer func() {
		GlobalHub.Unregister(userID)
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {

			break
		}
	}
}
