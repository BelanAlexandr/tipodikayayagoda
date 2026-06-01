package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func ProductShow(c *gin.Context) {
	tmpl, err := template.ParseFiles("internal/templates/product.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки шаблона"})
		return
	}

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	uid, _ := userID.(int)
	role, _ := userRole.(int)

	data := map[string]any{
		"UserID":   uid,
		"IsAdmin":  role == models.Roles.AdminID,
		"IsSeller": role == models.Roles.SellerID,
		"CanBuy":   role == models.Roles.ClientID,
	}

	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template error"})
		return
	}
}

func Product(c *gin.Context) {

	idStr := c.Param("id")
	idd, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	uid, _ := userID.(int)
	role, _ := userRole.(int)

	product, err := service.GetProdPoID(idd, role, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
