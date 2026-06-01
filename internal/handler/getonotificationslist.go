package handler

import (
	"net/http"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetNotificationsList(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	if !existsID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	notificationsFromDB, err := repository.GetNotificationPoId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": notificationsFromDB})
}
