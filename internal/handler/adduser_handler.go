package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"
)

func AdminRegisterShow(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/registr.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"IsAdmin": "true",
	})

	return
}

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	var req models.User

	json.NewDecoder(r.Body).Decode(&req)
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	fmt.Println("Registering user:", user)
	err := service.Register(req, user.Role)
	if err != nil {
		http.Error(w, "Error registering user", 500)
		return
	}
	http.Redirect(w, r, "/index", 302)

}
