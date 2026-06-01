package handler

import (
	"encoding/json"
	"net/http"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetNotificationsList(c *gin.Context) {
	user, _ := r.Context().Value(middleware.UserKey).(middleware.UserContext)
	if user.ID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notificationsFromDB, err := repository.GetNotificationPoId(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": notificationsFromDB,
	})
}
