package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/service"
)

func SellerOfferShow(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("internal/templates/addseller.html")
	t.Execute(w, nil)

}
func SellerOffer(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(middelware.UserKey).(middelware.UserContext)

	idStr := strings.TrimPrefix(r.URL.Path, "/api/addseller/")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
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

	fmt.Printf("Received Price: %f, Count: %d\n", req.Price, req.Count)

	err = service.AddOffer(productID, req.Count, req.Price, user.ID)
	if err != nil {
		http.Error(w, "Ошибка добавления", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "product created",
	})
}
