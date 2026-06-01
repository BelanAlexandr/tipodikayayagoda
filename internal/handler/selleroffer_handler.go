package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func SellerOfferShow(c *gin.Context) {

	t, _ := template.ParseFiles("internal/templates/addseller.html")
	t.Execute(w, nil)

}
func SellerOffer(c *gin.Context) {
	user, _ := r.Context().Value(middleware.UserKey).(middleware.UserContext)

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
