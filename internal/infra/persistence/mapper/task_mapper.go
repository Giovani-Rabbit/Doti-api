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

func SqlcTaskListToDomain(sqlcTaskList *[]sqlc.Task) (taskList []taskdomain.Task) {
	if sqlcTaskList == nil {
		return []taskdomain.Task{}
	}

	taskList = make([]taskdomain.Task, 0, len(*sqlcTaskList))

	for _, t := range *sqlcTaskList {
		taskList = append(taskList, taskdomain.NewFromDB(
			t.ID,
			t.ModuleID,
			t.Name,
			t.IsCompleted,
			t.Position,
			t.CreatedAt,
			t.UpdatedAt,
		))
	}
	return
}
