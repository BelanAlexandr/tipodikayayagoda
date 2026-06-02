package handler

import (
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AddProductHandlerShow(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userrole, ok := userRoleValue.(int)
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/addproduct.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]any{
		"UserID":            userID,
		"IsAdmin":           userrole == models.RoleAdmin,
		"IsSeller":          userrole == models.RoleSeller,
		"CanEditAnyProduct": userrole == models.RoleAdmin,
	}

	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func AddProductHandler(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		CategoryID  int    `json:"category_id"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.CategoryID <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "должна быть выбрана категория"})
		return
	}
	err := service.Addproduct(
		req.Name,
		req.Description,
		req.CategoryID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "prdocut_created",
	})
}
