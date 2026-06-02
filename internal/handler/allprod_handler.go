package handler

import (
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AllProd(c *gin.Context) {
	searchQuery := c.Query("search")
	prod, err := service.AllProd(searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prod == nil {
		prod = []models.Product{}
	}
	c.JSON(http.StatusOK, prod)
}
