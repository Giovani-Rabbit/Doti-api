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
	ListByModuleId(ctx context.Context, moduleId int32) ([]taskdomain.Task, error)
	GetTaskByID(ctx context.Context, taskId int32) (taskdomain.Task, error)
	GetTaskByPosition(ctx context.Context, moduleId, position int32) (taskdomain.Task, error)
	UpdateTaskPosition(ctx context.Context, taskId, position int32) error
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
	tx, err := tr.db.Begin()
	if err != nil {
		return nil, err
	}

	var committed bool
	q := tr.queries.WithTx(tx)

	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	t, err := q.CreateTask(ctx, sqlc.CreateTaskParams{
		ModuleID:    task.ModuleId(),
		Name:        task.Name(),
		IsCompleted: task.IsCompleted(),
		Position:    task.Position(),
		CreatedAt:   task.CreatedAt(),
		UpdatedAt:   task.UpdatedAt(),
	})
	if err != nil {
		return nil, err
	}

	err = q.CreateTaskDetails(ctx, sqlc.CreateTaskDetailsParams{
		TaskID: t.ID,
		Description: sql.NullString{
			String: "",
			Valid:  false,
		},
	})
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	committed = true

	return mapper.SqlcTaskToDomain(t), nil
}

func (tr *taskRepository) ListByModuleId(
	ctx context.Context, moduleId int32,
) ([]taskdomain.Task, error) {
	tasks, err := tr.queries.ListTasksByModuleId(ctx, moduleId)
	if err != nil {
		return nil, err
	}

	return mapper.SqlcTaskListToDomain(&tasks), nil
}

func (tr *taskRepository) GetTaskByID(
	ctx context.Context,
	taskId int32,
) (taskdomain.Task, error) {
	task, err := tr.queries.GetTaskbyId(ctx, taskId)
	if err != nil {
		return nil, err
	}

	taskEntity := mapper.SqlcTaskToDomain(task)

	return taskEntity, nil
}

func (tr *taskRepository) GetTaskByPosition(
	ctx context.Context,
	moduleId, position int32,
) (taskdomain.Task, error) {
	task, err := tr.queries.GetTaskByPosition(ctx, sqlc.GetTaskByPositionParams{
		ModuleID: moduleId,
		Position: position,
	})
	if err != nil {
		return nil, err
	}

	taskEntity := mapper.SqlcTaskToDomain(task)

	return taskEntity, nil
}

func (tr *taskRepository) UpdateTaskPosition(
	ctx context.Context, taskId, position int32,
) error {
	err := tr.queries.UpdateTaskPosition(ctx, sqlc.UpdateTaskPositionParams{
		ID:       taskId,
		Position: position,
	})
	if err != nil {
		return err
	}

	return nil
}
