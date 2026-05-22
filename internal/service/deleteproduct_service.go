package service

import (
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func DeleteProd(id int, userRole string) error {
	if userRole == "admin" {
		img := repository.GetImageURL(id)
		utils.DeleteFileByURL(img)
		return repository.DeleteProd(id)
	}
	return nil
}
