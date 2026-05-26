package handler

import (
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/middelware"
	"tipodikayayagoda/internal/repository"
)

func MarkSingleNotificationRead(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(middelware.UserKey).(middelware.UserContext)
	if user.ID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	notifIDStr := pathParts[len(pathParts)-1]
	notifID, err := strconv.Atoi(notifIDStr)
	if err != nil {
		http.Error(w, "Неверный ID уведомления", http.StatusBadRequest)
		return
	}

	err = repository.MarkNotificationAsRead(notifID, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
