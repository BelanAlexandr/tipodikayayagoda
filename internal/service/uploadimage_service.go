package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/storage" // Пакет, где инициализирован MinioClient и BucketName

	"github.com/minio/minio-go/v7"
)

func UploadImage(id int, file multipart.File, header *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	oldURL := repository.GetImageURL(id)
	if oldURL != "" {

		parts := strings.Split(oldURL, "/")
		if len(parts) > 0 {
			oldObjectName := parts[len(parts)-1]
			// Удаляем старый объект из бакета MinIO
			_ = storage.MinioClient.RemoveObject(ctx, storage.BucketName, oldObjectName, minio.RemoveObjectOptions{})
		}
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%d%s", id, time.Now().UnixNano(), ext)

	contentType := header.Header.Get("Content-Type")
	_, err := storage.MinioClient.PutObject(ctx, storage.BucketName, filename, file, header.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки в MinIO: %w", err)
	}

	newURL := fmt.Sprintf("http://localhost:9000/%s/%s", storage.BucketName, filename)

	err = repository.Updateimage(newURL, id)
	if err != nil {
		return "", fmt.Errorf("ошибка обновления картинки в БД: %w", err)
	}

	return newURL, nil
}
