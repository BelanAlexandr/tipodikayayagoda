package handler

import (
	"encoding/json"
	"net/http"
	"tipodikayayagoda/internal/service"
)

func AllProd(w http.ResponseWriter, r *http.Request) {
	prod := service.AllProd()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
	return
}
