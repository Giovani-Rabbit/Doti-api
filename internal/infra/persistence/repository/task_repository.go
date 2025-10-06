package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/mapper"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(ctx context.Context, task taskdomain.Task) (taskdomain.Task, error)
	CheckExists(ctx context.Context, taskId int32) (bool, error)
	Delete(ctx context.Context, taskId int32) error
	FindById(ctx context.Context, taskId int32) (taskdomain.Task, error)
	FindOwnerIdByTaskId(ctx context.Context, taskId int32) (uuid.UUID, error)
	PositionExists(ctx context.Context, moduleId, position int32) (bool, error)
	ListByModuleId(ctx context.Context, moduleId int32) ([]taskdomain.Task, error)
	UpdatePosition(ctx context.Context, tasks []taskdomain.TaskPositionParams) error
	UpdateCompletion(ctx context.Context, taskId int32, isComplete bool) error
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

func (tr *taskRepository) CheckExists(
	ctx context.Context, taskId int32,
) (bool, error) {
	exists, err := tr.queries.CheckTaskExists(ctx, taskId)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (tr *taskRepository) Delete(
	ctx context.Context, taskId int32,
) error {
	err := tr.queries.DeleteTask(ctx, taskId)
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) FindById(
	ctx context.Context, taskId int32,
) (taskdomain.Task, error) {
	task, err := tr.queries.FindTaskById(ctx, taskId)
	if err != nil {
		return nil, err
	}

	return mapper.SqlcTaskToDomain(task), nil
}

func (tr *taskRepository) FindOwnerIdByTaskId(
	ctx context.Context, taskId int32,
) (uuid.UUID, error) {
	userId, err := tr.queries.FindOwnerIdByTaskId(ctx, taskId)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userId, nil
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

func (tr *taskRepository) UpdateCompletion(
	ctx context.Context, taskId int32, isComplete bool,
) error {
	err := tr.queries.UpdateTaskCompletion(ctx, sqlc.UpdateTaskCompletionParams{
		ID:          taskId,
		IsCompleted: isComplete,
	})
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) UpdatePosition(
	ctx context.Context, movedTasks []taskdomain.TaskPositionParams,
) error {
	params := []any{}
	values := []string{}

	for i, task := range movedTasks {
		idx := i*2 + 1
		values = append(values, fmt.Sprintf("($%d::int, $%d::int)", idx, idx+1))
		params = append(params, task.Id, task.Position)
	}

	query := fmt.Sprintf(`
		UPDATE tasks t
		SET position = v.position
		FROM (VALUES %s) AS v(id, position)
		WHERE t.id = v.id
	`, strings.Join(values, ","))

	_, err := tr.db.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}
