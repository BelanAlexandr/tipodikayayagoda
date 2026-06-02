package routes

import (
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/middleware"

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

	auth := r.Group("/")
	auth.Use(middleware.RoleMiddleware())
	{
		auth.GET("/index", handler.IndexHandlerShow)
		auth.GET("/api/index", handler.IndexHandler)

		auth.GET("/product/:id", handler.ProductShow)
		auth.GET("/api/product/:id", handler.Product)

		auth.GET("/api/notifications/list", handler.GetNotificationsList)
		auth.POST("/api/notifications/read", handler.MarkSingleNotificationRead)
		auth.GET("/ws", handler.WebConn)

		auth.GET("/addproduct", handler.AddProductHandlerShow)
		auth.POST("/api/addproduct", handler.AddProductHandler)
		auth.GET("/adduser", handler.AdminRegisterShow)
		auth.POST("/api/adduser", handler.AdminRegister)

		auth.PUT("/api/product/edit/:id", handler.UpdateProductHandler)
		auth.DELETE("/api/product/delete/:id", handler.DeleteProductHandler)
		auth.POST("/api/uploadimage/:id", handler.UploadImageHandler)
		auth.POST("/api/category/add", handler.AddCategoryHandler)
		auth.GET("/api/sellers", handler.GetSeller)

		auth.GET("/addseller", handler.SellerOfferShow)

		auth.POST("/api/addseller", handler.SellerOffer)
		auth.GET("/api/addseller/all", handler.AllProd)
		auth.PUT("/api/offer/update", handler.OfferUpdate)

		auth.GET("/api/categories", handler.CategoriesListHandler)

		auth.POST("/api/product/buy/:id", handler.BuyProductHandler)
	}

	return r
}
