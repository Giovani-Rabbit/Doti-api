package dicase

import (
	detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type TaskDetailsUseCase struct {
	UpdateDescription    detailscase.UpdateDescription
	UpdatePomodoroTarget detailscase.UpdatePomodoroTarget
}

func NewTaskDetailsUseCase(
	td repository.TaskDetailsRepository,
) *TaskDetailsUseCase {
	return &TaskDetailsUseCase{
		UpdateDescription:    detailscase.NewTaskDetailsUseCase(td),
		UpdatePomodoroTarget: detailscase.NewTaskDetailsUpdatePomodoroTarget(td),
	}
}
