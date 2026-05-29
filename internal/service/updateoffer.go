package service

import "tipodikayayagoda/internal/repository"

func UpdateOffer(id int, price float64, count int, user_id int) error {
	return repository.UpdateOffer(id, price, count, user_id)
}
