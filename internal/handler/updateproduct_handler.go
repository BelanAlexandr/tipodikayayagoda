package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/service"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/edit/")

	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var req struct {
		Name        string `json:"name"`
		CategoryID  int    `json:"category_id"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	err = service.UpdateProd(
		productID,
		req.Name,
		req.Description,
		req.CategoryID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}
