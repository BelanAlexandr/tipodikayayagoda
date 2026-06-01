package handler

import (
	"encoding/json"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AllProd(c *gin.Context) {
	prod := service.AllProd()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}
