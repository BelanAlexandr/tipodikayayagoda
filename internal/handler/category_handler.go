package handler

import (
	"encoding/json"
	"net/http"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AddCategoryHandler(c *gin.Context) {
	user := r.Context().Value(middleware.UserKey).(middleware.UserContext)

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	err := service.AddCategory(user.Role, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func CategoriesListHandler(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
