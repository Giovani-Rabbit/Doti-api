package taskdto

import taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"

type UpdatePositionDTO struct {
	MovedTasks []taskdomain.TaskPositionParams `json:"movedTasks"`
}
