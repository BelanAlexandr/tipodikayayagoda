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

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}
	fmt.Println("Deleting product with ID:", id)
	err = service.DeleteProd(id, user.Role)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted",
	})
}
