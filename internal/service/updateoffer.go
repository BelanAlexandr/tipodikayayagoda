package service

import "tipodikayayagoda/internal/repository"

func UpdateOffer(id int, price float64, count int) error {
	return repository.UpdateOffer(id, price, count)
}
