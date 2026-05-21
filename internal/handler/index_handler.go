package handler

import (
	"html/template"
	"net/http"
	"tipodikayayagoda/internal/middelware"
)

func IndexHandlerShow(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value(middelware.RoleKey).(string)
	if !ok {
		role = "unknown"
	}

	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, map[string]any{
		"Role": role,
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}
