package service

import "tipodikayayagoda/internal/repository"

func DeleteProd(id int, userRole string) error {
	if userRole == "admin" {
		return repository.DeleteProd(id)
	}
	return nil
}
