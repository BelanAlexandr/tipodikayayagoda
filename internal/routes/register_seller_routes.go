package routes

import (
	"tipodikayayagoda/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterSellerRoutes(api *gin.RouterGroup) {
	api.GET("/sellers", handler.GetSeller)

	sellers := api.Group("/addseller")

	sellers.POST("/", handler.SellerOffer)
	sellers.GET("/all", handler.AllProd)

	api.PUT("/offer/update/:id", handler.OfferUpdate)
	api.POST("/adduser", handler.AdminRegister)
}
