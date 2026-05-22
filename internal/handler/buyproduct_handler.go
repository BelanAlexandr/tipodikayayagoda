package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/service"
)

func BuyProductHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/buy/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}
	fmt.Println("Buying product with ID:", id)
	var req struct {
		Count int `json:"count"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}
	err = service.BuyProduct(id, user.Role, req.Count)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "bought",
	})
}
