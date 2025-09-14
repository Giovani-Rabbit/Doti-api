package taskhdl

import (
	"net/http"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type Create interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type create struct {
	createTask taskcase.Create
}

func NewCreateTaskHandler(
	createTask taskcase.Create,
) Create {
	return &create{
		createTask: createTask,
	}
}

func (ct *create) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var taskFields taskdto.CreateTaskDTO
	if !res.DecodeJSONBody(r, &taskFields) {
		return
	}

	taskEntity := taskdomain.New(
		taskFields.ModuleId,
		taskFields.TaskName,
		taskFields.Position,
	)

	createdTask, err := ct.createTask.Execute(r.Context(), taskEntity)
	if err != nil {
		res.Error(err)
		return
	}

	taskResponse := taskdto.NewTaskDTOFromDomain(createdTask)
	res.AddBody(taskResponse)
	res.Write(http.StatusCreated)
}
