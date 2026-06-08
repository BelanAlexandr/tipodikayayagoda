package handler

import (
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AllProd(c *gin.Context) {
	_, existsRole := c.Get("userRole")
	if !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	searchQuery := c.Query("search")
	prod, err := service.AllProd(searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prod == nil {
		prod = []models.Product{}
	}

	for i := range prod {
		categoryName, err := repository.GetCategoryID(prod[i].Category_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		prod[i].Category = categoryName
	}

	c.JSON(http.StatusOK, prod)
}
