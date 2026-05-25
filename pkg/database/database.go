package database

import (
	"database/sql"
	"log"
	"tipodikayayagoda/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Conn(cfg *config.Config) *sql.DB {
	connStr := cfg.ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Не удалось создать драйвер миграций: %v", err)
	}

	// Инициализируем мигратор.
	// "file://migrations" указывает на папку внутри Docker-контейнера
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Ошибка инициализации мигратора: %v", err)
	}

	// Накатываем миграции до самой последней версии
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}
	return db
}
