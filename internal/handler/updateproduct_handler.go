package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"
)

func UpdateProductHandlerShow(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/addproduct.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"Role":   user.Role,
		"UserID": user.ID,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
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
	fmt.Println("Updating product with ID:", product)
	err = service.UpdateProd(product, user.ID, user.Role)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "updated",
	})
}
