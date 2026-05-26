package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/service"
)

func LoginShow(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)
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
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	token, err := service.Login(req.Login, req.Password)
	if err != nil {
		http.Error(w, "error logging in user", http.StatusUnauthorized)
		fmt.Println("Login error:", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
		Path:  "/",
	})

	http.Redirect(w, r, "/index", http.StatusFound)
}
