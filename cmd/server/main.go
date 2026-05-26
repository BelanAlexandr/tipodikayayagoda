package main

import (
	"fmt"
	"net/http"
	"tipodikayayagoda/internal/config"
	"tipodikayayagoda/internal/handler"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/rout"
	"tipodikayayagoda/internal/utils"
	"tipodikayayagoda/pkg/database"
)

func main() {
	handler.GlobalHub = handler.NewHub()
	fmt.Println("Starting server at port 8080")
	rout.Routes()
	cfg := config.LoadConfig()
	db := database.Conn(cfg)
	utils.Init(cfg.JwtSecret)
	repository.Init(db)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
