package handler

import "net/http"

func IndexHandlerShow(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}
