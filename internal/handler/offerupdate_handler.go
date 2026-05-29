package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/service"
)

func OfferUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/offer/update/")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID товара", http.StatusBadRequest)
		return
	}
	var req struct {
		Price float64 `json:"price"`
		Count int     `json:"count"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка разбора JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdateOffer(productID, req.Price, req.Count)
	if err != nil {
		http.Error(w, "Ошибка обновления в базе данных: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Предложение успешно обновлено",
	})
}
