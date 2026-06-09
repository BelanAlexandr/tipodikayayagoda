package routes

import (
	"tipodikayayagoda/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(api *gin.RouterGroup) {
	products := api.Group("/product")
	{
		products.GET("/", handler.Product)
		products.PUT("/edit/:id", handler.UpdateProductHandler)
		products.DELETE("/delete/:id", handler.DeleteProductHandler)
		products.POST("/buy/:id", handler.BuyProductHandler)
	}
	api.GET("/index", handler.IndexHandler)
	api.GET("/categories", handler.CategoriesListHandler)
	api.POST("/category/add", handler.AddCategoryHandler)
	api.POST("/addproduct", handler.AddProductHandler)
	api.POST("/uploadimage/:id", handler.UploadImageHandler)
	api.GET("/analytics", handler.GetAnalyticsData)

}
