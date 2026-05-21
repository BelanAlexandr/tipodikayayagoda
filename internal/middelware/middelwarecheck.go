package middelware

import (
	"context"
	"net/http"
	"tipodikayayagoda/internal/utils"
)

const RoleKey string = "role"

type UserContext struct {
	ID   int
	Role string
}

func RoleMiddleware(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("token")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			claims, err := utils.ValidateToken(cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			role, ok := claims["role"].(string)
			if !ok {
				http.Error(w, "Invalid role", http.StatusForbidden)
				return
			}
			id := int(claims["id"].(float64))

			allowed := false
			for _, r := range allowedRoles {
				if role == r {
					allowed = true
					break
				}
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			var us UserContext
			us.ID = id
			us.Role = role
			ctx := context.WithValue(r.Context(), RoleKey, us)
			next(w, r.WithContext(ctx))
		}
	}
}
