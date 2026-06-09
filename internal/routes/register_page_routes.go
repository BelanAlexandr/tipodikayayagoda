package routes

import (
	"tipodikayayagoda/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPageRoutes(rg *gin.RouterGroup) {
	rg.GET("/index", handler.IndexHandlerShow)
	rg.GET("/product", handler.ProductShow)
	rg.GET("/addproduct", handler.AddProductHandlerShow)
	rg.GET("/adduser", handler.AdminRegisterShow)
	rg.GET("/addseller", handler.SellerOfferShow)
	rg.GET("/analytics", handler.AnalyticsHandlerShow)

}
