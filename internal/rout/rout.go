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
	http.HandleFunc("/index", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.IndexHandlerShow))
	http.HandleFunc("/api/index", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.IndexHandler))
	http.HandleFunc("/addproduct", middelware.RoleMiddleware(models.Roles.AdminID)(handler.AddProductHandlerShow))
	http.HandleFunc("/api/addproduct", middelware.RoleMiddleware(models.Roles.AdminID)(handler.AddProductHandler))
	http.HandleFunc("/adduser", middelware.RoleMiddleware(models.Roles.AdminID)(handler.AdminRegisterShow))
	http.HandleFunc("/api/adduser", middelware.RoleMiddleware(models.Roles.AdminID)(handler.AdminRegister))
	http.HandleFunc("/product/", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.ProductShow))
	http.HandleFunc("/api/product/", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.Product))
	http.HandleFunc("/api/product/edit/", middelware.RoleMiddleware(models.Roles.AdminID)(handler.UpdateProductHandler))
	http.HandleFunc("/api/product/delete/", middelware.RoleMiddleware(models.Roles.AdminID)(handler.DeleteProductHandler))
	http.HandleFunc("/api/product/buy/", middelware.RoleMiddleware(models.Roles.ClientID)(handler.BuyProductHandler))
	http.HandleFunc("/api/uploadimage/", middelware.RoleMiddleware(models.Roles.AdminID)(handler.UploadImageHandler))
	http.HandleFunc("/api/categories", middelware.RoleMiddleware(models.Roles.AdminID, models.Roles.SellerID)(handler.CategoriesListHandler))
	http.HandleFunc("/api/category/add", middelware.RoleMiddleware(models.Roles.AdminID)(handler.AddCategoryHandler))
	http.HandleFunc("/api/sellers", middelware.RoleMiddleware(models.Roles.AdminID)(handler.GetSeller))
	http.HandleFunc("/addseller", middelware.RoleMiddleware(models.Roles.SellerID)(handler.SellerOfferShow))
	http.HandleFunc("/api/addseller/", middelware.RoleMiddleware(models.Roles.SellerID)(handler.SellerOffer))
	http.HandleFunc("/api/addseller", middelware.RoleMiddleware(models.Roles.SellerID)(handler.AllProd))
	http.HandleFunc("/api/offer/update/", middelware.RoleMiddleware(models.Roles.SellerID)(handler.OfferUpdate))

	http.HandleFunc("/api/notifications/list", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.GetNotificationsList))
	http.HandleFunc("/api/notifications/read/", middelware.RoleMiddleware(models.Roles.ClientID, models.Roles.SellerID, models.Roles.AdminID)(handler.MarkSingleNotificationRead))
	http.HandleFunc("/ws", middelware.RoleMiddleware(models.Roles.AdminID, models.Roles.ClientID, models.Roles.SellerID)(handler.WebConn))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/logout", handler.LogoutHandler)
}
