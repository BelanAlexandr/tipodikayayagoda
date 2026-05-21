package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"
)

func AddProductHandlerShow(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/templates/addproduct.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = service.Addproduct(product, user.Role, user.ID)
	if err != nil {
		http.Error(w, "error adding product", http.StatusInternalServerError)
		return
	}
}
