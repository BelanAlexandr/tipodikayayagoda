package handler

import (
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func UpdateProductHandler(c *gin.Context) {

	idStr := c.Param("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Name        string `json:"name"`
		CategoryID  int    `json:"category_id"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad body"})
		return
	}

	err = service.UpdateProd(
		productID,
		req.Name,
		req.Description,
		req.CategoryID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}
