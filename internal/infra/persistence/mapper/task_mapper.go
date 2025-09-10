package mapper

import (
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

func SqlcTaskToDomain(sqlcTask sqlc.Task) taskdomain.Task {
	return taskdomain.NewFromDB(
		sqlcTask.ID,
		sqlcTask.ModuleID,
		sqlcTask.Name,
		sqlcTask.IsCompleted,
		sqlcTask.Position,
		sqlcTask.CreatedAt,
		sqlcTask.UpdatedAt,
	)
}
