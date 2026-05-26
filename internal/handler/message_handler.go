package handler

import (
	"time"
	"tipodikayayagoda/internal/repository"
)

func Message(user_id int) {
	time.Sleep(5 * time.Second)
	repository.AddNotification(user_id, "Заказ отправлен")
	GlobalHub.SendToUser(user_id, []byte("Заказ отправлен"))
	time.Sleep(5 * time.Second)
	repository.AddNotification(user_id, "Заказ в пути")
	GlobalHub.SendToUser(user_id, []byte("Заказ в пути"))
	time.Sleep(5 * time.Second)
	repository.AddNotification(user_id, "Заказ на пункте выдачи")
	GlobalHub.SendToUser(user_id, []byte("Заказ на пункте выдачи"))
}
