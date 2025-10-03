package dicase

import (
	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type TaskUseCase struct {
	Create           taskcase.Create
	Delete           taskcase.Delete
	UpdatePosition   taskcase.UpdatePosition
	UpdateCompletion taskcase.UpdateCompletion
}

func NewTaskUseCase(
	tr repository.TaskRepository,
	mr repository.ModuleRepository,
) *TaskUseCase {
	return &TaskUseCase{
		Create:           taskcase.NewCreateTaskUseCase(tr, mr),
		Delete:           taskcase.NewDeleteTaskUseCase(tr),
		UpdatePosition:   taskcase.NewUpdateTaskPosition(tr),
		UpdateCompletion: taskcase.NewUpdateCompletion(tr),
	}
}
