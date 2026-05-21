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

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := service.Login(req.Login, req.Password)
	if err != nil {
		http.Error(w, "error logging in user", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	})

	http.Redirect(w, r, "/index", http.StatusFound)
}
