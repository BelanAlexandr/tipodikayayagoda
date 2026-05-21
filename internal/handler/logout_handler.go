package handler

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // удалить cookie
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
