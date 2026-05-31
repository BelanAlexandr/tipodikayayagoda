package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnectionString string
	JwtSecret        string
	Endpoint         string
	AccesKey         string
	SecretAccesKey   string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		ConnectionString: os.Getenv("DB_CONN_STR"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		Endpoint:         os.Getenv("MINIO_ENDPOINT"),
		AccesKey:         os.Getenv("ACCES_KEY"),
		SecretAccesKey:   os.Getenv("SECRET_ACCES_KEY"),
	}
}
