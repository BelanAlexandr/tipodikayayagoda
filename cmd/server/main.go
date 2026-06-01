package main

import (
	"fmt"
	"log"
	"net/http"
	"tipodikayayagoda/internal/config"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/rout"
	"tipodikayayagoda/internal/storage"
	"tipodikayayagoda/internal/utils"
	"tipodikayayagoda/pkg/database"
)

func main() {
	handler.GlobalHub = handler.NewHub()

	fmt.Println("Starting server at port 8080")

	cfg := config.LoadConfig()
	storage.InitMinio(cfg)
	db := database.Conn(cfg)
	if err := models.InitRoles(db); err != nil {
		log.Fatalf("Не удалось инициализировать роли: %v", err)
	}
	utils.Init(cfg.JwtSecret)
	repository.Init(db)
	rout.Routes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
