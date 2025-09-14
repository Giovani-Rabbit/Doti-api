package taskdto

import (
	"time"

	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
)

type TaskDTO struct {
	Id        int32     `json:"id"`
	ModuleId  int32     `json:"module_id"`
	Name      string    `json:"name"`
	Position  int32     `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskDTOFromDomain(domain taskdomain.Task) *TaskDTO {
	return &TaskDTO{
		Id:        domain.ID(),
		ModuleId:  domain.ModuleId(),
		Name:      domain.Name(),
		Position:  domain.Position(),
		CreatedAt: domain.CreatedAt(),
		UpdatedAt: domain.UpdatedAt(),
	}
}
