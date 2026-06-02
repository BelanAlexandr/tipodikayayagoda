package handler

import (
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterShow(c *gin.Context) {
	tmpl, err := template.ParseFiles("internal/templates/registr.html")
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, map[string]any{
		"IsAdmin": "false",
	})
}

func Register(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON: " + err.Error()})
		return
	}

	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)

	if req.Role == models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You cannot register as admin"})
		return
	}

	err := service.Register(req, models.RoleClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}
