package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {

	http.SetCookie(w, &http.Cookie{
		Name:   "tokenn",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // удалить cookie
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
