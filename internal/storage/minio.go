package storage

import (
	"context"
	"log"
	"tipodikayayagoda/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Глобальная переменная для работы с MinIO из других пакетов
var MinioClient *minio.Client

const BucketName = "images"

func InitMinio(cfg *config.Config) {
	endpoint := cfg.Endpoint
	accessKeyID := cfg.AccesKey
	secretAccessKey := cfg.SecretAccesKey
	useSSL := false

	// Инициализируем клиент
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("Ошибка инициализации MinIO:", err)
	}

	MinioClient = client
	log.Println("Успешное подключение к MinIO")

	// Автоматически создаем бакет, если его не существует
	ctx := context.Background()
	err = MinioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Проверяем, может бакет уже создан
		exists, errBucketExists := MinioClient.BucketExists(ctx, BucketName)
		if errBucketExists == nil && exists {
			log.Printf("Бакет '%s' уже существует\n", BucketName)
		} else {
			log.Fatalln("Ошибка создания бакета:", err)
		}
	} else {
		log.Printf("Бакет '%s' успешно создан\n", BucketName)
	}

	// Устанавливаем политику "По публичной ссылке можно только читать файлы"
	// Это нужно, чтобы тег <img src="..."> на сайте смог отобразить картинку
	// Устанавливаем политику "Анонимные пользователи могут только читать файлы"
	policy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {"AWS": ["*"]},
            "Action": ["s3:GetObject"],
            "Resource": ["arn:aws:s3:::` + BucketName + `/*"]
        }
    ]
}`

	err = MinioClient.SetBucketPolicy(ctx, BucketName, policy)
	if err != nil {
		log.Println("Ошибка применения политики бакета:", err)
	} else {
		log.Println("Бакет успешно переведен в публичный режим (Read-Only)")
	}
	err = MinioClient.SetBucketPolicy(ctx, BucketName, policy)
	if err != nil {
		log.Println("Предупреждение: не удалось установить политику бакета:", err)
	}
}
