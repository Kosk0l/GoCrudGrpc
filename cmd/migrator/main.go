package main

import (
	"errors"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// go run ./cmd/migrator --dsn="postgres://postgresCrud:qwerty@localhost:5433/postgresCrud?sslmode=disable" --migrations-path=./migrations
func main() {
	var (
		dsn            string
		migrationsPath string
	)

	flag.StringVar(&dsn, "dsn", "", "Postgres DSN connection string")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations folder")
	flag.Parse()

	if dsn == "" {
		log.Fatal("missing required flag: --dsn")
	}
	if migrationsPath == "" {
		log.Fatal("missing required flag: --migrations-path")
	}

	m, err := migrate.New("file://"+migrationsPath, dsn)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply — database already up to date")
			return
		}
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migrations applied successfully")
}