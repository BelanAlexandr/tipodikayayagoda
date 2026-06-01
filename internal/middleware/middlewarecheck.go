package middleware

import (
	"net/http"

	"tipodikayayagoda/internal/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Cookie("tokenn")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		claims, err := utils.ValidateToken(cookie)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
			c.Abort()
			return
		}

		userRole := claims["role"]

		hasAccess := false
		for _, role := range allowedRoles {
			if userRole == role {
				hasAccess = true
				break
			}
		}

		if !hasAccess {
			c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет доступа к этому ресурсу"})
			c.Abort()
			return
		}

		c.Set("userID", claims["id"])
		c.Set("userRole", claims["role"])

		c.Next()
	}
}
