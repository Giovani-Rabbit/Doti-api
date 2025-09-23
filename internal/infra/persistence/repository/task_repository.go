package repository

import (
	"context"
	"database/sql"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/mapper"
)

type TaskRepository interface {
	Create(ctx context.Context, task taskdomain.Task) (taskdomain.Task, error)
	PositionExists(ctx context.Context, moduleId, position int32) (bool, error)
	ListByModuleId(ctx context.Context, moduleId int32) ([]taskdomain.Task, error)
	UpdatePosition(ctx context.Context, tasks []taskdto.TaskPositionParams) error
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

func (tr *taskRepository) PositionExists(
	ctx context.Context, moduleId, position int32,
) (bool, error) {
	exists, err := tr.queries.TaskPositionExists(ctx, sqlc.TaskPositionExistsParams{
		ModuleID: moduleId,
		Position: position,
	})
	if err != nil {
		return false, err
	}

	return exists, nil
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

func (tr *taskRepository) UpdatePosition(
	ctx context.Context, tasks []taskdto.TaskPositionParams,
) error {
	err := tr.queries.SwapTaskPosition(ctx, sqlc.SwapTaskPositionParams{
		Position: tasks[0].TaskId, Position_2: tasks[0].Position, // taskId, position
		Position_3: tasks[1].TaskId, Position_4: tasks[1].Position, // taskId, position
	})

	if err != nil {
		return err
	}

	return nil
}
