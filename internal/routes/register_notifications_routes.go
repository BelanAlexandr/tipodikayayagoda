package routes

import (
	"tipodikayayagoda/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(api *gin.RouterGroup) {

	notifications := api.Group("/notifications")
	{
		notifications.GET("/list", handler.GetNotificationsList)
		notifications.POST("/read/:id", handler.MarkSingleNotificationRead)
	}
}
