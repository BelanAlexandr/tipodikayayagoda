package handler

import (
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func AnalyticsHandlerShow(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	userID, ok1 := userIDValue.(int)
	userRole, ok2 := userRoleValue.(int)
	if !ok1 || !ok2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID или роли пользователя"})
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/analytics.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, map[string]any{
		"UserID":   userID,
		"IsAdmin":  userRole == models.RoleAdmin,
		"IsSeller": userRole == models.RoleSeller,
	})
}

func GetAnalyticsData(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	userID, ok1 := userIDValue.(int)
	userRole, ok2 := userRoleValue.(int)
	if !ok1 || !ok2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID или роли пользователя"})
		return
	}

	if userRole == models.RoleAdmin {
		stats, err := repository.GetAdminAnalytics()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки статистики админа"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"type": "admin", "data": stats})
		return
	}

	if userRole == models.RoleSeller {
		stats, err := repository.GetSellerAnalytics(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"type": "seller", "data": stats})
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
}
