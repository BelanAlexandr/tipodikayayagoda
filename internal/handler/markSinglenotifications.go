package handler

import (
	"log"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func MarkSingleNotificationRead(c *gin.Context) {
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

	notifIDStr := c.Param("id")

	notifID, err := strconv.Atoi(notifIDStr)
	if err != nil || notifID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID уведомления"})
		return
	}

	err = repository.MarkNotificationAsRead(notifID, userID)
	if err != nil {
		log.Println("Ошибка при чтении уведомления в БД:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при обновлении статуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
