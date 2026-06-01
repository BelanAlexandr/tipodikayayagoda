package handler

import (
	"net/http"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func GetSeller(c *gin.Context) {
	sellers, err := service.GetSellerId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sellers)

}
