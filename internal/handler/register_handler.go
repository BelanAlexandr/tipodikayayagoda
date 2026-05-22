package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/service"
)

func RegisterShow(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/registr.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"IsAdmin": "false",
	})

	return
}

func Register(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	json.NewDecoder(r.Body).Decode(&req)
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	if req.Role == "admin" {
		http.Error(w, "You cannot register as admin", 403)
		return
	}
	err := service.Register(req.Login, req.Password, req.Role, "")
	if err != nil {
		http.Error(w, "Error registering user", 500)
		return
	}
	http.Redirect(w, r, "/login", 302)

}
