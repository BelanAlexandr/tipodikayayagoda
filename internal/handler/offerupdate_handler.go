package handler

import (
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func OfferUpdate(c *gin.Context) {
	userRoleValue, _ := c.Get("userRole")
	userIDValue, exists := c.Get("userID")
	if !exists || userRoleValue.(int) != models.RoleSeller {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	idStr := c.Param("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID товара"})
		return
	}

	var req struct {
		Price float64 `json:"price" binding:"required"`
		Count int     `json:"count" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка разбора JSON: " + err.Error()})
		return
	}

	err = service.UpdateOffer(productID, req.Price, req.Count, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления в базе данных: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Предложение успешно обновлено",
	})
}
