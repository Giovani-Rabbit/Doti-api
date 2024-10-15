package database

import (
	"database/sql"
	"fmt"

	"github.com/Giovani-Coelho/Doti-API/config"
)

type postgresDB struct {
	Db *sql.DB
}

func NewPostgresDB() (*postgresDB, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Env.DbUser,
		config.Env.DbPassword,
		config.Env.DbServer,
		config.Env.DbPort,
		config.Env.DbDatabase,
	)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return &postgresDB{Db: db}, nil
}
