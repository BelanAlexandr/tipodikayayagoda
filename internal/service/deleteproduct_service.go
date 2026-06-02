package service

import (
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"

	"github.com/minio/minio-go/v7" // Не забудь импортировать официальный пакет
)

func DeleteProd(minioClient *minio.Client, id int, userRole int) error {
	if userRole == models.RoleAdmin {

		img := repository.GetImageURL(id)

		utils.DeleteFileByURL(minioClient, img)

		return repository.DeleteProd(id)
	}

	return nil
}
