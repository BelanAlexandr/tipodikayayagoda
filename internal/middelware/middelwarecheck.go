package middelware

import (
	"fmt"
	"net/http"
)

func RoleMiddleware(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("token")
			fmt.Println(cookie, err)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			role, err := AuthCheck(cookie)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			for _, allowed := range allowedRoles {
				if role == allowed {
					next(w, r)
					return
				}
			}

			http.Error(w, "У вас нет прав доступа", http.StatusForbidden)
		}
	}
}
