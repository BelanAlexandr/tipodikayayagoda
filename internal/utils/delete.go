package utils

import (
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
)

func DeleteFileByURL(minioClient *minio.Client, fileURL string) {
	if fileURL == "" || strings.Contains(fileURL, "default.png") {
		return
	}

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Printf("Ошибка парсинга URL картинки: %v", err)
		return
	}

	path := strings.TrimPrefix(parsedURL.Path, "/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) < 2 {
		log.Printf("Некорректный путь к объекту в URL: %s", path)
		return
	}

	bucketName := parts[0]
	objectName := parts[1]

	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Printf("Ошибка удаления файла %s из MinIO бакета %s: %v", objectName, bucketName, err)
		return
	}

	log.Printf("Файл %s успешно удален из MinIO бакета %s", objectName, bucketName)
}
