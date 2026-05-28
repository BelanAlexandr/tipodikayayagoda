package handler

import (
	"encoding/json"
	"time"
	"tipodikayayagoda/internal/repository"
)

type NotificationResponse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func Message(user_id int) {
	time.Sleep(5 * time.Second)
	notifID, _ := repository.AddNotification(user_id, "Заказ отправлен")
	resp, _ := json.Marshal(NotificationResponse{ID: notifID, Text: "Заказ отправлен"})
	GlobalHub.SendToUser(user_id, resp)
	time.Sleep(5 * time.Second)
	notifID, _ = repository.AddNotification(user_id, "Заказ в пути")
	resp, _ = json.Marshal(NotificationResponse{ID: notifID, Text: "Заказ в пути"})
	GlobalHub.SendToUser(user_id, resp)
	time.Sleep(5 * time.Second)
	notifID, _ = repository.AddNotification(user_id, "Заказ на пункте")
	resp, _ = json.Marshal(NotificationResponse{ID: notifID, Text: "Заказ на пункте"})
	GlobalHub.SendToUser(user_id, resp)
}
