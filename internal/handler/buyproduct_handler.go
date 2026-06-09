package handler

import (
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func BuyProductHandler(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole || userRoleValue.(int) != models.RoleClient {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userrole, ok := userRoleValue.(int)
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Count    int `json:"count"`
		SellerID int `json:"seller_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.BuyProduct(id, userrole, req.Count, req.SellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go repository.TrackEvent(models.AnalyticsEvent{
		UserID:    &userID,
		ProductID: &id,
		SellerID:  &req.SellerID,
		EventType: models.EventPurchase,
		Quantity:  req.Count,
	})
	go Message(userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "product_bought",
	})
}
