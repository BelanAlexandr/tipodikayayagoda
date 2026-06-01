package handler

import (
	"encoding/json"
	"net/http"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func GetSeller(c *gin.Context) {
	sellers, err := service.GetSellerId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sellers)
}
