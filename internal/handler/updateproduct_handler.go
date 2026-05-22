package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/edit/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "bad body", 400)
		return
	}

	product.ID = id
	err = service.UpdateProd(product, user.ID, user.Role)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}
