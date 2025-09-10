package repository

import (
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

type TaskRepository interface{}

type taskRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}
