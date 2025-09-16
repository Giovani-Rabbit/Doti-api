package taskdto

import (
	"time"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
)

type TaskDTO struct {
	Id          int32     `json:"id"`
	ModuleId    int32     `json:"module_id"`
	Name        string    `json:"name"`
	IsCompleted bool      `json:"is_completed"`
	Position    int32     `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TasksResponse struct {
	Tasks []TaskDTO `json:"tasks"`
}

func NewTaskDTOFromDomain(domain taskdomain.Task) *TaskDTO {
	return &TaskDTO{
		Id:          domain.ID(),
		ModuleId:    domain.ModuleId(),
		Name:        domain.Name(),
		IsCompleted: domain.IsCompleted(),
		Position:    domain.Position(),
		CreatedAt:   domain.CreatedAt(),
		UpdatedAt:   domain.UpdatedAt(),
	}
}

func NewTaskListDTOFromDomain(domains []taskdomain.Task) *TasksResponse {
	taskList := make([]TaskDTO, len(domains))

	for i, t := range domains {
		taskList[i] = *NewTaskDTOFromDomain(t)
	}

	return &TasksResponse{Tasks: taskList}
}
