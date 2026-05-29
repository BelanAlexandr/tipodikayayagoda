package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"
)

func ProductShow(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("internal/templates/product.html")

	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	data := map[string]any{
		"UserID":   user.ID,
		"IsAdmin":  user.Role == models.RoleAdmin,
		"IsSeller": user.Role == models.RoleSeller,
		"CanBuy":   user.Role == models.RoleClient,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
func Product(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	idStr = strings.Trim(idStr, "/")
	idd, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	product, err := service.GetProdPoID(idd, user.Role, user.ID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
