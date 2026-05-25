package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/service"
)

func IndexHandlerShow(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if !ok {
		user.Role = 0
	}

	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"Role":   user.Role,
		"UserID": user.ID,
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)

	products, err := service.GetProducts(user.Role, user.ID)
	if err != nil {
		http.Error(w, "error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
