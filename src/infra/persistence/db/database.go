package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Giovani-Coelho/Doti-API/config"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Env.DbUser,
		config.Env.DbPassword,
		config.Env.DbServer,
		config.Env.DbPort,
		config.Env.DbDatabase,
	)

	conn, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	return conn, err
}
