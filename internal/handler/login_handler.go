package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/service"
)

func LoginShow(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)

	return
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)
	err := service.Login(req.Login, req.Password)

	if err != nil {
		http.Error(w, "Error logging in user", 500)
		return
	}
	http.Redirect(w, r, "/login", 302)

}
