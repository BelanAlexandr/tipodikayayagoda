package handler

import (
	"encoding/json"
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
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"IsAdmin": "false",
	})
}

func Register(c *gin.Context) {

	var req models.User

	json.NewDecoder(r.Body).Decode(&req)
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)

	if req.Role == models.Roles.AdminID {
		http.Error(w, "You cannot register as admin", 403)
		return
	}
	err := service.Register(req, models.Roles.ClientID)
	if err != nil {
		http.Error(w, "Error registering user", 500)
		return
	}
	http.Redirect(w, r, "/login", 302)

}
