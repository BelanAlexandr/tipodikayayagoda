package rout

import (
	"net/http"
	"tipodikayayagoda/internal/handler"
)

func Routes() {
	http.HandleFunc("/register", handler.RegisterShow)
	http.HandleFunc("/api/register", handler.Register)
	http.HandleFunc("/login", handler.LoginShow)
	http.HandleFunc("/api/login", handler.Login)
}
