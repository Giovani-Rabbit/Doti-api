package repository

import (
	"context"
	"database/sql"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/mapper"
)

type TaskRepository interface {
	Create(ctx context.Context, task taskdomain.Task) (taskdomain.Task, error)
}

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

func (tr *taskRepository) Create(
	ctx context.Context,
	task taskdomain.Task,
) (taskdomain.Task, error) {
	createdTask, err := tr.queries.CreateTask(ctx, sqlc.CreateTaskParams{
		ModuleID:    task.ModuleID(),
		Name:        task.Name(),
		IsCompleted: task.IsCompleted(),
		Position:    task.Position(),
		CreatedAt:   task.CreatedAt(),
		UpdatedAt:   task.UpdatedAt(),
	})
	if err != nil {
		return nil, err
	}

	return mapper.SqlcTaskToDomain(createdTask), nil
}
