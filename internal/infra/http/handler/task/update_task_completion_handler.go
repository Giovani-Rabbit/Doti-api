package taskhdl

import (
	"net/http"
	"strconv"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type UpdateCompletion interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type updateCompletion struct {
	updateTaskCompletion taskcase.UpdateCompletion
}

func NewUpdateCompletionHandler(
	updateTaskCompletion taskcase.UpdateCompletion,
) UpdateCompletion {
	return &updateCompletion{
		updateTaskCompletion: updateTaskCompletion,
	}
}

func (uc *updateCompletion) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	taskId, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		res.Error(resp.NewInvalidPathValue(err))
		return
	}

	var body taskdto.UpdateCompletionDTO
	if !res.DecodeJSONBody(r, &body) {
		return
	}

	err = uc.updateTaskCompletion.Execute(r.Context(), int32(taskId), body.IsComplete)
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
