package main

import (
	"fmt"
	"log"
	"tipodikayayagoda/internal/config"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/routes"
	"tipodikayayagoda/internal/storage"
	"tipodikayayagoda/internal/utils"
	"tipodikayayagoda/pkg/database"
)

func main() {
	handler.GlobalHub = handler.NewHub()

	cfg := config.LoadConfig()
	storage.InitMinio(cfg)
	db := database.Conn(cfg)
	utils.Init(cfg.JwtSecret)
	repository.Init(db)

	router := routes.Routes()

	fmt.Println("Starting Gin server at port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
