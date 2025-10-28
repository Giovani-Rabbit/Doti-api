package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	detailshdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task_details"
)

type TaskDetailsHandler struct {
	UpdateDescription detailshdl.UpdateDescription
}

func NewTaskDetailsHandler(
	detailscases *dicase.TaskDetailsUseCase,
) *TaskDetailsHandler {
	return &TaskDetailsHandler{
		UpdateDescription: detailshdl.NewUpdateDescriptionHandler(
			detailscases.UpdateDescription,
		),
	}
}
