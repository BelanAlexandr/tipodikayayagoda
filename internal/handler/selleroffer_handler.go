package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func SellerOfferShow(c *gin.Context) {
	_, exists := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !exists || !existsRole || userRoleValue.(int) != models.RoleSeller {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	t, err := template.ParseFiles("internal/templates/addseller.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки шаблона"})
		return
	}

	t.Execute(c.Writer, nil)
}

func SellerOffer(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleSeller {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	idStr := c.Query("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Price float64 `json:"price"`
		Count int     `json:"count"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка разбора JSON: " + err.Error()})
		return
	}

	err = service.AddOffer(productID, req.Count, req.Price, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка добавления"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product created",
	})
}
