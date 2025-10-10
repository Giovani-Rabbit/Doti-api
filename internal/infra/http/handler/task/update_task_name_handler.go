package taskhdl

import (
	"net/http"
	"strconv"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type UpdateName interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type updateName struct {
	updateName taskcase.UpdateName
}

func NewUpdateTaskNameHandler(
	updateNameUseCase taskcase.UpdateName,
) UpdateName {
	return &updateName{
		updateName: updateNameUseCase,
	}
}

func (un *updateName) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)
	taskIdStr := r.PathValue("id")

	taskId, err := strconv.ParseInt(taskIdStr, 10, 32)
	if err != nil {
		res.Error(resp.NewInvalidPathValue(err))
		return
	}

	var body taskdto.UpdateTaskNameHttpBody
	if !res.DecodeJSONBody(r, &body) {
		return
	}

	err = un.updateName.Execute(r.Context(), int32(taskId), body.TaskName)
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
