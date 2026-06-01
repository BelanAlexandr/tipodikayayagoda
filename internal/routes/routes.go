package routes

import (
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/models"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/register", handler.RegisterShow)
	r.POST("/api/register", handler.Register)
	r.GET("/login", handler.LoginShow)
	r.POST("/api/login", handler.Login)
	r.POST("/logout", handler.LogoutHandler)

	client := models.Roles.ClientID
	seller := models.Roles.SellerID
	admin := models.Roles.AdminID
	allRoles := r.Group("/")
	allRoles.Use(middleware.RoleMiddleware(client, seller, admin))
	{
		allRoles.GET("/index", handler.IndexHandlerShow)
		allRoles.GET("/api/index", handler.IndexHandler)

		allRoles.GET("/product", handler.ProductShow)
		allRoles.GET("/api/product", handler.Product)

		allRoles.GET("/api/notifications/list", handler.GetNotificationsList)
		allRoles.POST("/api/notifications/read", handler.MarkSingleNotificationRead)
		allRoles.GET("/ws", handler.WebConn)
	}

	adminOnly := r.Group("/")
	adminOnly.Use(middleware.RoleMiddleware(admin))
	{
		adminOnly.GET("/addproduct", handler.AddProductHandlerShow)
		adminOnly.POST("/api/addproduct", handler.AddProductHandler)
		adminOnly.GET("/adduser", handler.AdminRegisterShow)
		adminOnly.POST("/api/adduser", handler.AdminRegister)

		adminOnly.PUT("/api/product/edit", handler.UpdateProductHandler)
		adminOnly.DELETE("/api/product/delete/:id", handler.DeleteProductHandler)
		adminOnly.POST("/api/uploadimage", handler.UploadImageHandler)
		adminOnly.POST("/api/category/add", handler.AddCategoryHandler)
		adminOnly.GET("/api/sellers", handler.GetSeller)
	}

	sellerOnly := r.Group("/")
	sellerOnly.Use(middleware.RoleMiddleware(seller))
	{
		sellerOnly.GET("/addseller", handler.SellerOfferShow)

		sellerOnly.POST("/api/addseller", handler.SellerOffer)
		sellerOnly.GET("/api/addseller/all", handler.AllProd)
		sellerOnly.PUT("/api/offer/update", handler.OfferUpdate)
	}

	adminSeller := r.Group("/")
	adminSeller.Use(middleware.RoleMiddleware(admin, seller))
	{
		adminSeller.GET("/api/categories", handler.CategoriesListHandler)
	}

	clientOnly := r.Group("/")
	clientOnly.Use(middleware.RoleMiddleware(client))
	{
		clientOnly.POST("/api/product/buy/:id", handler.BuyProductHandler)
	}

	return r
}
