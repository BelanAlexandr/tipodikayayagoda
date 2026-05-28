package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
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
		"UserID":        user.ID,
		"IsAdmin":       user.Role == models.RoleAdmin,
		"IsSeller":      user.Role == models.RoleSeller,
		"CanBuy":        user.Role == models.RoleClient,
		"CanAddUser":    user.Role == models.RoleAdmin,
		"CanAddProduct": user.Role == models.RoleAdmin || user.Role == models.RoleSeller,
	})
}

type ProductsResponse struct {
	Products   []models.ProductCard `json:"products"`
	TotalCount int                  `json:"totalCount"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)

	search := r.URL.Query().Get("search")
	sort := r.URL.Query().Get("sort")

	categoryID, err := strconv.Atoi(r.URL.Query().Get("category"))
	if err != nil || categoryID < 0 {
		categoryID = 0
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 6
	}

	products, totalCount, err := service.GetProducts(user.Role, user.ID, search, page, limit, sort, categoryID)
	if err != nil {
		http.Error(w, "error loading products", 500)
		return
	}

	response := ProductsResponse{
		Products:   products,
		TotalCount: totalCount,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
