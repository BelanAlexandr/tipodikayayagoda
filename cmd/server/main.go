package main

import (
	"fmt"
	"net/http"
	"tipodikayayagoda/internal/config"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/rout"
	"tipodikayayagoda/pkg/database"
)

func main() {
	fmt.Println("Starting server at port 8080")
	rout.Routes()
	cfg := config.LoadConfig()
	db := database.Conn(cfg)
	repository.Init(db)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
