package handler

import (
	"net/http"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AllProd(c *gin.Context) {
	prod := service.AllProd()
	c.JSON(http.StatusOK, prod)
}
