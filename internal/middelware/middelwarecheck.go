package middelware

import (
	"context"
	"net/http"
	"tipodikayayagoda/internal/utils"
)

const UserKey string = "user"

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

			roleVal, ok := claims["role"].(string)
			if !ok {
				http.Error(w, "Invalid role", http.StatusForbidden)
				return
			}

			idVal, ok := claims["id"]
			if !ok {
				http.Error(w, "Invalid token (no id)", http.StatusForbidden)
				return
			}

			idFloat, ok := idVal.(float64)
			if !ok {
				http.Error(w, "Invalid id type", http.StatusForbidden)
				return
			}

			id := int(idFloat)

			allowed := false
			for _, r := range allowedRoles {
				if roleVal == r {
					allowed = true
					break
				}
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			user := UserContext{
				ID:   id,
				Role: roleVal,
			}

			ctx := context.WithValue(r.Context(), UserKey, user)
			next(w, r.WithContext(ctx))
		}
	}
}
