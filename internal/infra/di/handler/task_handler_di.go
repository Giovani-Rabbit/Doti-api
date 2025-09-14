package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
)

type TaskHandlers struct {
	Create taskhdl.Create
}

func NewTaskHandler(taskCases *dicase.TaskUseCase) *TaskHandlers {
	return &TaskHandlers{
		Create: taskhdl.NewCreateTaskHandler(taskCases.Create),
	}
}
