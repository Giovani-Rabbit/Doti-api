package main

import (
	"log"
	"os"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	postgresDB, err := database.NewPostgresDB()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(postgresDB.Db, &postgres.Config{})

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
