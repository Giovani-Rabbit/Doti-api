package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type GetTasksByModule interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type getTasksByModule struct {
	getTasks modulecase.GetTasks
}

func NewGetTasksByModuleHandler(
	g modulecase.GetTasks,
) GetTasksByModule {
	return &getTasksByModule{
		getTasks: g,
	}
}

func (gt *getTasksByModule) Execute(w http.ResponseWriter, r *http.Request) {
	moduleId := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	tasks, err := gt.getTasks.Execute(r.Context(), moduleId)
	if err != nil {
		res.Error(err)
		return
	}

	modulesListResponse := taskdto.NewTaskListDTOFromDomain(tasks)
	res.AddBody(modulesListResponse)
	res.Write(200)
}
