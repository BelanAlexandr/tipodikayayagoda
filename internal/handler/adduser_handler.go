package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func AdminRegisterShow(c *gin.Context) {

	tmpl, err := template.ParseFiles("internal/templates/registr.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"IsAdmin": "true",
	})
}

func AdminRegister(c *gin.Context) {
	user := r.Context().Value(middleware.UserKey).(middleware.UserContext)
	var req models.User

	json.NewDecoder(r.Body).Decode(&req)
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)

	err := service.Register(req, user.Role)
	if err != nil {
		http.Error(w, "Error registering user", 500)
		return
	}
	http.Redirect(w, r, "/index", 302)

}
