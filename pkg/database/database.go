package database

import (
	"database/sql"
	"tipodikayayagoda/internal/config"

	_ "github.com/lib/pq"
)

func Conn(cfg *config.Config) *sql.DB {
	connStr := cfg.ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
