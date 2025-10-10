package taskdto

import taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"

type UpdatePositionDTO struct {
	MovedTasks []taskdomain.TaskPositionParams `json:"movedTasks"`
}

type UpdateCompletionDTO struct {
	IsComplete bool `json:"isComplete"`
}

type UpdateTaskNameHttpBody struct {
	TaskName string `json:"taskName"`
}
