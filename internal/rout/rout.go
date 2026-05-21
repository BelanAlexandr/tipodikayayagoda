package rout

import (
	"net/http"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/middelware"
)

func Routes() {
	http.HandleFunc("/register", handler.RegisterShow)
	http.HandleFunc("/api/register", handler.Register)
	http.HandleFunc("/login", handler.LoginShow)
	http.HandleFunc("/api/login", handler.Login)
	http.HandleFunc("/index", middelware.RoleMiddleware("client", "seller", "admin")(handler.IndexHandlerShow))
}
