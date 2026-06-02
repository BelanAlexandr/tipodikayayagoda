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

	RegisterPageRoutes(auth)

	auth.GET("/ws", handler.WebConn)

	api := auth.Group("/api")

	RegisterProductRoutes(api)
	RegisterSellerRoutes(api)
	RegisterNotificationRoutes(api)
	return r
}
