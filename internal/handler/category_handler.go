package handler

import (
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AddCategoryHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userrole, ok := userRoleValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.AddCategory(userrole, req.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}
func CategoriesListHandler(c *gin.Context) {
	_, existsRole := c.Get("userRole")
	if !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	categories, err := service.GetCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
