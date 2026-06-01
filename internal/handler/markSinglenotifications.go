package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middleware"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func MarkSingleNotificationRead(c *gin.Context) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	user, _ := r.Context().Value(middleware.UserKey).(middleware.UserContext)
	if user.ID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	notifIDStr := strings.TrimPrefix(r.URL.Path, "/api/notifications/read/")
	notifIDStr = strings.TrimSpace(notifIDStr)

	notifIDStr = strings.TrimSuffix(notifIDStr, "/")

	notifID, err := strconv.Atoi(notifIDStr)
	if err != nil || notifID <= 0 {
		http.Error(w, "Неверный ID уведомления", http.StatusBadRequest)
		return
	}

	err = repository.MarkNotificationAsRead(notifID, user.ID)
	if err != nil {
		log.Println("Ошибка при чтении уведомления в БД:", err)
		http.Error(w, "Ошибка сервера при обновлении статуса", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
