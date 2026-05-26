package handler

import (
	"encoding/json"
	"net/http"
	"tipodikayayagoda/internal/service"
)

func GetSeller(w http.ResponseWriter, r *http.Request) {
	sellers, err := service.GetSellerId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sellers)
}
