package rout

import (
	"net/http"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/models"
)

func Routes() {
	http.HandleFunc("/register", handler.RegisterShow)
	http.HandleFunc("/api/register", handler.Register)
	http.HandleFunc("/login", handler.LoginShow)
	http.HandleFunc("/api/login", handler.Login)
	http.HandleFunc("/index", middelware.RoleMiddleware(models.RoleClient, models.RoleSeller, models.RoleAdmin)(handler.IndexHandlerShow))
	http.HandleFunc("/api/index", middelware.RoleMiddleware(models.RoleClient, models.RoleSeller, models.RoleAdmin)(handler.IndexHandler))
	http.HandleFunc("/addproduct", middelware.RoleMiddleware(models.RoleSeller, models.RoleAdmin)(handler.AddProductHandlerShow))
	http.HandleFunc("/api/addproduct", middelware.RoleMiddleware(models.RoleSeller, models.RoleAdmin)(handler.AddProductHandler))
	http.HandleFunc("/adduser", middelware.RoleMiddleware(models.RoleAdmin)(handler.AdminRegisterShow))
	http.HandleFunc("/api/adduser", middelware.RoleMiddleware(models.RoleAdmin)(handler.AdminRegister))
	http.HandleFunc("/product/", middelware.RoleMiddleware(models.RoleClient, models.RoleSeller, models.RoleAdmin)(handler.ProductShow))
	http.HandleFunc("/api/product/", middelware.RoleMiddleware(models.RoleClient, models.RoleSeller, models.RoleAdmin)(handler.Product))
	http.HandleFunc("/api/product/edit/", middelware.RoleMiddleware(models.RoleSeller, models.RoleAdmin)(handler.UpdateProductHandler))
	http.HandleFunc("/api/product/delete/", middelware.RoleMiddleware(models.RoleAdmin)(handler.DeleteProductHandler))
	http.HandleFunc("/api/product/buy/", middelware.RoleMiddleware(models.RoleClient)(handler.BuyProductHandler))
	http.HandleFunc("/api/uploadimage/", middelware.RoleMiddleware(models.RoleSeller, models.RoleAdmin)(handler.UploadImageHandler))
	http.HandleFunc("/api/categories", middelware.RoleMiddleware(models.RoleAdmin, models.RoleSeller)(handler.CategoriesListHandler))
	http.HandleFunc("/api/category/add", middelware.RoleMiddleware(models.RoleAdmin)(handler.AddCategoryHandler))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/logout", handler.LogoutHandler)
}
