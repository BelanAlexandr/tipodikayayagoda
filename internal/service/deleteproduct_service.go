package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func DeleteProd(id int, userRole int) error {
	if userRole == models.RoleAdmin {
		img := repository.GetImageURL(id)
		utils.DeleteFileByURL(img)
		return repository.DeleteProd(id)
	}
	return nil
}
