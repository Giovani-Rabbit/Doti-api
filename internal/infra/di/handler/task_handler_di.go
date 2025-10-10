package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
)

type TaskHandlers struct {
	Create           taskhdl.Create
	Delete           taskhdl.Delete
	UpdateName       taskhdl.UpdateName
	UpdatePosition   taskhdl.UpdatePosition
	UpdateCompletion taskhdl.UpdateCompletion
}

func NewTaskHandler(taskCases *dicase.TaskUseCase) *TaskHandlers {
	return &TaskHandlers{
		Create:           taskhdl.NewCreateTaskHandler(taskCases.Create),
		Delete:           taskhdl.NewDeleteTaskHandler(taskCases.Delete),
		UpdateName:       taskhdl.NewUpdateTaskNameHandler(taskCases.UpdateName),
		UpdatePosition:   taskhdl.NewUpdatePositionHandler(taskCases.UpdatePosition),
		UpdateCompletion: taskhdl.NewUpdateCompletionHandler(taskCases.UpdateCompletion),
	}
}
