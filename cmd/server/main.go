package main

import (
	"fmt"
	"log"
	"tipodikayayagoda/internal/config"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/routes" // Убедитесь, что импортируете пакет с новой функцией SetupRouter
	"tipodikayayagoda/internal/storage"
	"tipodikayayagoda/internal/utils"
	"tipodikayayagoda/pkg/database"
)

func main() {
	handler.GlobalHub = handler.NewHub()

	cfg := config.LoadConfig()
	storage.InitMinio(cfg)
	db := database.Conn(cfg)

	if err := models.InitRoles(db); err != nil {
		log.Fatalf("Не удалось инициализировать роли: %v", err)
	}

	utils.Init(cfg.JwtSecret)
	repository.Init(db)

	router := routes.Routes()

	fmt.Println("Starting Gin server at port 8080")

	// 2. Запускаем сервер через Gin.
	// Метод Run сам под капотом вызывает http.ListenAndServe(":8080", router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
