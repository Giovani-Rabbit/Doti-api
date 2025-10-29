package repository

import (
	"context"
	"database/sql"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

type TaskDetailsRepository interface {
	UpdateDescription(ctx context.Context, taskId int32, description string) error
	UpdatePomodoroTarget(ctx context.Context, taskId int32, target int) error
}

type taskDetailsRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewTaskDetailsRepository(
	db *sql.DB,
) TaskDetailsRepository {
	return &taskDetailsRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

func (td *taskDetailsRepository) UpdateDescription(
	ctx context.Context, taskId int32, description string,
) error {
	rows, err := td.queries.UpdateTaskDetailsDescription(ctx,
		sqlc.UpdateTaskDetailsDescriptionParams{
			TaskID: taskId,
			Description: sql.NullString{
				String: description,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return err
	}

	if rows == 0 {
		return taskdomain.ErrCouldNotFindTask()
	}

	return nil
}

func (td *taskDetailsRepository) UpdatePomodoroTarget(
	ctx context.Context, taskId int32, target int,
) error {
	rows, err := td.queries.UpdateTaskDetailsPomodoroTarget(ctx,
		sqlc.UpdateTaskDetailsPomodoroTargetParams{
			TaskID: taskId,
			PomodoroTarget: sql.NullInt32{
				Int32: int32(target),
				Valid: true,
			},
		},
	)

	if err != nil {
		return err
	}

	if rows == 0 {
		return taskdomain.ErrCouldNotFindTask()
	}

	return nil
}
