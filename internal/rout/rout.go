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
	http.HandleFunc("/api/index", middelware.RoleMiddleware("client", "seller", "admin")(handler.IndexHandler))
	http.HandleFunc("/addproduct", middelware.RoleMiddleware("seller", "admin")(handler.AddProductHandlerShow))
	http.HandleFunc("/api/addproduct", middelware.RoleMiddleware("seller", "admin")(handler.AddProductHandler))
	http.HandleFunc("/adduser", middelware.RoleMiddleware("admin")(handler.AdminRegisterShow))
	http.HandleFunc("/api/adduser", middelware.RoleMiddleware("admin")(handler.AdminRegister))
	http.HandleFunc("/product/", middelware.RoleMiddleware("client", "seller", "admin")(handler.ProductShow))
	http.HandleFunc("/api/product/", middelware.RoleMiddleware("client", "seller", "admin")(handler.Product))
	http.HandleFunc("/api/product/edit/", middelware.RoleMiddleware("seller", "admin")(handler.UpdateProductHandler))
	http.HandleFunc("/api/product/delete/", middelware.RoleMiddleware("admin")(handler.DeleteProductHandler))
	http.HandleFunc("/api/product/buy/", middelware.RoleMiddleware("client")(handler.BuyProductHandler))
	http.HandleFunc("/logout", handler.LogoutHandler)
}
