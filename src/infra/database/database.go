package database

import (
	"database/sql"
)

type postgresDB struct {
	Db *sql.DB
}

func NewPostgresDB() (*postgresDB, error) {
	connection := "postgres://postgres:1234@localhost:5432/doti_db?sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return &postgresDB{Db: db}, nil
}
