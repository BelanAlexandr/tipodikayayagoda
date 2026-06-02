package middleware

import (
	"net/http"

	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Cookie("tokenn")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		claims, err := utils.ValidateToken(cookie)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
			c.Abort()
			return
		}
		idFloat, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный формат ID в токене"})
			c.Abort()
			return
		}
		role, err := repository.GetUserByID(int(idFloat))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
			c.Abort()
			return
		}
		c.Set("userID", int(idFloat))
		c.Set("userRole", role)

		c.Next()
	}
}
