package handler

import (
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AdminRegisterShow(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole || userRoleValue.(int) != models.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	tmpl, err := template.ParseFiles("internal/templates/registr.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, map[string]any{
		"IsAdmin": "true",
	})
}

func AdminRegister(c *gin.Context) {
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
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	_, err := service.Register(req, userrole)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, "/index")

}
