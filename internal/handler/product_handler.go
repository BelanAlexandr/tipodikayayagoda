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
	id := strings.TrimPrefix(r.URL.Path, "/api/product/")
	idd, err := strconv.Atoi(id)

	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	product := service.GetProdPoID(idd, user.Role, user.ID)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(product)
}
