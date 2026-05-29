package service

import "tipodikayayagoda/internal/repository"

func AddOffer(id, count int, price float64, user_id int) error {
	return repository.AddOffer(id, count, price, user_id)
}
