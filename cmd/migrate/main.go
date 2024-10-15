package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/doti_db?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/infra/database/store/pgstore/migrations/", // migrations path
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	switch cmd := os.Args[len(os.Args)-1]; cmd {
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		log.Fatalf("Invalid option: %s", cmd)
	}
}
